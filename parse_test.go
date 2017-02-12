package gexpr

import "testing"

func TestParse(t *testing.T) {
	if v, err := Parse("1 * (2 + 3) + var1 / var2"); err != nil {
		t.Error("Parse failed:", err)
	} else if s := v.String(); s != "(+ (* 1 (+ 2 3)) (/ var1 var2))" {
		t.Error("Wrong parse result:", s)
	}
}
