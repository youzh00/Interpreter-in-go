package parser

import (
	"github.com/youzh00/Interpreter-in-go/ast"
	"github.com/youzh00/Interpreter-in-go/lexer"
	"github.com/youzh00/Interpreter-in-go/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: &lexer.Lexer{}}

	p.nextToken()
	p.nextToken()

	return p
}

// helper method that advances both curToken and peekToken
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
