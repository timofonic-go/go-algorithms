package main

import "fmt"

type Queue struct {
	queue []int
	len   int
}

func (s *Queue) Push(v int) {
	s.queue = append(s.queue, v)
	s.len++
}

func (s *Queue) Pop() {

	if s.len > 0 {
		s.queue = s.queue[1:]
		s.len--
	}
}

func (s Queue) Peek() int {
	return s.queue[0]
}

func main() {
	s := Queue{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println("Original queue: ", s.queue)

	s.Pop()
	s.Pop()

	el := s.Peek()

	fmt.Println("Final queue: ", s.queue)
	fmt.Println("Peeked el: ", el)
}
