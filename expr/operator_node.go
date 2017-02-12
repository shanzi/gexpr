package expr

import (
	"fmt"
	"go/token"

	"github.com/shanzi/gexpr/values"
)

type _expr_operator_node struct {
	op_code    int
	parameter1 ExprNode
	parameter2 ExprNode
}

func (self *_expr_operator_node) Value(context ExprContext) values.Value {
	op := context.Operators()
	v1 := self.parameter1.Value(context)
	v2 := self.parameter2.Value(context)
	switch op_code {

	case token.ADD:
		return op.ADD(v1, v2)
	case token.SUB:
		return op.SUB(v1, v2)
	case token.MUL:
		return op.MUL(v1, v2)
	case token.QUO:
		return op.DIV(v1, v2)
	case token.REM:
		return op.REM(v1, v2)

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

	panic(fmt.Sprintf("Unknown operator code: %d", self.op_code))
}
