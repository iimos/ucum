package ucum

import (
	"encoding/json"
	"testing"
)

func FuzzParse(f *testing.F) {
	seeds := []string{
		"m",
		"km/h",
		"1/(mol/kg)",
		"h{ annot }",
		"m3",
		"kg/m2",
		"",
		" ",
		"{}",
		"1/0",
		"m/m",
		"m0",
		"Cel",
		"[iU]",
		"[IU]",
		"[m/s2/Hz^(1/2)]",
		"m[H2O]",
		"10*",
		`m{ann"ot}`,
		`m{a\b}`,
		"m{a\tb}",
		// invalid UTF-8 must be rejected, not silently replaced with U+FFFD
		"{\xff}",
		"m{\xf4}",
		"m\xc3",
	}

	for _, seed := range seeds {
		f.Add([]byte(seed))
	}

	f.Fuzz(func(t *testing.T, input []byte) {
		unit, err := Parse(input)

		if err == nil {
			// If parsing succeeded, the String() representation should be parseable again
			str := unit.String()
			unit2, err2 := Parse([]byte(str))
			if err2 != nil {
				t.Errorf("Re-parsing of successful parse failed: original=%q, stringified=%q, error=%v",
					string(input), str, err2)
			}

			// The string representations should match
			if str != unit2.String() {
				t.Errorf("String representations don't match after re-parse: first=%q, second=%q",
					str, unit2.String())
			}
		}

		// Test JSON marshaling/unmarshaling for valid units
		if err == nil {
			jsonBytes, err := unit.MarshalJSON()
			if err != nil {
				t.Errorf("MarshalJSON failed for valid unit %q: %v", string(input), err)
			}

			// The output must be valid JSON, not just something that our own
			// UnmarshalJSON happens to accept.
			if !json.Valid(jsonBytes) {
				t.Errorf("MarshalJSON returned invalid JSON for unit %q: %s", string(input), jsonBytes)
			}

			var str string
			if err := json.Unmarshal(jsonBytes, &str); err != nil {
				t.Errorf("json.Unmarshal failed for unit %q: %v", string(input), err)
			} else if str != unit.String() {
				t.Errorf("JSON string round-trip mismatch: original=%q, after=%q", unit.String(), str)
			}

			var unit2 Unit
			err = unit2.UnmarshalJSON(jsonBytes)
			if err != nil {
				t.Errorf("UnmarshalJSON failed for valid unit %q: %v", string(input), err)
			}

			if unit.String() != unit2.String() {
				t.Errorf("JSON round-trip produced different string: original=%q, after=%q",
					unit.String(), unit2.String())
			}
		}
	})
}
