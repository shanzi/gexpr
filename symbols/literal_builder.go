package symbols

import (
	"fmt"
	"strconv"

	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type _literal_builder struct{}

var literal_builder_singleton = &_literal_builder{}

func LiteralBuilder() values.LiteralBuilderInterface {
	return literal_builder_singleton
}

func (self *_literal_builder) Integer(literal string) values.Value {
	if _, err := strconv.ParseInt(literal, 10, 64); err != nil {
		panic(fmt.Sprintf("Can not construct INTEGER from '%s': %s", literal, err))
	} else {
		return GetSymbol(types.INTEGER)
	}
}

func (self *_literal_builder) Float(literal string) values.Value {
	if _, err := strconv.ParseFloat(literal, 64); err != nil {
		panic(fmt.Sprintf("Can not construct FLOAT from '%s': %s", literal, err))
	} else {
		return GetSymbol(types.FLOAT)
	}
}

func (self *_literal_builder) Boolean(literal string) values.Value {
	if _, err := strconv.ParseBool(literal); err != nil {
		panic(fmt.Sprintf("Can not construct BOOLEAN from '%s': %s", literal, err))
	} else {
		return GetSymbol(types.BOOLEAN)
	}
}

func (self *_literal_builder) String(literal string) values.Value {
	return GetSymbol(types.STRING)
}
