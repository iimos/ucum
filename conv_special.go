package ucum

import (
	"math/big"
)

// specialConverter is a converter for special units.
// Convertation algorithm: y = to(from(multiplyBefore * x)) / divideAfter
type specialConverter struct {
	from           PairConverter // converts from source unit to intermediate one
	to             PairConverter // converts from intermediate unit to target one
	multiplyBefore *big.Rat
	divideAfter    *big.Rat
}

var _ PairConverter = (*specialConverter)(nil)

func (c *specialConverter) ConvRat(val *big.Rat) *big.Rat {
	v1 := new(big.Rat).Mul(val, c.multiplyBefore)
	v2 := c.from.ConvRat(v1)
	v3 := c.to.ConvRat(v2)
	v4 := v3.Quo(v3, c.divideAfter)
	return v4
}

func (c *specialConverter) ConvBigInt(val *big.Int) (converted *big.Int, exact bool) {
	rat := new(big.Rat).SetInt(val)
	result := c.ConvRat(rat)
	quo, rem := new(big.Int).QuoRem(result.Num(), result.Denom(), new(big.Int))
	return quo, rem.Cmp(bigZero) == 0
}

func (c *specialConverter) ConvFloat64(val float64) float64 {
	multiplyBefore, _ := c.multiplyBefore.Float64()
	divideAfter, _ := c.divideAfter.Float64()
	v1 := val * multiplyBefore
	v2 := c.from.ConvFloat64(v1)
	v3 := c.to.ConvFloat64(v2)
	v4 := v3 / divideAfter
	return v4
}
