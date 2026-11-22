package ucumdecimal

import (
	"testing"

	"github.com/iimos/ucum"
	"github.com/iimos/ucum/internal/testcases"
	"github.com/shopspring/decimal"
)

// sharedPrecision is the precision used to run the shared conversion matrix.
// It must be large enough to represent every expected value exactly: the
// longest one has 24 decimal places.
const sharedPrecision = 40

// TestSharedConversions runs the conversion matrix shared by all packages, see
// internal/testcases. Both the package-level function and a converter instance
// are checked, and every conversion must be reversible.
func TestSharedConversions(t *testing.T) {
	cases, err := testcases.Conversions()
	if err != nil {
		t.Fatalf("load shared conversion cases: %v", err)
	}

	type convFunc func(val decimal.Decimal, from string, to string, precision int32) (decimal.Decimal, error)

	runTests := func(t *testing.T, label string, fn convFunc) {
		for _, tc := range cases {
			t.Run(label+"/"+tc.Name(), func(t *testing.T) {
				value, err := decimal.NewFromString(tc.Value)
				if err != nil {
					t.Fatalf("decimal.NewFromString(%q): %v", tc.Value, err)
				}
				want, err := decimal.NewFromString(tc.Want)
				if err != nil {
					t.Fatalf("decimal.NewFromString(%q): %v", tc.Want, err)
				}

				got, err := fn(value, tc.From, tc.To, sharedPrecision)
				if err != nil {
					t.Fatalf("ConvDecimal() error = %v", err)
				}
				// The matrix holds exact finite decimals, so the result must
				// match exactly rather than within some tolerance.
				if !got.Equal(want) {
					t.Errorf("ConvDecimal() = %s, want %s", got, want)
				}

				// Converting the result back must yield the original value.
				back, err := fn(got, tc.To, tc.From, sharedPrecision)
				if err != nil {
					t.Fatalf("ConvDecimal() inverse error = %v", err)
				}
				if !back.Equal(value) {
					t.Errorf("ConvDecimal() inverse = %s, want %s", back, value)
				}
			})
		}
	}

	t.Run("Package", func(t *testing.T) {
		runTests(t, "pkg", ConvDecimal)
	})
	t.Run("Instance", func(t *testing.T) {
		runTests(t, "instance", NewConverter(ucum.NewConverter()).ConvDecimal)
	})
}
