# UCUM

[UCUM](https://ucum.org) (Unified Code for Units of Measure) is a code system intended to cover all units of measures.
It provides a formalism to express units in an unambiguous way suitable for electronic communication.

Features
--------
- **Comprehensive Support**: Handles conversions for `shopspring/decimal.Decimal`, `big.Rat`, `big.Int`, `float` types. Also it's easyly extensible for a new types (contributions welcome).
- **Full UCUM Coverage**: Supports all types of UCUM units, ensuring that any unit described by the UCUM standard can be converted.
- **Code Generation**: Conversion tables are built using code generation for accuracy and efficiency.
- **Internal Caching**: Efficient caching mechanism for quick retrieval of previously converted units.

Installation
------------
```bash
go get github.com/iimos/ucum
go get github.com/iimos/ucum/ucumdecimal
```

Basic Usage
-----
```go
import "github.com/iimos/ucum"

res, err := ucum.ConvFloat64(100, "km", "m")

rat := big.NewRat(100, 30)
res, err := ucum.ConvBigRat(rat, "km", "m")

bigInt := big.NewInt(100)
converted, exact, err := ucum.ConvBigInt(bigInt, "km", "m")

// ---------------------------

import "github.com/iimos/ucum/ucumdecimal"

d := decimal.NewFromFloat(100)
res, err := ucumdecimal.ConvDecimal(d, "km", "m", 10)
```

License
-------
This project is licensed under the Apache 2.0 License - see the [LICENSE](./LICENSE) file for details.

Contributing
------------
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
