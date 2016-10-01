package cci

import (
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	actual := []int{14, 33, 27, 10, 35, 19, 42, 44}
	MergeSort(&actual)

	if !sort.IntsAreSorted(actual) {
		t.Errorf("MergeSort failed. Actual result: %#v", actual)
	}
}

func TestQuickSort(t *testing.T) {
	given := []int{14, 33, 27, 10, 35, 19, 42, 44}
	actual := QuickSort(given)

	if !sort.IntsAreSorted(actual) {
		t.Errorf("QuickSort failed. Actual result: %#v", actual)
	}
}

func TestBubbleSort(t *testing.T) {
	given := []int{14, 33, 27, 10, 35, 19, 42, 44}
	actual := BubbleSort(given)

	if !sort.IntsAreSorted(actual) {
		t.Errorf("BubbleSort failed. Actual result: %#v", actual)
	}
}

func TestBinarySearch(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	searchFor := 6
	expected := 5
	actual := BinarySearch(arr, searchFor)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
