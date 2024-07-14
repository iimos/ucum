// Code generated; DO NOT EDIT.
package data

import "math/big"
import "github.com/iimos/ucum/internal/types"

var Conv = map[string]*types.Unit{
	// % = 1/100
	"%": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x64})), Components: map[types.ComponentKey]int{}},
	// ' = 10471975511965977461542144610931676280657231331250352736583148641/36000000000000000000000000000000000000000000000000000000000000000000⋅rad
	"'": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x19, 0x74, 0xb9, 0xf2, 0xbb, 0x34, 0x51, 0x5d, 0x36, 0x61, 0xcd, 0x1e, 0xbf, 0x4, 0xd3, 0x59, 0x37, 0x7a, 0x93, 0x61, 0xf1, 0x7e, 0xd9, 0xe, 0x41, 0xe4, 0x61}), (&big.Int{}).SetBytes([]byte{0x1, 0x55, 0xd7, 0x27, 0x0, 0x13, 0xab, 0x3, 0x42, 0xd4, 0xc8, 0x7a, 0xde, 0x0, 0x8a, 0x58, 0xc3, 0xc4, 0x53, 0xfe, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "rad", Annotation: ""}: 1}},
	// ” = 10471975511965977461542144610931676280657231331250352736583148641/2160000000000000000000000000000000000000000000000000000000000000000000⋅rad
	"''": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x19, 0x74, 0xb9, 0xf2, 0xbb, 0x34, 0x51, 0x5d, 0x36, 0x61, 0xcd, 0x1e, 0xbf, 0x4, 0xd3, 0x59, 0x37, 0x7a, 0x93, 0x61, 0xf1, 0x7e, 0xd9, 0xe, 0x41, 0xe4, 0x61}), (&big.Int{}).SetBytes([]byte{0x50, 0x1e, 0x6d, 0x24, 0x4, 0x9c, 0x14, 0xc3, 0xa9, 0xde, 0xfc, 0xcc, 0x8, 0x20, 0x6c, 0xcd, 0xe2, 0x3, 0xaf, 0x8b, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "rad", Annotation: ""}: 1}},
	// 10* = 10
	"10*": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xa}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// 10^ = 10
	"10^": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xa}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// A = C⋅s⁻¹
	"A": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// AU = 149597870691⋅m
	"AU": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x22, 0xd4, 0xba, 0x5a, 0x63}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// Ao = 1/10000000000⋅m
	"Ao": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x2, 0x54, 0xb, 0xe4, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// Bd = s⁻¹
	"Bd": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// Bi = 10⋅C⋅s⁻¹
	"Bi": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xa}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// Bq = s⁻¹
	"Bq": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// By = 8⋅bit
	"By": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "bit", Annotation: ""}: 1}},
	// Ci = 37000000000⋅s⁻¹
	"Ci": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x8, 0x9d, 0x5f, 0x32, 0x0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// F = 1/1000⋅C²⋅s²⋅g⁻¹⋅m⁻²
	"F": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 2, types.ComponentKey{AtomCode: "g", Annotation: ""}: -1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -2, types.ComponentKey{AtomCode: "s", Annotation: ""}: 2}},
	// G = 1/10⋅g⋅m⁰⋅C⁻¹⋅s⁻¹
	"G": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0xa})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -1, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 0, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// Gal = 1/100⋅m⋅s⁻²
	"Gal": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x64})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// Gb = 25000000000000000000000000000000000000000000000000000000000000000/31415926535897932384626433832795028841971693993751058209749445923⋅C⋅m⁰⋅s⁻¹
	"Gb": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3c, 0xc5, 0x89, 0xc7, 0x1f, 0xf0, 0xe4, 0x22, 0xa2, 0xfb, 0xd1, 0x93, 0x8e, 0x51, 0x7b, 0xde, 0x89, 0x4d, 0x82, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 0, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// Gy = m²⋅g⁰⋅s⁻²
	"Gy": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 0, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// H = 1000⋅m²⋅g⋅s⁰⋅C⁻²
	"H": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -2, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: 0}},
	// Hz = s⁻¹
	"Hz": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// J = 1000⋅m²⋅g⋅s⁻²
	"J": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// Ky = 100⋅m⁻¹
	"Ky": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x64}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: -1}},
	// L = 1/1000⋅m³
	"L": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// Lmb = 100000000000000000000000000000000000000000000000000000000000000000000/31415926535897932384626433832795028841971693993751058209749445923⋅cd⋅m⁻²
	"Lmb": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xb5, 0x8e, 0x88, 0xc7, 0x53, 0x13, 0xec, 0x9d, 0x32, 0x9e, 0xaa, 0xa1, 0x8f, 0xb9, 0x2f, 0x75, 0x21, 0x5b, 0x17, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "cd", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -2}},
	// Mx = 1/100000⋅m²⋅g⋅C⁻¹⋅s⁻¹
	"Mx": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1, 0x86, 0xa0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -1, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// N = 1000⋅g⋅m⋅s⁻²
	"N": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// Oe = 2500000000000000000000000000000000000000000000000000000000000000000/31415926535897932384626433832795028841971693993751058209749445923⋅C⋅m⁻¹⋅s⁻¹
	"Oe": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x17, 0xbd, 0x29, 0xd1, 0xc8, 0x7a, 0x19, 0x1d, 0x87, 0xaa, 0x5d, 0xdd, 0xa3, 0x97, 0xd4, 0x62, 0xed, 0xa2, 0x46, 0xfa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// Ohm = 1000⋅m²⋅g⋅s⁻¹⋅C⁻²
	"Ohm": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -2, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// P = 100⋅g⋅m⁻¹⋅s⁻¹
	"P": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x64}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// Pa = 1000⋅g⋅m⁻¹⋅s⁻²
	"Pa": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// R = 129/500000000⋅C⋅g⁻¹
	"R": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x81}), (&big.Int{}).SetBytes([]byte{0x1d, 0xcd, 0x65, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 1, types.ComponentKey{AtomCode: "g", Annotation: ""}: -1}},
	// RAD = 1/100⋅m²⋅g⁰⋅s⁻²
	"RAD": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x64})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 0, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// REM = 1/100⋅m²⋅g⁰⋅s⁻²
	"REM": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x64})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 0, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// S = 1/1000⋅C²⋅s⋅g⁻¹⋅m⁻²
	"S": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 2, types.ComponentKey{AtomCode: "g", Annotation: ""}: -1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -2, types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// St = 1/10000⋅m²⋅s⁻¹
	"St": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x27, 0x10})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// Sv = m²⋅g⁰⋅s⁻²
	"Sv": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 0, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// T = 1000⋅g⋅m⁰⋅C⁻¹⋅s⁻¹
	"T": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -1, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 0, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// U = 30110703800000000/3⋅s⁻¹
	"U": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6a, 0xf9, 0x86, 0x8b, 0xef, 0x9e, 0x0}), (&big.Int{}).SetBytes([]byte{0x3})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// V = 1000⋅m²⋅g⋅C⁻¹⋅s⁻²
	"V": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -1, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// W = 1000⋅m²⋅g⋅s⁻³
	"W": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -3}},
	// Wb = 1000⋅m²⋅g⋅C⁻¹⋅s⁻¹
	"Wb": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0xe8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -1, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [APL'U] = 1
	"[APL'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [AU] = 1
	"[AU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [Amb'a'1'U] = 1
	"[Amb'a'1'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [BAU] = 1
	"[BAU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [Btu] = 1054350⋅m²⋅g⋅s⁻²
	"[Btu]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x16, 0x8e}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [Btu_39] = 1059670⋅m²⋅g⋅s⁻²
	"[Btu_39]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x2b, 0x56}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [Btu_59] = 1054800⋅m²⋅g⋅s⁻²
	"[Btu_59]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x18, 0x50}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [Btu_60] = 1054680⋅m²⋅g⋅s⁻²
	"[Btu_60]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x17, 0xd8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [Btu_IT] = 52752792631/50000⋅m²⋅g⋅s⁻²
	"[Btu_IT]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xc, 0x48, 0x4f, 0xbc, 0x37}), (&big.Int{}).SetBytes([]byte{0xc3, 0x50})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [Btu_m] = 1055870⋅m²⋅g⋅s⁻²
	"[Btu_m]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x1c, 0x7e}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [Btu_th] = 1054350⋅m²⋅g⋅s⁻²
	"[Btu_th]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x16, 0x8e}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [CCID_50] = 1
	"[CCID_50]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [CFU] = 1
	"[CFU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [Cal] = 4184000⋅m²⋅g⋅s⁻²
	"[Cal]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3f, 0xd7, 0xc0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [Ch] = 1/3000⋅m
	"[Ch]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0xb, 0xb8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [D'ag'U] = 1
	"[D'ag'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [EID_50] = 1
	"[EID_50]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [ELU] = 1
	"[ELU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [EU] = 1
	"[EU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [FEU] = 1
	"[FEU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [FFU] = 1
	"[FFU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [FNU] = 1
	"[FNU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [GPL'U] = 1
	"[GPL'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [G] = 66743/1000000000000000000⋅m³⋅g⁻¹⋅s⁻²
	"[G]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x4, 0xb7}), (&big.Int{}).SetBytes([]byte{0xd, 0xe0, 0xb6, 0xb3, 0xa7, 0x64, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: -1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 3, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [HPF] = 1
	"[HPF]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [HP] = 37284993579113511/50000000000⋅m²⋅g⋅s⁻³
	"[HP]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x84, 0x76, 0x81, 0xd, 0xbc, 0x5c, 0x27}), (&big.Int{}).SetBytes([]byte{0xb, 0xa4, 0x3b, 0x74, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -3}},
	// [IR] = 1
	"[IR]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [IU] = [iU]
	"[IU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "[iU]", Annotation: ""}: 1}},
	// [LPF] = 100
	"[LPF]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x64}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [Lf] = 1
	"[Lf]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [MET] = 7/120000000000⋅m³⋅g⁻¹⋅s⁻¹
	"[MET]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7}), (&big.Int{}).SetBytes([]byte{0x1b, 0xf0, 0x8e, 0xb0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: -1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 3, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [MPL'U] = 1
	"[MPL'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [NTU] = 1
	"[NTU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [PFU] = 1
	"[PFU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [PNU] = 1
	"[PNU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [PRU] = 133322000000⋅g⋅s⁻¹⋅m⁻⁴
	"[PRU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1f, 0xa, 0x9c, 0x46, 0x80}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -4, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [S] = 1/10000000000000⋅s
	"[S]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x9, 0x18, 0x4e, 0x72, 0xa0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// [TCID_50] = 1
	"[TCID_50]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [USP'U] = 1
	"[USP'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [acr_br] = 15808008005469801/3906250000000⋅m²
	"[acr_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x38, 0x29, 0x4c, 0xad, 0xc5, 0xc6, 0x69}), (&big.Int{}).SetBytes([]byte{0x3, 0x8d, 0x7e, 0xa4, 0xc6, 0x80})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [acr_us] = 62726400000/15499969⋅m²
	"[acr_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xe, 0x9a, 0xc8, 0xe8, 0x0}), (&big.Int{}).SetBytes([]byte{0xec, 0x82, 0xc1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [anti'Xa'U] = 1
	"[anti'Xa'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [arb'U] = 1
	"[arb'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [bbl_us] = 9936705933/62500000000⋅m³
	"[bbl_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0x50, 0x46, 0x19, 0x8d}), (&big.Int{}).SetBytes([]byte{0xe, 0x8d, 0x4a, 0x51, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [bdsk'U] = 1
	"[bdsk'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [beth'U] = 1
	"[beth'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [bf_i] = 18435447/7812500000⋅m³
	"[bf_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x19, 0x4d, 0x77}), (&big.Int{}).SetBytes([]byte{0x1, 0xd1, 0xa9, 0x4a, 0x20})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [bu_br] = 454609/12500000⋅m³
	"[bu_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0xbe, 0xbc, 0x20})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [bu_us] = 220244188543/6250000000000⋅m³
	"[bu_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x33, 0x47, 0x93, 0x9d, 0x7f}), (&big.Int{}).SetBytes([]byte{0x5, 0xaf, 0x31, 0x7, 0xa4, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [c] = 299792458⋅m⋅s⁻¹
	"[c]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x11, 0xde, 0x78, 0x4a}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [car_Au] = 1/24
	"[car_Au]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x18})), Components: map[types.ComponentKey]int{}},
	// [car_m] = 1/5⋅g
	"[car_m]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x5})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [cft_i] = 55306341/1953125000⋅m³
	"[cft_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0x4b, 0xe8, 0x65}), (&big.Int{}).SetBytes([]byte{0x74, 0x6a, 0x52, 0x88})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [ch_br] = 125729901/6250000⋅m
	"[ch_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0x7e, 0x7c, 0x6d}), (&big.Int{}).SetBytes([]byte{0x5f, 0x5e, 0x10})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [ch_us] = 79200/3937⋅m
	"[ch_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x35, 0x60}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [cicero] = 203/45000⋅m
	"[cicero]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xcb}), (&big.Int{}).SetBytes([]byte{0xaf, 0xc8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [cin_i] = 2048383/125000000000⋅m³
	"[cin_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1f, 0x41, 0x7f}), (&big.Int{}).SetBytes([]byte{0x1d, 0x1a, 0x94, 0xa2, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [cml_i] = 506707479097497751431639751289151020192161452425210817865048813292067/1000000000000000000000000000000000000000000000000000000000000000000000000000000⋅m²
	"[cml_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x12, 0xcb, 0x79, 0xa6, 0x62, 0xd, 0xd1, 0xae, 0xc5, 0x3d, 0xd7, 0x95, 0x1a, 0x62, 0xcd, 0x17, 0x58, 0x1, 0x25, 0x3f, 0x9, 0x70, 0x7a, 0xbc, 0x11, 0xd0, 0x61, 0x4a, 0x23}), (&big.Int{}).SetBytes([]byte{0x8, 0xa2, 0xdb, 0xf1, 0x42, 0xdf, 0xcc, 0x7a, 0xb6, 0xe3, 0x56, 0x93, 0x26, 0xc7, 0x84, 0x33, 0x72, 0xa9, 0xf4, 0xd2, 0x50, 0x5e, 0x3a, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [cr_i] = 884901456/244140625⋅m³
	"[cr_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x34, 0xbe, 0x86, 0x50}), (&big.Int{}).SetBytes([]byte{0xe, 0x8d, 0x4a, 0x51})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [crd_us] = 884901456/244140625⋅m³
	"[crd_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x34, 0xbe, 0x86, 0x50}), (&big.Int{}).SetBytes([]byte{0xe, 0x8d, 0x4a, 0x51})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [cup_m] = 3/12500⋅m³
	"[cup_m]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3}), (&big.Int{}).SetBytes([]byte{0x30, 0xd4})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [cup_us] = 473176473/2000000000000⋅m³
	"[cup_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0x1, 0xd1, 0xa9, 0x4a, 0x20, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [cyd_i] = 1493271207/1953125000⋅m³
	"[cyd_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x59, 0x1, 0x82, 0xa7}), (&big.Int{}).SetBytes([]byte{0x74, 0x6a, 0x52, 0x88})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [degR] = 5/9⋅K
	"[degR]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x5}), (&big.Int{}).SetBytes([]byte{0x9})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "K", Annotation: ""}: 1}},
	// [den] = 1/9000⋅g⋅m⁻¹
	"[den]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x23, 0x28})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1}},
	// [didot] = 203/540000⋅m
	"[didot]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xcb}), (&big.Int{}).SetBytes([]byte{0x8, 0x3d, 0x60})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [diop] = m⁻¹
	"[diop]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: -1}},
	// [dpt_us] = 220244188543/400000000000000⋅m³
	"[dpt_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x33, 0x47, 0x93, 0x9d, 0x7f}), (&big.Int{}).SetBytes([]byte{0x1, 0x6b, 0xcc, 0x41, 0xe9, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [dqt_us] = 220244188543/200000000000000⋅m³
	"[dqt_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x33, 0x47, 0x93, 0x9d, 0x7f}), (&big.Int{}).SetBytes([]byte{0xb5, 0xe6, 0x20, 0xf4, 0x80, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [dr_ap] = 19439673/5000000⋅g
	"[dr_ap]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x28, 0xa0, 0x39}), (&big.Int{}).SetBytes([]byte{0x4c, 0x4b, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [dr_av] = 45359237/25600000⋅g
	"[dr_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xb4, 0x20, 0x85}), (&big.Int{}).SetBytes([]byte{0x1, 0x86, 0xa0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [drp] = 1/20000000⋅m³
	"[drp]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1, 0x31, 0x2d, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [dye'U] = 1
	"[dye'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [e] = 801088317/5000000000000000000000000000⋅C
	"[e]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2f, 0xbf, 0xa3, 0x3d}), (&big.Int{}).SetBytes([]byte{0x10, 0x27, 0xe7, 0x2f, 0x1f, 0x12, 0x81, 0x30, 0x88, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 1}},
	// [eps_0] = 8854187817/1000000000000000000000000⋅C²⋅s²⋅g⁻¹⋅m⁻³
	"[eps_0]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xf, 0xc0, 0x2f, 0x29}), (&big.Int{}).SetBytes([]byte{0xd3, 0xc2, 0x1b, 0xce, 0xcc, 0xed, 0xa1, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 2, types.ComponentKey{AtomCode: "g", Annotation: ""}: -1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -3, types.ComponentKey{AtomCode: "s", Annotation: ""}: 2}},
	// [fdr_br] = 454609/128000000000⋅m³
	"[fdr_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0x1d, 0xcd, 0x65, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [fdr_us] = 473176473/128000000000000⋅m³
	"[fdr_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0x74, 0x6a, 0x52, 0x88, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [foz_br] = 454609/16000000000⋅m³
	"[foz_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0x3, 0xb9, 0xac, 0xa0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [foz_m] = 3/100000⋅m³
	"[foz_m]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3}), (&big.Int{}).SetBytes([]byte{0x1, 0x86, 0xa0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [foz_us] = 473176473/16000000000000⋅m³
	"[foz_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0xe, 0x8d, 0x4a, 0x51, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [ft_br] = 3809997/12500000⋅m
	"[ft_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3a, 0x22, 0xcd}), (&big.Int{}).SetBytes([]byte{0xbe, 0xbc, 0x20})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [ft_i] = 381/1250⋅m
	"[ft_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x7d}), (&big.Int{}).SetBytes([]byte{0x4, 0xe2})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [ft_us] = 1200/3937⋅m
	"[ft_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4, 0xb0}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [fth_br] = 11429991/6250000⋅m
	"[fth_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xae, 0x68, 0x67}), (&big.Int{}).SetBytes([]byte{0x5f, 0x5e, 0x10})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [fth_i] = 1143/625⋅m
	"[fth_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4, 0x77}), (&big.Int{}).SetBytes([]byte{0x2, 0x71})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [fth_us] = 7200/3937⋅m
	"[fth_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x20}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [fur_us] = 792000/3937⋅m
	"[fur_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xc, 0x15, 0xc0}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [g] = 196133/20000⋅m⋅s⁻²
	"[g]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xfe, 0x25}), (&big.Int{}).SetBytes([]byte{0x4e, 0x20})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [gal_br] = 454609/100000000⋅m³
	"[gal_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0x5, 0xf5, 0xe1, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [gal_us] = 473176473/125000000000⋅m³
	"[gal_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0x1d, 0x1a, 0x94, 0xa2, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [gal_wi] = 220244188543/50000000000000⋅m³
	"[gal_wi]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x33, 0x47, 0x93, 0x9d, 0x7f}), (&big.Int{}).SetBytes([]byte{0x2d, 0x79, 0x88, 0x3d, 0x20, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [gil_br] = 454609/3200000000⋅m³
	"[gil_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0xbe, 0xbc, 0x20, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [gil_us] = 473176473/4000000000000⋅m³
	"[gil_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0x3, 0xa3, 0x52, 0x94, 0x40, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [gr] = 6479891/100000000⋅g
	"[gr]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x62, 0xe0, 0x13}), (&big.Int{}).SetBytes([]byte{0x5, 0xf5, 0xe1, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [h] = 132521403/200000000000000000000000000000000000000⋅m²⋅g⋅s⁻¹
	"[h]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0xe6, 0x1d, 0xbb}), (&big.Int{}).SetBytes([]byte{0x96, 0x76, 0x99, 0x50, 0xb5, 0xd, 0x88, 0xf4, 0x13, 0x14, 0x44, 0x80, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [hd_i] = 127/1250⋅m
	"[hd_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f}), (&big.Int{}).SetBytes([]byte{0x4, 0xe2})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [hnsf'U] = 1
	"[hnsf'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [hp_C] = 1
	"[hp_C]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [hp_M] = 1
	"[hp_M]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [hp_Q] = 1
	"[hp_Q]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [hp_X] = 1
	"[hp_X]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [iU] = 1
	"[iU]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [in_br] = 1269999/50000000⋅m
	"[in_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x13, 0x60, 0xef}), (&big.Int{}).SetBytes([]byte{0x2, 0xfa, 0xf0, 0x80})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [in_i'H2O] = 24908891/100⋅g⋅m⁻¹⋅s⁻²
	"[in_i'H2O]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x7c, 0x14, 0x5b}), (&big.Int{}).SetBytes([]byte{0x64})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [in_i'Hg] = 16931894/5⋅g⋅m⁻¹⋅s⁻²
	"[in_i'Hg]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x2, 0x5c, 0x36}), (&big.Int{}).SetBytes([]byte{0x5})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [in_i] = 127/5000⋅m
	"[in_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f}), (&big.Int{}).SetBytes([]byte{0x13, 0x88})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [in_us] = 100/3937⋅m
	"[in_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x64}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [k] = 1380649/100000000000000000000000000⋅m²⋅g⋅K⁻¹⋅s⁻²
	"[k]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x15, 0x11, 0x29}), (&big.Int{}).SetBytes([]byte{0x52, 0xb7, 0xd2, 0xdc, 0xc8, 0xc, 0xd2, 0xe4, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "K", Annotation: ""}: -1, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [ka'U] = 1
	"[ka'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [kn_br] = 8043327/15625000⋅m⋅s⁻¹
	"[kn_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7a, 0xbb, 0x3f}), (&big.Int{}).SetBytes([]byte{0xee, 0x6b, 0x28})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [kn_i] = 463/900⋅m⋅s⁻¹
	"[kn_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0xcf}), (&big.Int{}).SetBytes([]byte{0x3, 0x84})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [knk'U] = 1
	"[knk'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [kp_C] = 1
	"[kp_C]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [kp_M] = 1
	"[kp_M]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [kp_Q] = 1
	"[kp_Q]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [kp_X] = 1
	"[kp_X]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [lb_ap] = 58319019/156250⋅g
	"[lb_ap]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0x79, 0xe0, 0xab}), (&big.Int{}).SetBytes([]byte{0x2, 0x62, 0x5a})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [lb_av] = 45359237/100000⋅g
	"[lb_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xb4, 0x20, 0x85}), (&big.Int{}).SetBytes([]byte{0x1, 0x86, 0xa0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [lb_tr] = 58319019/156250⋅g
	"[lb_tr]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0x79, 0xe0, 0xab}), (&big.Int{}).SetBytes([]byte{0x2, 0x62, 0x5a})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [lbf_av] = 8896443230521/2000000000⋅g⋅m⋅s⁻²
	"[lbf_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x8, 0x17, 0x5d, 0x56, 0xa9, 0x39}), (&big.Int{}).SetBytes([]byte{0x77, 0x35, 0x94, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [lcwt_av] = 317514659/6250⋅g
	"[lcwt_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x12, 0xec, 0xe3, 0xa3}), (&big.Int{}).SetBytes([]byte{0x18, 0x6a})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [ligne] = 203/90000⋅m
	"[ligne]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xcb}), (&big.Int{}).SetBytes([]byte{0x1, 0x5f, 0x90})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [lk_br] = 125729901/625000000⋅m
	"[lk_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0x7e, 0x7c, 0x6d}), (&big.Int{}).SetBytes([]byte{0x25, 0x40, 0xbe, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [lk_us] = 792/3937⋅m
	"[lk_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0x18}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [lne] = 127/60000⋅m
	"[lne]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f}), (&big.Int{}).SetBytes([]byte{0xea, 0x60})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [lton_av] = 635029318/625⋅g
	"[lton_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x25, 0xd9, 0xc7, 0x46}), (&big.Int{}).SetBytes([]byte{0x2, 0x71})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [ly] = 9460730472580800⋅m⋅s⁰
	"[ly]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x21, 0x9c, 0x7b, 0xf7, 0x22, 0x46, 0xc0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: 0}},
	// [m_e] = 91093837139/100000000000000000000000000000000000000⋅g
	"[m_e]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x15, 0x35, 0x9d, 0xa5, 0x53}), (&big.Int{}).SetBytes([]byte{0x4b, 0x3b, 0x4c, 0xa8, 0x5a, 0x86, 0xc4, 0x7a, 0x9, 0x8a, 0x22, 0x40, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [m_p] = 33452438519/20000000000000000000000000000000000⋅g
	"[m_p]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0xc9, 0xeb, 0xb3, 0xf7}), (&big.Int{}).SetBytes([]byte{0x3, 0xda, 0x13, 0x7d, 0x5b, 0xf, 0x80, 0x6f, 0x1b, 0x1c, 0xc8, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [mclg'U] = 1
	"[mclg'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [mesh_i] = 5000/127⋅m⁻¹
	"[mesh_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x13, 0x88}), (&big.Int{}).SetBytes([]byte{0x7f})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: -1}},
	// [mi_br] = 125729901/78125⋅m
	"[mi_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0x7e, 0x7c, 0x6d}), (&big.Int{}).SetBytes([]byte{0x1, 0x31, 0x2d})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [mi_i] = 201168/125⋅m
	"[mi_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0x11, 0xd0}), (&big.Int{}).SetBytes([]byte{0x7d})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [mi_us] = 6336000/3937⋅m
	"[mi_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x60, 0xae, 0x0}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [mil_i] = 127/5000000⋅m
	"[mil_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f}), (&big.Int{}).SetBytes([]byte{0x4c, 0x4b, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [mil_us] = 1/39370⋅m
	"[mil_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x99, 0xca})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [min_br] = 454609/7680000000000⋅m³
	"[min_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0x6, 0xfc, 0x23, 0xac, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [min_us] = 157725491/2560000000000000⋅m³
	"[min_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x9, 0x66, 0xb3, 0x33}), (&big.Int{}).SetBytes([]byte{0x9, 0x18, 0x4e, 0x72, 0xa0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [mu_0] = 31415926535897932384626433832795028841971693993751058209749445923/25000000000000000000000000000000000000000000000000000000000000000000⋅g⋅m⋅s⁰⋅C⁻²
	"[mu_0]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23}), (&big.Int{}).SetBytes([]byte{0xed, 0x63, 0xa2, 0x31, 0xd4, 0xc4, 0xfb, 0x27, 0x4c, 0xa7, 0xaa, 0xa8, 0x63, 0xee, 0x4b, 0xdd, 0x48, 0x56, 0xc5, 0xc4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: -2, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: 0}},
	// [nmi_br] = 144779886/78125⋅m
	"[nmi_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x8, 0xa1, 0x2a, 0x6e}), (&big.Int{}).SetBytes([]byte{0x1, 0x31, 0x2d})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [nmi_i] = 1852⋅m
	"[nmi_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0x3c}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [oz_ap] = 19439673/625000⋅g
	"[oz_ap]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x28, 0xa0, 0x39}), (&big.Int{}).SetBytes([]byte{0x9, 0x89, 0x68})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [oz_av] = 45359237/1600000⋅g
	"[oz_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xb4, 0x20, 0x85}), (&big.Int{}).SetBytes([]byte{0x18, 0x6a, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [oz_m] = 28⋅g
	"[oz_m]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [oz_tr] = 19439673/625000⋅g
	"[oz_tr]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x28, 0xa0, 0x39}), (&big.Int{}).SetBytes([]byte{0x9, 0x89, 0x68})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [pc_br] = 3809997/5000000⋅m
	"[pc_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3a, 0x22, 0xcd}), (&big.Int{}).SetBytes([]byte{0x4c, 0x4b, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [pca] = 127/30000⋅m
	"[pca]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f}), (&big.Int{}).SetBytes([]byte{0x75, 0x30})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [pca_pr] = 5271897/1250000000⋅m
	"[pca_pr]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x50, 0x71, 0x59}), (&big.Int{}).SetBytes([]byte{0x4a, 0x81, 0x7c, 0x80})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [pi] = 31415926535897932384626433832795028841971693993751058209749445923/10000000000000000000000000000000000000000000000000000000000000000
	"[pi]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23}), (&big.Int{}).SetBytes([]byte{0x18, 0x4f, 0x3, 0xe9, 0x3f, 0xf9, 0xf4, 0xda, 0xa7, 0x97, 0xed, 0x6e, 0x38, 0xed, 0x64, 0xbf, 0x6a, 0x1f, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{}},
	// [pied] = 203/625⋅m
	"[pied]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xcb}), (&big.Int{}).SetBytes([]byte{0x2, 0x71})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [pk_br] = 454609/50000000⋅m³
	"[pk_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0x2, 0xfa, 0xf0, 0x80})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [pk_us] = 220244188543/25000000000000⋅m³
	"[pk_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x33, 0x47, 0x93, 0x9d, 0x7f}), (&big.Int{}).SetBytes([]byte{0x16, 0xbc, 0xc4, 0x1e, 0x90, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [pnt] = 127/360000⋅m
	"[pnt]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f}), (&big.Int{}).SetBytes([]byte{0x5, 0x7e, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [pnt_pr] = 1757299/5000000000⋅m
	"[pnt_pr]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1a, 0xd0, 0x73}), (&big.Int{}).SetBytes([]byte{0x1, 0x2a, 0x5, 0xf2, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [pouce] = 203/7500⋅m
	"[pouce]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xcb}), (&big.Int{}).SetBytes([]byte{0x1d, 0x4c})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [ppb] = 1/1000000000
	"[ppb]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3b, 0x9a, 0xca, 0x0})), Components: map[types.ComponentKey]int{}},
	// [ppm] = 1/1000000
	"[ppm]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0xf, 0x42, 0x40})), Components: map[types.ComponentKey]int{}},
	// [ppth] = 1/1000
	"[ppth]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{}},
	// [pptr] = 1/1000000000000
	"[pptr]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0xe8, 0xd4, 0xa5, 0x10, 0x0})), Components: map[types.ComponentKey]int{}},
	// [psi] = 8896443230521/1290320⋅g⋅m⁻¹⋅s⁻²
	"[psi]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x8, 0x17, 0x5d, 0x56, 0xa9, 0x39}), (&big.Int{}).SetBytes([]byte{0x13, 0xb0, 0x50})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// [pt_br] = 454609/800000000⋅m³
	"[pt_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0x2f, 0xaf, 0x8, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [pt_us] = 473176473/1000000000000⋅m³
	"[pt_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0xe8, 0xd4, 0xa5, 0x10, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [pwt_tr] = 19439673/12500000⋅g
	"[pwt_tr]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x28, 0xa0, 0x39}), (&big.Int{}).SetBytes([]byte{0xbe, 0xbc, 0x20})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [qt_br] = 454609/400000000⋅m³
	"[qt_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xef, 0xd1}), (&big.Int{}).SetBytes([]byte{0x17, 0xd7, 0x84, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [qt_us] = 473176473/500000000000⋅m³
	"[qt_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0x74, 0x6a, 0x52, 0x88, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [rch_us] = 120000/3937⋅m
	"[rch_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0xd4, 0xc0}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [rd_br] = 125729901/25000000⋅m
	"[rd_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0x7e, 0x7c, 0x6d}), (&big.Int{}).SetBytes([]byte{0x1, 0x7d, 0x78, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [rd_us] = 19800/3937⋅m
	"[rd_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4d, 0x58}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [rlk_us] = 1200/3937⋅m
	"[rlk_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4, 0xb0}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [sc_ap] = 6479891/5000000⋅g
	"[sc_ap]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x62, 0xe0, 0x13}), (&big.Int{}).SetBytes([]byte{0x4c, 0x4b, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [sct] = 40144896000000/15499969⋅m²
	"[sct]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x24, 0x82, 0xf6, 0x44, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0xec, 0x82, 0xc1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [scwt_av] = 45359237/1000⋅g
	"[scwt_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xb4, 0x20, 0x85}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [sft_i] = 145161/1562500⋅m²
	"[sft_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0x37, 0x9}), (&big.Int{}).SetBytes([]byte{0x17, 0xd7, 0x84})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [sin_i] = 16129/25000000⋅m²
	"[sin_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3f, 0x1}), (&big.Int{}).SetBytes([]byte{0x1, 0x7d, 0x78, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [smgy'U] = 1
	"[smgy'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [smi_us] = 40144896000000/15499969⋅m²
	"[smi_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x24, 0x82, 0xf6, 0x44, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0xec, 0x82, 0xc1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [smoot] = 8509/5000⋅m
	"[smoot]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x21, 0x3d}), (&big.Int{}).SetBytes([]byte{0x13, 0x88})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [srd_us] = 392040000/15499969⋅m²
	"[srd_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x17, 0x5e, 0xe, 0x40}), (&big.Int{}).SetBytes([]byte{0xec, 0x82, 0xc1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [ston_av] = 45359237/50⋅g
	"[ston_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xb4, 0x20, 0x85}), (&big.Int{}).SetBytes([]byte{0x32})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [stone_av] = 317514659/50000⋅g
	"[stone_av]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x12, 0xec, 0xe3, 0xa3}), (&big.Int{}).SetBytes([]byte{0xc3, 0x50})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// [syd_i] = 1306449/1562500⋅m²
	"[syd_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x13, 0xef, 0x51}), (&big.Int{}).SetBytes([]byte{0x17, 0xd7, 0x84})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [tb'U] = 1
	"[tb'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [tbs_m] = 3/200000⋅m³
	"[tbs_m]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3}), (&big.Int{}).SetBytes([]byte{0x3, 0xd, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [tbs_us] = 473176473/32000000000000⋅m³
	"[tbs_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1c, 0x34, 0x19, 0x99}), (&big.Int{}).SetBytes([]byte{0x1d, 0x1a, 0x94, 0xa2, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [todd'U] = 1
	"[todd'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// [tsp_m] = 1/200000⋅m³
	"[tsp_m]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xd, 0x40})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [tsp_us] = 157725491/32000000000000⋅m³
	"[tsp_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x9, 0x66, 0xb3, 0x33}), (&big.Int{}).SetBytes([]byte{0x1d, 0x1a, 0x94, 0xa2, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// [twp] = 1445216256000000/15499969⋅m²
	"[twp]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x5, 0x22, 0x6a, 0xa1, 0x90, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0xec, 0x82, 0xc1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// [wood'U] = 7999320000⋅g⋅s⁻¹⋅m⁻⁴
	"[wood'U]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0xdc, 0xcb, 0xef, 0xc0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -4, types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// [yd_br] = 11429991/12500000⋅m
	"[yd_br]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xae, 0x68, 0x67}), (&big.Int{}).SetBytes([]byte{0xbe, 0xbc, 0x20})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [yd_i] = 1143/1250⋅m
	"[yd_i]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4, 0x77}), (&big.Int{}).SetBytes([]byte{0x4, 0xe2})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// [yd_us] = 3600/3937⋅m
	"[yd_us]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xe, 0x10}), (&big.Int{}).SetBytes([]byte{0xf, 0x61})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// a = 31557600⋅s
	"a": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0xe1, 0x87, 0xe0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// a_g = 31556952⋅s
	"a_g": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0xe1, 0x85, 0x58}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// a_j = 31557600⋅s
	"a_j": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0xe1, 0x87, 0xe0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// a_t = 3944615652/125⋅s
	"a_t": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xeb, 0x1e, 0xe, 0xe4}), (&big.Int{}).SetBytes([]byte{0x7d})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// ar = 100⋅m²
	"ar": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x64}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// atm = 101325000⋅g⋅m⁻¹⋅s⁻²
	"atm": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6, 0xa, 0x18, 0xc8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// att = 98066500⋅g⋅m⁻¹⋅s⁻²
	"att": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x5, 0xd8, 0x60, 0x44}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// b = 1/10000000000000000000000000000⋅m²
	"b": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x20, 0x4f, 0xce, 0x5e, 0x3e, 0x25, 0x2, 0x61, 0x10, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 2}},
	// bar = 100000000⋅g⋅m⁻¹⋅s⁻²
	"bar": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x5, 0xf5, 0xe1, 0x0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// bit = 1
	"bit": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// cal = 4184⋅m²⋅g⋅s⁻²
	"cal": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x58}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// cal_IT = 20934/5⋅m²⋅g⋅s⁻²
	"cal_IT": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x51, 0xc6}), (&big.Int{}).SetBytes([]byte{0x5})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// cal_[15] = 20929/5⋅m²⋅g⋅s⁻²
	"cal_[15]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x51, 0xc1}), (&big.Int{}).SetBytes([]byte{0x5})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// cal_[20] = 41819/10⋅m²⋅g⋅s⁻²
	"cal_[20]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xa3, 0x5b}), (&big.Int{}).SetBytes([]byte{0xa})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// cal_m = 209501/50⋅m²⋅g⋅s⁻²
	"cal_m": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3, 0x32, 0x5d}), (&big.Int{}).SetBytes([]byte{0x32})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// cal_th = 4184⋅m²⋅g⋅s⁻²
	"cal_th": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x10, 0x58}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// circ = 31415926535897932384626433832795028841971693993751058209749445923/5000000000000000000000000000000000000000000000000000000000000000⋅rad
	"circ": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23}), (&big.Int{}).SetBytes([]byte{0xc, 0x27, 0x81, 0xf4, 0x9f, 0xfc, 0xfa, 0x6d, 0x53, 0xcb, 0xf6, 0xb7, 0x1c, 0x76, 0xb2, 0x5f, 0xb5, 0xf, 0x80, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "rad", Annotation: ""}: 1}},
	// d = 86400⋅s
	"d": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0x51, 0x80}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// deg = 10471975511965977461542144610931676280657231331250352736583148641/600000000000000000000000000000000000000000000000000000000000000000⋅rad
	"deg": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x19, 0x74, 0xb9, 0xf2, 0xbb, 0x34, 0x51, 0x5d, 0x36, 0x61, 0xcd, 0x1e, 0xbf, 0x4, 0xd3, 0x59, 0x37, 0x7a, 0x93, 0x61, 0xf1, 0x7e, 0xd9, 0xe, 0x41, 0xe4, 0x61}), (&big.Int{}).SetBytes([]byte{0x5, 0xb2, 0x84, 0xea, 0xaa, 0xfe, 0x95, 0x63, 0x3f, 0x47, 0x9b, 0xa5, 0xd5, 0x57, 0xa3, 0x9c, 0xdc, 0xdf, 0x44, 0x3c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "rad", Annotation: ""}: 1}},
	// dyn = 1/100⋅g⋅m⋅s⁻²
	"dyn": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x64})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// eV = 801088317/5000000000000000000000000⋅m²⋅g⋅C⁰⋅s⁻²
	"eV": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2f, 0xbf, 0xa3, 0x3d}), (&big.Int{}).SetBytes([]byte{0x4, 0x22, 0xca, 0x8b, 0xa, 0x0, 0xa4, 0x25, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 0, types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// eq = 602214076000000000000000
	"eq": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f, 0x86, 0x17, 0x29, 0x5f, 0x14, 0x5c, 0xc6, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// erg = 1/10000⋅m²⋅g⋅s⁻²
	"erg": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x27, 0x10})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 2, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// g% = 10000⋅g⋅m⁻³
	"g%": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x27, 0x10}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -3}},
	// gf = 196133/20000⋅g⋅m⋅s⁻²
	"gf": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x2, 0xfe, 0x25}), (&big.Int{}).SetBytes([]byte{0x4e, 0x20})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: 1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// gon = 31415926535897932384626433832795028841971693993751058209749445923/2000000000000000000000000000000000000000000000000000000000000000000⋅rad
	"gon": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23}), (&big.Int{}).SetBytes([]byte{0x12, 0xfd, 0xbb, 0xe, 0x39, 0xfb, 0x47, 0x4a, 0xd2, 0xee, 0xb1, 0x7e, 0x1c, 0x79, 0x76, 0xb5, 0x8a, 0xe8, 0x38, 0xc8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "rad", Annotation: ""}: 1}},
	// h = 3600⋅s
	"h": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xe, 0x10}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// kat = 602214076000000000000000⋅s⁻¹
	"kat": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f, 0x86, 0x17, 0x29, 0x5f, 0x14, 0x5c, 0xc6, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: -1}},
	// l = 1/1000⋅m³
	"l": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// lm = rad²⋅cd
	"lm": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "cd", Annotation: ""}: 1, types.ComponentKey{AtomCode: "rad", Annotation: ""}: 2}},
	// lx = rad²⋅cd⋅m⁻²
	"lx": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "cd", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -2, types.ComponentKey{AtomCode: "rad", Annotation: ""}: 2}},
	// m[H2O] = 9806650⋅g⋅m⁻¹⋅s⁻²
	"m[H2O]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x95, 0xa3, 0x3a}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// m[Hg] = 133322000⋅g⋅m⁻¹⋅s⁻²
	"m[Hg]": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7, 0xf2, 0x55, 0x10}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1, types.ComponentKey{AtomCode: "s", Annotation: ""}: -2}},
	// mho = 1/1000⋅C²⋅s⋅g⁻¹⋅m⁻²
	"mho": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "C", Annotation: ""}: 2, types.ComponentKey{AtomCode: "g", Annotation: ""}: -1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -2, types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// min = 60⋅s
	"min": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x3c}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// mo = 2629800⋅s
	"mo": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x28, 0x20, 0xa8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// mo_g = 2629746⋅s
	"mo_g": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x28, 0x20, 0x72}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// mo_j = 2629800⋅s
	"mo_j": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x28, 0x20, 0xa8}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// mo_s = 318930372/125⋅s
	"mo_s": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x13, 0x2, 0x7d, 0xc4}), (&big.Int{}).SetBytes([]byte{0x7d})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
	// mol = 602214076000000000000000
	"mol": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f, 0x86, 0x17, 0x29, 0x5f, 0x14, 0x5c, 0xc6, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// osm = 602214076000000000000000
	"osm": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x7f, 0x86, 0x17, 0x29, 0x5f, 0x14, 0x5c, 0xc6, 0x0, 0x0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{}},
	// pc = 30856780000000000⋅m
	"pc": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x6d, 0xa0, 0x13, 0xf2, 0xcf, 0xf8, 0x0}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 1}},
	// ph = 1/10000⋅rad²⋅cd⋅m⁻²
	"ph": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x27, 0x10})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "cd", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -2, types.ComponentKey{AtomCode: "rad", Annotation: ""}: 2}},
	// sb = 10000⋅cd⋅m⁻²
	"sb": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x27, 0x10}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "cd", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -2}},
	// sph = 31415926535897932384626433832795028841971693993751058209749445923/2500000000000000000000000000000000000000000000000000000000000000⋅rad²
	"sph": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x4c, 0x5e, 0x2d, 0xd8, 0x31, 0x9c, 0xf4, 0x17, 0xa3, 0x25, 0x67, 0x5c, 0x3d, 0xe, 0x7a, 0xb, 0xa6, 0x6f, 0xba, 0x25, 0xd4, 0x7c, 0x8b, 0x2a, 0xc5, 0xad, 0x23}), (&big.Int{}).SetBytes([]byte{0x6, 0x13, 0xc0, 0xfa, 0x4f, 0xfe, 0x7d, 0x36, 0xa9, 0xe5, 0xfb, 0x5b, 0x8e, 0x3b, 0x59, 0x2f, 0xda, 0x87, 0xc0, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "rad", Annotation: ""}: 2}},
	// sr = rad²
	"sr": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "rad", Annotation: ""}: 2}},
	// st = m³
	"st": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "m", Annotation: ""}: 3}},
	// t = 1000000⋅g
	"t": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0xf, 0x42, 0x40}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// tex = 1/1000⋅g⋅m⁻¹
	"tex": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1}), (&big.Int{}).SetBytes([]byte{0x3, 0xe8})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1, types.ComponentKey{AtomCode: "m", Annotation: ""}: -1}},
	// u = 8302695333/5000000000000000000000000000000000⋅g
	"u": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x1, 0xee, 0xe1, 0x13, 0xa5}), (&big.Int{}).SetBytes([]byte{0xf6, 0x84, 0xdf, 0x56, 0xc3, 0xe0, 0x1b, 0xc6, 0xc7, 0x32, 0x0, 0x0, 0x0, 0x0})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "g", Annotation: ""}: 1}},
	// wk = 604800⋅s
	"wk": &types.Unit{Coeff: (&big.Rat{}).SetFrac((&big.Int{}).SetBytes([]byte{0x9, 0x3a, 0x80}), (&big.Int{}).SetBytes([]byte{0x1})), Components: map[types.ComponentKey]int{types.ComponentKey{AtomCode: "s", Annotation: ""}: 1}},
}
