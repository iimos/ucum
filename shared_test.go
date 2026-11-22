package ucum

import (
	"math/big"
	"testing"

	"github.com/iimos/ucum/internal/testcases"
)

// convCase is a test case for conversion functions.
type convCase struct {
	A, B    string
	value   *big.Rat
	want    *big.Rat
	inexact bool
}

// sharedConvCases returns the exact conversion matrix shared with the ucumapd
// and ucumdecimal modules, see internal/testcases.
func sharedConvCases(t *testing.T) []convCase {
	t.Helper()

	cases, err := testcases.Conversions()
	if err != nil {
		t.Fatalf("load shared conversion cases: %v", err)
	}

	ret := make([]convCase, 0, len(cases))
	for _, c := range cases {
		if _, err := Parse([]byte(c.From)); err != nil {
			t.Fatalf("shared case %s: invalid from unit: %v", c.Name(), err)
		}
		if _, err := Parse([]byte(c.To)); err != nil {
			t.Fatalf("shared case %s: invalid to unit: %v", c.Name(), err)
		}
		ret = append(ret, convCase{
			A:     c.From,
			B:     c.To,
			value: mustParseRat(t, c.Value),
			want:  mustParseRat(t, c.Want),
		})
	}
	return ret
}

func mustParseRat(t *testing.T, s string) *big.Rat {
	t.Helper()
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		t.Fatalf("cannot parse %q as a rational number", s)
	}
	return r
}
