package main

import "fmt"

type Element struct {
	Value int
}

type ElementI interface {
	GenerateParantheses(int, int, string, *[]string)
}

func newElem(val int) ElementI {
	return Element{
		Value: val,
	}
}

func (e Element) GenerateParantheses(start, end int, str string, res *[]string) {
	if start == end && start == e.Value && end == e.Value {
		*res = append(*res, str)
		return
	}

	if start < e.Value {
		str += "("
		e.GenerateParantheses(start+1, end, str, res)
		str = str[:len(str)-1]
	}

	if end < start {
		str += ")"
		e.GenerateParantheses(start, end+1, str, res)
		str = str[:len(str)-1]
	}
}

func main() {
	obj := newElem(3)
	var res []string
	obj.GenerateParantheses(0, 0, "", &res)
	fmt.Println(res)
}
