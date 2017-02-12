package nodes

import (
	"fmt"

	"github.com/shanzi/gexpr"
	"github.com/shanzi/gexpr/values"
)

type _expr_value_node struct {
	name string
}

func NewValueNode(name string) ExprNode {
	return &_expr_value_node{name}
}

func (self *_expr_value_node) Value(context gexpr.ExprContext) values.Value {
	params := context.Params()
	if value, ok := params[self.name]; value {
		return value
	}
	panic(fmt.Sprintf("Undefined variable: %s", self.name))
}
