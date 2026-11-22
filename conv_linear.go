package ucum

import "math/big"

type linearConverter struct {
	from, to   Unit
	ratio      big.Rat
	ratioFloat float64
}

func (c *linearConverter) ConvBigRat(val *big.Rat) *big.Rat {
	return new(big.Rat).Mul(&c.ratio, val)
}

func (c *linearConverter) ConvBigInt(val *big.Int) (converted *big.Int, exact bool) {
	num := new(big.Int).Mul(val, c.ratio.Num())
	if c.ratio.IsInt() {
		return num, true
	}
	quo, rem := new(big.Int), new(big.Int)
	quo.QuoRem(num, c.ratio.Denom(), rem)
	return quo, rem.Sign() == 0
}

func (c *linearConverter) ConvFloat64(val float64) float64 {
	return c.ratioFloat * val
}

func (c *linearConverter) String() string {
	return "ucum.PairConverter(" + c.from.String() + "->" + c.to.String() + ")"
}
