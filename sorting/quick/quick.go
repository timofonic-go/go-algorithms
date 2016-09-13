package main

/**
https://en.wikipedia.org/wiki/Quicksort
https://github.com/mpmlj/go-algorithms/blob/master/sorting/quick_sort.go
*/

import (
	"fmt"
	"math/rand"
)

var (
	arr = []int{1, 6, 4, 6, 7, 9, 3, 5, 0, 0}
)

func main() {

	fmt.Println("Unsorted array: ", arr)

	res := QuickSort(arr)

	fmt.Println("Sorted array: ", res)
}

func QuickSort(arr []int) []int {

	// Restrict sort depth.
	if len(arr) <= 1 {
		return arr
	}

	// Set pivot value to a random value.
	pivot := arr[rand.Intn(len(arr))]

	// Init buckets
	left := make([]int, 0, len(arr))
	mid := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	// Compare each array element to a pivot.
	for _, v := range arr {

		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			mid = append(mid, v)
		case v > pivot:
			right = append(right, v)
		}

	}

	// Continue recursively sorting each of the "side-buckets".
	left = QuickSort(left)
	right = QuickSort(right)

	// Merge all buckets into a final value.
	left = append(left, mid...)
	left = append(left, right...)

	return left
}
