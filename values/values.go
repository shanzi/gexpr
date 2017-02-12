package values

import "github.com/shanzi/gexpr/types"

type Value interface {
	Type() types.Type

	Int64() int64
	Float64() float64
	Bool() bool
	String() string
}
