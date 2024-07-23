package heap

import (
	"container/heap"
	"fmt"
)

type IntHeapSE []int

func (h IntHeapSE) Len() int {
	return len(h)
}

func (h IntHeapSE) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeapSE) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *IntHeapSE) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeapSE) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(arr []int, n, k int) int {
	h := &IntHeapSE{}
	heap.Init(h)

	for _, val := range arr {
		heap.Push(h, val)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

func KthSmallestElement() {
	arr := []int{4, 3, 7, 6, 5}
	n := len(arr)
	k := 5

	result := kthSmallest(arr, n, k)
	fmt.Printf("The %d-th smallest element is %d\n", k, result)
}
