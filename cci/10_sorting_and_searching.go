package cci

/**

Bubble Sort | Runtime: 0( n2 ) average and worst case. Memory: 0( 1) .
In bubble sort, we start at the beginning of the array and swap the first two elements if the first is greater
than the second. Then, we go to the next pair, and so on, continuously making sweeps of the array until it is
sorted. In doing so, the smaller items slowly "bubble" up to the beginning of the list.

Quick Sort | Runtime: O(n log( n)) average, O(n2 ) worst case. Memory: 0( log(n)).
In quick sort we pick a random element and partition the array, such that all numbers that are less than the
partitioning element come before all elements that are greater than it. The partitioning can be performed
efficiently through a series of swaps (see below).

*/

/**

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
	remaining := middle - leftCursor
	for i := 0; i <= remaining; i++ {
		(*arr)[current+i] = helper[leftCursor+i]
	}
}
