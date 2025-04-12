package types

import (
	"math/big"
	"slices"
	"strings"
)

// CanonicalString returns the canonical string representation of the unit. It's used in tests only.
//
// Example: kg.m/s2 -> 1000⋅g⋅m⋅s⁻²
func (u *Unit) CanonicalString() string {
	if len(u.Components) == 0 {
		return u.Coeff.RatString()
	}

	var ret string
	if u.Coeff != nil && u.Coeff.Cmp(big.NewRat(1, 1)) != 0 {
		ret += u.Coeff.RatString()
	}

	keys := componentsList(u)
	slices.SortFunc(keys, func(a, b ComponentKey) int {
		exponentsDiff := u.Components[b] - u.Components[a]
		if exponentsDiff != 0 {
			return exponentsDiff
		}
		c := strings.Compare(a.AtomCode, b.AtomCode)
		if c != 0 {
			return c
		}
		return strings.Compare(a.Annotation, b.Annotation)
	})

	for i, key := range keys {
		exponent := u.Components[key]
		if key.AtomCode != "" {
			if i > 0 || ret != "" {
				ret += "⋅"
			}
			ret += componentString(key.AtomCode, exponent, key.Annotation)
		} else {
			ret += key.Annotation
		}
	}
	return ret
}

// componentString returns the UTF-8 string representation of the expression component.
func componentString(atomCode string, exponent int, annotation string) string {
	ret := make([]rune, 0, len(atomCode)+len(annotation)+2) // +2 for exponent
	ret = append(ret, []rune(atomCode)...)
	if exponent != 1 {
		ret = appendExponent(ret, exponent)
	}
	ret = append(ret, []rune(annotation)...)
	return string(ret)
}

var superscriptNums = [...]rune{'⁰', '¹', '²', '³', '⁴', '⁵', '⁶', '⁷', '⁸', '⁹'}

func appendExponent(dst []rune, exp int) []rune {
	if exp < 0 {
		dst = append(dst, '⁻')
		exp = -exp
	}
	if exp < 10 {
		dst = append(dst, superscriptNums[exp])
		return dst
	}

	digits := make([]rune, 0, 4)
	for exp > 0 {
		digits = append(digits, superscriptNums[exp%10])
		exp = exp / 10
	}
	slices.Reverse(digits)
	return append(dst, digits...)
}

func componentsList(u *Unit) []ComponentKey {
	r := make([]ComponentKey, 0, len(u.Components))
	for k := range u.Components {
		r = append(r, k)
	}
	return r
}
