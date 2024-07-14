package ucum

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := map[string]string{
		"1":                "1",
		"100":              "100",
		"{annot}":          "1",
		"{}":               "1",
		"kg10/2":           "500000000000000000000000000000⋅g¹⁰",
		"kg-10.2":          "1/500000000000000000000000000000⋅g⁻¹⁰",
		"min":              "60⋅s",
		"min2":             "3600⋅s²",
		"min2/s":           "3600⋅s",
		"m[H2O]":           "9806650⋅g⋅m⁻¹⋅s⁻²",
		"[in_us]":          "100/3937⋅m",
		"[gr]":             "6479891/100000000⋅g",
		"360.3600.''/[pi]": "2⋅rad",
		"180.deg/[pi]":     "rad",
		"2.Cel":            "2⋅Cel", // special units are not normalizable
		"2.kCel/3":         "2000/3⋅Cel",
	}
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			inputUnit, err := Parse([]byte(input))
			if err != nil {
				t.Fatal(err)
			}
			got := Normalize(inputUnit)
			gotCanonical := got.u.CanonicalString()
			if gotCanonical != want {
				t.Errorf("Normalize(%q) = %q, want %s", input, gotCanonical, want)
			}
		})
	}
}
