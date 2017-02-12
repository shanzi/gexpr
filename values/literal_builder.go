package values

import (
	"fmt"
	"strconv"
)

type LiteralBuilderInterface interface {
	Integer(literal string) Value
	Float(literal string) Value
	Boolean(literal string) Value
	String(literal string) Value
}

type _literal_builder struct{}

var literal_builder_singleton = &_literal_builder{}

func LiteralBuilder() LiteralBuilderInterface {
	return literal_builder_singleton
}

func (self *_literal_builder) Integer(literal string) Value {
	if value, err := strconv.ParseInt(literal, 10, 64); err != nil {
		panic(fmt.Sprintf("Can not construct INTEGER from '%s': %s", literal, err))
	} else {
		return Integer(value)
	}
}

func (self *_literal_builder) Float(literal string) Value {
	if value, err := strconv.ParseFloat(literal, 64); err != nil {
		panic(fmt.Sprintf("Can not construct FLOAT from '%s': %s", literal, err))
	} else {
		return Float(value)
	}
}

func (self *_literal_builder) Boolean(literal string) Value {
	if value, err := strconv.ParseBool(literal); err != nil {
		panic(fmt.Sprintf("Can not construct BOOLEAN from '%s': %s", literal, err))
	} else {
		return Boolean(value)
	}
}

func (self *_literal_builder) String(literal string) Value {
	return String(literal)
}
