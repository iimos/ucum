package ucumapd

import (
	"testing"

	"github.com/cockroachdb/apd/v3"
	"github.com/iimos/ucum"
)

// TestConvDecimalNonFinite checks that non-finite and nil inputs are rejected
// with an error instead of a panic.
func TestConvDecimalNonFinite(t *testing.T) {
	ctx := (&apd.BaseContext).WithPrecision(20)
	conv := NewConverter(ucum.NewConverter())

	tests := []struct {
		name string
		val  *apd.Decimal
	}{
		{name: "nil", val: nil},
		{name: "NaN", val: &apd.Decimal{Form: apd.NaN}},
		{name: "Infinite", val: &apd.Decimal{Form: apd.Infinite}},
		{name: "NegativeInfinite", val: &apd.Decimal{Form: apd.Infinite, Negative: true}},
	}

	for _, tt := range tests {
		t.Run("Package/"+tt.name, func(t *testing.T) {
			if got, err := ConvDecimal(tt.val, "km", "m", ctx); err == nil {
				t.Errorf("ConvDecimal() = %v, want an error", got)
			}
		})
		t.Run("Instance/"+tt.name, func(t *testing.T) {
			if got, err := conv.ConvDecimal(tt.val, "km", "m", ctx); err == nil {
				t.Errorf("Conv.ConvDecimal() = %v, want an error", got)
			}
		})
	}
}

// TestConvDecimalNilContext documents how a nil context is handled.
//
// ensureContext falls back to apd.BaseContext, which has Precision 0, and apd
// rejects division with such a context, so every conversion returns an error.
func TestConvDecimalNilContext(t *testing.T) {
	val := mustNewDecimal(t, "1")

	if got, err := ConvDecimal(val, "km", "m", nil); err == nil {
		t.Errorf("ConvDecimal() = %v, want an error for a nil context", got)
	}

	conv := NewConverter(ucum.NewConverter())
	if got, err := conv.ConvDecimal(val, "km", "m", nil); err == nil {
		t.Errorf("Conv.ConvDecimal() = %v, want an error for a nil context", got)
	}
}

// TestConvDecimalExponent checks values with positive and negative decimal
// exponents (both branches of apdDecimalToRat).
func TestConvDecimalExponent(t *testing.T) {
	ctx := (&apd.BaseContext).WithPrecision(30)

	tests := []struct {
		input    string
		from, to string
		want     string
	}{
		{input: "1E+3", from: "m", to: "km", want: "1"},
		{input: "0.001", from: "m", to: "km", want: "0.000001"},
		{input: "1E-3", from: "km", to: "m", want: "1"},
		{input: "-1E+3", from: "m", to: "km", want: "-1"},
	}

	for _, tt := range tests {
		t.Run(tt.input+" "+tt.from+"->"+tt.to, func(t *testing.T) {
			got, err := ConvDecimal(mustNewDecimal(t, tt.input), tt.from, tt.to, ctx)
			if err != nil {
				t.Fatalf("ConvDecimal() error = %v", err)
			}
			want := mustNewDecimal(t, tt.want)
			if got.Cmp(want) != 0 {
				t.Errorf("ConvDecimal() = %s, want %s", got, want)
			}
		})
	}
}
