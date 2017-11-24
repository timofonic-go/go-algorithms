package bubble

import (
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	arr := []int{3, 2, 5, 6, 1}

	actual := BubbleSort(arr)

	if !sort.IntsAreSorted(actual) {
		t.Error("Bubble sort error!")
	}
	t.Logf("\n %v \n", actual)

}
