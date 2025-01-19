package ucum

import (
	"math/big"
	"sync"
)

var defaultConverter = NewConverter()

func ConvRat(val *big.Rat, from, to string) (*big.Rat, error) {
	return defaultConverter.ConvRat(val, from, to)
}
func ConvBigInt(val *big.Int, from, to string) (converted *big.Int, exact bool, err error) {
	return defaultConverter.ConvBigInt(val, from, to)
}
func ConvFloat64(val float64, from, to string) (float64, error) {
	return defaultConverter.ConvFloat64(val, from, to)
}

type Conv interface {
	ConvRat(val *big.Rat, from, to string) (*big.Rat, error)
	ConvBigInt(val *big.Int, from, to string) (converted *big.Int, exact bool, err error)
	ConvFloat64(val float64, from, to string) (float64, error)
}

func NewConverter() Conv {
	return &conv{
		mu:    &sync.RWMutex{},
		cache: make(map[string]Unit),
	}
}

type conv struct {
	mu    *sync.RWMutex
	cache map[string]Unit
}

func (c *conv) ConvRat(val *big.Rat, from, to string) (*big.Rat, error) {
	pairConv, err := c.pairConv(from, to)
	if err != nil {
		return nil, err
	}
	return pairConv.ConvRat(val), nil
}

func (c *conv) ConvBigInt(val *big.Int, from, to string) (converted *big.Int, exact bool, err error) {
	pairConv, err := c.pairConv(from, to)
	if err != nil {
		return nil, false, err
	}
	converted, exact = pairConv.ConvBigInt(val)
	return converted, exact, nil
}

func (c *conv) ConvFloat64(val float64, from, to string) (float64, error) {
	pairConv, err := c.pairConv(from, to)
	if err != nil {
		return 0, err
	}
	return pairConv.ConvFloat64(val), nil
}

func (c *conv) parse(unit string) (Unit, error) {
	c.mu.RLock()
	u, exists := c.cache[unit]
	c.mu.RUnlock()
	if exists {
		return u, nil
	}

	u, err := Parse([]byte(unit))
	if err != nil {
		return Unit{}, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[unit] = u
	return u, nil
}

func (c *conv) pairConv(from, to string) (PairConverter, error) {
	fromUnit, err := c.parse(from)
	if err != nil {
		return nil, err
	}
	toUnit, err := c.parse(to)
	if err != nil {
		return nil, err
	}
	return NewPairConverter(fromUnit, toUnit)
}

func (c *conv) withRLock(f func()) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	f()
}
