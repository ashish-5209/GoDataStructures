package heap

import (
	"container/heap"
	"fmt"
)

// IntHeapLE is a min-heap of integers.
type IntHeapLE []int

func (h IntHeapLE) Len() int            { return len(h) }
func (h IntHeapLE) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeapLE) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeapLE) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeapLE) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthLargest(arr []int, k int) int {
	h := &IntHeapLE{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

func KthLasgestElement() {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Printf("The %d-th largest element is %d\n", k, kthLargest(arr, k))
}
