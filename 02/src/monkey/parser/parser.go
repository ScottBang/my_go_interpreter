package parser

import (
	"monkey02/lexer"
	"monkey02/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}
