package binary

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}

	expected := 3 // Index of digit "4"
	actual := BinarySearch(arr, 4)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestBinarySearch_Fail(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	expected := -1 // Not found
	actual := BinarySearch(arr, 6)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
