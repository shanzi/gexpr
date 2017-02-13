package gexpr

import (
	"fmt"
	"testing"
)

func TestParseArthmeticExpr(t *testing.T) {
	if v, err := Parse("1 * (2 + 3) + var1 / var2"); err != nil {
		t.Error("Parse failed:", err)
	} else if s := v.String(); s != "(+ (* 1 (+ 2 3)) (/ var1 var2))" {
		fmt.Println(s)
		t.Error("Wrong parse result:", s)
	}
}

func TestParseBooleanExpr(t *testing.T) {
	if v, err := Parse("(!var1) && (1 > 2) || (10 <= var2)"); err != nil {
		t.Error("Parse failed:", err)
	} else if s := v.String(); s != "(|| (&& (! var1) (> 1 2)) (<= 10 var2))" {
		fmt.Println(s)
		t.Error("Wrong parse result:", s)
	}
}
