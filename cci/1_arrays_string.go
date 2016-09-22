package cci

import (
	"fmt"
	"github.com/mpmlj/go-algorithms/util"
	"math"
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

// 1.5 OneAway
// 1. Three types of edits possible: insert, remove, replace
// 2. One or zero edits away
/*
Example:

Replacement
pale -> bale
we expect only 1 symbol on the same position to not match

Insertion/Removal
pale -> ple
same as ple -> pale


Algorithm.

1. If both strings are equal, this means zero away case, return true.

2. After #1 check, if both strings are of equal, then as per the task, we expect only 1 replacement,
i.e. we expect only 1 different symbol.

2. After #2, if string sizes are different, that means that:
- either an extra symbol was inserted to a beginning or an end of one string
- or it was inserted in the middle

In case it was inserted in the end, then smaller string should fit in to beginning of the larger string, with just a different symbol remaining in the end.
[pal]
[pal](e)

In case, a symbol was inserted in the beginning or a middle, and because all symbols are in the same order,
if skipping that symbol, the remaining parts of strings should match.
(p)[ale]
(b)[ale]
or
p(a)[le]
p(i)[le]

So the logic is to reach the different symbol looping through a smaller string.
If all symbols are still equals - we have case one.
If we stumbler upon a mismatching symbol, we compare the remaining parts.

s1 := ple
s2 := pale, reference

i=0, s1[0] = s2[0], p = p
i=1, s1[1] != s2[1], l != a
OK, as far we assume that we have only 1 mismatch, then after throwing away this mismatching symbol,
remaining parts should match again.

i=2, s1[1:] = s2[2:], i.e. "le" == "le" (we skipped "a" from the reference string)


*/
func IsOneAway(s1, s2 string) bool {

	// Case when zero edits away
	if s1 == s2 {
		return true
	}

	s1len := len(s1)
	s2len := len(s2)

	// At this point:
	// If s1len == s2len - then we expect 1 replacement only
	// If s2 > s1 - then we expect string 2 to have 1 extra character inserted at any position

	if s1len == s2len {

		diffCnt := 0
		for i := 0; i < s2len-1; i++ {
			if s2[i] != s1[i] {
				diffCnt++
				if diffCnt > 1 {
					return false
				}
			}
		}

		return true

	} else {

		// Validate strings, to be only 1 symbol away maximum.
		if math.Abs(float64(s1len-s2len)) > 1 {
			return false
		}

		// Arrange strings into a working and reference strings,
		// i.e. s1 - always a small string, s2 - always a large string.
		if s1len > s2len {
			s1, s2 = s2, s1
			s1len, s2len = s2len, s1len
		}

		// Logic recap:
		// 1. Find first breaking point
		// 2. If not reached - we are good: pal -> pale
		// 3. If reached - remaining parts should match: pale -> bale, matching: ale and ale
		match := true

		for i := 0; i < s1len-1; i++ {
			if s1[i] != s2[i] {
				// Does not match.
				// We assume that we should just jump over a mismatching symbol and the rest of the string should match.
				if s1[i:] != s2[i+1:] {
					match = false
					break
				}
			}
		}

		// If remaining parts are mismatching, the return false.
		if !match {
			return false
		}

		// OK. We reached the end of a smaller string and no mismatching symbols are found yet.
		// As per the task we assume this is the one remaining symbol,
		// i.e only 1 symbol away, so we can return true.
		return true
	}

	return false
}
