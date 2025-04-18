package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curTok    token.Token
	peekTok   token.Token
	errors    []string
	prefixFns map[token.TokenType]prefixParseFn
	infixFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixFns = make(map[token.TokenType]prefixParseFn)
	p.infixFns = make(map[token.TokenType]infixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	// ... register other prefixes
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.curTok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case token.SOW:
		return p.parseSowStatement()
	case token.HARVEST:
		return p.parseHarvestStatement()
	default:
		return p.parseExpressionStatement()
	}
}
