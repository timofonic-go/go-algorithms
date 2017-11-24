package hanoi_towers

import "fmt"

// HanoiTower implements CCI 8.6 Hanoi Tower problem.
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

	// Init array of towers
	n := 3
	towers := make([]*Tower, n)

	for i := 0; i < 3; i++ {
		towers[i] = NewTower(i)
	}

	// Fill in a first tower with disks in ascending order
	for i := n - 1; i >= 0; i-- {
		towers[0].add(i)
	}

	// Start the process of moving disks
	towers[0].moveDisks(n, towers[2], towers[1])

	return towers
}

type Tower struct {
	disks Stack // contains data as []int to store disks
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

// moveToTop takes a disks from a top of the current tower to a new tower
// need to move remaining, larger disks into a bottom level of new towers
func (t *Tower) moveToTop(twr *Tower) {
	top := t.disks.pop()
	twr.add(top)
}

// moveDisks implements logic of moving disks
func (src *Tower) moveDisks(n int, destination, buffer *Tower) {
	if n > 0 {

		// Move all top disks except for the one on the bottom
		// Moves from from the origin t to the buffer, using destination as a buffer
		src.moveDisks(n-1, buffer, destination)

		// Move remaining, largest disk into destination
		src.moveToTop(destination)

		// Move remaining n-1 disks, from the top of the buffer to the destination.
		// We are moving disks from the current buffer, hence buffer.moveDisks(...
		// We are moving disks using current tower as a buffer, hence src in ...,destination, src)
		buffer.moveDisks(n-1, destination, src) // NOTE! Always forgetting to use src here!

	}
}
