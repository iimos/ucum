package ucumparser

import (
	"github.com/iimos/ucum/internal/types"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]string{
		"1":                     "1",
		"100":                   "100",
		"{annot}":               "{annot}",
		"{}":                    "{}",
		"kg10/2":                "500000000000000000000000000000⋅g¹⁰",
		"kg-10.2":               "1/500000000000000000000000000000⋅g⁻¹⁰",
		"kg.m/s2":               "1000⋅g⋅m⋅s⁻²",
		"10*":                   "10*",
		"10*6":                  "10*⁶",
		"10^78":                 "10^⁷⁸",
		"((((((((100))))))))":   "100",
		"((((((((100)))))))).2": "200",
		"(100).m":               "100⋅m",
		"ng/(24.h)":             "1/24000000000⋅g⋅h⁻¹",
		"g/h/m2":                "g⋅h⁻¹⋅m⁻²",
		"kcal/kg":               "cal⋅g⁻¹",
		"kcal/kg/(24.h)":        "1/24⋅cal⋅g⁻¹⋅h⁻¹",
		"((kg)/(m.(s)))":        "1000⋅g⋅m⁻¹⋅s⁻¹",
		"m4/m2":                 "m²",
		"m4.m2.m":               "m⁷",
		"/m":                    "m⁻¹",
		"/(/m)":                 "m",
		"m3{annot1}/m2{annot2}": "m³{annot1}⋅m⁻²{annot2}", // different annotations are not mixed together
		"m3{annot1}/m2{annot1}": "m{annot1}",
		"u[IU]":                 "1/1000000⋅[IU]",
		"mg":                    "1/1000⋅g",
		"TiBy":                  "1099511627776⋅By",
		"YBy":                   "1000000000000000000000000⋅By",
		"yBy":                   "1/1000000000000000000000000⋅By",
		"L/L":                   "L⁰",
		"m[H2O]":                "m[H2O]",
		"m[H2O].[in_i]/m":       "[in_i]⋅m[H2O]⋅m⁻¹",
		"cm[Hg]":                "1/100⋅m[Hg]",
		"%[slope]":              "%[slope]",
		"10.[m/s2/Hz^(1/2)]":    "10⋅[m/s2/Hz^(1/2)]",
		"kCel":                  "1000⋅Cel", // special units are allowed to have prefix
	}
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			unit, err := Parse([]byte(input))
			if err != nil {
				t.Errorf("Parse(%q) error = %s", input, err)
				return
			}
			gotCanonical := unit.CanonicalString()
			if gotCanonical != want {
				t.Errorf("Parse(%q) = %s, want %s", input, gotCanonical, want)
			}
			if unit.Orig != input {
				t.Errorf("unit.Orig = %q, want %q", unit.Orig, input)
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	// some test cases are taken from https://github.com/dalito/ucumvert/blob/0beec522041d086f4ed5e1eb0259b0e183ad7a73/tests/test_parser.py#L48-L62
	tests := map[string]string{
		// input -> error
		"":                                 `ucum: empty unit`,
		" ":                                `ucum: spaces are not allowed`,
		"unknown":                          `ucum: unknown unit "unknown" at position 0`,
		"2.unknown":                        `ucum: unknown unit "unknown" at position 2`,
		"{unclosed annotation":             `ucum: unterminated annotation, "}" expected at position 20`,
		")":                                `ucum: unexpected symbol ")" at position 0`,
		"}":                                `ucum: unexpected symbol "}" at position 0`,
		"m//s":                             `ucum: unexpected symbol "/" at position 2`,
		"m/.s":                             `ucum: unexpected symbol "." at position 2`,
		"da":                               `ucum: unknown unit "da" at position 0`, // a is not metric
		"1{annot1}{annot2}":                `ucum: unexpected symbol "{" at position 9`,
		"(m/s)2":                           `ucum: unexpected symbol "2" at position 5`, // invalid since UCUM v 1.9
		"m(s)":                             `ucum: unexpected symbol "(" at position 1`,
		"{annotation at wrong position}kg": `ucum: unexpected symbol "k" at position 30`, // missing operator
		"{annotation}2":                    `ucum: unexpected symbol "2" at position 12`, // exponent after annotation
		"m/":                               `ucum: unexpected end`,
		"m.":                               `ucum: unexpected end`,
		"(m":                               `ucum: unexpected end, missing ")"`,
		".m":                               `ucum: unexpected symbol "." at position 0`,
		"2mg":                              `ucum: unknown unit "2mg" at position 0`, // missing operator
		"9999999999999999999999999":        `ucum: number too large at position 0`,   // overflow
		"m999999999999999999999999":        `ucum: number too large at position 1`,   // overflow
		"m-99999999999999999999999":        `ucum: number too large at position 2`,   // overflow
		"1/99999999999999999999999":        `ucum: number too large at position 2`,   // overflow
		"m[H2O":                            `ucum: unclosed square bracket at position 1`,
		"[":                                `ucum: unclosed square bracket at position 0`,
		"]":                                `ucum: unknown unit "]" at position 0`,
		"%[slope].m":                       `ucum: invalid unit: non-ratio unit '%[slope]' cannot be combined with other units`,
		"Cel/s":                            `ucum: invalid unit: non-ratio unit 'Cel' cannot be combined with other units`,
		"Cel2":                             `ucum: invalid unit: non-ratio unit 'Cel' cannot be raised to a power`,
		"Cel-1":                            `ucum: invalid unit: non-ratio unit 'Cel' cannot be raised to a power`,
	}

	for input, wantErr := range tests {
		t.Run(input, func(t *testing.T) {
			unit, err := Parse([]byte(input))
			if err == nil {
				t.Errorf("no error, want %q", wantErr)
				return
			}
			if err.Error() != wantErr {
				t.Errorf("Parse() error = %q, want %q", err, wantErr)
				return
			}
			if !reflect.DeepEqual(unit, types.Unit{}) {
				t.Errorf("Parse() returned error with nonempty Unit: %#v", unit)
			}
		})
	}
}
