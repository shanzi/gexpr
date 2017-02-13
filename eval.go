package gexpr

import (
	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/symbols"
	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

func Evaluate(exp expr.Expr, params map[string]interface{}) (interface{}, error) {
	if pm, err := values.PackMap(params); err != nil {
		return nil, err
	} else {
		context := expr.NewContext(pm, values.Operators(), values.LiteralBuilder())
		if value, err := exp.Eval(context); err != nil {
			return nil, err
		} else {
			return values.Unpack(value)
		}
	}
}

func EvaluateType(exp expr.Expr, params map[string]interface{}) (types.Type, error) {
	if pm, err := symbols.PackMap(params); err != nil {
		return nil, err
	} else {
		context := expr.NewContext(pm, symbols.Operators(), symbols.LiteralBuilder())
		if value, err := exp.Eval(context); err != nil {
			return nil, err
		} else {
			return symbols.Unpack(value)
		}
	}
}
