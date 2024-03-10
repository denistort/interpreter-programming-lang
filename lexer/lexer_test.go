package lexer

import (
	"fmt"
	"interpreter/token"
	"testing"
)

type ExpectedToken struct {
	expectedType    token.Type
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	textStream := "=+(){},;"

	tests := []ExpectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	fmt.Print(textStream, tests)
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
