package ucumapd

import (
	"testing"

	"github.com/cockroachdb/apd/v3"
	"github.com/iimos/ucum"
	"github.com/iimos/ucum/internal/testcases"
)

// TestSharedConversions runs the conversion matrix shared by all packages, see
// internal/testcases. Both the package-level function and a converter instance
// are checked, and every conversion must be reversible.
func TestSharedConversions(t *testing.T) {
	ctx := (&apd.BaseContext).WithPrecision(100)

	cases, err := testcases.Conversions()
	if err != nil {
		t.Fatalf("load shared conversion cases: %v", err)
	}

	type convFunc func(val *apd.Decimal, from string, to string, ctx *apd.Context) (*apd.Decimal, error)

	runTests := func(t *testing.T, label string, fn convFunc) {
		for _, tc := range cases {
			t.Run(label+"/"+tc.Name(), func(t *testing.T) {
				value := mustNewDecimal(t, tc.Value)
				want := mustNewDecimal(t, tc.Want)

				got, err := fn(value, tc.From, tc.To, ctx)
				if err != nil {
					t.Fatalf("ConvDecimal() error = %v", err)
				}
				// The matrix holds exact finite decimals, so the result must
				// match bit for bit rather than within some tolerance.
				if got.Cmp(want) != 0 {
					t.Errorf("ConvDecimal() = %s, want %s", got, want)
				}

				// Converting the result back must yield the original value.
				back, err := fn(got, tc.To, tc.From, ctx)
				if err != nil {
					t.Fatalf("ConvDecimal() inverse error = %v", err)
				}
				if back.Cmp(value) != 0 {
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
