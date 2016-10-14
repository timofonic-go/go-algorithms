package stacks

type StackOfPlates struct {
	stacks []*Stack
	last   int
}

func (s *StackOfPlates) NewStack() {
	s.stacks = append(s.stacks, &Stack{})
}

type Stack struct {
	data []int
}

func (s *Stack) push(i int) {
	s.data = append(s.data, i)
}

func (s *Stack) getCapacity() int {
	return len(s.data)
}

func (s *Stack) pop() int {
	el := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return el
}

const maxCapacity = 3

func (s *StackOfPlates) push(i int) {

	if len(s.stacks) == 0 {
		s.stacks = make([]*Stack, 0)
		s.NewStack()
		s.last = 0
	}

	if s.stacks[s.last].getCapacity() == maxCapacity {
		s.last++
		s.NewStack()
	}

	s.stacks[s.last].push(i)
}

// pop() returns a value from the top of the last stack.
func (s *StackOfPlates) pop() int {

	if s.stacks[s.last].getCapacity() == 0 {
		s.last--
	}

	// first column
	if s.last < 0 {
		panic("Cannot pop() from a first empty stack!")
	}

	return s.stacks[s.last].pop()
}

// pop() from a specific stack
func (s *StackOfPlates) popAt(i int) int {

	if i > s.last {
		panic("Asked to pop() from a non-existent stack!")
	}

	if s.last == i {
		return s.pop()
	}

	// not first, not last
	el := s.stacks[i].pop()

	// waterfall rest of el's to this stack

	for n := i; n < s.last; n++ {
		// pop from next stack
		e := s.stacks[n+1].pop()

		// add to current
		s.stacks[n].push(e)
	}

	// Cleanup empty last stack
	if s.stacks[s.last].getCapacity() == 0 {
		s.last--
	}

	return el
}
