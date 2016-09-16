package binary

// BinarySearch implements binary search algo.
func BinarySearch(arr []int, searchFor int) int {

	min := 0
	max := len(arr) - 1

	return binarySearch(arr, searchFor, min, max)
}

// binarySearch function variation for recursive runs.
// Returns an index of an item found or -1 if not found.
func binarySearch(arr []int, searchFor, min, max int) int {

	if max < min {
		return -1
	}

	// Determine the middle.
	mid := (min + max) / 2

	// Recursively divide array in halves until either value is found or impossible to divide further.
	switch {
	case searchFor < arr[mid]:
		return binarySearch(arr, searchFor, min, mid-1)
	case searchFor > arr[mid]:
		return binarySearch(arr, searchFor, mid+1, max)
	default:
		return mid
	}
}
