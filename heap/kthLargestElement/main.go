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
	GetKthLargetElement() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
	}
}

type MinHeap []int

func (mh MinHeap) Len() int { return len(mh) }

func (mh MinHeap) Less(i, j int) bool { return mh[i] < mh[j] }

func (mh MinHeap) Swap(i, j int) { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	temp := old[n-1]
	*mh = old[:n-1]
	return temp
}

func (e Element) GetKthLargetElement() int {
	h := MinHeap(e.arr[:e.k])
	heap.Init(&h)
	for i := e.k; i < len(e.arr); i++ {
		heap.Push(&h, e.arr[i])
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

func main() {
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	obj := newElem(arr, k)
	fmt.Println(obj.GetKthLargetElement())
}
