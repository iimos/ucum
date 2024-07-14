package main

import (
	"bytes"
	"fmt"
	"github.com/iimos/ucum/internal/types"
	"go/format"
	"log"
	"math/big"
)

type Generator struct {
	buf         bytes.Buffer
	packageName string
	units       []namedUnit
}

func (g *Generator) Printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(&g.buf, format, args...)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) Generate() {
	g.Printf("// Code generated; DO NOT EDIT.\n")
	g.Printf("package %s\n", g.packageName)
	g.Printf("import \"math/big\"\n")
	g.Printf("import \"github.com/iimos/ucum/internal/types\"\n")
	g.Printf("\n")
	g.Printf("var Conv = map[string]*types.Unit{\n")
	for _, u := range g.units {
		g.Printf("// %s = %s\n", u.name, u.unit.CanonicalString())
		g.Printf("%q: %s,\n", u.name, gocodeUnit(u.unit))
	}
	g.Printf("}\n")
}

// Format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) Format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		log.Printf("error: internal error: invalid Go generated: %s", err)
		log.Printf("error: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

func gocodeUnit(u *types.Unit) string {
	return fmt.Sprintf("&types.Unit{Coeff: %s, Components: %#v}", gocodeRat(u.Coeff), u.Components)
}

func gocodeRat(r *big.Rat) string {
	a, b := r.Num(), r.Denom()
	return fmt.Sprintf("(&big.Rat{}).SetFrac((&big.Int{}).SetBytes(%#v), (&big.Int{}).SetBytes(%#v))", a.Bytes(), b.Bytes())
}
