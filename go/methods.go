package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

// *rect is a "receiver"
func (r *rect) area() int {
	return r.width * r.height
}

// receiver can be either a pointer or a value
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{
		width:  10,
		height: 5,
	}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	r2 := &r

	fmt.Println("area: ", r2.area())
	fmt.Println("perim: ", r2.perim())
}
