package ast

import "github.com/guerilla/token"

// Node interface is implemented by every node in our abstract syntax tree
type Node interface {
	TokenLiteral() string
}

// Statement -
type Statement interface {
	Node
	statementNode()
}

// Expression -
type Expression interface {
	Node
	ExpressionNode()
}

// Program -
type Program struct {
	Statements []Statement
}

//TokenLiteral -
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement -
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {
}

//TokenLiteral -
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

//Identifier -
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {
}

//TokenLiteral -
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
