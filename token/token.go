package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)
