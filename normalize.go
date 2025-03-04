package ucum

import (
	"math/big"

	"github.com/iimos/ucum/internal/data"
	"github.com/iimos/ucum/internal/types"
)

// Normalize builds a normalized version of the given unit. Normalized unit consists of base and special units only.
func _Normalize(unit Unit) Unit {
	// It's not ready to be exported while we cant generate string representation fot normalised units
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
	return ret
}

func pow(num *big.Rat, exp int) *big.Rat {
	cpy := new(big.Rat).Set(num)
	if exp < 0 {
		cpy = cpy.Inv(cpy)
		exp = -exp
	} else if exp == 0 {
		cpy.SetFrac64(0, 1)
	}
	multiplier := new(big.Rat).Set(cpy)
	for i := exp; i > 1; i-- {
		cpy.Mul(cpy, multiplier)
	}
	return cpy
}
