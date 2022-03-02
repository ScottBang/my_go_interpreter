package lexer

import (
	"testing"

	"monkey01/token"
)

type dbg_token struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	// tests := []struct {
	// 	expectedType    token.TokenType
	// 	expectedLiteral string
	// }{
	// 	{token.ASSIGN, "="},
	// 	{token.PLUS, "+"},
	// 	{token.LPAREN, "("},
	// 	{token.RPAREN, ")"},
	// 	{token.LBRACE, "{"},
	// 	{token.RBRACE, "}"},
	// 	{token.COMMA, ","},
	// 	{token.SEMICOLON, ";"},
	// 	{token.EOF, ""},
	// }

	// input 의 문자, 단어(토큰)과 일치하게 만들어야 테스트가 통과 됨.
	tests2 := []dbg_token{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests2 {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
