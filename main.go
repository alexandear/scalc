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
		log.Panic(err)
	}

	log.Printf("%+v", expr)

	fa, err := os.Open("test/a.txt")
	if err != nil {
		log.Panicf("open file: %v", err)
	}

	defer func() {
		if errClose := fa.Close(); errClose != nil {
			log.Printf("failed to close file: %v", errClose)
		}
	}()

	intsA, err := ReadInts(fa)
	if err != nil {
		log.Panic(err)
	}

	log.Println(intsA)
}
