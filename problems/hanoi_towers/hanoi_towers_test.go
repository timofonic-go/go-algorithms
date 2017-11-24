package hanoi_towers

import (
	"reflect"
	"testing"
)

// Hanoi Tower (CCI 8.6)
func TestHanoiRun(t *testing.T) {

	expected := []int{2, 1, 0}

	towers := HanoiTower()
	actual := towers[2].disks.data

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Failed!")
		t.Log("Expected:\n")
		t.Logf("%+v\n", expected)

		t.Log("Got:\n")
		t.Logf("%+v\n", actual)
	}
}
