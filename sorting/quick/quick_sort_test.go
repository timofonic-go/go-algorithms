package quick

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {

	arr := []int{2, 3, 5, 0, 1}
	actual := QuickSort(arr)

	if !sort.IntsAreSorted(actual) {
		t.Error("Quick sort failed!")
		t.Logf("Actual value: %v", actual)
	}

}
