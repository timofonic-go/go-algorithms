package main

import "fmt"

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {

	if len(s) == 0 {
		return s, -1
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() int {
	return s[len(s)-1]
}

func main() {
	s := make(stack, 0)
	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)

	s, p := s.Pop()

	el := s.Peek()

	fmt.Println("Stack: ", s)
	fmt.Println("Popped element: ", p)
	fmt.Println("Peeked el: ", el)
}
