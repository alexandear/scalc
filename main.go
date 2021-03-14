package main

import (
	"fmt"
	"os"
	"strings"
)

func usage(appName, reason string) string {
	var res strings.Builder

	res.WriteString(reason)
	res.WriteString("\nUsage:\n")
	res.WriteString("   " + appName + " expression\n")
	res.WriteString("Where:\n")
	res.WriteString("   expression := [ operator N sets ]\n")
	res.WriteString("   sets := set | set sets\n")
	res.WriteString("   set := file | expression\n")
	res.WriteString("   operator := EQ | LE | GR\n")
	res.WriteString("   file is a file with sorted integers\n")
	res.WriteString("   N is a positive integer\n")
	res.WriteString("Example 1:\n")
	res.WriteString("   " + appName + " [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]\n")
	res.WriteString("Example 2:\n")
	res.WriteString("   " + appName + " [ LE 2 a.txt [ GR 1 b.txt c.txt ] ]\n")

	return res.String()
}

func main() {
	appArg := os.Args[0]

	if len(os.Args) <= 1 {
		exit(usage(appArg, "Missing expression"))
	}

	exprArg := strings.Join(os.Args[1:], " ")
	parser := NewParser()

	expr, err := parser.Parse(exprArg)
	if err != nil {
		exit(usage(appArg, fmt.Sprintf("Wrong expression \"%s\":\n   %v", exprArg, err)))
	}

	for result := Evaluate(expr); ; {
		v, ok := result.Next()
		if !ok {
			break
		}

		println(v)
	}
}

func exit(usage string) {
	_, _ = fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}
