package ucumdecimal

import (
	"github.com/iimos/ucum"
	"github.com/shopspring/decimal"
)

// ConverterDecimal converts shopspring/decimal.Decimal values from one UCUM unit to another.
type ConverterDecimal interface {
	ConvDecimal(val decimal.Decimal, from, to string, precision int32) (decimal.Decimal, error)
}

// ConverterDecimal creates converter for shopspring/decimal.Decimal values from one UCUM unit to another.
func NewConverter(conv ucum.ConverterBigRat) *Conv {
	return &Conv{
		conv: conv,
	}
}

// Conv converts shopspring/decimal.Decimal values from one UCUM unit to another.
// It implements ConverterDecimal.
type Conv struct {
	conv ucum.ConverterBigRat
}

var _ ConverterDecimal = (*Conv)(nil)

// ConvDecimal converts shopspring/decimal.Decimal values from one UCUM unit to another with a given precision.
func (c *Conv) ConvDecimal(val decimal.Decimal, from, to string, precision int32) (decimal.Decimal, error) {
	res, err := c.conv.ConvBigRat(val.Rat(), from, to)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return decimal.NewFromBigRat(res, precision), nil
}

// ConvDecimal converts shopspring/decimal.Decimal values from one UCUM unit to another with a given precision.
func ConvDecimal(val decimal.Decimal, from, to string, precision int32) (decimal.Decimal, error) {
	conv := NewConverter(ucum.DefaultConverter)
	return conv.ConvDecimal(val, from, to, precision)
}
