package ucum

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

// localConvCases returns exact conversions that are intentionally kept out of
// the shared matrix: their expected value is a 45-digit decimal that pins the
// float64 result of math.Tan/math.Atan bit for bit, which is too fragile (and
// too ugly) to share with the adapter modules.
func localConvCases() []convCase {
	want := big.NewRat(5479641287827929, 35184372088832) // 100*math.Tan(1), ≈ 155.7407724654902
	return []convCase{
		{A: "rad", B: "[p'diop]", value: big.NewRat(1, 1), want: want},
		{A: "deg/[pi]", B: "[p'diop]", value: big.NewRat(180, 1), want: want},
	}
}

// inexactConvCases returns conversions that cannot live in the shared exact
// matrix in internal/testcases/conversions.json. They fall into two groups:
//
//   - the mathematical result is exact, but the implementation computes it with
//     float64 math and returns noise: 45 deg -> %[slope] must be 100 but comes
//     out as 3518437208883199/35184372088832, 2 [hp'_X] must be 1/100 but comes
//     out as 5764607523034235/576460752303423488, and so on;
//   - the result is irrational: Np, B -> Np, B[SPL] at 5, B[V] at 3.
//
// Both groups are compared with a relative tolerance, see nearlyEqual.
//
// [pi] is kept here as well, although [pi] is an exact rational in the
// generated conversion table: its 63-digit decimal expansion is an artifact of
// the table and would make the shared matrix depend on the exact number of
// digits in it, while ucumapd and ucumdecimal already cover high precision in
// their own precision tests.
func inexactConvCases() []convCase {
	return []convCase{
		{
			A:       "[pi]",
			B:       "1",
			value:   big.NewRat(10, 1),
			want:    new(big.Rat).SetFloat64(10 * math.Pi),
			inexact: true,
		},
		{
			A:       "deg",
			B:       "%[slope]",
			value:   big.NewRat(45, 1),
			want:    big.NewRat(100, 1),
			inexact: true,
		},
		{
			A:       "[hp'_X]",
			B:       "1",
			value:   big.NewRat(2, 1),
			want:    big.NewRat(1, 100),
			inexact: true,
		},
		{
			A:       "[hp'_C]",
			B:       "1",
			value:   big.NewRat(2, 1),
			want:    big.NewRat(1, 100*100),
			inexact: true,
		},
		{
			A:       "[hp'_M]",
			B:       "1",
			value:   big.NewRat(2, 1),
			want:    big.NewRat(1, 1000*1000),
			inexact: true,
		},
		{
			A:       "[hp'_Q]",
			B:       "1",
			value:   big.NewRat(2, 1),
			want:    big.NewRat(1, 50000*50000),
			inexact: true,
		},
		{
			A:       "[pH]",
			B:       "mol/l",
			value:   big.NewRat(3, 1),
			want:    big.NewRat(1, 1000),
			inexact: true,
		},
		{
			A:       "2.[pH]",
			B:       "mol/l",
			value:   big.NewRat(3, 1),
			want:    big.NewRat(1, 1000000),
			inexact: true,
		},
		{
			A:       "Np",
			B:       "1",
			value:   big.NewRat(3, 1),
			want:    new(big.Rat).SetFloat64(20.0855369231877),
			inexact: true,
		},
		{
			A:       "B",
			B:       "Np",
			value:   big.NewRat(10, 1),
			want:    new(big.Rat).SetFloat64(23.0258509299405),
			inexact: true,
		},
		{
			A:       "B[SPL]",
			B:       "Pa",
			value:   big.NewRat(5, 1),
			want:    new(big.Rat).SetFloat64(0.00632455532033676),
			inexact: true,
		},
		{
			A:       "B[V]",
			B:       "V",
			value:   big.NewRat(3, 1),
			want:    new(big.Rat).SetFloat64(31.6227766016838),
			inexact: true,
		},
	}
}

