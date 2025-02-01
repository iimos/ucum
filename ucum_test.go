package ucum

import (
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
