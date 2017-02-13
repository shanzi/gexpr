package nodes

import (
	"fmt"
	"strings"

	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type _expr_func_node struct {
	name expr.ExprNode
	args []expr.ExprNode
}

func NewFuncNode(name expr.ExprNode, args []expr.ExprNode) expr.ExprNode {
	return &_expr_func_node{name, args}
}

func (self *_expr_func_node) NameValue(context expr.ExprContext) values.Func {
	namev := self.name.Value(context)
	if f, ok := namev.(values.Func); ok {
		return f
	}
	panic(fmt.Sprintf("%s is not a function", self.name.String()))
}

func (self *_expr_func_node) ArgValues(context expr.ExprContext) []values.Value {
	ret := make([]values.Value, 0, len(self.args))
	for _, v := range self.args {
		ret = append(ret, v.Value(context))
	}
	return ret
}

func (self *_expr_func_node) ArgTypes(values []values.Value) []types.Type {
	ret := make([]types.Type, 0, len(values))
	for _, v := range values {
		ret = append(ret, v.Type())
	}
	return ret
}

func (self *_expr_func_node) Value(context expr.ExprContext) values.Value {
	funcv := self.NameValue(context)
	argsv := self.ArgValues(context)
	argst := self.ArgTypes(argsv)

	if !funcv.Accept(argsv) {
		panic(fmt.Sprint("Illegal function call parameters: ", self))
	}

	res := funcv.Call(argsv)
	rest := funcv.CallType(argst)

	if !rest.Match(res.Type()) {
		panic(fmt.Sprint("Mismatched return type: ", res.Type()))
	}

	return res
}

func (self *_expr_func_node) String() string {
	argvalues := make([]string, 0, len(self.args)+1)
	argvalues = append(argvalues, self.name.String())
	for _, v := range self.args {
		argvalues = append(argvalues, v.String())
	}
	return fmt.Sprintf("(%s)", strings.Join(argvalues, " "))
}
