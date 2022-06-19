package goparser

import (
	"errors"
	"fmt"
	"strings"

	lex "github.com/tvanriel/go-lexer"
)

type NodeType int

type AST struct {
	Children    []*AST
	ValueType   NodeType
	ValueString string
	Parent      *AST
}

func (a *AST) Print(i int) string {
	if a == nil {
		return ""
	}
	var sb strings.Builder

	sb.WriteString(strings.Repeat(" ", i))
	sb.WriteString(a.ValueString)
	sb.WriteString("\n")
	for _, child := range a.Children {
		sb.WriteString(child.Print(i + 4))
	}
	return sb.String()

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

func (p *Parser) Dump() {
	fmt.Print(p.AST.Print(0))
}
