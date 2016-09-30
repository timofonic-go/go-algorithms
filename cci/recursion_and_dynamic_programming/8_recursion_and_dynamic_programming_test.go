package recursion_and_dynamic_programming

import (
	"sort"
	"testing"
)

/**
Find kth fibonacci number, top-down implementation with memorization.

idx (from 1)	1  2  3  4  5  6  7
fib. numbers	0, 1, 2, 3, 5, 8, 13

*/
func TestFindFibonacciTopDown(t *testing.T) {

	expected := 13
	actual := FindFibonacciTopDown(7)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

// 8.1 Triple Step.
var CountWaysTests = []struct {
	steps    int
	expected int
}{
	{
		3,
		4,
	},
	{
		4,
		7,
	},
	{
		5,
		13,
	},
}

func TestCountWays(t *testing.T) {

	for i, test := range CountWaysTests {
		actual := CountWays(test.steps)
		if test.expected != actual {
			t.Errorf("Test %d. Expected: %v, got: %v", i, test.expected, actual)
		}
	}
}

// 8.6 Hanoi Tower
func TestHanoiRun(t *testing.T) {

	//src := []int{2, 3, 1}

	actual := HanoiTower()

	if !sort.IntsAreSorted(actual.disks.data) {
		t.Error("Failed!")
	}
}
