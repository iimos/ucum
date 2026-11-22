package ucum

import (
	"strings"
	"testing"
)

func TestParsedUnitString(t *testing.T) {
	tests := []struct {
		unit string
		want string
	}{
		{"m", "m"},
		{"km", "km"},
		{"km/km", "km/km"},
		{"1/(mol/kg)", "1/(mol/kg)"},
		{"h", "h"},
		{"h{ annot }", "h{ annot }"},
	}

	for _, test := range tests {
		u := MustParse([]byte(test.unit))
		got := u.String()
		if got != test.want {
			t.Errorf("Unit.String mismatch: want %s; got %s", test.want, got)
		}
	}
}

func TestUnitJSON(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		var u Unit
		if err := u.UnmarshalJSON([]byte("null")); err != nil {
			t.Fatalf("UnmarshalJSON(null) error = %v", err)
		}
		if u.String() != "" {
			t.Errorf("UnmarshalJSON(null) changed the unit: got %q, want an empty unit", u.String())
		}

		u = MustParse([]byte("km"))
		if err := u.UnmarshalJSON([]byte("null")); err != nil {
			t.Fatalf("UnmarshalJSON(null) error = %v", err)
		}
		if u.String() != "km" {
			t.Errorf("UnmarshalJSON(null) overwrote the unit: got %q, want %q", u.String(), "km")
		}
	})

	t.Run("non-string", func(t *testing.T) {
		inputs := []string{
			"123",
			"1.5",
			"true",
			"{}",
			"[]",
			`["m"]`,
		}
		for _, input := range inputs {
			t.Run(input, func(t *testing.T) {
				var u Unit
				err := u.UnmarshalJSON([]byte(input))
				if err == nil {
					t.Fatalf("UnmarshalJSON(%s) = %q, want an error", input, u.String())
				}
				if !strings.Contains(err.Error(), "ucum: error decoding string") {
					t.Errorf("UnmarshalJSON(%s) error = %q, want it to mention decoding failure", input, err)
				}
			})
		}
	})

	t.Run("invalid unit", func(t *testing.T) {
		var u Unit
		err := u.UnmarshalJSON([]byte(`"unknown"`))
		if err == nil {
			t.Fatalf("UnmarshalJSON(%q) = %q, want an error", "unknown", u.String())
		}
		if !strings.Contains(err.Error(), "unknown unit") {
			t.Errorf("UnmarshalJSON() error = %q, want it to mention the unknown unit", err)
		}
	})

	t.Run("valid", func(t *testing.T) {
		var u Unit
		if err := u.UnmarshalJSON([]byte(`"km"`)); err != nil {
			t.Fatalf("UnmarshalJSON() error = %v", err)
		}
		if u.String() != "km" {
			t.Errorf("UnmarshalJSON() = %q, want %q", u.String(), "km")
		}
	})
}

// func TestNormalisedUnitString(t *testing.T) {
// 	tests := []struct {
// 		unit string
// 		want string
// 	}{
// 		{"m", "m"},
// 		{"km", "km"},
// 		{"km/km", "m/m"},
// 		{"1/(mol/kg)", "?"},
// 		{"h", "h"},
// 		{"h{ annot }", "?"},
// 	}

// 	for _, test := range tests {
// 		u := _Normalize(MustParse([]byte(test.unit)))
// 		got := u.String()
// 		if got != test.want {
// 			t.Errorf("Unit.String mismatch: want %s; got %s", test.want, got)
// 		}
// 	}
// }
