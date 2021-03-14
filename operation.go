package main

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

func PerformOperationInef(opFunc OpFunc, n uint, iters []Iterator) Iterator {
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

type ValueIdx struct {
	Value int
	Idx   int
}

func PerformOperation(opFunc OpFunc, n uint, iters []Iterator) Iterator {
	minValueIdxes := make([]ValueIdx, 0, len(iters))
	for i := range iters {
		minValueIdxes = append(minValueIdxes, ValueIdx{Idx: i})
	}

	result := make([]int, 0)

	var valueIdxHeap ValueIdxHeap

	heap.Init(&valueIdxHeap)

	for len(minValueIdxes) != 0 {
		for _, p := range minValueIdxes {
			v, ok := iters[p.Idx].Next()
			if !ok {
				continue
			}

			heap.Push(&valueIdxHeap, ValueIdx{
				Value: v,
				Idx:   p.Idx,
			})
		}

		if valueIdxHeap.Len() == 0 {
			return NewIterableSlice(result)
		}

		min := valueIdxHeap[0]

		var cnt uint

		for ; valueIdxHeap.Len() != 0 && valueIdxHeap[0].Value == min.Value; cnt++ {
			minValueIdxes = append(minValueIdxes, heap.Pop(&valueIdxHeap).(ValueIdx))
		}

		if opFunc(cnt, n) {
			result = append(result, min.Value)
		}
	}

	return NewIterableSlice(result)
}
