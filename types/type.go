package types

type Type interface {
	Name() string
	Match(that Type) bool
	Equals(that Type) bool
}
