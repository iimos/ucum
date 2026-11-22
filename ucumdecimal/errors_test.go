package ucumdecimal

import (
	"testing"

	"github.com/iimos/ucum"
	"github.com/shopspring/decimal"
)

func TestConvDecimalErrors(t *testing.T) {
	tests := [][2]string{
		{"m", "s"},          // incompatible dimensions
		{"m", "m/s"},        // incompatible dimensions
		{"m", "not_a_unit"}, // invalid target unit
		{"not_a_unit", "m"}, // invalid source unit
		{"", "m"},           // empty source unit
		{"m", ""},           // empty target unit
	}

	conv := NewConverter(ucum.NewConverter())
	val := decimal.NewFromInt(1)

	for _, tt := range tests {
		tname := tt[0] + " -> " + tt[1]
		t.Run("ConvDecimal "+tname, func(t *testing.T) {
			if got, err := ConvDecimal(val, tt[0], tt[1], 18); err == nil {
				t.Errorf("ConvDecimal() = %s, want an error", got)
			}
		})
		t.Run("Conv.ConvDecimal "+tname, func(t *testing.T) {
			if got, err := conv.ConvDecimal(val, tt[0], tt[1], 18); err == nil {
				t.Errorf("Conv.ConvDecimal() = %s, want an error", got)
			}
		})
	}
}

// TestConvDecimalPrecision checks that the requested precision is honoured for
// a result that has an infinite decimal representation.
func TestConvDecimalPrecision(t *testing.T) {
	conv := NewConverter(ucum.NewConverter())
	val := decimal.NewFromInt(1)

	tests := []struct {
		precision int32
		want      string
	}{
		{precision: 2, want: "0.02"},
		{precision: 5, want: "0.01667"},
		{precision: 18, want: "0.016666666666666667"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got, err := conv.ConvDecimal(val, "min", "h", tt.precision)
			if err != nil {
				t.Fatalf("ConvDecimal() error = %v", err)
			}
			if got.String() != tt.want {
				t.Errorf("ConvDecimal() = %s, want %s", got, tt.want)
			}
		})
	}
}
