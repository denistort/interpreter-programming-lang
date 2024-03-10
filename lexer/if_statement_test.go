package lexer

import (
	"interpreter/token"
	"testing"
)

func TestIfStatement(t *testing.T) {
	textStream := `
	let weqwe;
	if (1 > 2) {
		weqwe = 2;
	}

	if (1 != 2) {
		weqwe = 3;
	}
	`
	tests := []ExpectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "weqwe"},
		{token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.GraterThan, ">"},
		{token.INT, "2"},
		{token.RPAREN, ")"},

		{token.LBRACE, "{"},
		{token.IDENTIFIER, "weqwe"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.NOT_EQUALS, "!="},
		{token.INT, "2"},
		{token.RPAREN, ")"},

		{token.LBRACE, "{"},
		{token.IDENTIFIER, "weqwe"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

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
