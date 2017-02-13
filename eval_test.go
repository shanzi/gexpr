package gexpr

import "testing"

func TestEvaluateInteger(t *testing.T) {
	exp, _ := Parse("((10 + 2) - 4 * 5) / 2")

	if value, err := Evaluate(exp, nil); err == nil {
		if v, ok := value.(int64); !ok || v != -4 {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestEvaluateFloat(t *testing.T) {
	exp, _ := Parse("((10.0 + 2.0) - 4.0 * 5.0) / 2.0")

	if value, err := Evaluate(exp, nil); err == nil {
		if v, ok := value.(float64); !ok || v != -4.0 {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestEvaluateBooleanTrue(t *testing.T) {
	exp, _ := Parse("(10 > 5) || (10 <= 5)")

	if value, err := Evaluate(exp, nil); err == nil {
		if v, ok := value.(bool); !ok || !v {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestEvaluateBooleanFalse(t *testing.T) {
	exp, _ := Parse("(10 > 5) && (10 <= 5)")

	if value, err := Evaluate(exp, nil); err == nil {
		if v, ok := value.(bool); !ok || v {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}
