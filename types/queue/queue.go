package main

import "fmt"

type Queue struct {
	queue []int
	len   int
}

// Insert element into the end of the queue.
func (q *Queue) Enq(v int) {
	q.queue = append(q.queue, v)
	q.len++
}

// Remove & return an element from the queue.
func (q *Queue) Deq() int {
	el := q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return el
}

func main() {
	q := Queue{}
	q.Enq(1)
	q.Enq(4)
	q.Enq(2)

	fmt.Printf("\n\nStack: %v: ", q)

	el := q.Deq()
	fmt.Printf("\n\nDequeued el: %v: ", el)
	fmt.Printf("\n\nStack after dequeueing: %v: ", q)
}
