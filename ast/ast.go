package ast

import (
	"fmt"
	"strings"

	"github.com/lukeomalley/glocks/token"
)

// Node is the root class of ast nodes
type Node interface {
	String() string
}

// Expr is the root class of expression nodes
type Expr interface {
	Node
}

// Binary is used for binary oprations. Ex: 2 + 2
type Binary struct {
	Expr
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (b *Binary) String() string {
	var out strings.Builder
	out.WriteString("(")
	out.WriteString(b.Left.String())
	out.WriteString(" ")
	out.WriteString(b.Operator.Lexeme)
	out.WriteString(" ")
	out.WriteString(b.Right.String())
	out.WriteString(")")
	return out.String()
}

// Grouping is used for grouping expresisons. Ex: (1 + 2) * 3
type Grouping struct {
	Expr
	Expression Expr
}

func (g *Grouping) String() string {
	var out strings.Builder
	out.WriteString("(")
	out.WriteString(g.Expression.String())
	out.WriteString(")")

	return out.String()
}

// Unary is used for unary expressions. Ex: -1
type Unary struct {
	Expr
	Operator token.Token
	Right    Expr
}

func (u *Unary) String() string {
	var out strings.Builder
	out.WriteString(u.Operator.Lexeme)
	out.WriteString(u.Right.String())
	return out.String()
}

// Literal is used for literal values such as numbers and strings
type Literal struct {
	Expr
	Value interface{}
}

func (l *Literal) String() string {
	var out strings.Builder
	out.WriteString(fmt.Sprintf("%v", l.Value))
	return out.String()
}
