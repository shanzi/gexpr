package types

type _string string

const STRING _string = "string"

func (self _string) Name() string {
	return string(self)
}

func (self _string) Match(that Type) bool {
	return self.Name() == that.Name()
}

func (self _string) Equals(that Type) bool {
	return self == that
}
