// Package testcases provides the conversion matrix shared by the tests of the
// ucum package and of the ucumapd and ucumdecimal modules.
//
// The matrix lives here rather than in a testdata directory of the root module
// because the adapter modules are published separately: their module archives
// do not contain files of the parent module, so a relative path such as
// "../testdata" only works in a repository checkout.
package testcases

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed conversions.json
var conversionsJSON []byte

// Case is a single conversion case: converting Value from From to To must yield
// exactly Want. Value and Want are finite decimal numbers.
type Case struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Want  string `json:"want"`
}

// Name returns a human-readable name of the case suitable for a subtest name.
func (c Case) Name() string {
	return c.Value + " " + c.From + "->" + c.To
}

// Conversions returns the shared conversion cases.
func Conversions() ([]Case, error) {
	dec := json.NewDecoder(bytes.NewReader(conversionsJSON))
	dec.DisallowUnknownFields()

	var cases []Case
	if err := dec.Decode(&cases); err != nil {
		return nil, fmt.Errorf("testcases: decode conversions: %w", err)
	}
	if len(cases) == 0 {
		return nil, fmt.Errorf("testcases: conversions are empty")
	}
	for i, c := range cases {
		switch {
		case c.From == "":
			return nil, fmt.Errorf("testcases: case %d: empty from", i)
		case c.To == "":
			return nil, fmt.Errorf("testcases: case %d: empty to", i)
		case c.Value == "":
			return nil, fmt.Errorf("testcases: case %d: empty value", i)
		case c.Want == "":
			return nil, fmt.Errorf("testcases: case %d: empty want", i)
		}
	}
	return cases, nil
}
