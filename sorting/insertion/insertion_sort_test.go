package insertion

import (
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	arr := []int{2, 3, 5, 0, 1}
	actual := InsertionSort(arr)

	if !sort.IntsAreSorted(actual) {
		t.Error("Quick sort failed!")
		t.Logf("Actual value: %v", actual)
	}

}
