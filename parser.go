package goparser

import (
	lex "github.com/bbuck/go-lexer"
)

type NodeType int

type AST struct {
	Children    []*AST
	ValueType   NodeType
	ValueString string
	Parent      *AST
}

type Parser struct {
	Tokens []*lex.Token
	Cur    int
	AST    *AST
}

func (p *Parser) Current() *lex.Token {
	if p.Cur < len(p.Tokens) {
		return p.Tokens[p.Cur]
	}
	return nil
}
func (p *Parser) Next() {
	p.Cur++
}

func (p *Parser) HasTokens() bool {
	return p.Cur < len(p.Tokens)
}

func (p *Parser) AddChild(ast *AST) {
	ast.Parent = p.AST
	p.AST.Children = append(p.AST.Children, ast)
}

type ParseFunc func(*Parser) ParseFunc
