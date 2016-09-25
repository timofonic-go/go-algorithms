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

EXAMPLE:
[][][][][][][][][]

Total length = 9
Stacks Qty = 3

S0 start: (Total / Stacks Qty) * Stack idx = 9/3 * 0 = 0
S0 size: 9 / Stack Qty = 3

S1 start: (Total / Stacks Qty) * Stack idx = 9/3 * 1 = 3
S1 size: 9 / Stack Qty = 3

S2 start: (Total / Stacks Qty) * Stack idx = 9/3 * 2 = 6
S2 size: 9 / Stack Qty = 3

NOTES
1. Stacks should have a starting index
2. Should identify starting indexes in the beginning
3. For this we should know full array's length
4. Stacks should use the same array. Pass existing array by reference on init to a new Stack struct?
5. On Push should check if we overlap with the next stack. Can be implemented just defining max limit when initializing the task as we know array length from the beginning.
*/

type StackFixed struct {
	Capacity int
	Idx      int
}

type Arr struct {
	Stacks []StackFixed
	Values []int
	Sizes  []int
}

func (a *Arr) Push(stackNum, i int) {
	// If stack capacity not reached
	if a.Sizes[stackNum]+1 <= a.Stacks[stackNum].Capacity {

		// Increment stack size
		a.Sizes[stackNum]++

		// Set array value matching the top stacks value
		a.Values[a.GetStackTopIdx(stackNum)] = i
	}
}

func (a *Arr) GetStackTopIdx(stackNum int) int {
	// Get index lf last element of a given stack:
	// Starting index + stack size - 1 (to compensate for starting with 0)
	return a.Stacks[stackNum].Idx + a.Sizes[stackNum] - 1
}

// Due to a fixed array size, assume that the "not set" value is 0.
func (a *Arr) Pop(stackNum int) {
	a.Values[a.GetStackTopIdx(stackNum)] = 0
	a.Sizes[stackNum]--
}

func (a *Arr) Peek(stackNum int) int {
	return a.Values[a.GetStackTopIdx(stackNum)]
}

func (a *Arr) IsEmpty(stackNum int) bool {

	if a.Sizes[stackNum] > 0 {
		return true
	}

	return false
}
