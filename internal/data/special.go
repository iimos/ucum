package data

import (
	"math"
	"math/big"
)

var (
	bigZero    = big.NewInt(0)
	kelvinZero = big.NewRat(5463, 20) // 273.15 °C
)

type SpecialUnitConv struct {
	Unit    string
	To      func(*big.Rat) *big.Rat
	From    func(*big.Rat) *big.Rat
	Inexact bool
}

func (c SpecialUnitConv) ConvRat(val *big.Rat) *big.Rat {
	return c.To(val)
}

func (c SpecialUnitConv) ConvBigInt(val *big.Int) (converted *big.Int, exact bool) {
	rat := new(big.Rat).SetInt(val)
	result := c.To(rat)
	if result.IsInt() {
		return result.Num(), true
	}
	quo, rem := new(big.Int).QuoRem(result.Num(), result.Denom(), new(big.Int))
	return quo, rem.Cmp(bigZero) == 0
}

func (c SpecialUnitConv) ConvFloat64(val float64) float64 {
	rat := new(big.Rat).SetFloat64(val)
	f, _ := c.To(rat).Float64()
	return f
}

// Invert returns inverted convertor.
func (c SpecialUnitConv) Invert() SpecialUnitConv {
	return SpecialUnitConv{
		Unit: c.Unit,
		To:   c.From,
		From: c.To,
	}
}

var SpecialUnits = map[string]SpecialUnitConv{
	"Cel": { // degree Celsius
		Unit: "K",
		To: func(v *big.Rat) *big.Rat {
			return new(big.Rat).Set(v).Add(v, kelvinZero)
		},
		From: func(v *big.Rat) *big.Rat {
			return new(big.Rat).Set(v).Sub(v, kelvinZero)
		},
	},
	"[degF]": { // degree Fahrenheit
		Unit: "K",
		To: func(v *big.Rat) *big.Rat {
			// (Fahrenheit − 32) × 5/9 + 273.15
			ret := new(big.Rat).Set(v)
			ret.Sub(ret, big.NewRat(32, 1))
			ret.Mul(ret, big.NewRat(5, 9))
			ret.Add(ret, kelvinZero)
			return ret
		},
		From: func(v *big.Rat) *big.Rat {
			ret := new(big.Rat).Set(v)
			ret.Sub(ret, kelvinZero)
			ret.Quo(ret, big.NewRat(5, 9))
			ret.Add(ret, big.NewRat(32, 1))
			return ret
		},
	},
	"[degRe]": { // degree Rankine
		Unit: "K",
		To: func(v *big.Rat) *big.Rat {
			// Kelvin = (Réaumur * 1.25) + 273.15
			ret := new(big.Rat).Set(v)
			ret.Mul(ret, big.NewRat(5, 4))
			ret.Add(ret, kelvinZero)
			return ret
		},
		From: func(v *big.Rat) *big.Rat {
			ret := new(big.Rat).Set(v)
			ret.Sub(ret, kelvinZero)
			ret.Quo(ret, big.NewRat(5, 4))
			return ret
		},
	},
	"[p'diop]": {
		Unit:    "rad",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			// rad = atan(prism_diopter / 100)
			f, _ := v.Float64()
			tan := math.Atan(f / 100)
			return new(big.Rat).SetFloat64(tan)
		},
		From: func(v *big.Rat) *big.Rat {
			// prism diopter = 100 * tan(rad)
			f, _ := v.Float64()
			tan := math.Tan(f)
			return new(big.Rat).SetFloat64(100 * tan)
		},
	},
	"%[slope]": {
		Unit:    "rad",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			tan := math.Atan(f / 100)
			return new(big.Rat).SetFloat64(tan)
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			tan := math.Tan(f)
			return new(big.Rat).SetFloat64(100 * tan)
		},
	},
	"[hp'_X]": {
		Unit:    "1",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, -f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(-1 * math.Log10(f))
		},
	},
	"[hp'_C]": {
		Unit:    "1",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(100, -f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(-1 * math.Log(f) / math.Log(100))
		},
	},
	"[hp'_M]": {
		Unit:    "1",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(1000, -f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(-1 * math.Log(f) / math.Log(1000))
		},
	},
	"[hp'_Q]": {
		Unit:    "1",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(50000, -f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(-1 * math.Log(f) / math.Log(50000))
		},
	},
	"[pH]": {
		Unit:    "mol/l",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, -f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(-1 * math.Log10(f))
		},
	},
	"Np": {
		Unit:    "1",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(math.E, f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Log(f))
		},
	},
	"B": {
		Unit:    "1",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Log10(f))
		},
	},
	"B[SPL]": {
		Unit:    "10*-5.Pa",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f/2) * 2)
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(2 * math.Log10(f/2))
		},
	},
	"B[V]": {
		Unit:    "V",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f/2))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(2 * math.Log10(f))
		},
	},
	"B[mV]": {
		Unit:    "mV",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f/2))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(2 * math.Log10(f))
		},
	},
	"B[uV]": {
		Unit:    "uV",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f/2))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(2 * math.Log10(f))
		},
	},
	"B[10.nV]": {
		Unit:    "nV",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f/2) * 10)
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(2 * math.Log10(f/10))
		},
	},
	"B[W]": {
		Unit:    "W",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Log10(f))
		},
	},
	"B[kW]": {
		Unit:    "kW",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(10, f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Log10(f))
		},
	},
	"[m/s2/Hz^(1/2)]": {
		Unit:    "m2/s4/Hz",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			return new(big.Rat).Mul(v, v)
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Sqrt(f))
		},
	},
	"bit_s": {
		Unit:    "1",
		Inexact: true,
		To: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Pow(2, f))
		},
		From: func(v *big.Rat) *big.Rat {
			f, _ := v.Float64()
			return new(big.Rat).SetFloat64(math.Log2(f))
		},
	},
}
