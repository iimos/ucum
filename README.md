# UCUM

[UCUM](https://ucum.org) (Unified Code for Units of Measure) is a code system intended to cover all units of measures.
It provides a formalism to express units in an unambiguous way suitable for electronic communication.

This Go library provides means for convertation

Usage:

```go
import "github.com/iimos/ucum"

res, err := ucum.ConvFloat64(100, "km", "m")
fmt.Println(res) // 100000

rat := big.NewRat(100, 1)
res, err := ucum.ConvBigRat(rat, "km", "m")
fmt.Println(res) // 100000
```
