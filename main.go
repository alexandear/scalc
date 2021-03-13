package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

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

func ScanInt(scanner *bufio.Scanner) (int, error) {
	cont := scanner.Scan()
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner: %w", err)
	}

	if scanner.Text() == "" {
		return 0, io.EOF
	}

	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("not an integer: %w", err)
	}

	if !cont {
		return number, io.EOF
	}

	return number, nil
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var res []int

	for {
		n, err := ScanInt(scanner)
		if errors.Is(err, io.EOF) {
			return res, nil
		} else if err != nil {
			return []int{}, fmt.Errorf("scan: %w", err)
		}

		res = append(res, n)
	}
}

func main() {
	expr, err := Parse(`[ LE 3 a.txt b.txt c.txt [ GR 1 d.txt e.txt ] [ EQ 2 f.txt [ LE 1 g.txt h.txt ] ] ]`)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", expr)

	fa, err := os.Open("test/a.txt")
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer func() {
		if err := fa.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	intsA, err := ReadInts(fa)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(intsA)
}
