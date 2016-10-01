package cci

import (
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	actual := []int{2, 4, 3, 5, 1}

	MergeSort(&actual)

	if !sort.IntsAreSorted(actual) {
		t.Errorf("MergeSort failed. Actual result: %#v", actual)
	}
}
