package hanoi_tower

import (
	"testing"
	"sort"
)

func TestHanoiRun(t *testing.T) {

	src := []int{2, 3, 1}

	actual := HanoiRun(src)

	if !sort.IntsAreSorted(actual) {
		t.Error("Failed!")
	}
}
