package ucum

import (
	"math/big"
	"sync"
)

// ConverterBigRat converts big.Rat values from one unit to another.
type ConverterBigRat interface {
	ConvBigRat(val *big.Rat, from, to string) (*big.Rat, error)
}

// ConverterBigInt converts big.Int values from one unit to another.
type ConverterBigInt interface {
	ConvBigInt(val *big.Int, from, to string) (converted *big.Int, exact bool, err error)
}

// ConverterFloat64 converts float values from one unit to another.
type ConverterFloat64 interface {
	ConvFloat64(val float64, from, to string) (float64, error)
}

// DefaultConverter is a default instance of Conv.
var DefaultConverter = NewConverter()

// ConvBigRat converts a value from one unit to another using big.Rat.
// Returns the converted value as a big.Rat and an error if conversion fails.
func ConvBigRat(value *big.Rat, from, to string) (*big.Rat, error) {
	return DefaultConverter.ConvBigRat(value, from, to)
}

// ConvBigInt converts a value from one unit to another using big.Int.
// Returns the converted value as a big.Int, a boolean indicating if the conversion was exact, and an error if conversion fails.
func ConvBigInt(value *big.Int, from, to string) (converted *big.Int, exact bool, err error) {
	return DefaultConverter.ConvBigInt(value, from, to)
}

// ConvFloat64 converts a float64 value from one unit to another.
// Returns the converted value as a float64 and an error if conversion fails.
func ConvFloat64(value float64, from, to string) (float64, error) {
	return DefaultConverter.ConvFloat64(value, from, to)
}

// NewConverter creates a converter. Converter maintains an internal cache and is designed to be thread-safe.
//
// Number of distinct units used by one application is usualy not big
// so cache stores items forever and never evict keys.
func NewConverter() *Conv {
	return &Conv{
		mu:    &sync.RWMutex{},
		cache: make(map[string]Unit),
	}
}

// Conv is a converter. Conv maintains an internal cache and is designed to be thread-safe.
//
// Number of distinct units used by one application is usualy not big
// so cache stores items forever and never evict keys.
type Conv struct {
	mu    *sync.RWMutex
	cache map[string]Unit
}

var _ ConverterBigRat = &Conv{}
var _ ConverterBigInt = &Conv{}
var _ ConverterFloat64 = &Conv{}

// ConvBigRat converts a value from one unit to another using big.Rat.
// Returns the converted value as a big.Rat and an error if conversion fails.
func (c *Conv) ConvBigRat(val *big.Rat, from, to string) (*big.Rat, error) {
	pairConv, err := c.pairConv(from, to)
	if err != nil {
		return nil, err
	}
	return pairConv.ConvBigRat(val), nil
}

// ConvBigInt converts a value from one unit to another using big.Int.
// Returns the converted value as a big.Int, a boolean indicating if the conversion was exact, and an error if conversion fails.
func (c *Conv) ConvBigInt(val *big.Int, from, to string) (converted *big.Int, exact bool, err error) {
	pairConv, err := c.pairConv(from, to)
	if err != nil {
		return nil, false, err
	}
	converted, exact = pairConv.ConvBigInt(val)
	return converted, exact, nil
}

// ConvFloat64 converts a float64 value from one unit to another.
// Returns the converted value as a float64 and an error if conversion fails.
func (c *Conv) ConvFloat64(val float64, from, to string) (float64, error) {
	pairConv, err := c.pairConv(from, to)
	if err != nil {
		return 0, err
	}
	return pairConv.ConvFloat64(val), nil
}

func (c *Conv) parse(unit string) (Unit, error) {
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

func (c *Conv) pairConv(from, to string) (PairConverter, error) {
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
