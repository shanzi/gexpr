package gexpr

import (
	"fmt"
	"go/ast"
)

func Parse(exprStr string) (res Expr, err error) {
	defer func() {
		if r := recover(); r != nil {
			res = nil
			err = r
		}
	}()
	if expr, err := ast.Parse(exprStr); err != nil {
		return nil, err
	} else {
		root := transform(expr)
		return &_expr{root, exprStr}
	}
}

func transform(expr ast.Expr) ExprNode {
	switch typedexpr := expr.(type) {
	case ast.BasicLit:
		return NewLiteralNode(typedexpr.Kind, typedexpr.Value)
	case ast.BinaryExpr:
		return NewBinaryOperatorNode(
			typedexpr.Op,
			transform(typedexpr.X),
			transform(typedexpr.Y),
		)
	case ast.ParenExpr:
		return transform(typedexpr.X)
	case ast.Ident:
		return NewValueNode(typedexpr.Name)
	default:
		panic(fmt.Sprintf("Unsupported expr: %+v", expr))
	}
}
