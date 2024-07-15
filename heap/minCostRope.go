package heap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Len() int {
	return len(h)
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

func minCostRope(arr []int, n int) int {
	if n == 0 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)

	for _, val := range arr {
		heap.Push(h, val)
	}
	sum := 0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		heap.Push(h, a+b)
		sum += a + b
	}

	return sum
}

func main() {
	arr := []int{4, 2, 7, 6, 9}

	fmt.Println(minCostRope(arr, len(arr)))
}