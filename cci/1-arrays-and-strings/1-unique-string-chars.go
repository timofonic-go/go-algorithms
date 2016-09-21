// Is Unique: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?
package main

import "fmt"

func FindDupeNoExtraType(st string) bool {
	var c int
	var found bool
	for _, v := range st {
		s := string(v)
		c = 0
		for _, v2 := range st {
			s2 := string(v2)
			if s == s2 {
				c++
				if c > 1 {
					found = true
					break
				}
			}
		}
	}

	return found
}

func FindDupe(st string) bool {

	var found bool
	chars := make(map[string]struct{})

	for _, v := range st {
		s := string(v)
		_, ok := chars[s]
		if ok {
			found = true
			break
		}
		chars[s] = struct{}{}
	}

	return found
}

func main() {

	st := "abcdee"

	if FindDupe(st) {
		fmt.Println("Contains duplicates")
	} else {
		fmt.Println("All characters are unique")
	}
}
