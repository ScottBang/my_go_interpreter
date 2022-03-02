package lexer

import "monkey01/token"

type Lexer struct {
	input        string // full string
	position     int    // current char position
	readPosition int    // next char position
	ch           byte   // current char
}

func New(input string) *Lexer {
	// input 문자열로 Lexer 객체를 초기화 한다.
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// Lexer.ch 에 Lexer.position 이 가리키고 있는 문자를 지정한다.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	//
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	// 빈 문자를 skip 하는 처리.
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// func (l *Lexer) peekChar() byte {
// 	// readPosition의 문자를 byte로 리턴하는 처리.
// 	if l.readPosition >= len(l.input) {
// 		return 0
// 	} else {
// 		return l.input[l.readPosition]
// 	}
// }

// func isLetter(ch byte) bool {
// 	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
// }

// func isDigit(ch byte) bool {
// 	return '0' <= ch && ch <= '9'
// }

// func (l *Lexer) readIdentifier() string {
// 	position := l.position
// 	for isLetter(l.ch) {
// 		l.readChar()
// 	}
// 	return l.input[position:l.position]
// }

// func (l *Lexer) readNumber() string {
// 	position := l.position
// 	for isDigit(l.ch) {
// 		l.readChar()
// 	}
// 	return l.input[position:l.position]
// }
