package heap

import (
	"container/heap"
	"fmt"
)

type CharFrequency struct {
	char      byte
	frequency int
}

type MaxHeap []CharFrequency

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].frequency > h[j].frequency
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(CharFrequency))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rearrangeString(s string) string {
	frequencyMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		frequencyMap[s[i]]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for char, freq := range frequencyMap {
		heap.Push(h, CharFrequency{char, freq})
	}

	var result []byte
	var prev CharFrequency

	for h.Len() > 0 {
		curr := heap.Pop(h).(CharFrequency)
		result = append(result, curr.char)
		if prev.frequency > 0 {
			heap.Push(h, prev)
		}
		curr.frequency--
		prev = curr
	}

	if len(result) != len(s) {
		return ""
	}

	return string(result)
}

func main() {
	S := "aaab"
	result := rearrangeString(S)
	if result == "" {
		fmt.Println("Not possible to rearrange")
	} else {
		fmt.Println("Rearranged string:", result)
	}
}
