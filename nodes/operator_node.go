package nodes

import (
	"fmt"
	"go/token"

	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/values"
)

type _expr_binary_operator_node struct {
	kind int
	name string
	X    expr.ExprNode
	Y    expr.ExprNode
}

type _expr_unary_operator_node struct {
	kind int
	name string
	X    expr.ExprNode
}

func NewBinaryOperatorNode(kind int, a, b expr.ExprNode) expr.ExprNode {
	return &_expr_binary_operator_node{kind, GetOpName(kind), a, b}
}

func (self *_expr_binary_operator_node) Value(context expr.ExprContext) values.Value {
	op := context.Operators()
	v1 := self.X.Value(context)
	v2 := self.Y.Value(context)
	switch token.Token(self.kind) {

	case token.ADD:
		return op.ADD(v1, v2)
	case token.SUB:
		return op.SUB(v1, v2)
	case token.MUL:
		return op.MUL(v1, v2)
	case token.QUO:
		return op.DIV(v1, v2)
	case token.REM:
		return op.MOD(v1, v2)

	case token.AND:
		return op.AND(v1, v2)
	case token.OR:
		return op.OR(v1, v2)
	case token.XOR:
		return op.XOR(v1, v2)
	case token.SHL:
		return op.LEFT_SHIFT(v1, v2)
	case token.SHR:
		return op.RIGHT_SHIFT(v1, v2)
	case token.AND_NOT:
		return op.AND_NOT(v1, v2)

	case token.EQL:
		return op.EQUAL(v1, v2)
	case token.NEQ:
		return op.NOT_EQUAL(v1, v2)
	case token.GTR:
		return op.GREATER(v1, v2)
	case token.LSS:
		return op.LESS(v1, v2)
	case token.GEQ:
		return op.GEQ(v1, v2)
	case token.LEQ:
		return op.LEQ(v1, v2)

	case token.LAND:
		return op.BOOL_AND(v1, v2)
	case token.LOR:
		return op.BOOL_OR(v1, v2)
	}

	panic(fmt.Sprintf("Unknown operator code: %d", self.kind))
}

func (self *_expr_binary_operator_node) String() string {
	return fmt.Sprintf("(%s %s %s)", self.name, self.X.String(), self.Y.String())
}

func GetOpName(kind int) string {
	switch token.Token(kind) {
	default:
		panic(fmt.Sprintf("Unsupported binary operator kind: %d", kind))
	case token.ADD:
		return "+"
	case token.SUB:
		return "-"
	case token.MUL:
		return "*"
	case token.QUO:
		return "/"
	case token.REM:
		return "%"
	case token.AND:
		return "&"
	case token.OR:
		return "|"
	case token.XOR:
		return "^"
	case token.SHL:
		return "<<"
	case token.SHR:
		return ">>"
	case token.EQL:
		return "=="
	case token.NEQ:
		return "!="
	case token.GTR:
		return ">"
	case token.LSS:
		return "<"
	case token.GEQ:
		return ">="
	case token.LEQ:
		return "<="
	case token.LAND:
		return "&&"
	case token.LOR:
		return "||"
	case token.INC:
		return "++"
	case token.DEC:
		return "--"
	case token.NOT:
		return "!"

	}
}

func NewUnaryOperatorNode(kind int, a expr.ExprNode) expr.ExprNode {
	return &_expr_unary_operator_node{kind, GetOpName(kind), a}
}

func (self *_expr_unary_operator_node) Value(context expr.ExprContext) values.Value {
	op := context.Operators()
	v1 := self.X.Value(context)
	switch token.Token(self.kind) {

	case token.ADD:
		return op.POSITIVE(v1)
	case token.SUB:
		return op.NEGATIVE(v1)
	case token.XOR:
		return op.INV(v1)
	case token.INC:
		return op.INC(v1)
	case token.DEC:
		return op.DEC(v1)
	case token.NOT:
		return op.BOOL_NOT(v1)
	}

	panic(fmt.Sprintf("Unknown operator code: %d", self.kind))
}

func (self *_expr_unary_operator_node) String() string {
	return fmt.Sprintf("(%s %s)", self.name, self.X.String())
}
