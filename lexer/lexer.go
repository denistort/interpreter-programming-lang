package lexer

import (
	"interpreter/token"
	"regexp"
)

type Lexer struct {
	textStream       string
	position         int
	readPosition     int
	character        byte
	textStreamLength int
}

func New(textStream string) *Lexer {
	lexer := &Lexer{textStream: textStream, textStreamLength: len(textStream)}
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
	symbol := string(lexer.character)
	switch lexer.character {
	// Operators
	case '=':
		if lexer.peekChar() == '=' {
			previousSymbol := string(lexer.character)
			lexer.readCharacter()
			nextToken = newToken(token.EQUALS, previousSymbol+string(lexer.character))
		} else {
			nextToken = newToken(token.ASSIGN, symbol)
		}
	case '+':
		nextToken = newToken(token.PLUS, symbol)
	case '-':
		nextToken = newToken(token.MINUS, symbol)
	case '*':
		nextToken = newToken(token.MULTIPLY, symbol)
	case '/':
		nextToken = newToken(token.DIVIDE, symbol)
	case '>':
		nextToken = newToken(token.GraterThan, symbol)
	case '<':
		nextToken = newToken(token.LessThan, symbol)
	case '!':
		if lexer.peekChar() == '=' {
			previousCharacter := symbol
			lexer.readCharacter()
			nextToken = newToken(token.NOT_EQUALS, previousCharacter+string(lexer.character))
		} else {
			nextToken = newToken(token.NOT, symbol)
		}
	case ';':
		nextToken = newToken(token.SEMICOLON, symbol)
	case '(':
		nextToken = newToken(token.LPAREN, symbol)
	case ')':
		nextToken = newToken(token.RPAREN, symbol)
	case '{':
		nextToken = newToken(token.LBRACE, symbol)
	case '}':
		nextToken = newToken(token.RBRACE, symbol)
	case ',':
		nextToken = newToken(token.COMMA, symbol)
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
			nextToken = newToken(token.ILLEGAL, symbol)
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

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= lexer.textStreamLength {
		return 0
	} else {
		return lexer.textStream[lexer.readPosition]
	}
}

/**
utilities
*/

func newToken(tokenType token.Type, character string) token.Token {
	return token.Token{Type: tokenType, Literal: character}
}

func isLetter(char byte) bool {
	return regexp.MustCompile(`(?m)[a-zA-Z]`).MatchString(string(char)) || char == '_'
}

func IsDigit(character byte) bool {
	return regexp.MustCompile(`(?m)\d`).MatchString(string(character))
}
