package nodes

import (
	"fmt"

	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/values"
)

type _expr_value_node struct {
	name string
}

func NewValueNode(name string) expr.ExprNode {
	return &_expr_value_node{name}
}

func (self *_expr_value_node) Value(context expr.ExprContext) values.Value {
	if self.name == "true" {
		return context.LiteralBuilder().Boolean("true")
	}
	if self.name == "false" {
		return context.LiteralBuilder().Boolean("false")
	}

	params := context.Params()
	if value, ok := params[self.name]; ok {
		return value
	}
	panic(fmt.Sprintf("Undefined variable: %s", self.name))
}

func (self *_expr_value_node) String() string {
	return self.name
}
