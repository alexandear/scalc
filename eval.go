package main

import (
	"log"
	"os"
)

func Evaluate(expr *Expression) Iterator {
	var setIters []Iterator

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

func Calculate(operator Operator, n uint, setIters []Iterator) Iterator {
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

	return PerformOperation(opFunc, n, setIters)
}

func FileToIterator(file string) Iterator {
	fa, err := os.Open(file)
	if err != nil {
		log.Panicf("failed to open file: %v", err)
	}

	return NewIterableReader(fa)
}
