package ucum

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

//func TestDecimal(t *testing.T) {
//	d := decimal.NewFromBigRat(big.NewRat(1, 2), 18)
//	c := PairConverter{}
//	rat := c.ConvRat(d.Rat())
//	d := decimal.NewFromBigRat(rat, 18)
//}

func TestConvRat(t *testing.T) {
	type testCase struct {
		A, B    string
		value   *big.Rat
		want    *big.Rat
		inexact bool
	}
	tests := []testCase{
		{
			A:     "1",
			B:     "1",
			value: big.NewRat(100, 1),
			want:  big.NewRat(100, 1),
		},
		{
			A:       "[pi]",
			B:       "1",
			value:   big.NewRat(10, 1),
			want:    new(big.Rat).SetFloat64(10 * math.Pi),
			inexact: true,
		},
		{
			A:       "[hnsf'U]",
			B:       "1",
			value:   big.NewRat(10, 1),
			want:    big.NewRat(10, 1),
			inexact: true,
		},
		{
			A:     "m",
			B:     "m",
			value: big.NewRat(100, 1),
			want:  big.NewRat(100, 1),
		},
		{
			A:     "km",
			B:     "m",
			value: big.NewRat(1, 1),
			want:  big.NewRat(1000, 1),
		},
		{
			A:     "m",
			B:     "km",
			value: big.NewRat(1000, 1),
			want:  big.NewRat(1, 1),
		},
		{
			A:     "m",
			B:     "km",
			value: big.NewRat(100, 1),
			want:  big.NewRat(1, 10),
		},
		{
			A:     "m2",
			B:     "km2",
			value: big.NewRat(1000000, 1),
			want:  big.NewRat(1, 1),
		},
		{
			A:     "deg/[pi]",
			B:     "rad",
			value: big.NewRat(360, 1),
			want:  big.NewRat(2, 1),
		},
		{
			A:     "Cel",
			B:     "K",
			value: big.NewRat(100, 1),
			want:  big.NewRat(7463, 20), // = 373.15
		},
		{
			A:     "kCel",
			B:     "kK",
			value: big.NewRat(1, 1),
			want:  big.NewRat(127315, 100000), // = 1.27315
		},
		{
			A:     "kCel",
			B:     "K",
			value: big.NewRat(1, 1),
			want:  big.NewRat(127315, 100),
		},
		{
			A:     "Cel",
			B:     "[degF]",
			value: big.NewRat(100, 1),
			want:  big.NewRat(212, 1),
		},
		{
			// https://www.unitconverters.net/temperature/reaumur-to-celsius.htm
			A:     "Cel",
			B:     "[degRe]",
			value: big.NewRat(125, 1),
			want:  big.NewRat(100, 1),
		},
		{
			A:     "[degRe]",
			B:     "K",
			value: big.NewRat(15, 1),
			want:  big.NewRat(2919, 10),
		},
		{
			A:     "rad",
			B:     "[p'diop]",
			value: big.NewRat(1, 1),
			want:  big.NewRat(5479641287827929, 35184372088832), // 155.740772465, probably should be rechecked
		},
		{
			A:     "deg/[pi]",
			B:     "[p'diop]",
			value: big.NewRat(180, 1),
			want:  big.NewRat(5479641287827929, 35184372088832), // 155.740772465, probably should be rechecked
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
			A:       "[hp'_C]",
			B:       "[hp'_X]",
			value:   big.NewRat(1, 1),
			want:    big.NewRat(2, 1),
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
			B:       "1",
			value:   big.NewRat(3, 1),
			want:    big.NewRat(1000, 1),
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
			value:   big.NewRat(10, 1),
			want:    big.NewRat(2, 1),
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
			value:   big.NewRat(2, 1),
			want:    big.NewRat(10, 1),
			inexact: true,
		},
		{
			A:       "B[V]",
			B:       "V",
			value:   big.NewRat(3, 1),
			want:    new(big.Rat).SetFloat64(31.6227766016838),
			inexact: true,
		},
		{
			A:       "B[mV]",
			B:       "V",
			value:   big.NewRat(8, 1),
			want:    big.NewRat(10, 1),
			inexact: true,
		},
		{
			A:       "B[uV]",
			B:       "V",
			value:   big.NewRat(16, 1),
			want:    big.NewRat(100, 1),
			inexact: true,
		},
		{
			A:       "B[10.nV]",
			B:       "V",
			value:   big.NewRat(20, 1),
			want:    big.NewRat(100, 1),
			inexact: true,
		},
		{
			A:       "B[W]",
			B:       "W",
			value:   big.NewRat(2, 1),
			want:    big.NewRat(100, 1),
			inexact: true,
		},
		{
			A:       "B[kW]",
			B:       "W",
			value:   big.NewRat(2, 1),
			want:    big.NewRat(100000, 1),
			inexact: true,
		},
		{
			A:       "[m/s2/Hz^(1/2)]",
			B:       "m2/s4/Hz",
			value:   big.NewRat(2, 1),
			want:    big.NewRat(4, 1),
			inexact: true,
		},
		{
			A:       "bit_s",
			B:       "1",
			value:   big.NewRat(10, 1),
			want:    big.NewRat(1024, 1),
			inexact: true,
		},
	}

	runTest := func(t *testing.T, tt testCase) {
		tname := fmt.Sprintf("(%s)%s is (%s)%s", tt.value.String(), tt.A, tt.want.String(), tt.B)
		t.Run(tname, func(t *testing.T) {
			A := MustParse([]byte(tt.A))
			B := MustParse([]byte(tt.B))

			t.Run("PairConverter", func(t *testing.T) {
				converter, err := NewPairConverter(A, B)
				if err != nil {
					t.Fatalf("NewPairConverter() error = %v", err)
				}
				got := converter.ConvRat(tt.value)
				if tt.inexact {
					gotf, _ := got.Float64()
					wantf, _ := tt.want.Float64()
					if !nearlyEqual(gotf, wantf) {
						t.Errorf("ConvRat() got = %s, want %s; inexact=true", got.String(), tt.want.String())
					}
				} else {
					if got.Cmp(tt.want) != 0 {
						if pc, ok := converter.(*linearConverter); ok {
							t.Errorf("ConvRat() got = %s, want %s; ratio=%s", got.String(), tt.want.String(), pc.ratio.RatString())
						} else {
							t.Errorf("ConvRat() got = %s, want %s", got.String(), tt.want.String())
						}
					}
				}

				// Check that the same value is returned if the same value is passed.
				// It protects against occasional state mutations.
				got2 := converter.ConvRat(tt.value)
				if got2.Cmp(got) != 0 {
					t.Errorf("failed to reproduce results: first time got %s, second time got %s", got.String(), got2.String())
				}
			})
		})
	}

	for _, tt := range tests {
		invertedTest := testCase{
			A:       tt.B,
			B:       tt.A,
			value:   new(big.Rat).Set(tt.want),
			want:    new(big.Rat).Set(tt.value),
			inexact: tt.inexact,
		}
		runTest(t, tt)
		runTest(t, invertedTest)
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

			t.Run("PairConverter", func(t *testing.T) {
				converter, err := NewPairConverter(A, B)
				if err != nil {
					t.Fatalf("NewPairConverter() error = %v", err)
				}
				got, gotExact := converter.ConvBigInt(tt.value)
				if got.Cmp(tt.want) != 0 {
					if pc, ok := converter.(*linearConverter); ok {
						t.Errorf("ConvRat() got = %s, want %s; ratio=%s", got.String(), tt.want.String(), pc.ratio.RatString())
					} else {
						t.Errorf("ConvRat() got = %s, want %s", got.String(), tt.want.String())
					}
				}
				if gotExact != tt.wantExact {
					t.Errorf("ConvBigInt() gotExact = %t, want %t", gotExact, tt.wantExact)
				}

				// Check that the same value is returned if the same value is passed.
				// It protects against occasional state mutations.
				got2, _ := converter.ConvBigInt(tt.value)
				if got2.Cmp(got) != 0 {
					t.Errorf("failed to reproduce results: first time got %s, second time got %s", got.String(), got2.String())
				}
			})

			t.Run("ConvBigInt", func(t *testing.T) {
				got, gotExact, err := ConvBigInt(A, B, tt.value)
				if err != nil {
					t.Fatalf("ConvBigInt() error = %v", err)
				}
				if got.Cmp(tt.want) != 0 {
					t.Errorf("ConvBigInt() got = %s, want %s", got.String(), tt.want.String())
				}
				if gotExact != tt.wantExact {
					t.Errorf("ConvBigInt() gotExact = %t, want %t", gotExact, tt.wantExact)
				}
			})
		})
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
	}

	for input, wantErr := range tests {
		testName := input[0] + "->" + input[1]
		A := MustParse([]byte(input[0]))
		B := MustParse([]byte(input[1]))
		t.Run(testName, func(t *testing.T) {
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
