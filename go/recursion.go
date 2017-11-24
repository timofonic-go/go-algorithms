package main

import (
	"fmt"
)

// Factorial with a recursive function.
func fact(n int) int {

	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(4))
}
