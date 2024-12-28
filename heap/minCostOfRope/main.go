package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	arr []int
}

type ElementI interface {
	GetMinimumCost() int
}

func newElem(arr []int) ElementI {
	return Element{
		arr: arr,
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

func (e Element) GetMinimumCost() int {
	if len(e.arr) < 2 {
		return 0
	}
	h := &MinHeap{}
	heap.Init(h)

	for _, val := range e.arr {
		heap.Push(h, val)
	}
	res := 0
	for h.Len() > 1 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		res += a + b
		heap.Push(h, a+b)
	}
	return res
}

func main() {
	arr := []int{4, 2, 7, 6, 9}
	obj := newElem(arr)
	fmt.Println(obj.GetMinimumCost())
}
