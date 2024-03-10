package lexer

import (
	"interpreter/token"
	"regexp"
)

type Lexer struct {
	textStream   string
	position     int
	readPosition int
	character    byte
}

func New(textStream string) *Lexer {
	lexer := &Lexer{textStream: textStream}
	lexer.readCharacter()
	return lexer
}

func (lexer *Lexer) readCharacter() {
	if lexer.readPosition >= len(lexer.textStream) {
		lexer.character = 0
	} else {
		lexer.character = lexer.textStream[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}
func (lexer *Lexer) NextToken() token.Token {
	var nextToken token.Token
	lexer.skipWhiteSpace()
	switch lexer.character {
	// Operators
	case '=':
		nextToken = newToken(token.ASSIGN, lexer.character)
	case '+':
		nextToken = newToken(token.PLUS, lexer.character)
	case '-':
		nextToken = newToken(token.MINUS, lexer.character)
	case '*':
		nextToken = newToken(token.MULTIPLY, lexer.character)
	case '/':
		nextToken = newToken(token.DIVIDE, lexer.character)
	case '>':
		nextToken = newToken(token.GraterThan, lexer.character)
	case '<':
		nextToken = newToken(token.LessThan, lexer.character)
	case '!':
		nextToken = newToken(token.NOT, lexer.character)
	case ';':
		nextToken = newToken(token.SEMICOLON, lexer.character)
	case '(':
		nextToken = newToken(token.LPAREN, lexer.character)
	case ')':
		nextToken = newToken(token.RPAREN, lexer.character)
	case '{':
		nextToken = newToken(token.LBRACE, lexer.character)
	case '}':
		nextToken = newToken(token.RBRACE, lexer.character)
	case ',':
		nextToken = newToken(token.COMMA, lexer.character)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		if isLetter(lexer.character) {
			nextToken.Literal = lexer.readIdentifier()
			nextToken.Type = token.LookupIdent(nextToken.Literal)
			return nextToken
		}
		if IsDigit(lexer.character) {
			nextToken.Literal = lexer.readNumber()
			nextToken.Type = token.INT
			return nextToken
		} else {
			nextToken = newToken(token.ILLEGAL, lexer.character)
		}
	}
	lexer.readCharacter()
	return nextToken
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.character) {
		lexer.readCharacter()
	}
	return lexer.textStream[position:lexer.position]
}
func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for IsDigit(lexer.character) {
		lexer.readCharacter()
	}
	return lexer.textStream[position:lexer.position]
}

func (lexer *Lexer) skipWhiteSpace() {
	reg := regexp.MustCompile(`(?m)\s`)
	for reg.MatchString(string(lexer.character)) {
		lexer.readCharacter()
	}
}

/**
utilities
*/

func newToken(tokenType token.Type, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func isLetter(char byte) bool {
	return regexp.MustCompile(`(?m)[a-zA-Z]`).MatchString(string(char)) || char == '_'
}

func IsDigit(character byte) bool {
	return regexp.MustCompile(`(?m)\d`).MatchString(string(character))
}
