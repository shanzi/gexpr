package types

type _float string

const FLOAT _float = "float"

func (self _float) Name() string {
	return string(self)
}

func (self _float) Match(that Type) bool {
	return self.Name() == that.Name()
}

func (self _float) Equals(that Type) bool {
	return self == that
}
