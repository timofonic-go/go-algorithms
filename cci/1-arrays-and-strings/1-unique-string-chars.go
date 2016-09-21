// Is Unique: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?
package main

import "fmt"

func FindDupe(st string) bool {
	var c int
	var found bool
	for _, v := range st {
		s := string(v)
		c = 0
		//println(string(v))
		for _, v2 := range st {
			s2 := string(v2)
			if s == s2 {
				c++
				if c > 1 {
					found = true
					println(s)
					break
				}
			}
		}
	}

	return found
}

func main() {

	st := "abcdee"

	if FindDupe(st) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}
