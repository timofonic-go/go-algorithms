package main

import "fmt"

var (
	numbers   = []int{1, 2, 3, 4, 5, 6, 7, 8}
	searchFor = 5
)

func main() {

	fmt.Println("Numbers: ", numbers)
	fmt.Println("Searching for: ", searchFor)

	fmt.Println("Found at key: ", BinarySearch(numbers, searchFor))
}

// BinarySearch runs a binary search on provided array and returns an index of a value if found, or -1 if not found.
func BinarySearch(arr []int, searchFor int) int {

	return binarySearch(arr, searchFor, 0, len(arr)-1)
}

// binarySearch runs a binary search on an array slice from "min" to "max" index.
func binarySearch(arr []int, searchFor, min, max int) int {

	// Validation.
	if min > max {
		return -1
	}

	// Calculate a slice middle.
	mid := (min + max) / 2

	// Continue search recursion based on a bucket where result is found.
	switch {
	case searchFor < arr[mid]:
		return binarySearch(arr, searchFor, min, mid-1)
	case searchFor > arr[mid]:
		return binarySearch(arr, searchFor, mid+1, max)
	default:
		return mid
	}
}
