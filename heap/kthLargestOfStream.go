package heap

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	minHeap *MinHeap
	k       int
}

func NewKthLargest(k int) *KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	return &KthLargest{
		minHeap: h,
		k:       k,
	}
}

func (kl *KthLargest) Add(num int) int {
	if kl.minHeap.Len() < kl.k {
		heap.Push(kl.minHeap, num)
	} else if num > (*kl.minHeap)[0] {
		heap.Pop(kl.minHeap)
		heap.Push(kl.minHeap, num)
	}

	if kl.minHeap.Len() < kl.k {
		return -1
	}

	return (*kl.minHeap)[0]
}

func main() {
	k := 4
	arr := []int{1, 2, 3, 4, 5, 6}

	kl := NewKthLargest(k)
	for _, num := range arr {
		fmt.Print(kl.Add(num), " ")
	}
	fmt.Println()
}
