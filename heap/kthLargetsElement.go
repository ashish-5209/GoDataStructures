package heap

import (
	"container/heap"
	"fmt"
)

// IntHeap is a min-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthLargest(arr []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Printf("The %d-th largest element is %d\n", k, kthLargest(arr, k))
}
