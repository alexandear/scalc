package main

import (
	"fmt"
	"log"

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

var operators = map[string]Operator{
	"EQ": OpEQ,
	"LE": OpLE,
	"GR": OpGR,
}

func (o *Operator) Capture(s []string) error {
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
		{"Positive", `[0-9]\d*`, nil},
		{"File", `[a-zA-Z]\.\w*`, nil},
		{"Operator", LexerOperator(), nil},
		{"Bracket", `\[|\]`, nil},
		{"Whitespace", `[ \t\n\r]+`, nil},
	})
	parser := participle.MustBuild(&Expression{}, participle.Lexer(l), participle.Elide("Whitespace"))

	out := &Expression{}
	if err := parser.ParseString("", s, out); err != nil {
		return nil, err
	}

	return out, nil
}

func main() {
	expr, err := Parse(`[LE 2 a.txt b.txt c.txt [GR 2 d.txt] [EQ 3 e.txt]]`)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", expr)
}
