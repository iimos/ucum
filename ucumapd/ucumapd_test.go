package ucumapd

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/apd/v3"
	"github.com/iimos/ucum"
)

type testCase struct {
	name    string
	from    string
	to      string
	input   string
	want    string
	roundTo int32 // -1 means exact comparison
	wantErr bool
}

var tests = []testCase{
	{name: "Ident-100", from: "1", to: "1", input: "100", want: "100", roundTo: -1},
	{name: "Ident-Zero", from: "1", to: "1", input: "0", want: "0", roundTo: -1},
	{name: "Ident-Negative", from: "1", to: "1", input: "-100", want: "-100", roundTo: -1},

	{
		name:    "Pi-HighPrecision",
		from:    "[pi]",
		to:      "1",
		input:   "10",
		want:    "31.415926535897932384626433832795028841971693993751058209749445923",
		roundTo: 60,
	},
	{
		name:    "Pi-LowPrecision",
		from:    "[pi]",
		to:      "1",
		input:   "10",
		want:    "31.42",
		roundTo: 2,
	},

	{name: "MetersIdent", from: "m", to: "m", input: "100", want: "100", roundTo: 0},
	{name: "KmToM", from: "km", to: "m", input: "1", want: "1000", roundTo: 0},
	{name: "MToKm", from: "m", to: "km", input: "1000", want: "1", roundTo: 0},
	{name: "SquareMetersToSquareKm", from: "m2", to: "km2", input: "1000000", want: "1", roundTo: 0},

	{name: "CelsiusToKelvin", from: "Cel", to: "K", input: "100", want: "373.15", roundTo: 2},

	{
		name:    "GiByToBy",
		from:    "GiBy",
		to:      "By",
		input:   "0.00000000465661287",
		want:    "4.99999999669567488",
		roundTo: 17,
	},
	{
		name:    "ByToGiBy",
		from:    "By",
		to:      "GiBy",
		input:   "4.99999999669567488",
		want:    "0.00000000465661287",
		roundTo: 17,
	},

	{name: "DimensionMismatch", from: "m", to: "s", input: "10", wantErr: true},
	{name: "InvalidUnit", from: "m", to: "not_a_unit", input: "10", wantErr: true},
}

func TestConvDecimal(t *testing.T) {
	mathCtx := (&apd.BaseContext).WithPrecision(100)

	type convFunc func(val *apd.Decimal, from string, to string, ctx *apd.Context) (*apd.Decimal, error)

	runTests := func(t *testing.T, label string, fn convFunc) {
		for _, tc := range tests {
			tc := tc
			t.Run(fmt.Sprintf("%s/%s", label, tc.name), func(t *testing.T) {
				val := mustNewDecimal(t, tc.input)
				got, err := fn(val, tc.from, tc.to, mathCtx)

				if tc.wantErr {
					if err == nil {
						t.Fatalf("expected error, got nil")
					}
					return
				}
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				want := mustNewDecimal(t, tc.want)
				if tc.roundTo >= 0 {
					roundedGot := new(apd.Decimal)
					roundedWant := new(apd.Decimal)
					if _, err := mathCtx.Quantize(roundedGot, got, -tc.roundTo); err != nil {
						t.Fatalf("quantize got: %v", err)
					}
					if _, err := mathCtx.Quantize(roundedWant, want, -tc.roundTo); err != nil {
						t.Fatalf("quantize want: %v", err)
					}
					if roundedGot.Cmp(roundedWant) != 0 {
						t.Fatalf("got %s, want %s (rounded to %d places)", roundedGot, roundedWant, tc.roundTo)
					}
					return
				}

				if got.Cmp(want) != 0 {
					t.Fatalf("got %s, want %s", got, want)
				}
			})
		}
	}

	t.Run("Package", func(t *testing.T) {
		runTests(t, "pkg", ConvDecimal)
	})

	t.Run("Instance", func(t *testing.T) {
		conv := NewConverter(ucum.NewConverter())
		runTests(t, "instance", conv.ConvDecimal)
	})
}

func mustNewDecimal(t *testing.T, s string) *apd.Decimal {
	t.Helper()
	d, _, err := apd.NewFromString(s)
	if err != nil {
		t.Fatalf("apd.NewFromString(%q): %v", s, err)
	}
	return d
}
