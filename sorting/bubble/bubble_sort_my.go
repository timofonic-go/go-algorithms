package main

/*
 * Bubble sort - http://en.wikipedia.org/wiki/Bubble_sort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
	arr := utils.RandArray(10)
	//fmt.Println("Initial array is:", arr)
	//fmt.Println("")
	//
	tmp := 0
	//
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr)-1; j++ {
	//		if arr[j] > arr[j+1] {
	//			tmp = arr[j]
	//			arr[j] = arr[j+1]
	//			arr[j+1] = tmp
	//		}
	//	}
	//}
	//
	//fmt.Println("Sorted array is: ", arr)

	//arr = utils.RandArray(10)
	fmt.Println("Initial array 2 is:", arr)
	fmt.Println("")

	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(arr) - 1; i++ {
			if arr[i] > arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
                            swapped = true
			}
		}
	}

	fmt.Println("Sorted array 2 is: ", arr)
}
