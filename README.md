# UCUM

[UCUM](https://ucum.org/ucum) (Unified Code for Units of Measure) is a code system intended to cover all units of measures.
It provides a formalism to express units in an unambiguous way suitable for electronic communication.

Features
--------
- **Comprehensive Support**: Handles conversions for `shopspring/decimal.Decimal`, `cockroachdb/apd.Decimal`, `big.Rat`, `big.Int`, and `float64`. It is also easy to extend for new numeric types (contributions welcome).
- **Full UCUM Coverage**: Supports all types of UCUM units, ensuring that any unit described by the UCUM standard can be converted.
- **Typed Units**: Parse UCUM expressions into `ucum.Unit` values and serialize them as JSON strings.
- **Code Generation**: Conversion tables are built using code generation for accuracy and efficiency.
- **Internal Caching**: Efficient caching mechanism for quick retrieval of previously converted units.

Installation
------------
```bash
go get github.com/iimos/ucum
go get github.com/iimos/ucum/ucumdecimal # for shopspring/decimal.Decimal
go get github.com/iimos/ucum/ucumapd # for cockroachdb/apd.Decimal
```

Quick Start
-----------
```go
package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/iimos/ucum"
	"github.com/iimos/ucum/ucumdecimal"
	"github.com/shopspring/decimal"
)

func main() {
	// float64
	meters, err := ucum.ConvFloat64(100, "km", "m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(meters) // 100000

	// Exact rational arithmetic
	minutes, err := ucum.ConvBigRat(big.NewRat(1, 3), "h", "min")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(minutes) // 20/1

	// Integer arithmetic
	feet, exact, err := ucum.ConvBigInt(big.NewInt(10000), "[ft_i]", "m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feet, exact) // 3048 true

	// shopspring/decimal with 18 fractional digits
	hours, err := ucumdecimal.ConvDecimal(decimal.NewFromInt(1), "min", "h", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hours) // 0.016666666666666667
}
```

API Reference
-------------

### 1. Converting Values

These are the main entry points for unit conversion. Each function accepts a value, a source UCUM code, and a target UCUM code. Both units are parsed and validated on every call (the parsed units are cached internally), and an error is returned for an unknown unit or incompatible dimensions.

#### `ucum.ConvFloat64`

```go
func ConvFloat64(value float64, from, to string) (float64, error)
```

Converts a `float64`. This is the simplest option when floating-point precision is sufficient.

```go
meters, err := ucum.ConvFloat64(100, "km", "m")
// meters == 100000
```

#### `ucum.ConvBigRat`

```go
func ConvBigRat(value *big.Rat, from, to string) (*big.Rat, error)
```

Converts a `big.Rat` and keeps exact rational precision. For ordinary linear units, prefixes, and temperature scales the result is exact. Special units with logarithmic or trigonometric definitions (for example `[pH]`, `B`, `Np`, `[p'diop]`) are computed through `float64`, so their results only have `float64` precision even in `big.Rat` form.

```go
res, err := ucum.ConvBigRat(big.NewRat(1, 3), "h", "min")
// res == 20/1
```

#### `ucum.ConvBigInt`

```go
func ConvBigInt(value *big.Int, from, to string) (converted *big.Int, exact bool, err error)
```

Converts a `big.Int`. The result is truncated toward zero. The second return value reports whether the conversion was exact, i.e. whether the resulting value had no fractional part.

```go
res, exact, err := ucum.ConvBigInt(big.NewInt(1000), "[ft_i]", "m")
// res == 304, exact == false (1000 ft is 304.8 m)

res, exact, err = ucum.ConvBigInt(big.NewInt(10000), "[ft_i]", "m")
// res == 3048, exact == true
```

### 2. Decimal Adapters

For decimal types the conversion goes through the exact rational core and is then rounded to the requested decimal representation. Each adapter lives in its own module so the core package does not depend on third-party decimal libraries.

#### `shopspring/decimal`: `ucumdecimal`

```go
import "github.com/iimos/ucum/ucumdecimal"

func ConvDecimal(val decimal.Decimal, from, to string, precision int32) (decimal.Decimal, error)
```

`precision` is the number of digits after the decimal point. The result is rounded half away from zero (the same rounding as `shopspring/decimal`'s `DivRound`).

```go
d := decimal.NewFromInt(1)
res, err := ucumdecimal.ConvDecimal(d, "min", "h", 18)
// res == 0.016666666666666667
```

For repeated conversions with a dedicated cache, create an adapter with `ucumdecimal.NewConverter`:

```go
conv := ucumdecimal.NewConverter(ucum.NewConverter())
res, err := conv.ConvDecimal(decimal.NewFromInt(1), "min", "h", 18)
```

`NewConverter` accepts any `ucum.ConverterBigRat`, including `ucum.DefaultConverter`. The returned `*ucumdecimal.Conv` also implements `ucumdecimal.ConverterDecimal`, which is convenient for dependency injection.

#### `cockroachdb/apd`: `ucumapd`

```go
import "github.com/iimos/ucum/ucumapd"

func ConvDecimal(val *apd.Decimal, from, to string, ctx *apd.Context) (*apd.Decimal, error)
```

The precision and rounding are controlled by the supplied `*apd.Context`:

```go
ctx := (&apd.BaseContext).WithPrecision(20)
kelvins, err := ucumapd.ConvDecimal(apd.New(300, 0), "Cel", "K", ctx)
// kelvins == 573.15
```

Always pass a context with a positive `Precision`. A `nil` context falls back to `apd.BaseContext`, which has precision 0 and makes division fail, so conversions return an error.

For repeated conversions, create an adapter with `ucumapd.NewConverter`:

```go
conv := ucumapd.NewConverter(ucum.NewConverter())
res, err := conv.ConvDecimal(apd.New(300, 0), "Cel", "K", ctx)
```

Like the decimal adapter, `NewConverter` accepts any `ucum.ConverterBigRat`. The returned `*ucumapd.Conv` implements `ucumapd.Converter`.

### 3. Parsing UCUM Units

#### `ucum.Parse`

```go
func Parse(unit []byte) (Unit, error)
```

Parses a UCUM expression into a `Unit`.

```go
unit, err := ucum.Parse([]byte("mg/(24.h)"))
if err != nil {
	// invalid UCUM expression
}
fmt.Println(unit.String()) // mg/(24.h)
```

#### `ucum.MustParse`

```go
func MustParse(unit []byte) Unit
```

Works like `Parse` but panics on invalid input. Use it for package-level constants and other inputs that are known to be valid.

```go
unit := ucum.MustParse([]byte("km/h"))
```

#### `ucum.Unit`

`Unit` is a parsed UCUM unit. Its methods are:

```go
func (u *Unit) String() string
```

Returns the original UCUM string used at parse time:

```go
u := ucum.MustParse([]byte("mg/(24.h)"))
fmt.Println(u.String()) // mg/(24.h)
```

```go
func (u *Unit) MarshalJSON() ([]byte, error)
func (u *Unit) UnmarshalJSON(data []byte) error
```

The unit is encoded as a JSON string and decoded from a JSON string:

```go
u := ucum.MustParse([]byte("km/h"))
data, err := json.Marshal(&u) // "km/h"

var parsed ucum.Unit
err = json.Unmarshal([]byte(`"mg/L"`), &parsed)
fmt.Println(parsed.String()) // mg/L
```

`UnmarshalJSON` rejects non-string values and invalid UCUM units. A JSON `null` leaves the unit unchanged.

Note that `MarshalJSON` is defined on `*Unit`. When a `Unit` is a value field of a struct, marshal the struct by pointer so the method is used:

```go
type Measurement struct {
	Value float64   `json:"value"`
	Unit  ucum.Unit `json:"unit"`
}

data, err := json.Marshal(&m)
```

#### Exponent limits

```go
const (
	MaxExponent = 100
	MinExponent = -100
)
```

These constants define the allowed exponent range in parsed units. The parser rejects exponents outside this range with an `exponent too large` error.

### 4. Reusable Converter and Cache

```go
func NewConverter() *Conv
var DefaultConverter = NewConverter()
```

`Conv` is a thread-safe converter with an internal cache of parsed units. It provides the same three methods as the package-level functions:

```go
func (c *Conv) ConvFloat64(val float64, from, to string) (float64, error)
func (c *Conv) ConvBigRat(val *big.Rat, from, to string) (*big.Rat, error)
func (c *Conv) ConvBigInt(val *big.Int, from, to string) (converted *big.Int, exact bool, err error)
```

Use a dedicated converter when you want an isolated cache, or when you want to share one cache between the core package and the decimal adapters:

```go
conv := ucum.NewConverter()

meters, err := conv.ConvFloat64(1, "mi", "km")

decConv := ucumdecimal.NewConverter(conv)
apdConv := ucumapd.NewConverter(conv)
```

The package-level functions (`ucum.ConvFloat64`, `ucum.ConvBigRat`, `ucum.ConvBigInt`) delegate to `DefaultConverter`.

The cache stores parsed units forever and never evicts entries. This is intentional: a typical application uses a small, bounded set of distinct unit strings. `Conv` is safe for concurrent use.

### 5. Precomputed Pair Conversion

```go
func NewPairConverter(from, to Unit) (PairConverter, error)
```

`NewPairConverter` validates a source/target pair once and returns a `PairConverter` for that pair:

```go
type PairConverter interface {
	ConvBigRat(val *big.Rat) *big.Rat
	ConvBigInt(val *big.Int) (converted *big.Int, exact bool)
	ConvFloat64(val float64) float64
}
```

Its methods do not return errors because the pair has already been validated. Use it when converting many values between the same two units, especially in hot loops, to avoid repeated unit lookups.

```go
conv, err := ucum.NewPairConverter(
	ucum.MustParse([]byte("km/h")),
	ucum.MustParse([]byte("m/s")),
)
if err != nil {
	// incompatible units
}

speed := conv.ConvFloat64(36) // 10
```

The integer and rational methods follow the same semantics as the corresponding package-level functions.

### 6. Custom Numeric Types

The core package exposes interfaces for each numeric family:

```go
type ConverterBigRat interface {
	ConvBigRat(val *big.Rat, from, to string) (*big.Rat, error)
}

type ConverterBigInt interface {
	ConvBigInt(val *big.Int, from, to string) (converted *big.Int, exact bool, err error)
}

type ConverterFloat64 interface {
	ConvFloat64(val float64, from, to string) (float64, error)
}
```

`*ucum.Conv` implements all three. To support a new decimal-like type without adding it to this repository, convert it to `*big.Rat`, call `ConvBigRat`, and convert the exact result back:

```go
type MyDecimal struct {
	rat *big.Rat
}

func (d MyDecimal) Rat() *big.Rat { return d.rat }

func newMyDecimal(r *big.Rat) MyDecimal { return MyDecimal{rat: r} }

func ConvertMyDecimal(c ucum.ConverterBigRat, v MyDecimal, from, to string) (MyDecimal, error) {
	r, err := c.ConvBigRat(v.Rat(), from, to)
	if err != nil {
		return MyDecimal{}, err
	}
	return newMyDecimal(r), nil
}
```

This is also the pattern used by the `ucumdecimal` and `ucumapd` adapters. Both adapters accept any `ucum.ConverterBigRat` in their `NewConverter` constructors.

### 7. Conversion Rules and Caveats

- **Case-sensitive codes.** UCUM codes are case-sensitive: `mm` is millimeter, while `Mm` is megameter.
- **Errors instead of panics.** Conversion functions return an error for unknown units and for units with incompatible dimensions. Only `MustParse` panics.
- **Integer conversion loses fractions.** `ConvBigInt` truncates toward zero and reports `exact == false` when the mathematical result has a fractional part. Prefer `ConvBigRat`, `ConvFloat64`, or a decimal adapter when fractions matter.
- **Special units are supported.** Temperature (`Cel`, `[degF]`), logarithmic (`[pH]`, `Np`, `B`), trigonometric (`[p'diop]`, `%[slope]`), and other UCUM special units are converted with their non-linear definitions. For logarithmic and trigonometric units the implementation uses `float64` math internally.
- **Arbitrary units are isolated.** Procedure-defined units such as `[IU]` and `[hp_X]` cannot be converted to unrelated units. Conversion between the same arbitrary unit is allowed, including different metric prefixes.
- **Annotations are preserved.** Annotations in expressions such as `m{annot}` are kept in `Unit.String()` and JSON round-trips.
- **Concurrency.** `ucum.Conv` and `ucum.DefaultConverter` are safe for concurrent use.

License
-------
This project is licensed under the [Apache 2.0 License](./LICENSE).
