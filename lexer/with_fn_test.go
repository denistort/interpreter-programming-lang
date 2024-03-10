package lexer

import (
	"interpreter/token"
	"testing"
)

func TestWithFunctions(t *testing.T) {
	textStream := `
	let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	}
	let result = add(five, ten);
	`

	tests := []ExpectedToken{
		// first row
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// second row
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// third row
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		// last row
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
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
