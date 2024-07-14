package ucum

//go:generate go run ./internal/generate_atoms
//go:generate go run ./internal/generate_convertation_table

import (
	"encoding/json"
	"fmt"
	"github.com/iimos/ucum/internal/types"
	"github.com/iimos/ucum/internal/ucumparser"
)

// Unit is a UCUM unit of measure.
type Unit struct {
	u types.Unit
}

var (
	// todo
	//_ fmt.Formatter            = &Unit{}
	//_ fmt.Scanner              = &Unit{}
	_ fmt.Stringer     = &Unit{}
	_ json.Marshaler   = &Unit{}
	_ json.Unmarshaler = &Unit{}
	//_ encoding.TextUnmarshaler = &Unit{}
)

func (u *Unit) String() string {
	//todo: what if Orig==""?
	return u.u.Orig
}

func Parse(unit []byte) (Unit, error) {
	u, err := ucumparser.Parse(unit)
	if err != nil {
		return Unit{}, err
	}
	return Unit{u: u}, nil
}

func MustParse(unit []byte) Unit {
	u, err := Parse(unit)
	if err != nil {
		panic(err)
	}
	return u
}

// MarshalJSON implements the json.Marshaler interface.
func (u *Unit) MarshalJSON() ([]byte, error) {
	str := "\"" + u.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *Unit) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		return nil
	}
	unit, err := Parse(unquote(bytes))
	if err != nil {
		return fmt.Errorf("ucum: error decoding string '%s': %s", string(bytes), err)
	}
	*u = unit
	return nil
}

func unquote(bytes []byte) []byte {
	if len(bytes) >= 2 && bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		return bytes[1 : len(bytes)-1]
	}
	return bytes
}
