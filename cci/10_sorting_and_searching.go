package cci

import (
	"math/rand"
)

/**

http://www.cprogramming.com/tutorial/computersciencetheory/sortcomp.html


Merge Sort | Runtime: 0 ( n log ( n)) average and worst case. Memory: Depends.

Merge sort divides the array in half, sorts each of those halves, and then merges them back together.
Each of those halves has the same sorting algorithm applied to it.
Eventually, you are merging just two single element arrays.
It is the "merge" part that does all the heavy lifting.

The merge method operates by copying all the elements from the target array segment into a helper array,
keeping track of where the start of the left and right halves should be (helperleft and rightCursor).
We then iterate through helper, copying the smaller element from each half into the array. At the end, we
copy any remaining elements into the target array.

MORE INFO:
https://www.tutorialspoint.com/data_structures_algorithms/merge_sort_algorithm.htm
http://quiz.geeksforgeeks.org/merge-sort/

*/
func MergeSort(arr *[]int) {

	mergeSort(arr, 0, len(*arr)-1)
}

// Main function that sorts arr[l..r] using merge()
func mergeSort(arr *[]int, low, high int) {
	if low < high {

		// Find the middle point
		middle := (low + high) / 2

		// Sort first and second halves
		mergeSort(arr, low, middle)
		mergeSort(arr, middle+1, high)

		// Merge the sorted halves
		merge(arr, low, middle, high)
	}
}

// Merges two subarrays of arr[].
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]
func merge(arr *[]int, low, middle, high int) {

	// Note: arr will be totally overwritten by the end of the function
	// with sorted values from helper.

	helper := make([]int, len(*arr))

	// Copy both halves into a helper array
	for i := low; i <= high; i++ {
		helper[i] = (*arr)[i]
	}

	// Set cursors to initial positions to scan two halves of a provided array.
	leftCursor := low
	rightCursor := middle + 1
	current := low

	// Iterate through helper array.
	// Compare the left and right half, copying back the smaller element
	// from the two halves into the original array.
	for leftCursor <= middle && rightCursor <= high {
		if helper[leftCursor] <= helper[rightCursor] {
			(*arr)[current] = helper[leftCursor]
			leftCursor++
		} else {
			// If right element is smaller than left element
			(*arr)[current] = helper[rightCursor]
			rightCursor++
		}

		// Cursor for the main array (arr []int).
		current++
	}

	// Copy the rest of the left side of the array into the target array.
	// Left side = because we were filling in starting from the left side
	// and right elements already there and match.
	// We continue from where we left: current cursor for the arr, and leftCursor for the helper.
	remaining := middle - leftCursor
	for i := 0; i <= remaining; i++ {
		(*arr)[current+i] = helper[leftCursor+i]
	}
}

/**

Quick Sort | Runtime: O( n log(n) ) average, O( n2 ) worst case. Memory: O( log(n) ).

In quick sort we pick a random element and partition the array, such that all numbers that are less than the
partitioning element come before all elements that are greater than it. The partitioning can be performed
efficiently through a series of swaps (see below).

If we repeatedly partition the array (and its sub-arrays) around an element, the array will eventually become
sorted. However, as the partitioned element is not guaranteed to be the median (or anywhere near the
median), our sorting could be very slow. This is the reason for the O(n2) worst case runtime.

INFO
https://www.khanacademy.org/computing/computer-science/algorithms/quick-sort/a/overview-of-quicksort

So why think about quicksort when merge sort is at least as good?
That's because the constant factor hidden in the big-Θ notation for quicksort is quite good.
In practice, quicksort outperforms merge sort, and it significantly outperforms selection sort and insertion sort.

*/

func QuickSort(arr []int) []int {

	length := len(arr)

	if length < 2 {
		return arr
	}

	pivot := arr[rand.Intn(length)]

	left := make([]int, 0, length)
	middle := make([]int, 0, length)
	right := make([]int, 0, length)

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

	left, right = QuickSort(left), QuickSort(right)

	left = append(left, middle...)
	left = append(left, right...)

	return left
}

/**
Bubble Sort | Runtime: 0( n2 ) average and worst case. Memory: 0( 1) .

In bubble sort, we start at the beginning of the array and swap the first two elements if the first is greater
than the second. Then, we go to the next pair, and so on, continuously making sweeps of the array until it is
sorted. In doing so, the smaller items slowly "bubble" up to the beginning of the list.

http://www.algolist.net/Algorithms/Sorting/Bubble_sort
*/
func BubbleSort(arr []int) []int {

	swapMore := true

	for swapMore {
		swapMore = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				swapMore = true
			}
		}
	}

	return arr
}

/**
Binary Search

In binary search, we look for an element x in a __sorted array__ by first comparing x to the midpoint of the array.
If xis less than the midpoint, then we search the left half of the array. If x is greater than the midpoint, then
we search the right half of the array. We then repeat this process, treating the left and right halves as subarrays.

Again, we compare x to the midpoint of this subarray and then search either its left or right side. We
repeat this process until we either find x or the subarray has size 0.

INFO
https://www.khanacademy.org/computing/computer-science/algorithms/binary-search/a/implementing-binary-search-of-an-array
*/

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

// 10.1 Sorted Merge: You are given two sorted arrays, A and B, where A has a large enough buffer at the
// end to hold B. Write a method to merge B into A in sorted order.
// Note that you don't need to copy the contents of A after running out of elements in B. They are already in place.
func SortedMerge(a, b []int, lastA, lastB int) []int {
	indexA := lastA - 1              // Index of last element in array a
	indexB := lastB - 1              // Index of last element in array b
	indexMerged := lastB + lastA - 1 // End of merged array

	// Merge a and b, starting from the last element in each
	for indexB >= 0 {
		// end of a is > than end of b
		if indexA >= 0 && a[indexA] > b[indexB] {

			a[indexMerged] = a[indexA] // copy element
			indexA--
		} else {

			a[indexMerged] = b[indexB] // copy element
			indexB--
		}
		indexMerged--
	}

	return a
}
