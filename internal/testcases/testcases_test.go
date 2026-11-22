package testcases

import (
	"math/big"
	"testing"
)

func TestConversions(t *testing.T) {
	cases, err := Conversions()
	if err != nil {
		t.Fatalf("Conversions() error = %v", err)
	}

	seen := make(map[string]bool, len(cases))
	for _, c := range cases {
		name := c.Name()
		if seen[name] {
			t.Errorf("duplicate case name %q", name)
		}
		seen[name] = true

		if _, ok := new(big.Rat).SetString(c.Value); !ok {
			t.Errorf("%s: value %q is not a number", name, c.Value)
		}
		if _, ok := new(big.Rat).SetString(c.Want); !ok {
			t.Errorf("%s: want %q is not a number", name, c.Want)
		}
	}
}
