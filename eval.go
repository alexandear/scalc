package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

func Evaluate(expr *Expression) intIterator {
	var setIters []intIterator

	for _, s := range expr.Sets {
		switch {
		case s.File != nil:
			setIters = append(setIters, FileToIterator(*s.File))
		case s.SubExpression != nil:
			setIters = append(setIters, Evaluate(s.SubExpression))
		}
	}

	return Calculate(expr.Operator, expr.N, setIters)
}

func Calculate(operator Operator, n uint, setIters []intIterator) intIterator {
	opFunc := OpFuncEQ

	switch operator {
	case OpEQ:
		opFunc = OpFuncEQ
	case OpLE:
		opFunc = OpFuncLE
	case OpGR:
		opFunc = OpFuncGR
	default:
	}

	return PerformOperationEf(opFunc, n, setIters)
}

type iterableReader struct {
	scanner *bufio.Scanner
	err     error
}

func (r *iterableReader) Next() (value int, ok bool) {
	n, err := ScanInt(r.scanner)

	if errors.Is(err, io.EOF) {
		return n, false
	} else if err != nil {
		log.Panicf("failed to scan: %v", err)
	}

	return n, true
}

func (r *iterableReader) Err() error {
	return r.err
}

func newIterableReader(r io.Reader) *iterableReader {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	return &iterableReader{
		scanner: scanner,
	}
}

func FileToIterator(file string) intIterator {
	fa, err := os.Open(file)
	if err != nil {
		log.Panicf("failed to open file %s: %v", file, err)
	}

	return newIterableReader(fa)
}
