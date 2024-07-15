package heap

import (
	"container/heap"
	"fmt"
)

// IntHeap is a min heap of integers
type IntHeap []int

// Len returns the number of elements in the heap
func (h IntHeap) Len() int {
	return len(h)
}

// Less returns true if the element at index i is less than the element at index j.
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements at indices i and j.
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element from the heap.
func (h *IntHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// sort nearly sorted array
func NearlySorted(arr []int, k int) []int {
	var res []int
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range arr {
		heap.Push(h, num)
		if h.Len() > k+1 {
			res = append(res, heap.Pop(h).(int))

		}
	}

	// Pop the remaining elements from the heap
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

func main() {
	arr := []int{3, 1, 4, 2, 5}
	k := 2
	fmt.Println(NearlySorted(arr, k))
}
