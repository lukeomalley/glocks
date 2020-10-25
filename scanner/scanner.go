package scanner

import (
	"fmt"
	"strconv"

	"github.com/lukeomalley/glocks/parsererror"
	"github.com/lukeomalley/glocks/token"
)

// Scanner transforms the source into tokens
type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []token.Token
}

var keywords = map[string]token.Type{
	"and":      token.AND,
	"or":       token.OR,
	"class":    token.CLASS,
	"if":       token.IF,
	"else":     token.ELSE,
	"true":     token.TRUE,
	"false":    token.FALSE,
	"fun":      token.FUN,
	"for":      token.FOR,
	"nil":      token.NIL,
	"print":    token.PRINT,
	"return":   token.RETURN,
	"super":    token.SUPER,
	"this":     token.THIS,
	"var":      token.VAR,
	"while":    token.WHILE,
	"break":    token.BREAK,
	"continue": token.CONTINUE,
}

// NewScanner instantiates and returns a new scanner
func NewScanner(src string) *Scanner {
	return &Scanner{source: src, line: 1, tokens: make([]token.Token, 0)}
}

// ScanTokens transforms the source string into a slice of tokens
func (sc *Scanner) ScanTokens() []token.Token {
	for !sc.isAtEnd() {
		sc.start = sc.current // At the start of a new token
		sc.scanToken()
	}

	sc.tokens = append(sc.tokens, token.Token{Type: token.EOF})
	return sc.tokens
}

func (sc *Scanner) scanToken() {
	ch := sc.advance()

	switch ch {
	case '(':
		sc.addToken(token.RIGHTPAREN)
	case ')':
		sc.addToken(token.LEFTPAREN)
	case '{':
		sc.addToken(token.LEFTBRACE)
	case '}':
		sc.addToken(token.RIGHTBRACE)
	case ',':
		sc.addToken(token.COMMA)
	case '.':
		sc.addToken(token.DOT)
	case '-':
		sc.addToken(token.MINUS)
	case '+':
		sc.addToken(token.PLUS)
	case '?':
		sc.addToken(token.QMARK)
	case ':':
		sc.addToken(token.COLON)
	case ';':
		sc.addToken(token.SEMICOLON)
	case '*':
		if sc.nextCharIs('*') {
			sc.advance()
			sc.addToken(token.POWER)
		} else {
			sc.addToken(token.STAR)
		}
	case '>':
		if sc.nextCharIs('=') {
			sc.advance()
			sc.addToken(token.GREATEREQUAL)
		} else {
			sc.addToken(token.GREATER)
		}
	case '<':
		if sc.nextCharIs('=') {
			sc.advance()
			sc.addToken(token.LESSEQUAL)
		} else {
			sc.addToken(token.LESS)
		}
	case '=':
		if sc.nextCharIs('=') {
			sc.advance()
			sc.addToken(token.EQUALEQUAL)
		} else {
			sc.addToken(token.EQUAL)
		}
	case '!':
		if sc.nextCharIs('=') {
			sc.advance()
			sc.addToken(token.BANGEQUAL)
		} else {
			sc.addToken(token.BANG)
		}
	case '/':
		if sc.nextCharIs('/') {
			for !sc.nextCharIs('\n') || !sc.isAtEnd() {
				sc.advance()
			}
			sc.advance() // consume the newline char
		} else {
			sc.addToken(token.SLASH)
		}
	case '\n':
		sc.line++
	case ' ', '\t', '\r':
		// Ignore whitespace characters
	default:
		if sc.isDigit(ch) {
			sc.scanNumber()
		} else if sc.isAlpha(ch) {
			sc.scanIdentifier()
		} else {
			parsererror.LogMessage(sc.line, fmt.Sprintf("Unexpected character: %c", ch))
		}
	}
}

func (sc *Scanner) scanNumber() {
	// Scan the number
	for sc.isDigit(sc.peek()) {
		sc.advance()

		if sc.nextCharIs('.') && sc.isDigit(sc.peekNext()) {
			sc.advance() // consume the dot
			for sc.isDigit(sc.peek()) {
				sc.advance()
			}
		}
	}

	// Convert the number
	num, err := strconv.ParseFloat(sc.source[sc.start:sc.current], 64)
	if err != nil {
		panic("Invalid number format")
	}

	// Add the token
	sc.addTokenWithLiteral(token.NUMBER, num)
}

func (sc *Scanner) scanIdentifier() {
	// Scan the token
	for sc.isAlpha(sc.peek()) {
		sc.advance()
	}

	// Check if the lexeme is a keyword and add the token
	lexeme := sc.source[sc.start:sc.current]
	keyword, ok := keywords[lexeme]
	if ok {
		sc.addToken(keyword)
	} else {
		sc.addToken(token.IDENTIFIER)
	}
}

// =============================================================================
// Utility Functions
// =============================================================================

func (sc *Scanner) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (sc *Scanner) isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (sc *Scanner) addToken(t token.Type) {
	sc.addTokenWithLiteral(t, nil)
}

func (sc *Scanner) addTokenWithLiteral(t token.Type, literal interface{}) {
	sc.tokens = append(sc.tokens, sc.makeToken(t, literal))
}

func (sc *Scanner) makeToken(t token.Type, literal interface{}) token.Token {
	lexeme := sc.source[sc.start:sc.current]
	return token.Token{Type: t, Lexeme: lexeme, Literal: literal, Line: sc.line}
}

func (sc *Scanner) nextCharIs(ch byte) bool {
	return sc.source[sc.current+1] == ch
}

func (sc *Scanner) advance() byte {
	sc.current++
	return sc.source[sc.current-1]
}

func (sc *Scanner) peek() byte {
	if sc.isAtEnd() {
		return 0
	}

	return sc.source[sc.current]
}

func (sc *Scanner) peekNext() byte {
	if sc.current+1 > len(sc.source) {
		return 0
	}

	return sc.source[sc.current+1]
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.source)
}
