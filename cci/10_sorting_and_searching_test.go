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

// 10.1 Sorted Merge
func TestSortedMerge(t *testing.T) {

	a := make([]int, 6)
	a[0] = 1
	a[1] = 3
	a[2] = 5

	//t.Logf("Array a: %+v \n\n", a)

	b := []int{2, 4, 6}
	//t.Logf("Array b: %+v \n\n", b)

	actual := SortedMerge(a, b, 3, 3)

	//t.Logf("Sorted array: %+v \n\n", actual)

	if !sort.IntsAreSorted(actual) {
		t.Errorf("SortedMerge failed. Actual result: %#v", actual)
	}
}
