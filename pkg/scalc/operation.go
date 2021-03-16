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

type Iterator interface {
	Next() (value int, ok bool)
}

type Pair struct {
	Value int
	Idx   int
}

func Calculate(operator Operator, n uint, iters []Iterator) Iterator {
	res := make(chan int)

	go func() {
		calculate(operator, n, iters, res)
		close(res)
	}()

	return &iterableChannel{Chan: res}
}

type iterableChannel struct {
	Chan <-chan int
}

func (c *iterableChannel) Next() (value int, ok bool) {
	value, ok = <-c.Chan

	return value, ok
}

func calculate(operator Operator, n uint, iters []Iterator, res chan<- int) {
	pairHeap := make(PairHeap, len(iters))

	for idx, iter := range iters {
		v, ok := iter.Next()
		if !ok {
			continue
		}

		pairHeap[idx] = Pair{Idx: idx, Value: v}
	}

	heap.Init(&pairHeap)

	for pairHeap.Len() != 0 {
		min, _ := heap.Pop(&pairHeap).(Pair)
		nextIdxes := []int{min.Idx}

		for pairHeap.Len() > 0 && pairHeap[0].Value == min.Value {
			min, _ = heap.Pop(&pairHeap).(Pair)
			nextIdxes = append(nextIdxes, min.Idx)
		}

		if cnt := len(nextIdxes); operatorFn(operator)(cnt, n) {
			res <- min.Value
		}

		for _, idx := range nextIdxes {
			v, ok := iters[idx].Next()
			if !ok {
				continue
			}

			heap.Push(&pairHeap, Pair{Idx: idx, Value: v})
		}
	}
}

func operatorFn(operator Operator) func(cnt int, n uint) bool {
	switch operator {
	case OpLE:
		return func(cnt int, n uint) bool { return uint(cnt) < n }
	case OpGR:
		return func(cnt int, n uint) bool { return uint(cnt) > n }
	case OpEQ:
		return func(cnt int, n uint) bool { return uint(cnt) == n }
	default:
		return func(cnt int, n uint) bool { return uint(cnt) == n }
	}
}
