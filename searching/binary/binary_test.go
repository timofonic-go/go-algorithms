package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	searchFor := 5 // Search value.
	expected := 4  // Index of a search value (5).

	actual := BinarySearch(array, searchFor)

	if actual != expected {

		t.Log("Expected value: ", expected)
		t.Log("Actual value: ", actual)
		t.Error("Incorrect search result!")
	}
}
