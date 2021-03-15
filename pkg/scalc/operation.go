package scalc

import (
	"container/heap"
	"sort"
)

type Iterator interface {
	Next() (value int, ok bool)
}

type OpFunc func(cnt, n uint) bool

func OpFuncEQ(cnt, n uint) bool {
	return cnt == n
}

func OpFuncLE(cnt, n uint) bool {
	return cnt < n
}

func OpFuncGR(cnt, n uint) bool {
	return cnt > n
}

type Pair struct {
	Value int
	Idx   int
}

func Calculate(operator Operator, n uint, iters []Iterator) Iterator {
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

	return calculate(opFunc, n, iters)
}

func calculate(opFunc OpFunc, n uint, iters []Iterator) Iterator {
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

		if opFunc(cnt, n) {
			result = append(result, min.Value)
		}
	}

	return NewIterableSlice(result)
}

func calculateInefficient(opFunc OpFunc, n uint, iters []Iterator) Iterator {
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
