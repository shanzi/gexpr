package gexpr

import (
	"errors"
	"fmt"

	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/symbols"
	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

func packValuesMap(params map[string]interface{}) (map[string]values.Value, error) {
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

func packTypesMap(params map[string]types.Type) (map[string]values.Value, error) {
	ret := make(map[string]values.Value, len(params))
	for k, v := range params {
		if value, err := symbols.Pack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type for key: ", k))
		} else {
			ret[k] = value
		}
	}
	return ret, nil
}

func Evaluate(exp expr.Expr, params map[string]interface{}) (interface{}, error) {
	if pm, err := packValuesMap(params); err != nil {
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

func EvaluateType(exp expr.Expr, params map[string]types.Type) (types.Type, error) {
	if pm, err := packTypesMap(params); err != nil {
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
