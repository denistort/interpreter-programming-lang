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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
	ELSE     = "ELSE"
	IF       = "if"
	// Operators
	ASSIGN     = "="
	PLUS       = "+"
	GraterThan = ">"
	LessThan   = "<"
	NOT        = "!"
	MINUS      = "-"
	MULTIPLY   = "*"
	DIVIDE     = "/"
	EQUALS     = "EQUALS"
	NOT_EQUALS = "NOT_EQUALS"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

var Keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) Type {
	if tok, ok := Keywords[ident]; ok {
		return tok
	} else {
		return IDENTIFIER
	}
}
