package heap

import (
	"container/heap"
	"fmt"
)

// IntHeapNS is a min heap of integers
type IntHeapNS []int

// Len returns the number of elements in the heap
func (h IntHeapNS) Len() int {
	return len(h)
}

// Less returns true if the element at index i is less than the element at index j.
func (h IntHeapNS) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements at indices i and j.
func (h IntHeapNS) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *IntHeapNS) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element from the heap.
func (h *IntHeapNS) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// sort nearly sorted array
func NearlySorted(arr []int, k int) []int {
	var res []int
	h := &IntHeapNS{}
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

func NearlySortedArray() {
	arr := []int{3, 1, 4, 2, 5}
	k := 2
	fmt.Println(NearlySorted(arr, k))
}
