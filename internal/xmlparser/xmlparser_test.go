package xmlparser

import (
	"math/big"
	"reflect"
	"strings"
	"testing"
)

const xmlContent = `
<?xml version="1.0" encoding="ascii"?>
<root xmlns="http://unitsofmeasure.org/ucum-essence" version="2.2" revision="N/A" revision-date="2024-06-17">
   <prefix Code="Y" CODE="YA">
      <name>yotta</name>
      <printSymbol>Y</printSymbol>
      <value value="1e24">1 &#215; 10<sup>24</sup>
      </value>
   </prefix>
   <base-unit Code="m" CODE="M" dim="L">
      <name>meter</name>
      <printSymbol>m</printSymbol>
      <property>length</property>
   </base-unit>
    <unit Code="10*" CODE="10*" isMetric="no" class="dimless">
      <name>the number ten for arbitrary powers</name>
      <printSymbol>10</printSymbol>
      <property>number</property>
      <value Unit="1" UNIT="1" value="10">10</value>
   </unit>
</root>
`

func TestParse(t *testing.T) {
	got := Parse(strings.NewReader(xmlContent))
	yotta, ok := (&big.Rat{}).SetString("1e24")
	if !ok {
		t.Fatalf("failed to set yotta")
	}
	want := UCUMData{
		Units: []Unit{
			{Code: "m", FullCode: "m", Kind: "length", Metric: true, Magnitude: big.NewRat(1, 1)},
			{Code: "m", FullCode: "Ym", Kind: "length", Metric: true, Magnitude: yotta},
			{Code: "10*", FullCode: "10*", Kind: "number", Metric: false, Magnitude: big.NewRat(1, 1)},
		},
		XML: XMLRoot{
			Version:      "2.2",
			Revision:     "N/A",
			RevisionDate: "2024-06-17",
			Prefixes: []XMLPrefix{
				{
					XMLConcept: XMLConcept{Code: "Y", CodeUC: "YA", Name: []string{"yotta"}, PrintSymbol: "Y"},
					Value:      XMLValue{Value: yotta},
				},
			},
			BaseUnits: []XMLBaseUnit{
				{
					XMLConcept: XMLConcept{Code: "m", CodeUC: "M", Name: []string{"meter"}, PrintSymbol: "m"},
					Property:   "length",
				},
			},
			Units: []XMLUnit{
				{
					XMLConcept: XMLConcept{Code: "10*", CodeUC: "10*", Name: []string{"the number ten for arbitrary powers"}, PrintSymbol: "10"},
					Property:   "number",
					Value:      XMLValue{Unit: "1", Value: big.NewRat(10, 1)},
					Class:      "dimless",
					Metric:     "no",
				},
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %#v", got)
	}
}