func TestConvRatFloat(t *testing.T) {
	tests := sharedConvCases(t)
	tests = append(tests, localConvCases()...)
	tests = append(tests, inexactConvCases()...)

	runTest := func(t *testing.T, tt convCase) {
		tname := fmt.Sprintf("(%s)%s is (%s)%s", tt.value.String(), tt.A, tt.want.String(), tt.B)
		t.Run(tname, func(t *testing.T) {
			A := MustParse([]byte(tt.A))
			B := MustParse([]byte(tt.B))

			t.Run("PairConverter Rat", func(t *testing.T) {
				converter, err := NewPairConverter(A, B)
				if err != nil {
					t.Fatalf("NewPairConverter() error = %v", err)
				}
				got := converter.ConvBigRat(tt.value)
				if tt.inexact {
					gotf, _ := got.Float64()
					wantf, _ := tt.want.Float64()
					if !nearlyEqual(gotf, wantf) {
						t.Errorf("ConvBigRat() got = %s, want %s; inexact=true", got.String(), tt.want.String())
					}
				} else {
					if got.Cmp(tt.want) != 0 {
						if pc, ok := converter.(*linearConverter); ok {
							t.Errorf("ConvBigRat() got = %s, want %s; ratio=%s", got.String(), tt.want.String(), pc.ratio.RatString())
						} else {
							t.Errorf("ConvBigRat() got = %s, want %s", got.String(), tt.want.String())
						}
					}
				}

				// Check that the same value is returned if the same value is passed.
				// It protects against occasional state mutations.
				got2 := converter.ConvBigRat(tt.value)
				if got2.Cmp(got) != 0 {
					t.Errorf("failed to reproduce results: first time got %s, second time got %s", got.String(), got2.String())
				}
			})

			t.Run("Converter Rat", func(t *testing.T) {
				converter := NewConverter()
				got, err := converter.ConvBigRat(tt.value, tt.A, tt.B)
				if err != nil {
					t.Fatalf("NewConverter().ConvBigRat error = %v", err)
				}
				if tt.inexact {
					gotf, _ := got.Float64()
					wantf, _ := tt.want.Float64()
					if !nearlyEqual(gotf, wantf) {
						t.Errorf("ConvBigRat() got = %s, want %s; inexact=true", got.String(), tt.want.String())
					}
				} else {
					if got.Cmp(tt.want) != 0 {
						t.Errorf("got = %s, want %s", got.String(), tt.want.String())
					}
				}

				// Check that the same value is returned if the same value is passed.
				// It protects against occasional state mutations.
				got2, err := converter.ConvBigRat(tt.value, tt.A, tt.B)
				if err != nil {
					t.Fatalf("NewConverter().ConvBigRat error = %v", err)
				}
				if got2.Cmp(got) != 0 {
					t.Errorf("failed to reproduce results: first time got %s, second time got %s", got.String(), got2.String())
				}
			})

			t.Run("Converter Float", func(t *testing.T) {
				converter := NewConverter()
				f64, _ := tt.value.Float64()
				got, err := converter.ConvFloat64(f64, tt.A, tt.B)
				if err != nil {
					t.Fatalf("NewConverter().ConvBigRat error = %v", err)
				}
				wantf, _ := tt.want.Float64()
				if !nearlyEqual(got, wantf) {
					t.Errorf("ConvBigRat() got = %f, want %s", got, tt.want.String())
				}
			})

			t.Run("ConvBigRat", func(t *testing.T) {
				got, err := ConvBigRat(tt.value, tt.A, tt.B)
				if err != nil {
					t.Fatalf("error = %v", err)
				}
				if tt.inexact {
					gotf, _ := got.Float64()
					wantf, _ := tt.want.Float64()
					if !nearlyEqual(gotf, wantf) {
						t.Errorf("got = %s, want %s; inexact=true", got.String(), tt.want.String())
					}
				} else {
					if got.Cmp(tt.want) != 0 {
						t.Errorf("got = %s, want %s", got.String(), tt.want.String())
					}
				}
			})

			t.Run("ConvFloat64", func(t *testing.T) {
				f64, _ := tt.value.Float64()
				got, err := ConvFloat64(f64, tt.A, tt.B)
				if err != nil {
					t.Fatalf("error = %v", err)
				}
				wantf, _ := tt.want.Float64()
				if !nearlyEqual(got, wantf) {
					t.Errorf("got = %f, want %s", got, tt.want.String())
				}
			})
		})
	}

	for _, tt := range tests {
		invertedTest := convCase{
			A:       tt.B,
			B:       tt.A,
			value:   new(big.Rat).Set(tt.want),
			want:    new(big.Rat).Set(tt.value),
			inexact: tt.inexact,
		}
		t.Run("fwd", func(t *testing.T) { runTest(t, tt) })
		t.Run("inv", func(t *testing.T) { runTest(t, invertedTest) })
	}
}

