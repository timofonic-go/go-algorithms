package recursion_and_dynamic_programming

import "fmt"

// Find nth fibonacci number
/**
	I'll be using a top-to-bottom approach with memorization from the 8th chapter.

	EXAMPLE
	idx	1  2  3  4  5  6  7
	numbers	0, 1, 2, 3, 5, 8, 13

	Let's say I need 4th number (should return 5):


						fib(5)

				fib(4)					fib(3)

	        fib(3)		   	  fib(2)      		fib(2)		fib(1)
	fib(2)	       fib(1)	    fib(1)     fib(0)	   fib(1)   fib(0)
      fib(1)  fib(0)


*/
func FindFibonacciTopDown(i int) int {

	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	m := make(map[int]int)

	return fib(i-1, m) + fib(i-2, m)
}

func fib(i int, m map[int]int) int {

	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	_, ok := m[i]
	if !ok {
		m[i] = fib(i-1, m) + fib(i-2, m)
	}

	return m[i]
}

// 3.1 Triple Step
/**
EXAMPLE

How many steps hopping:
| 1 2 3
1 steps = 3: 1,1,1
2 steps = 0: _
3 steps = 1: 1
Total: 4

| 1 2 3 4
1 steps = 4: 1,1,1,1
2 steps = 2: 2, 2
3 steps = 0: _
Also: 1,3
Total: 7

| 1 2 3 4 5
1 steps = 5: 1,1,1,1,1
2 steps = 2: _
3 steps = 0: _
But also:
1,1,3
1,3,1
3,3,1
2,2,1
2,1,2
1,2,2
2,3
3,2
Total: 13


BRUTE-FORCE
Just loop through the whole array and sum up all routes.
*/

func CountWays(n int) int {
	if n < 0 {
		// Number of hops does not match the length
		return 0
	} else if n == 0 {
		// Number of hops allows to "land" on the last step
		return 1
	} else {
		// Have not reached the final step yet, continue hopping.
		return CountWays(n-1) + CountWays(n-2) + CountWays(n-3)
	}
}

// 8.6 Hanoi Tower
//
// In the classic problem of the Towers of Hanoi, you have 3 towers and N disks of
// different sizes which can slide onto any tower.
//
// The puzzle starts with disks sorted in ascending order of size from top to bottom
// (i.e., each disk sits on top of an even larger one). You have the following constraints:
//
// (1) Only one disk can be moved at a time.
// (2) A disk is slid off the top of one tower onto another tower.
// (3) A disk cannot be placed on top of a smaller disk.
//
// Write a program to move the disks from the first tower to the last using Stacks.

func HanoiTower() []*Tower {

	// rods
	n := 3
	towers := make([]*Tower, n)

	for i := 0; i < 3; i++ {
		towers[i] = NewTower(i)
	}

	for i := n - 1; i >= 0; i-- {
		towers[0].add(i)
	}

	towers[0].moveDisks(n, towers[2], towers[1])

	return towers
}

type Tower struct {
	disks Stack
	index int
}

func NewTower(i int) *Tower {
	return &Tower{
		index: i,
	}
}

func (t *Tower) add(d int) {
	if !t.disks.isEmpty() && t.disks.peek() <= d {
		fmt.Printf("Error placing disk %d ! \n", d)
	} else {
		t.disks.push(d)
	}
}

func (t *Tower) moveToTop(twr *Tower) {
	top := t.disks.pop()
	twr.add(top)
}

func (t *Tower) moveDisks(n int, destination, buffer *Tower) {
	if n > 0 {
		t.moveDisks(n-1, buffer, destination)
		t.moveToTop(destination)
		buffer.moveDisks(n-1, destination, t)
	}
}
