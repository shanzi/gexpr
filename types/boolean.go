package types

type _boolean string

const BOOLEAN _boolean = "boolean"

func (self _boolean) Name() string {
	return string(self)
}

func (self _boolean) Match(that Type) bool {
	return self.Name() == that.Name()
}

func (self _boolean) Equals(that Type) bool {
	return self == that
}
