package bubble

// BubbleSort implements bubble sort algorithm.
func BubbleSort(arr []int) []int {

	swapMore := true

	for swapMore {
		swapMore = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = tmp
				swapMore = true
			}
		}
	}

	return arr
}
