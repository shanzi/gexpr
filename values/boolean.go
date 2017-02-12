package values

import (
	"strconv"

	"github.com/shanzi/gexpr/types"
)

type Boolean bool

func (self Boolean) Type() types.Type {
	return types.BOOLEAN
}

func (self Boolean) Int64() int64 {
	if self.Bool() {
		return 1
	} else {
		return 0
	}
}

func (self Boolean) Float64() float64 {
	return float64(self.Int64())
}

func (self Boolean) Bool() bool {
	return bool(self)
}

func (self Boolean) String() string {
	return strconv.FormatBool(bool(self))
}
