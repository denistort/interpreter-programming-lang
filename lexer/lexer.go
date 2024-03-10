package lexer

import "interpreter/token"

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
	switch lexer.character {
	case '=':
		nextToken = newToken(token.ASSIGN, lexer.character)
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
	case '+':
		nextToken = newToken(token.PLUS, lexer.character)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	}
	lexer.readCharacter()
	return nextToken
}

func newToken(tokenType token.Type, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}
