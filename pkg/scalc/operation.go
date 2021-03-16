package scalc

import (
	"container/heap"
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

func operatorToFunc(operator Operator) opFunc {
	switch operator {
	case OpEQ:
		return opFuncEQ
	case OpLE:
		return opFuncLE
	case OpGR:
		return opFuncGR
	default:
		return opFuncEQ
	}
}

type Iterator interface {
	Next() (value int, ok bool)
}

type Pair struct {
	Value int
	Idx   int
}

func Calculate(operator Operator, n uint, iters []Iterator) Iterator {
	pairHeap := make(PairHeap, len(iters))

	for idx, iter := range iters {
		v, ok := iter.Next()
		if !ok {
			continue
		}

		pairHeap[idx] = Pair{Idx: idx, Value: v}
	}

	heap.Init(&pairHeap)

	var result []int

	for pairHeap.Len() != 0 {
		min, _ := heap.Pop(&pairHeap).(Pair)
		nextIdxes := []int{min.Idx}

		for pairHeap.Len() > 0 && pairHeap[0].Value == min.Value {
			min, _ = heap.Pop(&pairHeap).(Pair)
			nextIdxes = append(nextIdxes, min.Idx)
		}

		if operatorToFunc(operator)(uint(len(nextIdxes)), n) {
			result = append(result, min.Value)
		}

		for _, idx := range nextIdxes {
			v, ok := iters[idx].Next()
			if !ok {
				continue
			}

			heap.Push(&pairHeap, Pair{Idx: idx, Value: v})
		}
	}

	return NewIterableSlice(result)
}
