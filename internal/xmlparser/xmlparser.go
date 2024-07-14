package xmlparser

import (
	"encoding/xml"
	"io"
	"log"
	"math/big"
)

type UCUMData struct {
	Units []Unit
	XML   XMLRoot
}

type Unit struct {
	Code string
	// FullCode is a code with a prefix.
	FullCode  string
	Kind      string
	Metric    bool
	Magnitude *big.Rat
}

func Parse(reader io.Reader) UCUMData {
	decoder := xml.NewDecoder(reader)
	decoder.Strict = true
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset != "ascii" {
			panic("unexpected charset: " + charset)
		}
		return input, nil
	}
	data := &XMLRoot{}
	if err := decoder.Decode(data); err != nil {
		panic(err)
	}

	propertyMap := make(map[string]struct{})
	classMap := make(map[string]struct{})
	for _, u := range data.BaseUnits {
		propertyMap[u.Property] = struct{}{}
	}
	for _, u := range data.Units {
		propertyMap[u.Property] = struct{}{}
		classMap[u.Class] = struct{}{}
	}

	units := make([]Unit, 0, len(data.Units))
	unitsMap := make(map[string]Unit, len(data.Units))
	addUnit := func(u Unit) {
		if _, exists := unitsMap[u.FullCode]; exists {
			log.Fatalf("duplicate unit: %v", u.FullCode)
		}
		unitsMap[u.FullCode] = u
		units = append(units, u)
	}

	for _, u := range data.BaseUnits {
		addUnit(Unit{
			Code:      u.Code,
			FullCode:  u.Code,
			Kind:      u.Property,
			Metric:    true, // base units are always metric
			Magnitude: big.NewRat(1, 1),
		})
		for _, pref := range data.Prefixes {
			addUnit(Unit{
				Code:      u.Code,
				FullCode:  pref.Code + u.Code,
				Kind:      u.Property,
				Metric:    true,
				Magnitude: pref.Value.Value,
			})
		}
	}

	for _, u := range data.Units {
		addUnit(Unit{
			Code:      u.Code,
			FullCode:  u.Code,
			Kind:      u.Property,
			Metric:    u.Metric == "yes",
			Magnitude: big.NewRat(1, 1),
		})
		if u.Metric == "yes" {
			for _, pref := range data.Prefixes {
				addUnit(Unit{
					Code:      u.Code,
					FullCode:  pref.Code + u.Code,
					Kind:      u.Property,
					Metric:    true,
					Magnitude: pref.Value.Value,
				})
			}
		}
	}

	return UCUMData{
		Units: units,
		XML:   *data,
	}
}

type XMLRoot struct {
	Version      string        `xml:"version,attr"`
	Revision     string        `xml:"revision,attr"`
	RevisionDate string        `xml:"revision-date,attr"`
	Prefixes     []XMLPrefix   `xml:"prefix"`
	BaseUnits    []XMLBaseUnit `xml:"base-unit"`
	Units        []XMLUnit     `xml:"unit"`
}

type XMLConcept struct {
	Code        string   `xml:"Code,attr"`
	CodeUC      string   `xml:"CODE,attr"`
	Name        []string `xml:"name"`
	PrintSymbol string   `xml:"printSymbol"`
}

type XMLPrefix struct {
	XMLConcept
	Value XMLValue `xml:"value"`
}

type XMLBaseUnit struct {
	XMLConcept
	Property string `xml:"property"`
	Dim      rune   `xml:"dim"`
}

type XMLUnit struct {
	XMLConcept
	Class       string   `xml:"class,attr"`
	IsSpecial   string   `xml:"isSpecial,attr"`
	IsArbitrary string   `xml:"isArbitrary,attr"`
	Metric      string   `xml:"isMetric,attr"`
	Property    string   `xml:"property"`
	Value       XMLValue `xml:"value"`
}

type XMLValue struct {
	Unit     string      `xml:"Unit,attr"`
	Value    *big.Rat    `xml:"value,attr"`
	Function XMLFunction `xml:"function"`
}

func (v *XMLValue) IsFunctional() bool {
	return v.Function != XMLFunction{}
}

type XMLFunction struct {
	Name  string  `xml:"name,attr"`
	Value float64 `xml:"value,attr"`
	Unit  string  `xml:"Unit,attr"`
}
