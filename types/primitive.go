package types

import "sync"

var primitive_once sync.Once
var primitive_singleton Type

func Primitive() Type {
	primitive_once.Do(func() {
		primitive_singleton = NewUnion(INTEGER, FLOAT, BOOLEAN, STRING)
	})
	return primitive_singleton
}
