package ucum

import (
	"math/big"

	"github.com/iimos/ucum/internal/data"
	"github.com/iimos/ucum/internal/types"
)

// Normalize builds a normalized version of the given unit. Normalized unit consists of base and special units only.
func _Normalize(unit Unit) Unit {
	// It's not ready to be exported while we cant generate string representation for normalised units
	norm := Unit{u: normalize(unit.u)}
	// norm.u.Orig = ??? //todo
	return norm
}

func normalize(unit types.Unit) types.Unit {
	ret := types.Unit{
		Coeff:      (&big.Rat{}).Set(unit.Coeff),
		Components: make(map[types.ComponentKey]int, len(unit.Components)),
	}
	for key, exponent := range unit.Components {
		if key.AtomCode == "" {
			k := types.ComponentKey{AtomCode: "1"}
			ret.Components[k] += exponent
			continue
		}

		normed, ok := data.Conv[key.AtomCode]
		if !ok {
			k := types.ComponentKey{AtomCode: key.AtomCode} // strip annotation
			ret.Components[k] += exponent
			continue
		}

		coeff := pow(normed.Coeff, exponent)
		ret.Coeff = ret.Coeff.Mul(ret.Coeff, coeff)

		for key2, exponent2 := range normed.Components {
			ret.Components[key2] += exponent * exponent2
		}
	}

	// Components that cancel out must not stay in the component set: a term
	// like "m/m" is dimensionless and has to be compatible with "1".
	// Arbitrary units are the exception: even a cancelled arbitrary unit keeps
	// its term incommensurable with everything else (UCUM §24■3, §25■1).
	for key, exponent := range ret.Components {
		if exponent != 0 {
			continue
		}
		if atom, ok := data.Atoms[key.AtomCode]; ok && atom.Arbitrary {
			continue
		}
		delete(ret.Components, key)
	}
	return ret
}

func pow(num *big.Rat, exp int) *big.Rat {
	cpy := new(big.Rat).Set(num)
	if exp < 0 {
		cpy = cpy.Inv(cpy)
		exp = -exp
	} else if exp == 0 {
		cpy.SetFrac64(1, 1)
	}
	multiplier := new(big.Rat).Set(cpy)
	for i := exp; i > 1; i-- {
		cpy.Mul(cpy, multiplier)
	}
	return cpy
}
