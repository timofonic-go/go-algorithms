package main

/*
 * Quick sort - https://en.wikipedia.org/wiki/Quicksort
 */

import "fmt"
import "math/rand"

import "github.com/0xAX/go-algorithms"

func quick_sort2(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[rand.Intn(len(arr))]

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < pivot:
			low_part = append(low_part, item)
		case item == pivot:
			middle_part = append(middle_part, item)
		case item > pivot:
			high_part = append(high_part, item)
		}
	}

	low_part = quick_sort2(low_part)
	high_part = quick_sort2(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}

func main() {
	arr := utils.RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", quick_sort2(arr))
}

// 5 3 5 5 1 0

// pivot: 5

// 5: m[5]
// 3: l[3]
// 5: m[5, 5]
// 5: m[5, 5, 5]
// 1: l[3, 1]
// 0: l[3, 1, 0]

// low_part start: 3, 1, 0
// pivot: 0
// 3: r[3]
// 1: r[3, 1]
// 0: m[0]
// low_part: 0, 1, 3


// high_part start: []
