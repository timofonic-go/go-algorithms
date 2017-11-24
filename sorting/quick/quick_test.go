package main

import (
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	actual := QuickSort(arr)

	sorted := sort.IntsAreSorted(actual)

	if !sorted {
		t.Error("Insertion sort failed!")

		expected := make([]int, len(arr))
		copy(expected, arr)
		sort.Ints(expected)
		t.Log("Expected: ", expected)
		t.Log("Actual: ", actual)
	}
}

func BenchmarkInsertionSort(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_ = QuickSort(arr)

	}
}
