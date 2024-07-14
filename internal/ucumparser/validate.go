package ucumparser

import (
	"errors"
	"github.com/iimos/ucum/internal/data"
	"github.com/iimos/ucum/internal/types"
)

func validate(u types.Unit) error {
	// special units cannot take part in any algebraic operations involving other units
	// https://ucum.org/ucum#section-Special-Units-on-non-ratio-Scales
	for comp, exponent := range u.Components {
		if _, isSpecial := data.SpecialUnits[comp.AtomCode]; isSpecial {
			if len(u.Components) > 1 {
				return errors.New("ucum: invalid unit: non-ratio unit '" + comp.AtomCode + "' cannot be combined with other units")
			}
			if exponent != 1 {
				return errors.New("ucum: invalid unit: non-ratio unit '" + comp.AtomCode + "' cannot be raised to a power")
			}
		}
	}
	return nil
}
