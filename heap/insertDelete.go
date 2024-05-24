package heap

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(val int) {
	h.array = append(h.array, val)
	h.MaxHeapifyUp(len(h.array) - 1)
}

// Extracts returns the largets	key, and removes it from the heap
func (h *MaxHeap) ExtractMax() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	// take out the last element and put it in the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.MaxHeapifyDown(0)
	return extracted
}

// MaxHeapifyUp will heapify from the bottom
func (h *MaxHeap) MaxHeapifyUp(idx int) {
	for h.array[parent(idx)] < h.array[idx] {
		h.Swap(parent(idx), idx)
		idx = parent(idx)
	}
}

// Swap the index value
func (h *MaxHeap) Swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left chile
func left(i int) int {
	return 2*i + 1
}

// get the right child
func right(i int) int {
	return 2*i + 2
}

// MaxHeapifyDown will heapify top to bottom
func (h *MaxHeap) MaxHeapifyDown(idx int) {
	lastIndex := len(h.array) - 1
	l, r := left(idx), right(idx)
	childToCompare := 0

	// loop when index has atleast one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[idx] < h.array[childToCompare] {
			h.Swap(idx, childToCompare)
			idx = childToCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

func InsertDelete() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, val := range buildHeap {
		m.Insert(val)
		fmt.Println(*m)
	}

	for i := 0; i < 5; i++ {
		m.ExtractMax()
		fmt.Println(*m)
	}
}
