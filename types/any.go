package types

type _any string

const ANY _any = "any"

func (self _any) Name() string {
	return string(self)
}

func (self _any) Match(that Type) bool {
	return true
}

func (self _any) Equals(that Type) bool {
	return self == that
}
