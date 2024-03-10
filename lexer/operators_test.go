package lexer

import (
	"interpreter/token"
	"testing"
)

func TestOperators(t *testing.T) {
	textStream := `
	+
	-
	*
	/
	=
	>
	<
	!
	`

	tests := []ExpectedToken{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVIDE, "/"},
		{token.ASSIGN, "="},
		{token.GraterThan, ">"},
		{token.LessThan, "<"},
		{token.NOT, "!"},
		{token.EOF, ""},
	}
	lexer := New(textStream)
	for i, test := range tests {
		tokenFromLexer := lexer.NextToken()
		if tokenFromLexer.Type != test.expectedType {
			t.Fatalf("tests[%d] - token type wrong expected=%q got=%q", i, test.expectedType, tokenFromLexer.Type)
		}

		if tokenFromLexer.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong expected=%q got=%q", i, test.expectedLiteral, tokenFromLexer.Literal)
		}
	}
}
