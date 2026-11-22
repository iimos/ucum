package ucum

// This file contains regression tests for problems found in the
// implementation. Each test pins the behaviour required by UCUM and by the
// package documentation, so it failed before the corresponding bug was fixed
// and keeps failing if the bug is reintroduced.
//
// TestSpecialUnitDomainErrors covers a known, not yet fixed problem with values
// outside the domain of logarithmic and fractional special units; it is skipped
// until the fix lands.

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"testing"
)

// mustNotPanic runs f and reports an unexpected panic as a test failure.
func mustNotPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	f()
}

// TestConvBigIntNonIntegerRatio checks ConvBigInt for conversions whose ratio
// has a numerator other than 1.
//
// The result must be val*ratio.Num()/ratio.Denom(). A previous implementation
// dropped the numerator, dividing the original value by the denominator and
// silently returning a wrong result marked as exact.
func TestConvBigIntNonIntegerRatio(t *testing.T) {
	tests := []struct {
		from, to  string
		value     int64
		want      int64
		wantExact bool
	}{
		// 1 [in_us] = 100/3937 m
		{from: "[in_us]", to: "m", value: 3937, want: 100, wantExact: true},
		{from: "[in_us]", to: "m", value: 7874, want: 200, wantExact: true},
		{from: "[in_us]", to: "m", value: -3937, want: -100, wantExact: true},
		// 1 [ft_i] = 381/1250 m
		{from: "[ft_i]", to: "m", value: 10000, want: 3048, wantExact: true},
		{from: "[ft_i]", to: "m", value: 1000, want: 304, wantExact: false},
	}

	for _, tt := range tests {
		tname := fmt.Sprintf("%d %s -> %s", tt.value, tt.from, tt.to)
		t.Run(tname, func(t *testing.T) {
			value := big.NewInt(tt.value)

			got, exact, err := ConvBigInt(value, tt.from, tt.to)
			if err != nil {
				t.Fatalf("ConvBigInt() error = %v", err)
			}
			if got.Cmp(big.NewInt(tt.want)) != 0 {
				t.Errorf("ConvBigInt() = %s, want %d", got, tt.want)
			}
			if exact != tt.wantExact {
				t.Errorf("ConvBigInt() exact = %t, want %t", exact, tt.wantExact)
			}

			conv, err := NewPairConverter(MustParse([]byte(tt.from)), MustParse([]byte(tt.to)))
			if err != nil {
				t.Fatalf("NewPairConverter() error = %v", err)
			}
			got, exact = conv.ConvBigInt(value)
			if got.Cmp(big.NewInt(tt.want)) != 0 {
				t.Errorf("PairConverter.ConvBigInt() = %s, want %d", got, tt.want)
			}
			if exact != tt.wantExact {
				t.Errorf("PairConverter.ConvBigInt() exact = %t, want %t", exact, tt.wantExact)
			}
		})
	}
}

// TestPowZeroExponent checks that raising to the power 0 yields 1, not 0.
//
// Both the parser (for units like "m0") and normalize (for terms whose
// components cancel out, like "L/L") used to set the coefficient to 0.
func TestPowZeroExponent(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		for _, input := range []string{"m0", "L0", "kg0", "km0"} {
			t.Run(input, func(t *testing.T) {
				unit := MustParse([]byte(input))
				if got := unit.u.Coeff; got.Cmp(big.NewRat(1, 1)) != 0 {
					t.Errorf("Parse(%q): coefficient = %s, want 1", input, got)
				}
			})
		}
	})

	t.Run("Normalize", func(t *testing.T) {
		for _, input := range []string{"L/L", "m2/m2", "min/min", "mol/mol"} {
			t.Run(input, func(t *testing.T) {
				norm := _Normalize(MustParse([]byte(input)))
				if got := norm.u.Coeff; got.Cmp(big.NewRat(1, 1)) != 0 {
					t.Errorf("_Normalize(%q): coefficient = %s, want 1", input, got)
				}
			})
		}
	})
}

