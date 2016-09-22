package cci

import (
	"fmt"
	"github.com/mpmlj/go-algorithms/util"
	"strings"
)

// 1.1 Is Unique
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

// 1.1 Is Unique w/o additional data type
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

// 1.2 Check Permutation
func IfPerm(s1, s2 string) bool {
	s1Sort := util.SortString(s1)
	s2Sort := util.SortString(s2)

	if s1Sort == s2Sort {
		return true
	}

	return false
}

// 1.3 URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)
//
// s := "Mr John Smith     "
func URLify(s string) string {

	var out string
	arr := strings.Split(s, "")

	// Flag to implement right trim. Keeps true if right part still hold non-space values.
	currentSpace := true

	for i := len(arr) - 1; i >= 0; i-- {
		c := string(arr[i])
		if c == " " {
			if !currentSpace {
				out = "%20" + out
			}
		} else {
			currentSpace = false
			out = c + out
		}
	}

	return out
}

// 1.4 Palindrome Permutation:
// - Given a string, write a function to check if it is a permutation of a palindrome.
// - A palindrome is a word or phrase that is the same forwards and backwards.
// - A permutation is a rearrangement of letters.
// - The palindrome does not need to be limited to just dictionary words.
func PalindromePerm(s string) bool {

	// Algorithm:
	//
	// 1.	Palindromes may use capital letters and spaces.
	// 	To compare left and right sizes we need to normalize the string first,
	// 	i.e. convert a string to all-lowercase characters and remove spaces.
	//
	// 2.	If this string is really a permutation of a palindrome, and according to the task,
	// 	does not need to be limited by dictionary words,
	// 	then we should just check if the string can be symmetrical.
	//
	// 3. 	I assume that number of symbols >= 2.

	// Convert to lower case
	st := strings.ToLower(s)
	fmt.Println("lower case string: ", st)

	// If it was a palindrome, it should have a set of characters, that we can arrange symmetrically.
	// This is possible when:
	// 1. If even number of characters, quantity of each symbol should be even as well.
	// 2. If odd number of characters, one character should be odd, all others - odd.

	// Determine odd or even
	arrRaw := strings.Split(st, "")
	var arr []string
	for _, v := range arrRaw {
		if string(v) != " " {
			arr = append(arr, string(v))
		}
	}

	// Calculate quantity of each symbol
	symbMap := make(map[string]int)
	for _, v := range arr {
		_, ok := symbMap[v]
		if !ok {
			symbMap[v] = 1
		} else {
			symbMap[v]++
		}
	}

	fmt.Printf("symbMap: %v \n\n", symbMap)

	var result bool

	if len(arr)%2 == 0 {
		// Even. Quantity of each symbol should be even as well.
		allEven := true
		for _, v := range symbMap {
			if v%2 != 0 {
				allEven = false
				break
			}
		}

		if allEven {
			return true
		}

	} else {
		// Odd. One character should be odd, all others - odd.
		var qtyOdd int
		for _, v := range symbMap {
			if v%2 != 0 {
				qtyOdd++
				if qtyOdd >= 2 {
					break
				}
			}
		}

		if qtyOdd == 1 {
			return true
		}
	}

	return result
}
