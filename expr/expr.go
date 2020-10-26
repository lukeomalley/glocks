package expr

import (
	"github.com/lukeomalley/glocks/token"
)

// Visitor ...
type Visitor interface {
	VisitBinaryExpr(b *Binary) interface{}
	VisitUnaryExpr(u *Unary) interface{}
	VisitGroupingExpr(g *Grouping) interface{}
	VisitLiteralExpr(l *Literal) interface{}
}

// Expr is the root class of expression nodes
type Expr interface {
	Accept(v Visitor) interface{}
}

// Binary is used for binary oprations. Ex: 2 + 2
type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

// Accept routes the visitor to the correct type
func (b *Binary) Accept(v Visitor) interface{} {
	return v.VisitBinaryExpr(b)
}

// Grouping is used for grouping expresisons. Ex: (1 + 2) * 3
type Grouping struct {
	Expr Expr
}

// Accept routes the visitor to the correct type
func (g *Grouping) Accept(v Visitor) interface{} {
	return v.VisitGroupingExpr(g)
}

// Unary is used for unary expressions. Ex: -1
type Unary struct {
	Operator token.Token
	Right    Expr
}

// Accept routes the visitor to the correct type
func (u *Unary) Accept(v Visitor) interface{} {
	return v.VisitUnaryExpr(u)
}

// Literal is used for literal values such as numbers and strings
type Literal struct {
	Value interface{}
}

// Accept routes the visitor to the correct type
func (l *Literal) Accept(v Visitor) interface{} {
	return v.VisitLiteralExpr(l)
}
