package symbols

import (
	"fmt"

	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type Symbol interface {
	values.Value
}

type base_symbol struct {
	type_ types.Type
}

var integer_symbol = &base_symbol{types.INTEGER}
var float_symbol = &base_symbol{types.FLOAT}
var boolean_symbol = &base_symbol{types.BOOLEAN}
var string_symbol = &base_symbol{types.STRING}

func GetSymbol(typ types.Type) Symbol {
	if types.INTEGER.Equals(typ) {
		return integer_symbol
	}

	if types.FLOAT.Equals(typ) {
		return float_symbol
	}

	if types.BOOLEAN.Equals(typ) {
		return boolean_symbol
	}

	if types.STRING.Equals(typ) {
		return string_symbol
	}

	panic(fmt.Sprintf("Unsupported symbol type: %s", typ.Name()))
}

func (self *base_symbol) Type() types.Type {
	return self.type_
}

func (self *base_symbol) Int64() int64 {
	panic("Not implemented!")
}

func (self *base_symbol) Float64() float64 {
	panic("Not implemented!")
}

func (self *base_symbol) Bool() bool {
	panic("Not implemented!")
}

func (self *base_symbol) String() string {
	return fmt.Sprintf("symbol<%s>(%s)", self.Type().Name(), self.literal)
}
