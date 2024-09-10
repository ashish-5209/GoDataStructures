package main

import "fmt"

type Elem struct {
	str string
}

type ElemI interface {
	FindUniqueString() int
}

func newElem() Elem {
	return Elem{
		str: "",
	}
}

func (e Elem) FindUniqueString() int {
	if len(e.str) < 2 {
		return len(e.str)
	}

	maxlength := 0
	visited := make([]bool, 256)

	left, right := 0, 0
	for right < len(e.str) {
		for visited[e.str[right]] {
			visited[e.str[left]] = false
			left++
		}
		visited[e.str[right]] = true
		maxlength = max(maxlength, right-left+1)
		right++
	}
	return maxlength
}

func main() {
	e := newElem()
	e.str = "geeksforgeeks"
	val := e.FindUniqueString()
	fmt.Println(val)
}
