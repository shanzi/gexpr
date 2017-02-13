package nodes

import (
	"fmt"
	"go/token"

	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type _expr_literal_node struct {
	literalType   types.Type
	literalString string
}

func NewLiteralNode(kind int, value string) expr.ExprNode {
	var littype types.Type
	switch token.Token(kind) {
	case token.INT:
		littype = types.INTEGER
	case token.FLOAT:
		littype = types.FLOAT
	case token.STRING:
		littype = types.STRING
	default:
		panic(fmt.Sprintf("Unsupported literal kind: %d", kind))
	}

	return &_expr_literal_node{littype, value}
}

func (self *_expr_literal_node) Value(context expr.ExprContext) values.Value {
	lb := context.LiteralBuilder()

	if types.INTEGER.Match(self.literalType) {
		return lb.Integer(self.literalString)
	}

	if types.FLOAT.Match(self.literalType) {
		return lb.Float(self.literalString)
	}

	if types.BOOLEAN.Match(self.literalType) {
		return lb.Boolean(self.literalString)
	}

	if types.STRING.Match(self.literalType) {
		return lb.String(self.literalString)
	}

	panic(fmt.Sprintf("Can not construct literal of type %s", self.literalType.Name()))
}

func (self *_expr_literal_node) String() string {
	return self.literalString
}
