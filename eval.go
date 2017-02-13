package gexpr

import (
	"errors"
	"fmt"

	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/values"
)

func PackMap(params map[string]interface{}) (map[string]values.Value, error) {
	ret := make(map[string]values.Value, len(params))
	for k, v := range params {
		if value, err := values.Pack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type for key: ", k))
		} else {
			ret[k] = value
		}
	}
	return ret, nil
}

func Evaluate(exp expr.Expr, params map[string]interface{}) (interface{}, error) {
	if pm, err := PackMap(params); err != nil {
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
