package values

import (
	"errors"

	"github.com/shanzi/gexpr/types"
)

type Func interface {
	Value

	Accept(values []Value) bool
	AcceptTypes(tps []types.Type) bool

	Call(values []Value) Value
}

type _func struct {
	type_    types.Func
	callback func([]interface{}) interface{}
}

func PackFunc(tps []types.Type, ret types.Type, f func([]interface{}) interface{}) (Func, error) {
	if types.Primitive().Match(ret) {
		return &_func{types.NewFunc(tps, ret), f}, nil
	} else {
		return nil, errors.New("Can not pack func returns non primitive types")
	}
}

func (self *_func) Type() types.Type {
	return self.type_
}

func (self *_func) Int64() int64 {
	panic("Not Implemented!")
}

func (self *_func) Float64() float64 {
	panic("Not Implemented!")
}

func (self *_func) Bool() bool {
	panic("Not Implemented!")
}

func (self *_func) String() string {
	return self.type_.Name()
}

func (self *_func) Accept(values []Value) bool {
	argtypes := make([]types.Type, len(values))
	for _, t := range values {
		argtypes = append(argtypes, t.Type())
	}
	return self.AcceptTypes(argtypes)
}

func (self *_func) AcceptTypes(tps []types.Type) bool {
	return types.AssertSliceMatch(self.type_.ArgumentTypes(), tps)
}

func (self *_func) Call(values []Value) Value {
	if !self.Accept(values) {
		panic("Illegal arguments")
	}

	if args, err := UnpackSlice(values); err != nil {
		panic("Cannot unpack arguments")
	} else if ret, err := Pack(self.callback(args)); err != nil {
		panic("Function Calling failed")
	} else {
		return ret
	}
}
