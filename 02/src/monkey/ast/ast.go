/*
--- monkey language parser 만들기 ---
AST: 소스코드를 내부적으로 표현할 때 쓰는 자료구조를 추상구문트리(Abstract Syntax Tree) 라고 부른다.

*/

package ast

import "monkey02/token"

// 기본 노드 인터페이스
type Node interface {
	TokenLiteral() string
}

// 선언구 노드 인터페이스
type Statement interface {
	Node
	statementNode()
}

// 표현식 노드 인터페이스
type Expression interface {
	Node
	expressionNode()
}

// Program 구조체
type Program struct {
	Statements []Statement
}

// Program 구조체가  Node 인터페이스를 구현
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LET 선언구 구조체
type LetStatement struct {
	// token.LET 토큰
	Token token.Token
	Name  *Identifier
	Value Expression
}

// LetStatement 구조체가 Statement 인터페이스, Node 인터페이스 구현
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// 식별자 구조체
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// Identifier 구조체가 Expression 인터페이스, Node 인터페이스 구현
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
