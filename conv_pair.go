package ucum

import (
	"fmt"
	"math/big"

	"github.com/iimos/ucum/internal/data"
	"github.com/iimos/ucum/internal/types"
)

var (
	bigZero   = big.NewInt(0)
	bigRatOne = big.NewRat(1, 1)
)

// PairConverter makes conversion between two UCUM units.
type PairConverter interface {
	ConvBigRat(val *big.Rat) *big.Rat
	ConvBigInt(val *big.Int) (converted *big.Int, exact bool)
	ConvFloat64(val float64) float64
}

// NewPairConverter creates a new PairConverter.
func NewPairConverter(from, to Unit) (PairConverter, error) {
	a := _Normalize(from).u
	b := _Normalize(to).u

	if len(a.Components) != len(b.Components) {
		// Special units are not normalizable so if number of components doesn't match it might be that units are special
		specConv, ok := newSpecialConverter(a, b)
		if !ok {
			return nil, fmt.Errorf("ucum: %q cannot be converted to %q", from.String(), to.String())
		}
		return specConv, nil
	}

	for key, expA := range a.Components {
		expB, exists := b.Components[key] // normalized units are stripped from annotations, so we can look up directly by key
		if !exists {
			// Special units are not normalizable so try to interpret it as a special units if mismatched.
			specConv, ok := newSpecialConverter(a, b)
			if !ok {
				return nil, fmt.Errorf("ucum: %q cannot be converted to %q", from.String(), to.String())
			}
			return specConv, nil
		}
		if expA != expB {
			return nil, fmt.Errorf("ucum: %q cannot be converted to %q", from.String(), to.String())
		}
	}
	ratio := new(big.Rat).Quo(a.Coeff, b.Coeff)
	ratioFloat, _ := ratio.Float64()
	return &linearConverter{
		from:       from,
		to:         to,
		ratio:      *ratio,
		ratioFloat: ratioFloat,
	}, nil
}

// newSpecialConverter creates convertor for special units.
// It assumes that the special units are already normalized.
func newSpecialConverter(from, to types.Unit) (PairConverter, bool) {
	specialConv := func(u types.Unit) (conv data.SpecialUnitConv, ok bool) {
		if len(u.Components) != 1 {
			return data.SpecialUnitConv{}, false
		}
		for key, exp := range u.Components {
			if exp != 1 {
				// special units cannot be raised to a power
				return data.SpecialUnitConv{}, false
			}
			conv, ok = data.SpecialUnits[key.AtomCode]
			return conv, ok
		}
		return data.SpecialUnitConv{}, false
	}

	if fromConv, ok := specialConv(from); ok {
		interm := MustParse([]byte(fromConv.Unit))
		toConv, err := NewPairConverter(interm, Unit{u: to})
		if err != nil {
			return nil, false
		}
		return &specialConverter{
			multiplyBefore: from.Coeff,
			from:           fromConv,
			to:             toConv,
			divideAfter:    bigRatOne,
		}, true
	}

	if toConv, ok := specialConv(to); ok {
		interm := MustParse([]byte(toConv.Unit))
		fromConv, err := NewPairConverter(Unit{u: from}, interm)
		if err != nil {
			return nil, false
		}
		return &specialConverter{
			multiplyBefore: bigRatOne,
			from:           fromConv,
			to:             toConv.Invert(),
			divideAfter:    to.Coeff,
		}, true
	}

	return nil, false
}
