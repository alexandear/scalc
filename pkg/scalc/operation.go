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

func Calculate(operator Operator, n uint, iters []Iterator) Iterator {
	res := make(chan int)

	go func() {
		calculate(operator, n, iters, res)
		close(res)
	}()

	return &iterableChannel{Chan: res}
}

func calculate(operator Operator, n uint, iters []Iterator, res chan<- int) {
	queue := make(PriorityQueue, len(iters))

	for idx, iter := range iters {
		priority, ok := iter.Next()
		if !ok {
			continue
		}

		queue[idx] = &Item{
			value:    idx,
			priority: priority,
		}
	}

	heap.Init(&queue)

	for queue.Len() != 0 {
		min, _ := heap.Pop(&queue).(*Item)
		nextIts := []int{min.value}

		for queue.Len() > 0 && queue[0].priority == min.priority {
			min, _ = heap.Pop(&queue).(*Item)
			nextIts = append(nextIts, min.value)
		}

		if cnt := len(nextIts); operatorFn(operator)(cnt, n) {
			res <- min.priority
		}

		for _, idx := range nextIts {
			priority, ok := iters[idx].Next()
			if !ok {
				continue
			}

			heap.Push(&queue, &Item{priority: priority, value: idx})
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

type iterableChannel struct {
	Chan <-chan int
}

func (c *iterableChannel) Next() (value int, ok bool) {
	value, ok = <-c.Chan

	return value, ok
}