func TestConvBigInt(t *testing.T) {
	tests := []struct {
		A, B      string
		value     *big.Int
		want      *big.Int
		wantExact bool
	}{
		{
			A:     "m",
			B:     "m",
			value: big.NewInt(100),
			want:  big.NewInt(100), wantExact: true,
		},
		{
			A:     "km",
			B:     "m",
			value: big.NewInt(1),
			want:  big.NewInt(1000), wantExact: true,
		},
		{
			A:     "m",
			B:     "km",
			value: big.NewInt(1000),
			want:  big.NewInt(1), wantExact: true,
		},
		{
			A:     "m",
			B:     "km",
			value: big.NewInt(100),
			want:  big.NewInt(0), wantExact: false,
		},
	}
	for _, tt := range tests {
		tname := fmt.Sprintf("%s%s is %s%s", tt.value.String(), tt.A, tt.want.String(), tt.B)
		t.Run(tname, func(t *testing.T) {
			A := MustParse([]byte(tt.A))
			B := MustParse([]byte(tt.B))

			// check verifies a single conversion result against the expected
			// value and exactness. api names the call under test, details is
			// appended to failure messages.
			check := func(t *testing.T, api string, got *big.Int, exact bool, err error, details string) {
				t.Helper()
				if err != nil {
					t.Fatalf("%s() error = %v", api, err)
				}
				if got.Cmp(tt.want) != 0 {
					t.Errorf("%s() got = %s, want %s%s", api, got, tt.want, details)
				}
				if exact != tt.wantExact {
					t.Errorf("%s() exact = %t, want %t%s", api, exact, tt.wantExact, details)
				}
			}

			// checkRepeated verifies that a second call with the same arguments
			// returns the same value and exactness, which protects against
			// occasional state mutations.
			checkRepeated := func(t *testing.T, api string, got *big.Int, exact bool, got2 *big.Int, exact2 bool) {
				t.Helper()
				if got2.Cmp(got) != 0 {
					t.Errorf("%s(): failed to reproduce results: first time got %s, second time got %s", api, got, got2)
				}
				if exact2 != exact {
					t.Errorf("%s(): failed to reproduce exactness: first time got %t, second time got %t", api, exact, exact2)
				}
			}

			t.Run("PairConverter", func(t *testing.T) {
				converter, err := NewPairConverter(A, B)
				if err != nil {
					t.Fatalf("NewPairConverter() error = %v", err)
				}
				var details string
				if pc, ok := converter.(*linearConverter); ok {
					details = "; ratio=" + pc.ratio.RatString()
				}

				got, exact := converter.ConvBigInt(tt.value)
				check(t, "PairConverter.ConvBigInt", got, exact, nil, details)

				got2, exact2 := converter.ConvBigInt(tt.value)
				checkRepeated(t, "PairConverter.ConvBigInt", got, exact, got2, exact2)
			})

			t.Run("Converter", func(t *testing.T) {
				converter := NewConverter()
				got, exact, err := converter.ConvBigInt(tt.value, tt.A, tt.B)
				check(t, "Converter.ConvBigInt", got, exact, err, "")

				got2, exact2, err := converter.ConvBigInt(tt.value, tt.A, tt.B)
				if err != nil {
					t.Fatalf("Converter.ConvBigInt() error = %v", err)
				}
				checkRepeated(t, "Converter.ConvBigInt", got, exact, got2, exact2)
			})

			t.Run("ConvBigInt", func(t *testing.T) {
				got, exact, err := ConvBigInt(tt.value, tt.A, tt.B)
				check(t, "ConvBigInt", got, exact, err, "")

				got2, exact2, err := ConvBigInt(tt.value, tt.A, tt.B)
				if err != nil {
					t.Fatalf("ConvBigInt() error = %v", err)
				}
				checkRepeated(t, "ConvBigInt", got, exact, got2, exact2)
			})
		})
	}
}

