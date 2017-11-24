package insertion

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {

			tmp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = tmp

		}
	}
	return arr
}
