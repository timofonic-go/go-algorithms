package cci

import (
	"github.com/mpmlj/go-algorithms/util"
	"strings"
)

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

func IfPerm(s1, s2 string) bool {
	s1Sort := util.SortString(s1)
	s2Sort := util.SortString(s2)

	if s1Sort == s2Sort {
		return true
	}

	return false
}

// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
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
