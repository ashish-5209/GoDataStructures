package main

import "fmt"

type Fib struct {
	n int
}

type FibN interface {
	FindNthFib() int
}

func (n Fib) FindNthFib() int {
	if n.n < 2 {
		return n.n
	}
	a, b := 0, 1
	i := 0
	for i = range n.n {
		c := a + b
		a = b
		b = c
	}
	a = i
	return b
}

func main() {
	obj := Fib{
		n: 7,
	}

	val := obj.FindNthFib()
	fmt.Println(val)
}
