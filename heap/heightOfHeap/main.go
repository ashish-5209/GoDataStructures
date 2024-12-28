package main

import "fmt"

type Element struct {
	arr []int
}

type ElementI interface {
	GetHeapHeight() int
}

func newElem(arr []int) ElementI {
	return Element{
		arr: arr,
	}
}

func (e Element) GetHeapHeight() int {
	n := len(e.arr)
	if n == 1 {
		return 1
	}
	res := 0
	for n > 1 {
		res += 1
		n /= 2
	}
	return res
}

func main() {
	arr := []int{3, 6, 9, 2, 15, 10, 14, 5, 12}
	obj := newElem(arr)
	fmt.Println(obj.GetHeapHeight())
}
