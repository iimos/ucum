package ucumdecimal

import (
	"fmt"
	"testing"

	"github.com/iimos/ucum"
	"github.com/shopspring/decimal"
)

func TestConvDecimal(t *testing.T) {
	type testCase struct {
		A, B      string
		value     string
		want      string
		precision int32
	}

	// Exact conversions shared by all packages live in internal/testcases and
	// are covered by TestSharedConversions. This table keeps cases specific to
	// shopspring/decimal precision control.
	tests := []testCase{
		{
			A:         "[pi]",
			B:         "1",
			value:     "10",
			want:      "31.415926535897932384626433832795028841971693993751058209749445923",
			precision: 100,
		},
		{
			A:         "[pi]",
			B:         "1",
			value:     "10",
			want:      "31.42",
			precision: 2,
		},
		{
			A:         "GiBy",
			B:         "By",
			value:     "-0.00000000465661287",
			want:      "-4.99999999669567488",
			precision: 100,
		},
	}

	type ConvFunc func(val decimal.Decimal, from string, to string, precision int32) (decimal.Decimal, error)

	runTest := func(t *testing.T, tt testCase, funcToTest ConvFunc) {
		tname := fmt.Sprintf("ConvDecimal: (%s)%s is (%s)%s", tt.value, tt.A, tt.want, tt.B)
		t.Run(tname, func(t *testing.T) {
			val, err := decimal.NewFromString(tt.value)
			if err != nil {
				t.Fatalf("decimal.NewFromString(%s): %s", tt.A, err)
			}
			got, err := funcToTest(val, tt.A, tt.B, tt.precision)
			if err != nil {
				t.Fatalf("ConvDecimal(): %v", err)
			}
			if got.String() != tt.want {
				t.Errorf("ConvDecimal(): got %s, want %s", got.String(), tt.want)
			}
		})
	}

	for _, tt := range tests {
		invertedTest := testCase{
			A:         tt.B,
			B:         tt.A,
			value:     tt.want,
			want:      tt.value,
			precision: tt.precision,
		}

		t.Run("ConvDecimal", func(t *testing.T) {
			runTest(t, tt, ConvDecimal)
			runTest(t, invertedTest, ConvDecimal)
		})

		t.Run("Conv.ConvDecimal", func(t *testing.T) {
			conv := NewConverter(ucum.NewConverter())
			runTest(t, tt, conv.ConvDecimal)
			runTest(t, invertedTest, conv.ConvDecimal)
		})
	}
}
