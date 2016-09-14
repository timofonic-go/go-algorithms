// Package main provides a sorting using a "quick sort" algo
package main

/*
 * Quick sort - https://en.wikipedia.org/wiki/Quicksort
 */

import (
	"fmt"
	"math/rand"
)

func quick_sort3(arr []int) []int {

	// Limit sort depth
	if len(arr) <= 1 {
		return arr
	}

	// Set pivot
	pivot := arr[rand.Intn(len(arr))]

	// Init buckets
	left := make([]int, 0, len(arr))
	middle := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	// Passes
	for _, v := range arr {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		case v > pivot:
			right = append(right, v)
		}
	}

	// Continue sorting recursively
	left = quick_sort3(left)
	right = quick_sort3(right)

	// Combine the result
	left = append(left, middle...)
	left = append(left, right...)

	return left
}

func main() {
	arr := []int{1, 7, 4, 9, 0, 3, 5, 7, 6, 2}
	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", quick_sort3(arr))

	// Expected: 0, 1, 2, 3, 4, 5, 6, 7, 7, 9
}

// ORIG: Sorted array is:  [0 1 2 3 4 5 6 7 7 9]
