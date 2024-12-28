package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	GetKthSmallestElement() int
}

type MaxHeap []int

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	temp := old[n-1]
	*h = old[:n-1]
	return temp
}

func (e Element) GetKthSmallestElement() int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, val := range e.arr {
		heap.Push(h, val)
		if h.Len() > e.k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

func main() {
	arr := []int{4, 3, 7, 6, 5}
	k := 2
	obj := newElem(arr, k)
	fmt.Println(obj.GetKthSmallestElement())
}
