package problems

import (
	"reflect"
	"testing"
)

func TestNumberEmployeeUnderManager(t *testing.T) {

	// Map of direct employees of managers:
	// employee "A" is reporting to manager "C",
	// employee "B" is reporting to manager "C" and so on.
	given := map[string]string{
		"A": "C",
		"B": "C",
		"C": "F",
		"D": "E",
		"E": "F",
		"F": "F",
	}

	// Map of "manager" -> total count of his/her reporting employees.
	expected := map[string]int{
		"A": 0,
		"B": 0,
		"C": 2,
		"D": 0,
		"E": 1,
		"F": 5,
	}

	actual := NumberOfReports(given)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
