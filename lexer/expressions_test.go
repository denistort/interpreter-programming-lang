package lexer

import (
	"interpreter/token"
	"testing"
)

func TestExpressions(t *testing.T) {
	textStream := `
	let isMore = 5 < 2;
	`
	tests := []ExpectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "isMore"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.LessThan, "<"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
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
