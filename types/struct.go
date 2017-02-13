package types

import (
	"fmt"
	"sort"
	"strings"
)

type Struct interface {
	Type

	Size() int
	FieldMapping() map[string]Type
}

type _struct map[string]Type

func NewStruct(mapping map[string]Type) Struct {
	return _struct(mapping)
}

func (self _struct) Name() string {
	structmap := (map[string]Type)(self)
	fields := make([]string, 0, len(structmap))
	for k, v := range structmap {
		fields = append(fields, fmt.Sprintf("%s: %s", k, v.Name()))
	}
	sort.Strings(fields)
	return fmt.Sprintf("struct{%s}", strings.Join(fields, ", "))
}

func (self _struct) Match(that Type) bool {
	if thatstruct, ok := that.(Struct); ok {
		if self.Size() != thatstruct.Size() {
			return false
		}

		mapping := (map[string]Type)(self)
		thatmapping := thatstruct.FieldMapping()

		for n, t := range mapping {
			if thattype, ok := thatmapping[n]; ok {
				if !t.Match(thattype) {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}
	return false
}

func (self _struct) Equals(that Type) bool {
	if thatstruct, ok := that.(Struct); ok {
		if self.Size() != thatstruct.Size() {
			return false
		}

		thatmapping := thatstruct.FieldMapping()
		for n, t := range self.FieldMapping() {
			if !t.Equals(thatmapping[n]) {
				return false
			}
		}

		return true
	}

	return false
}

func (self _struct) Size() int {
	return len((map[string]Type)(self))
}

func (self _struct) FieldMapping() map[string]Type {
	return (map[string]Type)(self)
}
