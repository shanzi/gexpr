package values

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/shanzi/gexpr/types"
)

type OperatorInterface interface {
	// Arithmetic Operators
	ADD(a, b Value) Value // +
	SUB(a, b Value) Value // -
	MUL(a, b Value) Value // *
	DIV(a, b Value) Value // /
	MOD(a, b Value) Value // %

	// Binary Operators
	AND(a, b Value) Value         // &
	OR(a, b Value) Value          // |
	XOR(a, b Value) Value         // ^
	LEFT_SHIFT(a, b Value) Value  // <<
	RIGHT_SHIFT(a, b Value) Value // >>
	AND_NOT(a, b Value) Value     // &^

	// Logic Operators
	EQUAL(a, b Value) Value     // ==
	NOT_EQUAL(a, b Value) Value // !=
	GREATER(a, b Value) Value   // >
	LESS(a, b Value) Value      // <
	GEQ(a, b Value) Value       // >=
	LEQ(a, b Value) Value       // <=

	// Boolean Operators
	BOOL_AND(a, b Value) Value
	BOOL_OR(a, b Value) Value
}

type _operators struct{}

var operators_singleton = &_operators{}

func Operators() OperatorInterface {
	return operators_singleton
}

func (self *_operators) ADD(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() + b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Float(a.Float64() + b.Float64())
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		var buf bytes.Buffer
		buf.WriteString(a.String())
		buf.WriteString(b.String())
		return String(buf.String())
	}

	panic(fmt.Sprintf("Can not add %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) SUB(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() - b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Float(a.Float64() - b.Float64())
	}

	panic(fmt.Sprintf("Can not subtruct %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) MUL(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() * b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Float(a.Float64() * b.Float64())
	}

	panic(fmt.Sprintf("Can not multiply %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) DIV(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() / b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Float(a.Float64() / b.Float64())
	}

	panic(fmt.Sprintf("Can not divide %s with %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) MOD(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() / b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type()) && types.AssertMatch(types.INTEGER, b.Type()) {
		return Float(a.Float64() / b.Float64())
	}

	panic(fmt.Sprintf("Can not modulo %s with %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) AND(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() & b.Int64())
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) OR(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() | b.Int64())
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) XOR(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() ^ b.Int64())
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) LEFT_SHIFT(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() << uint(b.Int64()))
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) RIGHT_SHIFT(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() >> uint(b.Int64()))
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) AND_NOT(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Integer(a.Int64() &^ b.Int64())
	}

	panic(fmt.Sprintf("Can not apply binary operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) EQUAL(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Boolean(a.Int64() == b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Boolean(a.Float64() == b.Float64())
	}

	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return Boolean(a.Bool() == b.Bool())
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return Boolean(strings.Compare(a.String(), b.String()) == 0)
	}

	panic(fmt.Sprintf("Can not apply EQUAL operator on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) NOT_EQUAL(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Boolean(a.Int64() != b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Boolean(a.Float64() != b.Float64())
	}

	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return Boolean(a.Bool() != b.Bool())
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return Boolean(strings.Compare(a.String(), b.String()) != 0)
	}

	panic(fmt.Sprintf("Can not apply NOT_EQUAL operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) GREATER(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Boolean(a.Int64() > b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Boolean(a.Float64() > b.Float64())
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return Boolean(strings.Compare(a.String(), b.String()) > 0)
	}

	panic(fmt.Sprintf("Can not apply GREATER operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) LESS(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Boolean(a.Int64() < b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Boolean(a.Float64() < b.Float64())
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return Boolean(strings.Compare(a.String(), b.String()) < 0)
	}

	panic(fmt.Sprintf("Can not apply LESS operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) GEQ(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Boolean(a.Int64() >= b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Boolean(a.Float64() >= b.Float64())
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return Boolean(strings.Compare(a.String(), b.String()) >= 0)
	}

	panic(fmt.Sprintf("Can not apply GEQ operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) LEQ(a, b Value) Value {
	if types.AssertMatch(types.INTEGER, a.Type(), b.Type()) {
		return Boolean(a.Int64() <= b.Int64())
	}

	if types.AssertMatch(types.FLOAT, a.Type(), b.Type()) {
		return Boolean(a.Float64() <= b.Float64())
	}

	if types.AssertMatch(types.STRING, a.Type(), b.Type()) {
		return Boolean(strings.Compare(a.String(), b.String()) <= 0)
	}

	panic(fmt.Sprintf("Can not apply LEQ operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) BOOL_AND(a, b Value) Value {
	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return Boolean(a.Bool() && b.Bool())
	}

	panic(fmt.Sprintf("Can not apply BOOL_AND operators on %s and %s", a.Type().Name(), b.Type().Name()))
}

func (self *_operators) BOOL_OR(a, b Value) Value {
	if types.AssertMatch(types.BOOLEAN, a.Type(), b.Type()) {
		return Boolean(a.Bool() || b.Bool())
	}

	panic(fmt.Sprintf("Can not apply BOOL_OR operators on %s and %s", a.Type().Name(), b.Type().Name()))
}