func TestLinearConverterString(t *testing.T) {
	conv, err := NewPairConverter(MustParse([]byte("m")), MustParse([]byte("km")))
	if err != nil {
		t.Fatalf("NewPairConverter() error = %v", err)
	}
	linear, ok := conv.(*linearConverter)
	if !ok {
		t.Fatalf("NewPairConverter() = %T, want *linearConverter", conv)
	}
	if got, want := linear.String(), "ucum.PairConverter(m->km)"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestNewPairConverter(t *testing.T) {
	tests := map[[2]string]string{
		// input -> error
		{"m", "By"}:     `ucum: "m" cannot be converted to "By"`,
		{"m/s", "By/s"}: `ucum: "m/s" cannot be converted to "By/s"`,
		{"m", "m/s"}:    `ucum: "m" cannot be converted to "m/s"`,
		{"m", "1"}:      `ucum: "m" cannot be converted to "1"`,
		{"m", "m2"}:     `ucum: "m" cannot be converted to "m2"`,
		{"m", "1/m"}:    `ucum: "m" cannot be converted to "1/m"`,

		// Arbitrary units are not commensurable with any other unit (UCUM §24, §25).
		{"bit", "[iU]"}:         `ucum: "bit" cannot be converted to "[iU]"`,
		{"[iU]", "bit"}:         `ucum: "[iU]" cannot be converted to "bit"`,
		{"By", "[iU]"}:          `ucum: "By" cannot be converted to "[iU]"`,
		{"1", "[iU]"}:           `ucum: "1" cannot be converted to "[iU]"`,
		{"[iU]", "1"}:           `ucum: "[iU]" cannot be converted to "1"`,
		{"[CFU]", "1"}:          `ucum: "[CFU]" cannot be converted to "1"`,
		{"[hp_X]", "[hp_C]"}:    `ucum: "[hp_X]" cannot be converted to "[hp_C]"`,
		{"[iU]", "[arb'U]"}:     `ucum: "[iU]" cannot be converted to "[arb'U]"`,
		{"[iU]", "[iU]2"}:       `ucum: "[iU]" cannot be converted to "[iU]2"`,
		{"[iU]/mL", "[CFU]/mL"}: `ucum: "[iU]/mL" cannot be converted to "[CFU]/mL"`,
		{"[iU]/mL", "[iU]/l"}:   `ucum: "[iU]/mL" cannot be converted to "[iU]/l"`,
		{"[iU]/mL", "1/mL"}:     `ucum: "[iU]/mL" cannot be converted to "1/mL"`,
		{"[iU]/[iU]", "1"}:      `ucum: "[iU]/[iU]" cannot be converted to "1"`,
	}

	for input, wantErr := range tests {
		testName := input[0] + "->" + input[1]
		A := MustParse([]byte(input[0]))
		B := MustParse([]byte(input[1]))
		t.Run("PairConverter: "+testName, func(t *testing.T) {
			conv, err := NewPairConverter(A, B)
			if err == nil {
				t.Errorf("no error, want %q", wantErr)
				return
			}
			if err.Error() != wantErr {
				t.Errorf("NewPairConverter() error = %q, want %q", err, wantErr)
				return
			}
			if conv != nil {
				t.Errorf("NewPairConverter() returned error with non-nil converter")
			}
		})
		t.Run("Converter: "+testName, func(t *testing.T) {
			result, err := NewConverter().ConvFloat64(1, input[0], input[1])
			if err == nil {
				t.Errorf("no error, want %q", wantErr)
				return
			}
			if err.Error() != wantErr {
				t.Errorf("Converter() error = %q, want %q", err, wantErr)
				return
			}
			if result != 0 {
				t.Errorf("Converter() returned error with non-zero result")
			}
		})
	}
}

// nearlyEqual compares two float64s and returns whether they are equal, accounting for rounding errors.At worst, the
// result is correct to 7 significant digits.
// Taken from https://github.com/faiface/pixel/blob/0a251bc08bfcfd68720fc237ee6a6268fa3a12e6/vector.go#L40-L57
func nearlyEqual(a, b float64) bool {
	epsilon := 0.000001

	if a == b {
		return true
	}

	diff := math.Abs(a - b)

	if a == 0.0 || b == 0.0 || diff < math.SmallestNonzeroFloat64 {
		return diff < (epsilon * math.SmallestNonzeroFloat64)
	}

	absA := math.Abs(a)
	absB := math.Abs(b)

	return diff/math.Min(absA+absB, math.MaxFloat64) < epsilon
}
