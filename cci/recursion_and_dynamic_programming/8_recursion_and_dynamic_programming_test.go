package recursion_and_dynamic_programming

import (
	"reflect"
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

	expected := []int{2, 1, 0}

	actual := HanoiTower()

	if !reflect.DeepEqual(expected, actual[2].disks.data) {
		t.Error("Failed!")
		t.Log("Expected:\n")
		t.Logf("%+v\n", expected)

		t.Log("Got:\n")
		t.Logf("%+v\n", actual[2].disks.data)
	}
}
