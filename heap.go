package main

// ValueIdxHeap is a min-heap of ValueIdx values comparing by ValueIdx.Value.
type ValueIdxHeap []ValueIdx

func (h ValueIdxHeap) Len() int { return len(h) }

func (h ValueIdxHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }

func (h ValueIdxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ValueIdxHeap) Push(x interface{}) {
	*h = append(*h, x.(ValueIdx))
}

func (h *ValueIdxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
