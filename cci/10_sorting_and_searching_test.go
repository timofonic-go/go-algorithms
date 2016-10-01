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
