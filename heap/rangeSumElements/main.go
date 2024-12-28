package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	arr []int
	k1  int
	k2  int
}

type ElementI interface {
	GetSumOfRange() int
}

func newElem(arr []int, k1, k2 int) ElementI {
	return Element{
		arr: arr,
		k1:  k1,
		k2:  k2,
	}
}

type MaxHeap []int

func (mh MaxHeap) Len() int { return len(mh) }

func (mh MaxHeap) Less(i, j int) bool { return mh[i] > mh[j] }

func (mh MaxHeap) Swap(i, j int) { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MaxHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	temp := old[n-1]
	*mh = old[:n-1]
	return temp
}

func (e Element) GetSumOfRange() int {
	// Create independent slices for h1 and h2
	h1 := MaxHeap(append([]int{}, e.arr[:e.k1]...))   // Copy the first k1 elements for h1
	h2 := MaxHeap(append([]int{}, e.arr[:e.k2-1]...)) // Copy the first k2-1 elements for h2

	heap.Init(&h1)
	heap.Init(&h2)

	for i := e.k1; i < len(e.arr); i++ {
		heap.Push(&h1, e.arr[i])
		heap.Pop(&h1)
	}

	for i := e.k2 - 1; i < len(e.arr); i++ {
		heap.Push(&h2, e.arr[i])
		heap.Pop(&h2)
	}
	s1, s2 := 0, 0
	for h1.Len() > 0 {
		s1 += heap.Pop(&h1).(int)
	}

	for h2.Len() > 0 {
		s2 += heap.Pop(&h2).(int)
	}
	return s2 - s1
}

func main() {
	arr := []int{10, 2, 50, 12, 48, 13}
	k1, k2 := 2, 6
	obj := newElem(arr, k1, k2)
	fmt.Println(obj.GetSumOfRange())
}
