package heap

import (
	"container/heap"
	"fmt"
)

// MaxHeapMS is a max-heap of integers.
type MaxHeapMS []int

func (h MaxHeapMS) Len() int           { return len(h) }
func (h MaxHeapMS) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeapMS) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapMS) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeapMS) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeapMS is a min-heap of integers.
type MinHeapMS []int

func (h MinHeapMS) Len() int           { return len(h) }
func (h MinHeapMS) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapMS) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapMS) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapMS) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MedianFinder is a struct that helps find the median from a stream of integers.
type MedianFinder struct {
	left  *MaxHeapMS // max-heap for the left half
	right *MinHeapMS // min-heap for the right half
}

func NewMedianFinder() *MedianFinder {
	left := &MaxHeapMS{}
	right := &MinHeapMS{}
	heap.Init(left)
	heap.Init(right)
	return &MedianFinder{
		left:  left,
		right: right,
	}
}

// AddNum adds a number to the data structure.
func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == 0 || num <= (*mf.left)[0] {
		heap.Push(mf.left, num)
	} else {
		heap.Push(mf.right, num)
	}

	// Balance the heaps
	if mf.left.Len() > mf.right.Len()+1 {
		heap.Push(mf.right, heap.Pop(mf.left))
	} else if mf.right.Len() > mf.left.Len() {
		heap.Push(mf.left, heap.Pop(mf.right))
	}
}

// FindMedian returns the median of all numbers added so far.
func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() > mf.right.Len() {
		return float64((*mf.left)[0])
	}
	return float64((*mf.left)[0]+(*mf.right)[0]) / 2.0
}

func MedianOfStream() {
	// Example input
	X := []int{5, 10, 15}

	medianFinder := NewMedianFinder()

	for _, num := range X {
		medianFinder.AddNum(num)
		median := medianFinder.FindMedian()
		fmt.Println(median)
	}
}
