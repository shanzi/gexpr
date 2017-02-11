package types

import (
	"fmt"
	"sort"
	"strings"
)

type UnionTypes interface {
	Type

	Union(that UnionTypes) UnionTypes
	Subtract(that UnionTypes) UnionTypes
	Intersect(that UnionTypes) UnionTypes

	Size() int
	TypeList() []Type
	Contains(that Type) bool
}

type _union []Type

func NewUnion(subtypes ...Type) UnionTypes {
	types := make([]Type, len(subtypes))
	for _, subtype := range subtypes {
		if _, ok := subtype.(UnionTypes); ok {
			panic("UnionTypes can not be nested!")
		}
		if !_union(types).Contains(subtype) {
			types = append(types, subtype)
		}
	}
	return _union(types)
}

func (self _union) Name() string {
	types := ([]Type)(self)
	names := make([]string, len(types))
	for _, t := range types {
		names = append(names, t.Name())
	}
	sort.Strings(names)
	return fmt.Sprintf("union_type(%s)", strings.Join(names, ", "))
}

func (self _union) Match(that Type) bool {
	types := ([]Type)(self)
	if thatunion, ok := that.(UnionTypes); ok {
		// `that` is a union type
		for _, subtype := range thatunion.TypeList() {
			if !self.Match(subtype) {
				return false
			}
		}
		return true
	} else {
		// `that` is not a union type
		for _, t := range types {
			if t.Match(that) {
				return true
			}
		}

		return false
	}
}

func (self _union) Equals(that Type) bool {

	if thatunion, ok := that.(UnionTypes); ok {
		if self.Size() != thatunion.Size() {
			return false
		}

		for _, t := range self.TypeList() {
			if !thatunion.Contains(t) {
				return false
			}
		}

		return true
	}

	return false
}

func (self _union) Union(that UnionTypes) UnionTypes {
	types := make([]Type, self.Size()+that.Size())
	for _, t := range self.TypeList() {
		types = append(types, t)
	}
	for _, t := range that.TypeList() {
		if !self.Contains(t) {
			types = append(types, t)
		}
	}
	return _union(types)
}

func (self _union) Subtract(that UnionTypes) UnionTypes {
	types := make([]Type, self.Size())
	for _, t := range self.TypeList() {
		if !that.Contains(t) {
			types = append(types, t)
		}
	}
	return _union(types)
}

func (self _union) Intersect(that UnionTypes) UnionTypes {
	types := make([]Type, self.Size())
	for _, t := range self.TypeList() {
		if that.Contains(t) {
			types = append(types, t)
		}
	}
	return _union(types)
}

func (self _union) Size() int {
	return len(([]Type)(self))
}

func (self _union) TypeList() []Type {
	return ([]Type)(self)
}

func (self _union) Contains(that Type) bool {
	for _, t := range self.TypeList() {
		if t.Equals(that) {
			return true
		}
	}
	return false
}
