package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface{ Node }
type Expression interface{ Node }

type Program struct{ Statements []Statement }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type SowStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (s *SowStatement) TokenLiteral() string { return s.Token.Literal }
func (s *SowStatement) String() string {
	return s.TokenLiteral() + " " + s.Name.String() + " = " + s.Value.String() + ";"
}

// ... other AST nodes for ExpressionStatement, Identifier, IntegerLiteral,
// PrefixExpression, InfixExpression, IfExpression (SproutExpression),
// BlockStatement, FunctionLiteral (PlantLiteral), CallExpression, LoopExpression, HarvestStatement
