package ucumapd

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/cockroachdb/apd/v3"
	"github.com/iimos/ucum"
)

// Converter converts apd.Decimal values from one UCUM unit to another.
type Converter interface {
	ConvDecimal(val *apd.Decimal, from, to string, ctx *apd.Context) (*apd.Decimal, error)
}

// Conv converts apd.Decimal values from one UCUM unit to another. It implements Converter.
type Conv struct {
	conv ucum.ConverterBigRat
}

var _ Converter = (*Conv)(nil)

// NewConverter creates converter for apd.Decimal values from one UCUM unit to another.
func NewConverter(conv ucum.ConverterBigRat) *Conv {
	return &Conv{conv: conv}
}

// ConvDecimal converts apd.Decimal values from one UCUM unit to another using the provided context.
// If ctx is nil the conversion runs with apd.BaseContext (precision 0, RoundHalfUp). That context
// does not round repeating decimals, so Quo will return an error whenever the result requires
// infinitely many digits (for example converting 1/3).
func (c *Conv) ConvDecimal(val *apd.Decimal, from, to string, ctx *apd.Context) (*apd.Decimal, error) {
	rat, err := apdDecimalToRat(val)
	if err != nil {
		return nil, err
	}
	res, err := c.conv.ConvBigRat(rat, from, to)
	if err != nil {
		return nil, err
	}
	return bigRatToAPDDecimal(res, ctx)
}

// ConvDecimal converts apd.Decimal values from one UCUM unit to another.
// It uses ucum.DefaultConverter under the hood.
func ConvDecimal(val *apd.Decimal, from, to string, ctx *apd.Context) (*apd.Decimal, error) {
	conv := NewConverter(ucum.DefaultConverter)
	return conv.ConvDecimal(val, from, to, ctx)
}

func apdDecimalToRat(val *apd.Decimal) (*big.Rat, error) {
	if val == nil {
		return nil, errors.New("ucumapd: apd decimal is nil")
	}
	if val.Form != apd.Finite {
		return nil, fmt.Errorf("ucumapd: apd decimal must be finite, got %v", val.Form)
	}

	num := val.Coeff.MathBigInt()
	if val.Negative && num.Sign() != 0 {
		num.Neg(num)
	}
	den := big.NewInt(1)
	switch {
	case val.Exponent > 0:
		num.Mul(num, pow10Int(int64(val.Exponent)))
	case val.Exponent < 0:
		den = pow10Int(int64(-val.Exponent))
	}
	return new(big.Rat).SetFrac(num, den), nil
}

func bigRatToAPDDecimal(val *big.Rat, ctx *apd.Context) (*apd.Decimal, error) {
	if val == nil {
		return nil, errors.New("ucumapd: rational value is nil")
	}
	ctx = ensureContext(ctx)

	num := apd.NewWithBigInt(new(apd.BigInt).SetMathBigInt(val.Num()), 0)
	den := apd.NewWithBigInt(new(apd.BigInt).SetMathBigInt(val.Denom()), 0)
	res := new(apd.Decimal)
	if _, err := ctx.Quo(res, num, den); err != nil {
		return nil, fmt.Errorf("ucumapd: convert result: %w", err)
	}
	return res, nil
}

func ensureContext(ctx *apd.Context) *apd.Context {
	if ctx != nil {
		return ctx
	}
	return &apd.BaseContext
}

var bigTen = big.NewInt(10)

func pow10Int(n int64) *big.Int {
	if n < 0 {
		panic("pow10Int negative exponent")
	}
	if n == 0 {
		return big.NewInt(1)
	}
	return new(big.Int).Exp(bigTen, big.NewInt(n), nil)
}