// TestDimensionlessConversionToUnity checks that a term whose components cancel
// out is compatible with the unity "1" and with other dimensionless terms.
//
// Normalization used to keep zero-exponent components in the component set, so
// the conversion was rejected (and "L/L" even panicked with division by zero
// because its coefficient was 0).
func TestDimensionlessConversionToUnity(t *testing.T) {
	tests := []struct {
		from, to string
		value    *big.Rat
		want     *big.Rat
	}{
		{from: "m/m", to: "1", value: big.NewRat(5, 1), want: big.NewRat(5, 1)},
		{from: "L/L", to: "1", value: big.NewRat(5, 1), want: big.NewRat(5, 1)},
		{from: "m2/m2", to: "1", value: big.NewRat(5, 1), want: big.NewRat(5, 1)},
		{from: "L/L", to: "L2/L2", value: big.NewRat(5, 1), want: big.NewRat(5, 1)},
		{from: "min/min", to: "1", value: big.NewRat(5, 1), want: big.NewRat(5, 1)},
		{from: "1", to: "m/m", value: big.NewRat(5, 1), want: big.NewRat(5, 1)},
	}

	for _, tt := range tests {
		t.Run(tt.from+" -> "+tt.to, func(t *testing.T) {
			mustNotPanic(t, func() {
				got, err := ConvBigRat(new(big.Rat).Set(tt.value), tt.from, tt.to)
				if err != nil {
					t.Fatalf("ConvBigRat(%s, %q, %q) error = %v", tt.value, tt.from, tt.to, err)
				}
				if got.Cmp(tt.want) != 0 {
					t.Errorf("ConvBigRat() = %s, want %s", got, tt.want)
				}
			})
		})
	}
}

// TestSpecialUnitDomainErrors checks that logarithmic and fractional special
// units reject values outside their domain with an error instead of panicking.
//
// Known open problem: the conversion functions build a big.Rat from
// math.Log/math.Sqrt results without checking them, and big.Rat.SetFloat64
// returns nil for ±Inf/NaN, which leads to a nil pointer dereference. The test
// is skipped until this is fixed.
func TestSpecialUnitDomainErrors(t *testing.T) {
	t.Skip("todo: conversion from values outside the domain of logarithmic and fractional special units panics instead of returning an error")

	tests := []struct {
		from, to string
		value    float64
	}{
		// log10/log of zero is -Inf
		{from: "1", to: "[hp'_X]", value: 0},
		{from: "1", to: "[hp'_C]", value: 0},
		{from: "1", to: "[hp'_M]", value: 0},
		{from: "1", to: "[hp'_Q]", value: 0},
		{from: "1", to: "B", value: 0},
		{from: "1", to: "B[W]", value: 0},
		{from: "1", to: "bit_s", value: 0},
		// log of a negative number is NaN
		{from: "1", to: "B", value: -1},
		{from: "1", to: "Np", value: -1},
		{from: "1", to: "B[V]", value: -1},
		{from: "1", to: "B[SPL]", value: -1},
		{from: "1", to: "B[kW]", value: -1},
		// square root of a negative number is NaN
		{from: "m2/s4/Hz", to: "[m/s2/Hz^(1/2)]", value: -4},
	}

	for _, tt := range tests {
		tname := fmt.Sprintf("%g %s -> %s", tt.value, tt.from, tt.to)
		t.Run("ConvFloat64 "+tname, func(t *testing.T) {
			mustNotPanic(t, func() {
				got, err := ConvFloat64(tt.value, tt.from, tt.to)
				if err == nil {
					t.Errorf("ConvFloat64() = %v, want an error for a value outside the domain", got)
				}
			})
		})
		t.Run("ConvBigRat "+tname, func(t *testing.T) {
			mustNotPanic(t, func() {
				value := big.NewRat(1, 1)
				if tt.value != 1 {
					value.SetInt64(int64(tt.value))
				}
				got, err := ConvBigRat(value, tt.from, tt.to)
				if err == nil {
					t.Errorf("ConvBigRat() = %v, want an error for a value outside the domain", got)
				}
			})
		})
	}
}

