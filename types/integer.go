package types

type _integer string

const INTEGER _integer = "integer"

func (self _integer) Name() string {
	return string(self)
}

func (self _integer) Match(that Type) bool {
	return self.Name() == that.Name()
}

func (self _integer) Equals(that Type) bool {
	return self == that
}
