package values

import (
	"fmt"
	"strconv"

	"github.com/shanzi/gexpr/types"
)

type String string

func (self String) Type() types.Type {
	return types.STRING
}
func (self String) Int64() int64 {
	if v, e := strconv.ParseInt(self.String(), 10, 64); e == nil {
		return v
	} else {
		panic(fmt.Sprintf("Can not convert '%s' to int64", self.String()))
	}
}

func (self String) Float64() float64 {
	if v, e := strconv.ParseFloat(self.String(), 64); e == nil {
		return v
	} else {
		panic(fmt.Sprintf("Can not convert '%s' to float64", self.String()))
	}
}

func (self String) Bool() bool {
	return len(self.String()) > 0
}
func (self String) String() string {
	return string(self)
}
