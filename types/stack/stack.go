package main

import "fmt"

type stackStore struct {
	stack []int
	len   int
}

func (s *stackStore) Push(v int) {
	s.stack = append(s.stack, v)
	s.len++
}

func (s *stackStore) Pop() {

	if s.len > 0 {
		lastIdx := s.len - 1
		s.stack = s.stack[:lastIdx]
		s.len--
	}
}

func (s stackStore) Peek() int {
	return s.stack[s.len-1]
}

func main() {
	s := stackStore{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println("Original stack: ", s.stack)

	s.Pop()
	s.Pop()

	el := s.Peek()

	fmt.Println("Final stack: ", s.stack)
	fmt.Println("Peeked el: ", el)
}
