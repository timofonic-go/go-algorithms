package recursion_and_dynamic_programming

import "testing"

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
