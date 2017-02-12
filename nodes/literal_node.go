package nodes

import (
	"fmt"

	"github.com/shanzi/gexpr"
	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type _expr_literal_node struct {
	literalType   types.Type
	literalString string
}

func (self *_expr_literal_node) Value(context gexpr.ExprContext) values.Value {
	lb := ExprContext.LiteralBuilder()

	if types.INTEGER.match(self.literalType) {
		return lb.Integer(self.literalString)
	}

	if types.FLOAT.match(self.literalType) {
		return lb.Float(self.literalString)
	}

	if types.BOOL.match(self.literalType) {
		return lb.Boolean(self.literalString)
	}

	if types.String.match(self.literalType) {
		return lb.String(self.literalString)
	}

	panic(fmt.Sprintf("Can not construct literal of type %s", self.literalType.Name()))
}
