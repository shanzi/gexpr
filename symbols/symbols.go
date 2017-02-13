package symbols

import (
	"errors"
	"fmt"

	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type Symbol interface {
	values.Value

	SymbolName() string
}

type base_symbol struct {
	type_ types.Type
}

var integer_symbol = &base_symbol{types.INTEGER}
var float_symbol = &base_symbol{types.FLOAT}
var boolean_symbol = &base_symbol{types.BOOLEAN}
var string_symbol = &base_symbol{types.STRING}

func Pack(tp interface{}) (Symbol, error) {
	if s, ok := tp.(Symbol); ok {
		return s, nil
	}

	if t, ok := tp.(types.Type); ok {
		if types.INTEGER.Equals(t) {
			return integer_symbol, nil
		}

		if types.FLOAT.Equals(t) {
			return float_symbol, nil
		}

		if types.BOOLEAN.Equals(t) {
			return boolean_symbol, nil
		}

		if types.STRING.Equals(t) {
			return string_symbol, nil
		}
	}

	if f, ok := tp.(values.Func); ok {
		return PackFunc(f), nil
	}

	return nil, errors.New(fmt.Sprint("Unsupported type: ", tp))
}

func Unpack(v values.Value) (types.Type, error) {
	return v.Type(), nil
}

func PackMap(params map[string]interface{}) (map[string]values.Value, error) {
	ret := make(map[string]values.Value, len(params))
	for k, v := range params {
		if value, err := Pack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type for key: ", k))
		} else {
			ret[k] = value
		}
	}
	return ret, nil
}

func PackSlice(params []interface{}) ([]values.Value, error) {
	ret := make([]values.Value, 0, len(params))
	for _, v := range params {
		if value, err := Pack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type for key"))
		} else {
			ret = append(ret, value)
		}
	}
	return ret, nil
}

func UnpackSlice(params []values.Value) ([]types.Type, error) {
	ret := make([]types.Type, 0, len(params))
	for _, v := range params {
		if value, err := Unpack(v); err != nil {
			return nil, errors.New(fmt.Sprint("Unsupported value type for key"))
		} else {
			ret = append(ret, value)
		}
	}
	return ret, nil
}

func GetSymbol(typ types.Type) Symbol {
	if sym, err := Pack(typ); err == nil {
		return sym
	}
	panic(fmt.Sprintf("Unsupported symbol type: %s", typ.Name()))
}

func (self *base_symbol) Type() types.Type {
	return self.type_
}

func (self *base_symbol) Int64() int64 {
	panic("Not implemented!")
}

func (self *base_symbol) Float64() float64 {
	panic("Not implemented!")
}

func (self *base_symbol) Bool() bool {
	panic("Not implemented!")
}

func (self *base_symbol) String() string {
	return fmt.Sprintf("symbol<%s>", self.Type().Name())
}

func (self *base_symbol) SymbolName() string {
	return self.String()
}
