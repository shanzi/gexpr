package values

import (
	"errors"
	"fmt"

	"github.com/shanzi/gexpr/types"
)

type Value interface {
	Type() types.Type

	Int64() int64
	Float64() float64
	Bool() bool
	String() string
}

func Pack(v interface{}) (Value, error) {
	switch value := v.(type) {
	case Value:
		return value, nil
	case int:
		return Integer(value), nil
	case int16:
		return Integer(value), nil
	case int32:
		return Integer(value), nil
	case int64:
		return Integer(value), nil
	case float32:
		return Float(value), nil
	case float64:
		return Float(value), nil
	case bool:
		return Boolean(value), nil
	case string:
		return String(value), nil
	default:
		return nil, errors.New(fmt.Sprint("Unsupported value: ", v))
	}
}

func PackMap(params map[string]interface{}) (map[string]Value, error) {
	ret := make(map[string]Value, len(params))
	for k, v := range params {
		if value, err := Pack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type for key: ", k))
		} else {
			ret[k] = value
		}
	}
	return ret, nil
}

func PackSlice(params []interface{}) ([]Value, error) {
	ret := make([]Value, len(params))
	for _, v := range params {
		if value, err := Pack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type"))
		} else {
			ret = append(ret, value)
		}
	}
	return ret, nil
}

func Unpack(v Value) (interface{}, error) {
	tp := v.Type()
	if tp.Match(types.INTEGER) {
		return v.Int64(), nil
	}
	if tp.Match(types.FLOAT) {
		return v.Float64(), nil
	}
	if tp.Match(types.BOOLEAN) {
		return v.Bool(), nil
	}
	if tp.Match(types.STRING) {
		return v.String(), nil
	}
	return nil, errors.New(fmt.Sprint("Can not unpack value of type: ", tp.Name()))
}

func UnpackSlice(params []Value) ([]interface{}, error) {
	ret := make([]interface{}, len(params))
	for _, v := range params {
		if value, err := Unpack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type"))
		} else {
			ret = append(ret, value)
		}
	}
	return ret, nil
}
