package data

import (
	"testing"

	"github.com/iimos/ucum/internal/types"
)

// TestConvFullyExpanded checks that the generated conversion table is fully expanded:
// no entry may reference another entry of the table, otherwise only one level of
// expansion would be applied and the component sets of convertible units may differ.
func TestConvFullyExpanded(t *testing.T) {
	for name, u := range Conv {
		for key := range u.Components {
			if _, ok := Conv[key.AtomCode]; ok {
				t.Errorf("%q references %q; the table must be fully expanded", name, key.AtomCode)
			}
		}
	}
}

// TestArbitraryUnitsAreOpaque checks that arbitrary units (UCUM §24) are not expanded
// into other units: their definition as the unity "1" must not make them dimensionless,
// otherwise they would become commensurable with every dimensionless unit, which the
// standard forbids (§24■3, §25■1).
func TestArbitraryUnitsAreOpaque(t *testing.T) {
	for _, code := range []string{"[iU]", "[CFU]", "[hp_X]", "[arb'U]", "[USP'U]"} {
		if _, ok := Conv[code]; ok {
			t.Errorf("arbitrary unit %q must not be expanded in Conv", code)
		}
	}

	// [IU] is defined as [iU] by the standard, so it is the only arbitrary unit
	// that stays in the table, as an alias of [iU].
	iu, ok := Conv["[IU]"]
	if !ok {
		t.Fatalf(`Conv["[IU]"] must be an alias of [iU]`)
	}
	exp, ok := iu.Components[types.ComponentKey{AtomCode: "[iU]"}]
	if len(iu.Components) != 1 || !ok || exp != 1 {
		t.Errorf(`Conv["[IU]"] must reference [iU], got %+v`, iu)
	}
}
