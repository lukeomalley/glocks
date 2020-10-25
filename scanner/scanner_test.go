package scanner

import (
	"testing"

	"github.com/lukeomalley/glocks/token"
)

func TestScanner(t *testing.T) {
	input := `
		var a = 10;
		var b = 20;
		var c = a + b;
		print c;
	`

	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
	}{
		{token.VAR, "var"},
		{token.IDENTIFIER, "a"},
		{token.EQUAL, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "b"},
		{token.EQUAL, "="},
		{token.NUMBER, "20"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "c"},
		{token.EQUAL, "="},
		{token.IDENTIFIER, "a"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "b"},
		{token.SEMICOLON, ";"},
		{token.PRINT, "print"},
		{token.IDENTIFIER, "c"},
		{token.SEMICOLON, ";"},
	}

	scanner := NewScanner(input)
	tokens := scanner.ScanTokens()

	if len(tests) != len(tokens)-1 {
		t.Fatalf("tests - Number of tokens is wrong. Expected: %d, but got %d", len(tests), len(tokens))
	}

	for i, tt := range tests {
		if tt.expectedType != tokens[i].Type {
			t.Fatalf("[test %d] - Token type wrong. Expected: %q, but got %q", i, tt.expectedType, tokens[i].Type)
		}

		if tt.expectedLexeme != tokens[i].Lexeme {
			t.Fatalf("[test %d] - Token lexeme wrong. Expected: %q, but got %q", i, tt.expectedLexeme, tokens[i].Lexeme)
		}
	}

	if tokens[len(tokens)-1].Type != token.EOF {
		t.Fatalf("tests - Final token is wrong. Expected %q, but got %q", token.EOF, tokens[len(tokens)-1])
	}

}
