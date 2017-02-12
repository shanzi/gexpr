package nodes

import (
	"fmt"
	"go/token"

	"github.com/shanzi/gexpr"
	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type _expr_literal_node struct {
	literalType   types.Type
	literalString string
}

func NewLiteralNode(kind int, value string) ExprNode {
	var littype types.Type
	switch kind {
	case token.INT:
		littype = types.INTEGER
	case token.FLOAT:
		littype = types.FLOAT
	case token.CHAR:
		littype = types.STRING
	case token.STRING:
		littype = types.STRING
	default:
		panic(fmt.Sprintf("Unsupported literal kind: %d", kind))
	}

	return &_expr_literal_node{littype, value}
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
