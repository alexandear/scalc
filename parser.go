package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

type Expression struct {
	Operator Operator `"[" @Operator`
	N        uint     `@Positive`
	Sets     []*Set   `@@+ "]"`
}

type Operator int

const (
	OpEQ Operator = iota
	OpLE
	OpGR
)

var operators = map[string]Operator{ // nolint:gochecknoglobals // just once
	"EQ": OpEQ,
	"LE": OpLE,
	"GR": OpGR,
}

func (o *Operator) Capture(s []string) error { // nolint:unparam // not carry
	*o = operators[s[0]]

	return nil
}

func LexerOperator() string {
	var res string

	for str := range operators {
		res += str + "|"
	}

	return res[:len(res)-1]
}

type Set struct {
	File          *string     `  @File`
	SubExpression *Expression `| @@`
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

func Parse(s string) (*Expression, error) {
	l := stateful.MustSimple([]stateful.Rule{
		{Name: "Positive", Pattern: `[0-9]\d*`},
		{Name: "File", Pattern: `[a-zA-Z]\.\w*`},
		{Name: "Operator", Pattern: LexerOperator()},
		{Name: "Bracket", Pattern: `\[|\]`},
		{Name: "Whitespace", Pattern: `[ \t\n\r]+`},
	})

	parser := participle.MustBuild(&Expression{}, participle.Lexer(l), participle.Elide("Whitespace"))

	out := &Expression{}
	if err := parser.ParseString("", s, out); err != nil {
		return nil, fmt.Errorf("failed to parse string: %w", err)
	}

	return out, nil
}
