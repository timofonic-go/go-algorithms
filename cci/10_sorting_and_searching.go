package cci

/**

Bubble Sort | Runtime: 0( n2 ) average and worst case. Memory: 0( 1) .
In bubble sort, we start at the beginning of the array and swap the first two elements if the first is greater
than the second. Then, we go to the next pair, and so on, continuously making sweeps of the array until it is
sorted. In doing so, the smaller items slowly "bubble" up to the beginning of the list.

Merge Sort | Runtime: 0 ( n log ( n)) average and worst case. Memory: Depends.
Merge sort divides the array in half, sorts each of those halves, and then merges them back together. Each
of those halves has the same sorting algorithm applied to it. Eventually, you are merging just two singleelement
arrays. It is the "merge" part that does all the heavy lifting.

Quick Sort | Runtime: O(n log( n)) average, O(n2 ) worst case. Memory: 0( log(n)).
In quick sort we pick a random element and partition the array, such that all numbers that are less than the
partitioning element come before all elements that are greater than it. The partitioning can be performed
efficiently through a series of swaps (see below).

*/

func MergeSort(arr *[]int) {
	helper := make([]int, len(*arr))
	mergeSort(arr, &helper, 0, len(*arr)-1)
}

func mergeSort(arr, helper *[]int, low, high int) {
	if low < high {
		middle := (low + high) / 2
		mergeSort(arr, helper, low, middle)
		mergeSort(arr, helper, middle+1, high)
		merge(arr, helper, low, middle, high)
	}
}

func merge(arr, helper *[]int, low, middle, high int) {

	// Copy both halves into a helper array*/
	for i := low; i <= high; i++ {
		(*helper)[i] = (*arr)[i]
	}

	helperLeft := low
	helperRight := middle + 1
	current := low

	// Iterate through helper array.
	// Compare the left and right half, copying back the smaller element
	// from the two halves into the original array.
	for helperLeft <= middle && helperRight <= high {
		if (*helper)[helperLeft] <= (*helper)[helperRight] {
			(*arr)[current] = (*helper)[helperLeft]
			helperLeft++
		} else {
			// If right element is smaller than left element
			(*arr)[current] = (*helper)[helperRight]
			helperRight++
		}
		current++
	}

	// Copy the rest of the left side of the array into the target array
	remaining := middle - helperLeft
	for i := 0; i <= remaining; i++ {
		(*arr)[current+i] = (*helper)[helperLeft+i]
	}
}
