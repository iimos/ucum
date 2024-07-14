package ucumparser

import (
	"fmt"
	"github.com/iimos/ucum/internal/data"
	"github.com/iimos/ucum/internal/types"
	"math/big"
	"strconv"
)

func Parse(unit []byte) (types.Unit, error) {
	if len(unit) == 0 {
		return types.Unit{}, fmt.Errorf("ucum: empty unit")
	}
	p := newParser(unit)
	p.readTerm(false, 1)
	if p.error != nil {
		return types.Unit{}, p.error
	}
	u := types.Unit{
		Orig:       string(unit),
		Components: p.components,
		Coeff:      p.coef,
	}
	if err := validate(u); err != nil {
		return types.Unit{}, err
	}
	return u, nil
}

type parser struct {
	buf        []byte
	head       int
	tail       int
	components map[types.ComponentKey]int // <unit atom, annotation> -> exponent
	coef       *big.Rat
	error      error
}

func newParser(unit []byte) *parser {
	return &parser{
		buf:        unit,
		head:       0,
		tail:       len(unit),
		components: make(map[types.ComponentKey]int),
		coef:       big.NewRat(1, 1),
	}
}

// https://ucum.org/ucum#section-Syntax-Rules
func (p *parser) readTerm(insideBrackets bool, termExponent int) {
	start := p.head
	componentExponent := 1
	for p.error == nil {
		// The division operator can be used as a binary and unary operator, i.e. a leading solidus will invert the unit that directly follows it.
		if p.head == start && p.head < p.tail && p.buf[p.head] == '/' {
			componentExponent = -1
			p.head++
		}

		p.readComponent(termExponent * componentExponent)

		b := p.readByte()
		switch b {
		case '.':
			componentExponent = 1
		case '/':
			componentExponent = -1
		case ')':
			if insideBrackets {
				return
			}
			p.unreadByte(b)
			p.reportError(p.head, `unexpected ")"`)
			return
		case 0: // EOF
			if insideBrackets {
				p.reportError(-1, `unexpected end, missing ")"`)
			}
			return
		default:
			p.unreadByte(b)
			p.reportError(p.head, `unexpected symbol "%c"`, b)
		}
	}
}

// https://ucum.org/ucum#section-Syntax-Rules
func (p *parser) readComponent(exponent int) {
	b := p.readByte()
	switch b {
	case '(':
		p.readTerm(true, exponent)
		return
	case 0: // EOF
		p.reportError(-1, `unexpected end`)
		return
	default:
		p.unreadByte(b)
		p.readAnnotatable(exponent)
	}
}

// https://ucum.org/ucum#section-Syntax-Rules
func (p *parser) readAnnotatable(exponent int) {
	origHead := p.head
	if p.head == p.tail {
		p.reportError(p.head, `unexpected end`)
		return
	}

	coef := big.NewRat(1, 1)
	multiplierOk := false

	atom, atomOk := p.readAtom()
	if p.error != nil {
		return
	}
	if atomOk {
		if exp, ok := p.tryReadExponent(); ok {
			exponent *= exp
		}
		coef.Set(atom.Magnitude)
	} else {
		// exponent without unit is just a number
		if num, ok := p.readDigits(1); ok {
			multiplierOk = true
			coef = new(big.Rat).SetFrac64(int64(num), 1)
		}
	}

	annotation := p.readAnnotation()

	// if nothing meaningful was found it's an error
	if !multiplierOk && !atomOk && len(annotation) == 0 {
		p.reportError(origHead, `unexpected symbol "%c"`, p.buf[origHead])
		return
	}

	pow(coef, exponent)
	p.coef.Mul(p.coef, coef) // combine global coefficient with local one

	if atomOk || len(annotation) > 0 {
		key := types.ComponentKey{
			AtomCode:   atom.Code,
			Annotation: string(annotation),
		}
		p.components[key] += exponent
	}
}

func pow(base *big.Rat, exp int) {
	if exp < 0 {
		base = base.Inv(base)
		exp = -exp
	} else if exp == 0 {
		base.SetFrac64(0, 1)
	}
	cpy := new(big.Rat).Set(base)
	for i := exp; i > 1; i-- {
		base.Mul(base, cpy)
	}
}

func (p *parser) readAtom() (types.Atom, bool) {

	// A terminal unit symbol can not consist of only digits (‘0’–‘9’) because those digit strings
	// are interpreted as positive integer numbers. However, a symbol “10*” is allowed because it ends
	// with a non-digit allowed to be part of a symbol.

	from := p.head

	// skip digits at the beginning of the unit atom
	for p.head < p.tail {
		switch p.buf[p.head] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			p.head++
			continue
		}
		break
	}

	endOfDigits := p.head

loop:
	for p.head < p.tail {
		switch p.buf[p.head] {
		case '.', '/', '(', ')', '{', '}', '+', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			break loop
		case '[':
			p.skipSquareBrackets()
			continue
		case ' ':
			p.reportError(-1, "spaces are not allowed")
			return types.Atom{}, false
		}
		p.head++
	}

	// if the unit atom contains only digits, it is a number and not a unit so roll back
	if endOfDigits == p.head {
		p.head = from
		return types.Atom{}, false
	}

	unit, ok := data.Atoms[string(p.buf[from:p.head])]
	if !ok {
		p.reportError(from, "unknown unit %q", string(p.buf[from:p.head]))
	}
	return unit, true
}

func (p *parser) skipSquareBrackets() {
	openBracket := p.head
	for {
		switch p.readByte() {
		case ']':
			return
		case 0: // EOF
			p.reportError(openBracket, "unclosed square bracket")
			return
		}
	}
}

func (p *parser) tryReadExponent() (exp int, ok bool) {
	b := p.readByte()
	switch b {
	case '+':
		return p.readDigits(1)
	case '-':
		return p.readDigits(-1)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		p.unreadByte(b)
		return p.readDigits(1)
	default:
		p.unreadByte(b)
		return 0, false
	}
}

func (p *parser) reportError(position int, msg string, args ...any) {
	if p.error == nil {
		if len(args) > 0 {
			msg = fmt.Sprintf(msg, args...)
		}
		if position >= 0 {
			p.error = fmt.Errorf("ucum: %s at position %d", msg, position)
		} else {
			p.error = fmt.Errorf("ucum: %s", msg)
		}
	}
}

func (p *parser) readDigits(sign int) (num int, ok bool) {
	start := p.head
	for p.head < p.tail {
		d := p.buf[p.head]
		switch d {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if num > safeIntToMultiple10() {
				p.reportError(start, "number too large")
				return 0, false
			}
			num = 10*num + int(d-'0')
			ok = true
			p.head++
			continue
		}
		break
	}
	return sign * num, ok
}

func safeIntToMultiple10() int {
	if strconv.IntSize == 32 {
		return 0xffffffff/10 - 1
	}
	return 0xffffffffffffffff/10 - 1
}

// https://ucum.org/ucum#section-Syntax-Rules
func (p *parser) readAnnotation() []byte {
	from := p.head

	if b := p.readByte(); b != '{' {
		p.unreadByte(b)
		return nil
	}

	for {
		switch p.readByte() {
		case '}':
			return p.buf[from:p.head]
		case 0: // EOF
			p.reportError(p.head, "unterminated annotation, \"}\" expected")
			return nil
		}
	}
}

func (p *parser) readByte() byte {
	if p.head == p.tail {
		return 0
	}
	b := p.buf[p.head]
	p.head++
	return b
}

func (p *parser) unreadByte(c byte) {
	if p.error == nil && c != 0 { // c == 0 means EOF
		p.head--
	}
}
