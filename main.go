package main

import (
	"fmt"
	"os"
	"strings"
)

func Usage(app, reason string) string {
	var res strings.Builder

	res.WriteString(reason)
	res.WriteString("\nUsage:\n")
	res.WriteString("   " + app + " expression\n")
	res.WriteString("Where:\n")
	res.WriteString("   expression := [ operator N sets ]\n")
	res.WriteString("   sets := set | set sets\n")
	res.WriteString("   set := file | expression\n")
	res.WriteString("   operator := EQ | LE | GR\n")
	res.WriteString("   file is a file with sorted integers\n")
	res.WriteString("   N is a positive integer\n")
	res.WriteString("Example 1:\n")
	res.WriteString("   " + app + " [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]\n")
	res.WriteString("Example 2:\n")
	res.WriteString("   " + app + " [ LE 2 a.txt [ GR 1 b.txt c.txt ] ]\n")

	return res.String()
}

func main() {
	appArg := os.Args[0]

	if len(os.Args) <= 1 {
		println(Usage(appArg, "Missing expression"))
		os.Exit(1)
	}

	expressionArg := strings.Join(os.Args[1:], " ")
	parser := NewParser()

	expr, err := parser.Parse(expressionArg)
	if err != nil {
		println(Usage(appArg, fmt.Sprintf("Wrong expression \"%s\":\n   %v", expressionArg, err)))
		os.Exit(1)
	}

	for result := Evaluate(expr); ; {
		v, ok := result.Next()
		if !ok {
			break
		}

		println(v)
	}
}
