package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {

	// "dereferencing the pointer from its memory address to the current value at this address"
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("zeroptr: ", &i)

	a := 2
	b := 3
	fmt.Println(mult(&a, &b))
}

func mult(a, b *int) int {

	return *a * *b

}
