package heap

import (
	"container/heap"
	"fmt"
)

type MinHeapLS []int

func (h MinHeapLS) Len() int {
	return len(h)
}

func (h MinHeapLS) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeapLS) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *MinHeapLS) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapLS) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	minHeapLS *MinHeapLS
	k         int
}

func NewKthLargest(k int) *KthLargest {
	h := &MinHeapLS{}
	heap.Init(h)
	return &KthLargest{
		minHeapLS: h,
		k:         k,
	}
}

func (kl *KthLargest) Add(num int) int {
	if kl.minHeapLS.Len() < kl.k {
		heap.Push(kl.minHeapLS, num)
	} else if num > (*kl.minHeapLS)[0] {
		heap.Pop(kl.minHeapLS)
		heap.Push(kl.minHeapLS, num)
	}

	if kl.minHeapLS.Len() < kl.k {
		return -1
	}

	return (*kl.minHeapLS)[0]
}

func KthLargestStream() {
	k := 4
	arr := []int{1, 2, 3, 4, 5, 6}

	kl := NewKthLargest(k)
	for _, num := range arr {
		fmt.Print(kl.Add(num), " ")
	}
	fmt.Println()
}
