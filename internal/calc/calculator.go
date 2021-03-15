package calc

import (
	"fmt"
	"io"

	"go.uber.org/multierr"

	"github.com/alexandear/scalc/internal/parser"
	"github.com/alexandear/scalc/pkg/scalc"
)

type Parser interface {
	Parse(s string) (*parser.Expression, error)
}

type FileToIterator interface {
	Iterator(file string) (scalc.Iterator, io.Closer, error)
}

type Calculator struct {
	parser   Parser
	fileIter FileToIterator
	closers  []io.Closer
}

func NewCalculator(parser Parser, fileIter FileToIterator) *Calculator {
	return &Calculator{
		parser:   parser,
		fileIter: fileIter,
	}
}

func (c *Calculator) Calculate(expression string) (scalc.Iterator, error) {
	expr, err := c.parser.Parse(expression)
	if err != nil {
		return nil, fmt.Errorf("wrong expression \"%s\": %w", expression, err)
	}

	it, err := c.evaluate(expr)
	if err != nil {
		return nil, fmt.Errorf("evaluate expression \"%s\": %w", expression, err)
	}

	return it, nil
}

func (c *Calculator) evaluate(expr *parser.Expression) (scalc.Iterator, error) {
	var iters []scalc.Iterator

	for _, s := range expr.Sets {
		switch {
		case s.File != nil:
			fit, cl, err := c.fileIter.Iterator(*s.File)
			if err != nil {
				return nil, fmt.Errorf("convert file to iterator: %w", err)
			}

			c.closers = append(c.closers, cl)

			iters = append(iters, fit)
		case s.SubExpression != nil:
			eval, err := c.evaluate(s.SubExpression)
			if err != nil {
				return nil, fmt.Errorf("evaluate sub expression: %w", err)
			}

			iters = append(iters, eval)
		}
	}

	return scalc.Calculate(expr.Operator, expr.N, iters), nil
}

func (c *Calculator) Close() error {
	var err error

	for _, c := range c.closers {
		err = multierr.Append(err, c.Close())
	}

	return err // nolint:wrapcheck // as is
}
