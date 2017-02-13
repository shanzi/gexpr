package gexpr

import (
	"testing"

	"github.com/shanzi/gexpr/types"
	"github.com/shanzi/gexpr/values"
)

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

func TestEvaluateIntegerVars(t *testing.T) {
	params := map[string]interface{}{
		"var1": 10,
		"var2": 2,
	}
	exp, _ := Parse("((var1 + 2) - 4 * 5) / var2")

	if value, err := Evaluate(exp, params); err == nil {
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

func TestEvaluateFloatVars(t *testing.T) {
	params := map[string]interface{}{
		"var1": 10.0,
		"var2": 2.0,
	}
	exp, _ := Parse("((var1 + 2.0) - 4.0 * 5.0) / var2")

	if value, err := Evaluate(exp, params); err == nil {
		if v, ok := value.(float64); !ok || v != -4.0 {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestEvaluateStringVars(t *testing.T) {
	params := map[string]interface{}{
		"var1": "aaaa",
		"var2": "cccc",
	}
	exp, _ := Parse("var1 + `bbbb` + var2")

	if value, err := Evaluate(exp, params); err == nil {
		if v, ok := value.(string); !ok || v != "aaaabbbbcccc" {
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

func TestEvaluateTypeInteger(t *testing.T) {
	params := map[string]interface{}{
		"var1": types.INTEGER,
		"var2": types.INTEGER,
	}
	exp, _ := Parse("(var1 + 2) / (var2 - 10)")

	if value, err := EvaluateType(exp, params); err == nil {
		if !types.INTEGER.Match(value) {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestEvaluateTypeBoolean(t *testing.T) {
	params := map[string]interface{}{
		"var1": types.INTEGER,
		"var2": types.BOOLEAN,
	}
	exp, _ := Parse("(var1 > 2) || (var2 && false)")

	if value, err := EvaluateType(exp, params); err == nil {
		if !types.BOOLEAN.Match(value) {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestEvaluateTypeString(t *testing.T) {
	params := map[string]interface{}{
		"var1": types.STRING,
		"var2": types.STRING,
	}
	exp, _ := Parse("var1 + var2")

	if value, err := EvaluateType(exp, params); err == nil {
		if !types.STRING.Match(value) {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestCallFunction(t *testing.T) {
	sum, _ := values.PackFunc(
		[]types.Type{types.INTEGER, types.INTEGER, types.INTEGER},
		types.INTEGER,
		func(args []interface{}) interface{} {
			var res int64 = 0
			for _, v := range args {
				if n, ok := v.(int64); ok {
					res += n
				}
			}
			return res
		})
	params := map[string]interface{}{
		"sum": sum,
		"v1":  1,
		"v2":  2,
	}
	exp, _ := Parse("sum(v1, v2, 3) * 4")

	if value, err := Evaluate(exp, params); err == nil {
		if v, ok := value.(int64); !ok || v != 24 {
			t.Error("Incorrect value: ", value)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}

func TestCallFunctionType(t *testing.T) {
	sum, _ := values.PackFunc(
		[]types.Type{types.INTEGER, types.INTEGER, types.INTEGER},
		types.INTEGER,
		func(args []interface{}) interface{} {
			var res int64 = 0
			for _, v := range args {
				if n, ok := v.(int64); ok {
					res += n
				}
			}
			return res
		})
	params := map[string]interface{}{
		"sum": sum,
		"v1":  types.INTEGER,
		"v2":  types.INTEGER,
	}
	exp, _ := Parse("sum(v1, v2, 3) * 4")

	if tp, err := EvaluateType(exp, params); err == nil {
		if !types.INTEGER.Match(tp) {
			t.Error("Incorrect value: ", tp)
		}
	} else {
		t.Error("Cannot evaluate expression:", err)
	}
}
