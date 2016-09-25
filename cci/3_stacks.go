package cci

import "fmt"

type Stack struct {
	Values    []int
	Size      int
	MinValues []int
	Min       int
}

func (s *Stack) Pop() {
	if s.Size > 0 {
		// New stack size.
		s.Size--

		// Rewrite the stack with its "cut" size.
		s.Values = s.Values[0:s.Size]

		// On pop, check replace current Min value with a stored Min value for this position from MinValues.
		s.Min = s.MinValues[s.Size-1]

		s.MinValues = s.MinValues[0:s.Size]

	}
}

func (s *Stack) Push(i int) {
	s.Values = append(s.Values, i)
	// s.Min is initially 0, so on first push we need to set to the first pushed element.
	if s.Size == 0 {
		s.Min = i
	} else {
		if i < s.Min {
			s.Min = i
		}
	}
	// On array Push, fill a corresponding min value.
	s.MinValues = append(s.MinValues, s.Min)
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

// 3.2 Stack Min.
// How would you design a stack which, in addition to push and pop, has a function min,
// which returns the minimum element?
// Push, pop and min should all operate in 0(1) time.
/**
4, 6, 2, 3
BRUTE FORCE: Keep track of minimal value.
Add a MinEl property to the stack.
On every Push - compare new element to MinEl, and overwrite MinEl if a new element is lower.
On every Pop - compare new element to MinEl. If matches, replace with the "next" minimal element. However, we cannot run a loop due to the O(1) restriction.

In general, for every Pop we should know what minimum element exists for the popped element.
We can have a second stack that wil be storing current min value for each index.

1. Update the Stack implementation above.
2. Implement a min() method.
*/

func (s *Stack) GetMin() int {
	return s.Min
}

// 3.3 Stack Of Plates
// Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
// Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold.
// Implement a data structure SetOfStacks that mimics this.
// SetOfStacks should be composed of several stacks and should create a new stack once the previous one exceeds capacity.
// SetOfStacks. push() and SetOfStacks. pop() should behave identically to a single stack
// (that is, pop () should return the same values as it would if there were just a single stack).
// FOLLOW UP
// Implement a function popAt (int index) which performs a pop operation on a specific sub-stack.
/**
EXAMPLE
Let's say our capacity is 3 plates.
Start with a 1st stack:
Push() -> stack 1 : 1
Push() -> stack 1 : 2
Push() -> stack 1 : 3

Push() -> here we need to know that we reached the capacity limit.
If we reached, we need to push to the next stack.
We can differentiate stacks by their own index.
So, we capacity is reached, we increment current stack index by 1 and push an item to it.
Push() -> stack 2 : 4
etc.

Pop()
Should check if number of elements in the current stack is > 0
If > 0, then Pop()

Now, len(el) == 0
If size == 0, we need to start popping from a previous stack. For this, we need to decrement current stack index.

ALGORITHM:

currStackIdx := 0
stackCapacity := 3

Push(4) -> 0:4
Push(7) -> 0:7
Push(2) -> 0:9
Push(3) -> 1:3
Push(1) -> 1:1

Pop()
0:4
0:7
0:9
1:3

Pop()
0:4
0:7
0:9

Pop()
0:4
0:7

etc.


NOTES:
Pop() and Push() should change stack index if reaching lower and higher limits!

*/
type StacksOfPlates struct {
	CurrentStackIndex int
	Stacks            []*StackOfPlates
	Capacity          int
}

type StackOfPlates struct {
	Values []int
}

func (s *StacksOfPlates) Push(i int) {

	if len(s.Stacks) == 0 {
		// Initialize first stack.
		s.AddStack()
	}

	// Check if capacity is reached
	if s.GetCurrStackSize() >= s.Capacity {
		// Reached. Need to switch to a new stack and add plate there.
		s.CurrentStackIndex++

		// Add a stack
		s.AddStack()
	}
	s.Stacks[s.CurrentStackIndex].Values = append(s.Stacks[s.CurrentStackIndex].Values, i)
}

func (s *StacksOfPlates) GetCurrStackSize() int {
	return len(s.Stacks[s.CurrentStackIndex].Values)
}

func (s *StacksOfPlates) GetCurrStackIdx() int {
	return s.CurrentStackIndex
}

func (s *StacksOfPlates) Pop() {

	// Check if stack is empty. If empty, need to switch to a previous stack.
	if s.GetCurrStackSize() == 0 {
		s.CurrentStackIndex--
	}
	s.Stacks[s.CurrentStackIndex].Values = s.Stacks[s.CurrentStackIndex].Values[0 : s.GetCurrStackSize()-1]

	// If after Pop, the stack is empty, switch to a previous stack.
	if s.GetCurrStackSize() == 0 {
		s.CurrentStackIndex--
	}
}

func (s *StacksOfPlates) AddStack() {
	s.Stacks = append(s.Stacks, &StackOfPlates{})
}

func (s *StacksOfPlates) Peek() int {
	return s.Stacks[s.CurrentStackIndex].Values[s.GetCurrStackSize()-1]
}

// popAt
// Implement a function popAt (int index) which performs a pop operation on a specific sub-stack.
/**


8  4  3
5  2  8  2
1  9  7  1

if popAt(0), we will have:
_  4  3
5  2  8  2
1  9  7  1

Works as Pop if on a last stack
However, we need to fill in the empty slot, if not a last stack.
We can do this shifting all stack elements down. However, as far as this is a stack, we need to Peek & Pop element from the next stack. Thus moving 1 element only from each stack - top plates!

Example:
8  (_) < 3
5  2  8 < 2
1  9  7  1

8  3  2
5  2  8
1  9  7  1


ALGORITHM:
1st element of next stack got Pushed to the stack with a missing top el.
We switched to the next stack.
Removed 1st element
Repeat procedure, if this stack is not the last one.
*/
func (s *StacksOfPlates) PopAt(idx int) {
	// Validation if last stack.
	if idx == len(s.Stacks)-1 {
		// We need to pop an element from the last stack. No extra work required.
		s.Pop()
	} else {
		// Need to start the shifts.
		s.PopShift(idx)
	}
}

// PopShift iterates through all stacks starting from the given one,
// runs Pop() on a current stack,
// assigns top (Peek'ed) element from the next stack.
func (s *StacksOfPlates) PopShift(idx int) {

	// Loop through remaining stacks including the current one.
	for i := idx; i <= len(s.Stacks)-1; i++ {

		// Check for a case where last stack may become empty when a last plate is removed.
		if len(s.Stacks[i].Values) > 0 {
			// Remove the top plate from the current stack of plates.

			fmt.Println("Working on stack ", i)

			s.Stacks[i].Values = s.Stacks[i].Values[0 : len(s.Stacks[i].Values)-1]

			//fmt.Printf("\n\ncurrent els: %+v \n\n", s.Stacks[i].Values)

			// Take a plate from the next stack...

			// Check if next stack exists or we are at a current stack.
			if i+1 < len(s.Stacks) {

				el := s.GetTopEl(i + 1)

				//fmt.Printf("\n\nel: %+v \n\n", el)

				// ...and put over the current stack.
				s.Stacks[i].Values = append(s.Stacks[i].Values, el)

				//fmt.Printf("\n\n%+v \n\n", s.Stacks[i].Values)
			} else {
				fmt.Printf("Last stack (%d) reached.\n\n", i)
			}
		}

		// Loop end. Start working with a next stack of plates.
	}
}

func (s *StacksOfPlates) GetTopEl(idx int) int {
	return s.Stacks[idx].Values[len(s.Stacks[idx].Values)-1]
}

// 3.4 Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks.
/**
EXAMPLE

given:
3
2
1

add 4:
4
3
2
1


remove: == need to remove 1

cycle through arr1 to pop -push elements to arr2:
1
2
3
4

return 1 via peek
pop it
run cycle on remaining arr2 to re-create arr1:
4
3
2

*/
type MyQueue struct {
	s1 []int
	s2 []int
}

func (q *MyQueue) Add(i int) {
	q.s1 = append(q.s1, i)
}

func (q *MyQueue) Remove() {

	// re-create
	q.s2 = []int{}

	for _, v := range q.s1 {
		q.s2 = append(q.s2, v)
	}

	q.s2 = q.s2[0 : len(q.s2)-1]
	q.s1 = q.s2
}

func (q *MyQueue) Peek() int {
	return q.s2[len(q.s2)-1]
}
