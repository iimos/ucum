package types

import "math/big"

// Unit is a UCUM unit of measure.
type Unit struct {
	Orig       string
	Components map[ComponentKey]int // <unit atom, annotation> -> exponent
	Coeff      *big.Rat
}

type ComponentKey struct {
	AtomCode   string
	Annotation string
}

type Atom struct {
	Code      string
	Kind      string
	Magnitude *big.Rat
	Metric    bool
	// Arbitrary marks procedure defined units (UCUM §24). Arbitrary units are
	// not commensurable with other units, so they must not be expanded into
	// non-arbitrary units. An alias of another arbitrary unit (e.g. [IU] = [iU])
	// is the only definition that may still be expanded.
	Arbitrary bool
}
