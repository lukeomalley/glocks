package astprinter

import (
	"fmt"
	"strings"

	"github.com/lukeomalley/glocks/expr"
)

// AstPrinter pretty prints the ast structure
type AstPrinter struct{}

// Print prints the ast
func (a *AstPrinter) String(e expr.Expr) interface{} {
	return e.Accept(a)
}

// VisitBinaryExpr pretty prints a binary expression
func (a *AstPrinter) VisitBinaryExpr(b *expr.Binary) interface{} {
	return a.parenthesize(b.Operator.Lexeme, b.Left, b.Right)
}

// VisitUnaryExpr pretty prints a unary expression
func (a *AstPrinter) VisitUnaryExpr(u *expr.Unary) interface{} {
	return a.parenthesize(u.Operator.Lexeme, u.Right)
}

// VisitGroupingExpr pretty prints a grouping expression
func (a *AstPrinter) VisitGroupingExpr(g *expr.Grouping) interface{} {
	return a.parenthesize("group", g.Expr)
}

// VisitLiteralExpr pretty prints a literal expression
func (a *AstPrinter) VisitLiteralExpr(l *expr.Literal) interface{} {
	if l.Value == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", l.Value)
}

func (a *AstPrinter) parenthesize(name string, exprs ...expr.Expr) string {
	var out strings.Builder
	out.WriteString("( ")
	out.WriteString(name)

	for _, expr := range exprs {
		out.WriteString(" ")
		out.WriteString(expr.Accept(a).(string))
	}

	out.WriteString(" )")
	return out.String()
}
