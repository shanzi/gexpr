package symbols

import (
	"fmt"

	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type _operators struct{}

var operators_singleton = &_operators{}

func Operators() values.OperatorInterface {
	return operators_singleton
}

func (self *_operators) ADD(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return float_symbol
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return string_symbol
	}

	panic(fmt.Sprintf("Can not add %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) SUB(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return float_symbol
	}

	panic(fmt.Sprintf("Can not subtruct %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) MUL(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return float_symbol
	}

	panic(fmt.Sprintf("Can not multiply %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) DIV(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return float_symbol
	}

	panic(fmt.Sprintf("Can not divide %s with %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) MOD(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type()) && types.AssertMatch(types.INTEGER, b.Type()) {
		return float_symbol
	}

	panic(fmt.Sprintf("Can not modulo %s with %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) AND(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) OR(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) XOR(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) LEFT_SHIFT(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) RIGHT_SHIFT(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) AND_NOT(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return integer_symbol
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) EQUAL(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply EQUAL operator on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) NOT_EQUAL(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply NOT_EQUAL operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) GREATER(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply GREATER operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) LESS(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply LESS operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) GTE(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply GTE operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) LTE(a, b values.Value) values.Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return boolean_symbol
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply GTE operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) BOOL_AND(a, b values.Value) values.Value {
	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply BOOL_AND operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) BOOL_OR(a, b values.Value) values.Value {
	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return boolean_symbol
	}

	panic(fmt.Sprintf("Can not apply BOOL_OR operators on %s and %s", a.Type().Name(), b.Type().Name()))
}
