package main

import (
	"container/heap"
	"fmt"
)

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

type Element struct {
	arr []int
}

type ElementI interface {
	GetLastStoneWeight() int
}

func newElem(arr []int) ElementI {
	return Element{
		arr: arr,
	}
}

func (e Element) GetLastStoneWeight() int {
	h := MaxHeap(e.arr)
	heap.Init(&h)
	for h.Len() > 1 {
		x, y := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if (x - y) > 0 {
			heap.Push(&h, x-y)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(&h).(int)
}

func main() {
	arr := []int{2, 7, 4, 1, 8, 1}
	obj := newElem(arr)
	fmt.Println(obj.GetLastStoneWeight())
}
