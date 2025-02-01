package ucum

import (
	"fmt"
	"math/big"
)

func ExampleFloat() {
	res, err := ConvFloat64(100, "km", "m")
	if err != nil {
		// ...
	}
	fmt.Println(res) // 100000
}

func ExampleBigRat() {
	rat := big.NewRat(100, 1)
	res, err := ConvBigRat(rat, "km", "m")
	if err != nil {
		// ...
	}
	fmt.Println(res) // 100000
}
