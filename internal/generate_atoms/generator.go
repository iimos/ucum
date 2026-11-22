package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"math/big"

	"github.com/iimos/ucum/internal/xmlparser"
)

type Generator struct {
	buf         bytes.Buffer
	packageName string
}

func (g *Generator) Printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(&g.buf, format, args...)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) Generate(data xmlparser.UCUMData) {
	g.Printf("// Code generated; DO NOT EDIT.\n")
	g.Printf("package %s\n", g.packageName)
	g.Printf("import \"math/big\"\n")
	g.Printf("import \"github.com/iimos/ucum/internal/types\"\n")
	g.Printf("\n")
	g.Printf("var prefixes = [...]big.Rat{}\n")
	g.Printf("\n")
	g.Printf("var Atoms = map[string]types.Atom{\n")
	for _, unit := range data.Units {
		// Print the Arbitrary field only when it is set to minimize the diff
		arbitrary := ""
		if unit.Arbitrary {
			arbitrary = ", Arbitrary: true"
		}
		g.Printf("%q: {Code: %q, Kind: %q, Metric: %t%s, Magnitude: %s},\n",
			unit.FullCode, unit.Code, unit.Kind, unit.Metric, arbitrary, gocodeRat(unit.Magnitude))
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

func gocodeRat(r *big.Rat) string {
	a, b := r.Num(), r.Denom()
	return fmt.Sprintf("(&big.Rat{}).SetFrac((&big.Int{}).SetBytes(%#v), (&big.Int{}).SetBytes(%#v))", a.Bytes(), b.Bytes())
}
