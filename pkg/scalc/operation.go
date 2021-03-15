package scalc

import (
	"container/heap"
	"sort"
)

type Operator string

const (
	OpEQ Operator = "EQ"
	OpLE Operator = "LE"
	OpGR Operator = "GR"
)

type opFunc func(cnt, n uint) bool

func opFuncEQ(cnt, n uint) bool { return cnt == n }

func opFuncLE(cnt, n uint) bool { return cnt < n }

func opFuncGR(cnt, n uint) bool { return cnt > n }

type Iterator interface {
	Next() (value int, ok bool)
}

func Calculate(operator Operator, n uint, iters []Iterator) Iterator {
	opFn := opFuncEQ

	switch operator {
	case OpEQ:
		opFn = opFuncEQ
	case OpLE:
		opFn = opFuncLE
	case OpGR:
		opFn = opFuncGR
	default:
	}

	return calculate(opFn, n, iters)
}

type Pair struct {
	Value int
	Idx   int
}

func calculate(opFn opFunc, n uint, iters []Iterator) Iterator {
	operationLine := make([]Pair, 0, len(iters))
	for i := range iters {
		operationLine = append(operationLine, Pair{Idx: i})
	}

	var result []int

	var pairHeap PairHeap

	heap.Init(&pairHeap)

	for len(operationLine) != 0 {
		for _, line := range operationLine {
			v, ok := iters[line.Idx].Next()
			if !ok {
				continue
			}

			vi := Pair{
				Value: v,
				Idx:   line.Idx,
			}
			heap.Push(&pairHeap, vi)
		}

		if pairHeap.Len() == 0 {
			return NewIterableSlice(result)
		}

		min := pairHeap[0]

		var cnt uint

		for ; pairHeap.Len() != 0 && pairHeap[0].Value == min.Value; cnt++ {
			vi, _ := heap.Pop(&pairHeap).(Pair)
			operationLine = append(operationLine, vi)
		}

		if opFn(cnt, n) {
			result = append(result, min.Value)
		}
	}

	return NewIterableSlice(result)
}

func CalculateInefficient(operator Operator, n uint, iters []Iterator) Iterator {
	opFn := opFuncEQ

	switch operator {
	case OpEQ:
		opFn = opFuncEQ
	case OpLE:
		opFn = opFuncLE
	case OpGR:
		opFn = opFuncGR
	default:
	}

	return calculateInefficient(opFn, n, iters)
}

func calculateInefficient(opFunc opFunc, n uint, iters []Iterator) Iterator {
	counts := make(map[int]uint, len(iters))

	for _, iter := range iters {
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}

			counts[v]++
		}
	}

	res := make([]int, 0)

	for number, cnt := range counts {
		if opFunc(cnt, n) {
			res = append(res, number)
		}
	}

	sort.Ints(res)

	return NewIterableSlice(res)
}
