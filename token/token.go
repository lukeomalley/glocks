package token

import "fmt"

// Type describes the type of lexeme as a string
type Type string

const (
	// Single character tokens
	LEFTPAREN  = "("
	RIGHTPAREN = ")"
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"
	COMMA      = ","
	DOT        = "."
	MINUS      = "-"
	PLUS       = "+"
	STAR       = "*"
	SLASH      = "/"
	SEMICOLON  = ";"
	QMARK      = "?"
	COLON      = ":"
	EQUAL      = "="
	BANG       = "!"
	GREATER    = ">"
	LESS       = "<"

	// Two character tokens
	EQUALEQUAL   = "=="
	BANGEQUAL    = "!="
	GREATEREQUAL = ">="
	LESSEQUAL    = "<="
	POWER        = "**"

	// Literals
	IDENTIFIER = "IDENT"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// Keywords
	AND      = "and"
	OR       = "or"
	CLASS    = "class"
	IF       = "if"
	ELSE     = "else"
	TRUE     = "true"
	FALSE    = "false"
	FUN      = "fun"
	FOR      = "for"
	NIL      = "nil"
	PRINT    = "print"
	RETURN   = "return"
	SUPER    = "super"
	THIS     = "this"
	VAR      = "var"
	WHILE    = "while"
	BREAK    = "break"
	CONTINUE = "continue"
	EOF      = "eof"
	INVALID  = "__invalid__"
)

// Token contains the lexeme read by the scanner
type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t *Token) String() string {
	return fmt.Sprintf("%s, %s %v", t.Type, t.Lexeme, t.Literal)
}
