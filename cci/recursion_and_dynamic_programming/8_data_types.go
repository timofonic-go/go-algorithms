package recursion_and_dynamic_programming

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) peek() int {
	if len(s.data) == 0 {
		fmt.Println("Attempt to peek() on an empty stack!")
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) pop() int {
	if len(s.data) == 0 {
		fmt.Println("Attempt to pop() from an empty stack!")
		return -1
	}
	out := s.data[len(s.data)-1]

	s.data = s.data[0 : len(s.data)-1]

	return out
}

func (s *Stack) push(d int) {
	s.data = append(s.data, d)
}
