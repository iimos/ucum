package main

import (
	"fmt"
	"github.com/iimos/ucum/internal/types"
	"github.com/iimos/ucum/internal/ucumparser"
	"github.com/iimos/ucum/internal/xmlparser"
	"log"
	"math/big"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
)

const url = "https://raw.githubusercontent.com/ucum-org/ucum/main/ucum-essence.xml"

type namedUnit struct {
	name string
	unit *types.Unit
}

func main() {
	//xx := ucumparser.Parse("")
	//fmt.Printf("%s = %s\n", xx.String(), xx.CanonicalString())
	//return

	resp, err := (&http.Client{Timeout: 10 * time.Second}).Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data := xmlparser.Parse(resp.Body)

	conv := make(map[string]*types.Unit, len(data.XML.Units))

	for _, xu := range data.XML.Units {
		if !xu.Value.IsFunctional() {
			u, err2 := ucumparser.Parse([]byte(xu.Value.Unit))
			if err2 != nil {
				log.Fatalf("ucum.Parse(%s): %s", xu.Value.Unit, err)
			}
			u.Coeff.Mul(u.Coeff, xu.Value.Value)
			conv[xu.Code] = &u
		} else {
			fn := xu.Value.Function
			fmt.Printf("Func: %s(%g %s); IsSpecial=%s\n", fn.Name, fn.Value, fn.Unit, xu.IsSpecial)
		}
	}

	//xx := conv["[gr]"]
	//fmt.Printf("%s = %s\n", xx.String(), xx.CanonicalString())
	//return

	// Expand every unit until all unit consists only of primitive ones.
	for {
		nothingDone := true
		for _, u := range conv {
			cpy := make(map[types.ComponentKey]int, len(u.Components))
			for key, exp := range u.Components {
				u2, ok := conv[key.AtomCode]
				if !ok || isOne(u2) {
					cpy[key] += exp
				} else {
					nothingDone = false
					u.Coeff.Mul(u.Coeff, pow(u2.Coeff, exp))
					for key2, exp2 := range u2.Components {
						cpy[key2] += exp * exp2
					}
				}
			}
			//fmt.Printf("%s = %s\n", name, u.CanonicalString())
			u.Components = cpy
		}
		if nothingDone {
			break
		}
	}

	unitsSimplified := make([]namedUnit, 0, len(conv))
	for name, unit := range conv {
		unitsSimplified = append(unitsSimplified, namedUnit{name, unit})
	}
	slices.SortFunc(unitsSimplified, func(a, b namedUnit) int {
		return strings.Compare(a.name, b.name)
	})

	//baseUnitsSet := make(map[string]struct{})
	//for _, p := range unitsSimplified {
	//	//fmt.Printf("%s = %s\n", p.name, p.unit.CanonicalString())
	//	//fmt.Printf("%s\n", gocodeUnit(p.unit))
	//	for comp, _ := range p.unit.Components {
	//		baseUnitsSet[comp.AtomCode] = struct{}{}
	//	}
	//}
	//baseUnits := maps.Keys(baseUnitsSet)
	//slices.Sort(baseUnits)
	//fmt.Printf("\nBase units: %#v\n", baseUnits)

	gen := Generator{
		packageName: "data",
		units:       unitsSimplified,
	}
	gen.Generate()
	gocode := gen.Format()
	if err = os.WriteFile("./internal/data/conv.gen.go", gocode, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

func isOne(u *types.Unit) bool {
	if len(u.Components) != 0 {
		return false
	}
	return u.Coeff.Cmp(big.NewRat(1, 1)) == 0
}

func pow(num *big.Rat, exp int) *big.Rat {
	cpy := new(big.Rat).Set(num)
	if exp < 0 {
		cpy = cpy.Inv(cpy)
		exp = -exp
	} else if exp == 0 {
		cpy.SetFrac64(0, 1)
	}
	multiplier := new(big.Rat).Set(cpy)
	for i := exp; i > 1; i-- {
		cpy.Mul(cpy, multiplier)
	}
	return cpy
}
