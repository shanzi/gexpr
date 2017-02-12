package expr

import "github.com/shanzi/gexpr/values"

type Expr interface {
	// Evaluate(params map[string]values.Value) values.Value
	// EvaluateType(params map[string]types.Type) types.Type
	// Eval(context ExprContext) ExprContext

	String() string
}

type ExprContext interface {
	Params() map[string]values.Value
	Operators() values.OperatorInterface
	LiteralBuilder() values.LiteralBuilderInterface
}

type ExprNode interface {
	Value(context ExprContext) values.Value
	String() string
}

type _expr struct {
	root    ExprNode
	exprStr string
}

func New(root ExprNode, exprStr string) Expr {
	return &_expr{root, exprStr}
}

func (self *_expr) String() string {
	return self.root.String()
}
