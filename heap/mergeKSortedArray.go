package heap

import (
	"container/heap"
	"fmt"
)

type Element struct {
	value int
	row   int
	col   int
}

type MinHeap []Element

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKSortedArrays(arrays [][]int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		heap.Push(h, Element{arrays[i][0], i, 0})
	}

	var result []int

	for h.Len() > 0 {
		minElem := heap.Pop(h).(Element)
		result = append(result, minElem.value)

		if minElem.col+1 < len(arrays[minElem.row]) {
			heap.Push(h, Element{arrays[minElem.row][minElem.col+1], minElem.row, minElem.col + 1})
		}
	}

	return result
}

func main() {
	arrays := [][]int{
		{1, 5, 9},
		{2, 6, 8},
		{3, 7, 10},
	}
	k := len(arrays)

	result := mergeKSortedArrays(arrays, k)
	fmt.Println("Merged sorted array:", result)
}
