package main

import (
	"log"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

type OpType string

const (
	OpEQ OpType = "EQ"
	OpLE OpType = "LE"
	OpGR OpType = "GR"
)

func LexerOperator() string {
	operations := []OpType{OpEQ, OpLE, OpGR}

	var res string
	for i, op := range operations {
		res += string(op)

		if i != len(operations)-1 {
			res += "|"
		}
	}

	return res
}

type Expression struct {
	Set Set `"[" @@+ "]"`
}

type N struct {
	PositiveNumber int `@PositiveNumber`
}

type File struct {
	File string `@File`
}

type Operation struct {
	Type OpType `@Operator`
}

type Set struct {
	Operation Operation `@@`
	N         N         `@@`
	Files     []File    `@@+`
}

func main() {
	l := stateful.MustSimple([]stateful.Rule{
		{"PositiveNumber", `[0-9]\d*`, nil},
		{"File", `[a-zA-Z]\.\w*`, nil},
		{"Operator", LexerOperator(), nil},
		{"Bracket", `\[|\]`, nil},
		{"Whitespace", `[ \t\n\r]+`, nil},
	})
	parser := participle.MustBuild(&Expression{}, participle.Lexer(l), participle.Elide("Whitespace"))

	out := &Expression{}
	if err := parser.ParseString("", `[LE 2 d.txt [ GR 2 a.txt b.txt c.txt ] ]`, out); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", *out)
}
