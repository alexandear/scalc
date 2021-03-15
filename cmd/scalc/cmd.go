package scalc

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alexandear/scalc/pkg/scalc"
)

func Execute(args []string) (string, error) {
	appArg := args[0]

	if len(os.Args) <= 1 {
		return usage(appArg), errors.New("missing expression")
	}

	exprArg := strings.Join(os.Args[1:], " ")
	parser := scalc.NewParser()

	expr, err := parser.Parse(exprArg)
	if err != nil {
		return usage(appArg), fmt.Errorf("wrong expression \"%s\": %w", exprArg, err)
	}

	var res strings.Builder

	for ev := scalc.Evaluate(expr); ; {
		v, ok := ev.Next()
		if !ok {
			break
		}

		res.WriteString(strconv.Itoa(v))
		res.WriteRune('\n')
	}

	return res.String(), nil
}

func usage(appName string) string {
	var res strings.Builder

	res.WriteString("Usage:\n   ")
	res.WriteString(appName)
	res.WriteString(" expression\n")
	res.WriteString("Where expression:\n")
	res.WriteString("   expression := [ operator N sets ]\n")
	res.WriteString("   sets := set | set sets\n")
	res.WriteString("   set := file | expression\n")
	res.WriteString("   operator := EQ | LE | GR\n")
	res.WriteString("   file is a file with sorted integers\n")
	res.WriteString("   N is a positive integer\n")
	res.WriteString("Example 1:\n   ")
	res.WriteString(appName)
	res.WriteString(" [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]\n")
	res.WriteString("Example 2:\n   ")
	res.WriteString(appName)
	res.WriteString(" [ LE 2 a.txt [ GR 1 b.txt c.txt ] ]\n")

	return res.String()
}
