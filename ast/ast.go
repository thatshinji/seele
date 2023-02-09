package ast

import "seele/token"

//let <标识符> = 表达式

type Node interface {
	TokenLiteral() string
}

// Statement 语句节点
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式节点
type Expression interface {
	Node
	expressionNode()
}

// Program ast root element
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
