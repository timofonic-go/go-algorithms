package main

import "fmt"

type Stack struct {
	stack []int
	len   int
}

func (s *Stack) Push(v int) {
	s.stack = append(s.stack, v)
	s.len++
}

func (s *Stack) Pop() {
	// Take everything except for the head element.
	s.stack = s.stack[:s.Len()-1]
	s.len--
}

func (s Stack) Len() int {
	return len(s.stack)
}

func main() {

	s := Stack{}
	s.Push(1)
	s.Push(3)
	s.Push(4)

	fmt.Printf("%v", s)

	s.Pop()
	fmt.Printf("\n\n%v", s)
}
