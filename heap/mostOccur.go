package heap

import (
	"container/heap"
	"fmt"
)

type IntHeapMO []int

func (h IntHeapMO) Len() int {
	return len(h)
}

func (h IntHeapMO) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeapMO) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeapMO) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeapMO) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mostOccur(arr []int, n, k int) int {
	frequencyMap := make(map[int]int)
	for _, num := range arr {
		frequencyMap[num]++
	}

	frequencies := make([]int, 0, len(frequencyMap))

	for _, count := range frequencyMap {
		frequencies = append(frequencies, count)
	}

	h := &IntHeapMO{}
	heap.Init(h)

	for _, freq := range frequencies {
		heap.Push(h, freq)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	return sum
}

func MostOccurElement() {
	arr := []int{3, 1, 4, 4, 5, 2, 6, 1}
	n := len(arr)
	k := 2

	result := mostOccur(arr, n, k)
	fmt.Println("Sum of the frequencies of the k most frequent elements:", result)
}
