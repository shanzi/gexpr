package expr

import (
	"fmt"

	"github.com/shanzi/gexpr/values"
)

type _expr_value_node struct {
	name string
}

func (self *_expr_value_node) Value(context ExprContext) values.Value {
	params := context.Params()
	if value, ok := params[self.name]; value {
		return value
	}
	panic(fmt.Sprintf("Undefined variable: %s", self.name))
}
