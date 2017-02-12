package values

import (
	"strconv"

	"github.com/shanzi/gexpr/types"
)

type Integer int64

func (self Integer) Type() types.Type {
	return types.INTEGER
}

func (self Integer) Int64() int64 {
	return int64(self)
}

func (self Integer) Float64() float64 {
	return float64(self.Int64())
}

func (self Integer) Bool() bool {
	return self.Int64() != 0
}

func (self Integer) String() string {
	return strconv.FormatInt(int64(self), 10)
}
