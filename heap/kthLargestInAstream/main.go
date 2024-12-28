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
	GetKthLargetElementInStream() []int
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

func (e Element) GetKthLargetElementInStream() []int {
	res := []int{}
	h := MinHeap(e.arr[:e.k])
	heap.Init(&h)
	x := 0
	for x < e.k-1 {
		x++
		res = append(res, -1)
	}
	res = append(res, h[0])
	for i := e.k; i < len(e.arr); i++ {
		heap.Push(&h, e.arr[i])
		heap.Pop(&h)
		res = append(res, h[0])
	}
	return res
}

func main() {
	arr := []int{3, 4}
	k := 1
	obj := newElem(arr, k)
	fmt.Println(obj.GetKthLargetElementInStream())
}
