package main

import "fmt"

func fibonacci() func() int {

	// Init leading and next element.
	first, second := 0, 1

	return func() int {

		// Return current first element.
		ret := first

		// Recalculate 1st and 2nd values for the next step.
		first, second = second, first+second
		return ret
	}
}

func main() {

	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v \n", f())
	}
}
