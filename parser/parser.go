package parser

import (
	"github.com/guerilla/ast"
	"github.com/guerilla/lexer"
	"github.com/guerilla/token"
)

//Parser ...
type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

//New -
func New(lex *lexer.Lexer) *Parser {
	p := &Parser{
		l:      lex,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

//nextToken -
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//ParseProgram -
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}
