package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Element struct {
	arr []int
	k   int
}

type ElementI interface {
	GetMaximumGifts() int
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

func (e Element) GetMaximumGifts() int {
	h := MaxHeap(e.arr)
	heap.Init(&h)
	for e.k > 0 {
		e.k--
		val := int(math.Sqrt(float64(heap.Pop(&h).(int))))
		heap.Push(&h, val)
	}
	res := 0
	for h.Len() > 0 {
		res += heap.Pop(&h).(int)
	}
	return res
}

func main() {
	arr := []int{25, 64, 9, 4, 100}
	k := 4
	obj := newElem(arr, k)
	fmt.Println(obj.GetMaximumGifts())
}
