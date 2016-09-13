package main

/**
https://en.wikipedia.org/wiki/Insertion_sort
https://rosettacode.org/wiki/Sorting_algorithms/Insertion_sort#Go
*/

import (
	"fmt"
)

var (
	arr = []int{1, 6, 4, 6, 7, 9, 3, 5, 0, 0}
)

func main() {

	fmt.Println("Unsorted array: ", arr)

	res := InsertionSort(arr)

	fmt.Println("Sorted array: ", res)
}

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {

		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {

			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp

		}
	}

	return arr
}
