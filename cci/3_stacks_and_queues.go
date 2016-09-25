package cci

type Stack struct {
	Values []int
	Size   int
}

func (s *Stack) Pop() {
	if s.Size > 0 {
		s.Values = s.Values[0 : s.Size-1]
		s.Size--
	}
}

func (s *Stack) Push(i int) {
	s.Values = append(s.Values, i)
	s.Size++
}

func (s *Stack) Peek() int {
	return s.Values[s.Size-1]
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

// 3.1 Three in One: Describe how you could use a single array to implement three stacks.
/**
MY QUESTIONS FIRST
Has this array a limited length?
If limited, then:
1. We attach first stack to the beginning of the array
2. We attach the second stack to the end of the array and grow it backwards
3. We attach beginning of the 3rd stack to the middle and grow to the right.
Also, we add a verification to the Push operation to check if the new position going to overwrite another stack.
In case it can override, we shift entire stack to the left with an additional check before the shift, to avoid overriding the previous stack.

In case of unlimited length, we could just spread starting points at a certain range (maybe withing the integer range),
and grow to the right, until possible. And then either pop current stack or shift the second stack to the right.
Simplest case.

*/
