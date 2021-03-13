package main

import (
	"errors"
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

type Expression struct {
	Operator Operator `parser:"\"[\" @Operator"`
	N        uint     `parser:"      @Positive"`
	Sets     []*Set   `parser:"      @@+ \"]\""`
}

type Operator string

const (
	OpEQ Operator = "EQ"
	OpLE Operator = "LE"
	OpGR Operator = "GR"
)

type Set struct {
	File          *string     `parser:"  @File"`
	SubExpression *Expression `parser:"| @@"`
}

func (s *Set) String() string {
	var res string

	if s.File != nil {
		res += *s.File
	} else if s.SubExpression != nil {
		res += fmt.Sprintf("%+v", s.SubExpression)
	}

	return res
}

type Parser struct {
	parser *participle.Parser
}

func NewParser() *Parser {
	l := stateful.MustSimple([]stateful.Rule{
		{Name: "Positive", Pattern: `^[1-9]\d*`},
		{Name: "File", Pattern: `[a-zA-Z]\.\w*`},
		{Name: "Operator", Pattern: `EQ|LE|GR`},
		{Name: "Bracket", Pattern: `\[|\]`},
		{Name: "Whitespace", Pattern: `[ \t\n\r]+`},
	})

	parser := participle.MustBuild(&Expression{}, participle.Lexer(l), participle.Elide("Whitespace"))

	return &Parser{parser: parser}
}

func (p *Parser) Parse(s string) (*Expression, error) {
	if p.parser == nil {
		return nil, errors.New("parser must be created with NewParser func")
	}

	out := &Expression{}
	if err := p.parser.ParseString("", s, out); err != nil {
		return nil, fmt.Errorf("failed to parse string: %w", err)
	}

	return out, nil
}
