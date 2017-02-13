package expr

import "github.com/shanzi/gexpr/values"

type ExprContext interface {
	Params() map[string]values.Value
	Operators() values.OperatorInterface
	LiteralBuilder() values.LiteralBuilderInterface
}

type _expr_context struct {
	params          map[string]values.Value
	operators       values.OperatorInterface
	literal_builder values.LiteralBuilderInterface
}

func NewContext(params map[string]values.Value, op values.OperatorInterface, lb values.LiteralBuilderInterface) ExprContext {
	return &_expr_context{params, op, lb}
}

func (self *_expr_context) Params() map[string]values.Value {
	return self.params
}

func (self *_expr_context) Operators() values.OperatorInterface {
	return self.operators
}

func (self *_expr_context) LiteralBuilder() values.LiteralBuilderInterface {
	return self.literal_builder
}
