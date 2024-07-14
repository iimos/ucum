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
}
