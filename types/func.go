package types

import (
	"fmt"
	"strings"
)

type Func interface {
	Type

	ArgumentTypes() []Type
	ReturnType() Type
}

type _func struct {
	args []Type
	ret  Type
}

func NewFunc(argtypes []Type, rettype Type) Func {
	return &_func{argtypes, rettype}
}

func (self *_func) Name() string {
	argnames := make([]string, len(self.args))
	for _, v := range self.args {
		argnames = append(argnames, v.Name())
	}
	return fmt.Sprintf("func(%s) %s", strings.Join(argnames, ", "), self.ret.Name())
}

func (self *_func) Match(that Type) bool {
	if ftype, ok := that.(Func); ok {
		return AssertSliceEqual(self.args, ftype.ArgumentTypes()) && self.ret.Match(ftype.ReturnType())
	}
	return false
}

func (self *_func) Equals(that Type) bool {
	if ftype, ok := that.(Func); ok {
		return AssertSliceMatch(self.args, ftype.ArgumentTypes()) && self.ret.Equals(ftype.ReturnType())
	}
	return false
}

func (self *_func) ArgumentTypes() []Type {
	return self.args
}

func (self *_func) ReturnType() Type {
	return self.ret
}
