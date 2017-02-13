package expr

import (
	"errors"

	"github.com/shanzi/gexpr/values"
)

type Expr interface {
	Eval(context ExprContext) (values.Value, error)

	String() string
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

func (self *_expr) Eval(context ExprContext) (v values.Value, e error) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
			if msg, ok := r.(string); ok {
				e = errors.New(msg)
			} else {
				e = errors.New("Can not evaluate expression")
			}
		}
	}()
	return self.root.Value(context), nil
}

func (self *_expr) String() string {
	return self.root.String()
}