// TestMarshalJSONEscaping checks that Unit.MarshalJSON produces valid JSON.
//
// A previous implementation wrapped String() in quotes without escaping, so
// annotations and units containing quotes, backslashes or control characters
// produced invalid JSON.
func TestMarshalJSONEscaping(t *testing.T) {
	tests := []string{
		`m{ann"ot}`,
		`m{a\b}`,
		"m{a\tb}",
		`m{"}`,
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			unit := MustParse([]byte(input))

			data, err := unit.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON() error = %v", err)
			}
			if !json.Valid(data) {
				t.Fatalf("MarshalJSON() = %s, not valid JSON", data)
			}

			var got string
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatalf("json.Unmarshal(%s) error = %v", data, err)
			}
			if got != input {
				t.Errorf("JSON round-trip = %q, want %q", got, input)
			}

			var decoded Unit
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("json.Unmarshal into Unit error = %v", err)
			}
			if decoded.String() != input {
				t.Errorf("decoded unit = %q, want %q", decoded.String(), input)
			}
		})
	}
}

// TestParseRejectsInvalidUTF8 checks that invalid UTF-8 is rejected instead of
// being silently accepted.
//
// Unit.Orig keeps the parsed bytes verbatim, so accepting invalid UTF-8 would
// break the String/JSON round-trip: encoding/json replaces invalid sequences
// with U+FFFD, so the decoded unit would differ from the original one.
func TestParseRejectsInvalidUTF8(t *testing.T) {
	inputs := []string{
		"{\xf4}",      // lone continuation byte in an annotation
		"m{\xff}",     // invalid byte in an annotation
		"{\xc3}",      // truncated multi-byte sequence
		"m{\xf0\x9f}", // truncated 4-byte sequence (emoji prefix)
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
			if _, err := Parse([]byte(input)); err == nil {
				t.Errorf("Parse(%q) = nil error, want error", input)
			}
		})
	}

	// Valid UTF-8 annotations must still work and round-trip through JSON.
	for _, input := range []string{"m{°C}", "m{微}", "m{🌡}"} {
		t.Run(input, func(t *testing.T) {
			unit := MustParse([]byte(input))
			data, err := json.Marshal(&unit)
			if err != nil {
				t.Fatalf("json.Marshal() error = %v", err)
			}
			var decoded Unit
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("json.Unmarshal(%s) error = %v", data, err)
			}
			if decoded.String() != input {
				t.Errorf("JSON round-trip = %q, want %q", decoded.String(), input)
			}
		})
	}
}

// TestConvBigIntSpecialUnits covers integer conversion through special units.
func TestConvBigIntSpecialUnits(t *testing.T) {
	tests := []struct {
		from, to  string
		value     int64
		want      int64
		wantExact bool
	}{
		{from: "Cel", to: "K", value: 300, want: 573, wantExact: false},     // 573.15 K
		{from: "K", to: "Cel", value: 300, want: 26, wantExact: false},      // 26.85 °C
		{from: "kCel", to: "K", value: 300, want: 300273, wantExact: false}, // 300273.15 K
		{from: "[degF]", to: "K", value: 300, want: 422, wantExact: false},  // 422.038… K
	}

	for _, tt := range tests {
		tname := fmt.Sprintf("%d %s -> %s", tt.value, tt.from, tt.to)
		t.Run(tname, func(t *testing.T) {
			got, exact, err := ConvBigInt(big.NewInt(tt.value), tt.from, tt.to)
			if err != nil {
				t.Fatalf("ConvBigInt() error = %v", err)
			}
			if got.Cmp(big.NewInt(tt.want)) != 0 {
				t.Errorf("ConvBigInt() = %s, want %d", got, tt.want)
			}
			if exact != tt.wantExact {
				t.Errorf("ConvBigInt() exact = %t, want %t", exact, tt.wantExact)
			}
		})
	}
}

// TestConvCacheConcurrent checks that the converter cache is safe for concurrent
// use (run with -race to actually verify it).
func TestConvCacheConcurrent(t *testing.T) {
	conv := NewConverter()

	const workers = 8
	const iterations = 100

	var wg sync.WaitGroup
	errs := make(chan error, workers*iterations)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				got, err := conv.ConvBigRat(big.NewRat(1, 1), "km", "m")
				if err != nil {
					errs <- err
					return
				}
				if got.Cmp(big.NewRat(1000, 1)) != 0 {
					errs <- fmt.Errorf("ConvBigRat(1, km, m) = %s, want 1000", got)
					return
				}
			}
		}()
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		t.Error(err)
	}
}
