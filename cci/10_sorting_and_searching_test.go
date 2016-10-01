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
