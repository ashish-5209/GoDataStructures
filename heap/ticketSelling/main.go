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
	GetMaximumProfit() int
}

func newElem(arr []int, k int) ElementI {
	return Element{
		arr: arr,
		k:   k,
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

func (e Element) GetMaximumProfit() int {
	h := MaxHeap(e.arr)
	heap.Init(&h)
	res := 0
	for e.k > 0 {
		e.k--
		val := heap.Pop(&h).(int)
		res += val
		val--
		heap.Push(&h, val)
	}
	return res
}

func main() {
	arr := []int{2, 3}
	k := 3
	obj := newElem(arr, k)
	fmt.Println(obj.GetMaximumProfit())
}
