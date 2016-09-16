package quick

import (
	"math/rand"
)

func QuickSort(arr []int) []int {

	length := len(arr)

	// Restrict run depth.
	if length < 2 {
		return arr
	}

	// Set a pivot.
	pivot := arr[rand.Intn(length)]

	// Prepare buckets.
	left := make([]int, 0, length)
	middle := make([]int, 0, length)
	right := make([]int, 0, length)

	// Split values into the buckets.
	for i := 0; i < length-1; i++ {
		switch {
		case arr[i] < pivot:
			left = append(left, arr[i])
		case arr[i] > pivot:
			right = append(right, arr[i])
		case arr[i] == pivot:
			middle = append(middle, arr[i])
		}
	}

	// Recurrently continue splitting each of the buckets till the end.
	left, right = QuickSort(left), QuickSort(right)

	// Join sort results.
	left = append(left, middle...)
	left = append(left, right...)

	return left
}
