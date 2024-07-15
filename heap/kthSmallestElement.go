package heap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(arr []int, n, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, val := range arr {
		heap.Push(h, val)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

func main() {
	arr := []int{4, 3, 7, 6, 5}
	n := len(arr)
	k := 5

	result := kthSmallest(arr, n, k)
	fmt.Printf("The %d-th smallest element is %d\n", k, result)
}
