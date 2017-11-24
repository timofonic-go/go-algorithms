package main

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {

	unsorted := map[string][]string{
		"0": []string{"1", "4"},
		"1": []string{"3", "5"},
		"2": []string{"5"},
		"3": []string{"5", "7"},
		"4": []string{},
		"5": []string{"6"},
		"6": []string{"7"},
		"7": []string{},
	}

	// There can be 2 solutions for this graph case and both are correct as we have 2 entry points: "0" and "2".
	expected := []string{"2", "0", "4", "1", "3", "5", "6", "7"}
	expected2 := []string{"0", "4", "1", "3", "2", "5", "6", "7"}

	actual := TopologicalSort(unsorted)

	if !reflect.DeepEqual(expected, actual) && !reflect.DeepEqual(expected2, actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
