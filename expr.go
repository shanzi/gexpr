package gexpr

import (
	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type Expr interface {
	Evaluate(params map[string]values.Value) values.Value
	EvaluateType(params map[string]types.Type) types.Type
	Eval(context ExprContext) ExprContext

	String() string
}

type ExprContext interface {
	Params() map[string]values.Value
	Operators() values.OperatorInterface
	LiteralBuilder() values.LiteralBuilderInterface
}

type ExprNode interface {
	Value(context ExprContext) values.Value
}

type _expr struct {
	root    ExprNode
	exprStr string
}
