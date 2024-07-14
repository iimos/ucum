package data

import "math/big"

func parseBigRat(s string) *big.Rat {
	rat, ok := (&big.Rat{}).SetString(s)
	if !ok {
		panic("failed to parse big.Rat: \"" + s + "\"")
	}
	return rat
}
