package symbols

import (
	"fmt"

	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

type Func interface {
	values.Func
	SymbolName() string
}

type _func_symbol struct {
	func_obj values.Func
}

func PackFunc(f values.Func) Func {
	return &_func_symbol{f}
}

func (self *_func_symbol) Type() types.Type {
	return self.func_obj.Type()
}

func (self *_func_symbol) Int64() int64 {
	panic("Not implemented!")
}

func (self *_func_symbol) Float64() float64 {
	panic("Not implemented!")
}

func (self *_func_symbol) Bool() bool {
	panic("Not implemented!")
}

func (self *_func_symbol) String() string {
	return fmt.Sprintf("symbol(%s)", self.func_obj.String())
}

func (self *_func_symbol) Accept(values []values.Value) bool {
	return self.func_obj.Accept(values)
}

func (self *_func_symbol) AcceptTypes(tps []types.Type) bool {
	return self.func_obj.AcceptTypes(tps)
}

func (self *_func_symbol) Call(values []values.Value) values.Value {
	if !self.Accept(values) {
		panic("Illegal arguments")
	}

	if args, err := UnpackSlice(values); err != nil {
		panic("Cannot unpack arguments")
	} else if ft, ok := self.Type().(types.Func); !ok {
		panic("Can not check func return type")
	} else if ct := self.CallType(args); ft.Match(ct) {
		panic("Mismatched return type")
	} else if res, err := Pack(ct); err != nil {
		panic(err)
	} else {
		return res
	}
}

func (self *_func_symbol) CallType(tps []types.Type) types.Type {
	return self.func_obj.CallType(tps)
}

func (self *_func_symbol) SymbolName() string {
	return self.String()
}
