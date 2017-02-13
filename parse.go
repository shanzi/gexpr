package gexpr

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"

	"github.com/shanzi/gexpr/expr"
	"github.com/shanzi/gexpr/nodes"
)

func Parse(exprStr string) (res expr.Expr, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rerr, ok := r.(string); ok {
				err = errors.New(rerr)
			} else {
				err = errors.New("Can not parse expression")
			}
			res = nil
		}
	}()
	if exp, err := parser.ParseExpr(exprStr); err != nil {
		return nil, err
	} else {
		root := transform(exp)
		return expr.New(root, exprStr), nil
	}
}

func transform(expr ast.Expr) expr.ExprNode {
	switch typedexpr := expr.(type) {
	case *ast.BasicLit:
		return nodes.NewLiteralNode(int(typedexpr.Kind), typedexpr.Value)
	case *ast.BinaryExpr:
		return nodes.NewBinaryOperatorNode(
			int(typedexpr.Op),
			transform(typedexpr.X),
			transform(typedexpr.Y),
		)
	case *ast.UnaryExpr:
		return nodes.NewUnaryOperatorNode(
			int(typedexpr.Op),
			transform(typedexpr.X),
		)
	case *ast.CallExpr:
		return createFuncNode(typedexpr)
	case *ast.ParenExpr:
		return transform(typedexpr.X)
	case *ast.Ident:
		return nodes.NewValueNode(typedexpr.Name)
	default:
		panic(fmt.Sprintf("Unsupported expr: %+v", expr))
	}
}

func createFuncNode(exp *ast.CallExpr) expr.ExprNode {
	args := exp.Args
	argvalues := make([]expr.ExprNode, 0, len(args))
	for _, arg := range args {
		argvalues = append(argvalues, transform(arg))
	}
	return nodes.NewFuncNode(transform(exp.Fun), argvalues)
}
