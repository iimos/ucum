package ucum

import (
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
