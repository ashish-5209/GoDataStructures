package stack

import "fmt"

// Stack represent stack that holds a slice
type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Pop will remove a value at the end
// and Returns the required value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func InsertDelete() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
