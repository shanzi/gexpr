package values

import (
	"strconv"

	"github.com/shanzi/gexpr/types"
)

type Float float64

func (self Float) Type() types.Type {
	return types.FLOAT
}

func (self Float) Int64() int64 {
	return int64(self.Float64())
}

func (self Float) Float64() float64 {
	return float64(self)
}

func (self Float) Bool() bool {
	value := self.Float64()
	eps := 1E-6
	return value > -eps && value < eps
}

func (self Float) String() string {
	return strconv.FormatFloat(float64(self), 'g', -1, 64)
}
