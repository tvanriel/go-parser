package goparser

import (
	"errors"

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
	Error  error
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

func (p *Parser) Err() error {
	return p.Error
}

// Whether the next token can be parsed.
func (p *Parser) Continue() bool {
	return p.Err() == nil && p.HasTokens()
}

// Set the error flag on the parser.
func (p *Parser) SetError(err string) {
	p.Error = errors.New(err)
}

func (p *Parser) AddChild(ast *AST) {
	ast.Parent = p.AST
	p.AST.Children = append(p.AST.Children, ast)
}

type ParseFunc func(*Parser) ParseFunc
