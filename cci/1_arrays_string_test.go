package cci

import (
	"fmt"
	"testing"
)

func TestFindDupe(t *testing.T) {
	st := "abcdee"

	if !FindDupe(st) {
		t.Error("Expected true, got false")
	}
}

func TestFindDupeFail(t *testing.T) {
	st := "abcde"

	if FindDupe(st) {
		t.Error("Expected false, got true")
	}
}

func TestFindDupeNoExtraType(t *testing.T) {
	st := "abcdee"

	if !FindDupeNoExtraType(st) {
		t.Error("Expected true, got false")
	}
}

func TestFindDupeNoExtraTypeFail(t *testing.T) {
	st := "abcde"

	if FindDupeNoExtraType(st) {
		t.Error("Expected false, got true")
	}
}

func TestIfPermPerm(t *testing.T) {
	s1 := "abcde"
	s2 := "bcdea"

	if !IfPerm(s1, s2) {
		t.Error("Expected true, got false")
	}
}

func TestIfPermNoPerm(t *testing.T) {
	s1 := "abcde"
	s2 := "abcde12345"

	if IfPerm(s1, s2) {
		t.Error("Expected false, got true")
	}
}

// Test with spaces in the end of a string.
func TestURLify(t *testing.T) {

	s := "Mr John Smith     "
	expected := "Mr%20John%20Smith"
	actual := URLify(s)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// Test w/o spaces in the end of a string.
func TestURLifyNoTrimReq(t *testing.T) {

	s := "Mr John Smith"
	expected := "Mr%20John%20Smith"
	actual := URLify(s)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// 1.4 Palindrome Permutation
func TestPalindromePermEven(t *testing.T) {
	s := "aba b" // From palindrome: abba

	if !PalindromePerm(s) {
		t.Error("Incorrect 1.4 Palindrome Permutation task result.")
	}
}

func TestPalindromePermOdd(t *testing.T) {
	s := "abab d" // From palindrome: abdba

	if !PalindromePerm(s) {
		t.Error("Incorrect 1.4 Palindrome Permutation task result.")
	}
}

func TestPalindromePermFail(t *testing.T) {
	s := "abcd" // Cannot create a palindrome from these symbols

	if PalindromePerm(s) {
		t.Error("Incorrect 1.4 Palindrome Permutation task result.")
	}
}

// 1.5 OneAway
var IsOneAwayTests = []struct {
	s1       string
	s2       string
	expected bool
}{
	{
		"pale",
		"pale",
		true,
	},
	{
		"pale",
		"ple",
		true,
	},
	{
		"pales",
		"pale",
		true,
	},
	{
		"pale",
		"bale",
		true,
	},
	{
		"pale",
		"bake",
		false,
	},
	{
		"pale",
		"pale12",
		false,
	},
	{
		"",
		"",
		true,
	},
}

func TestIsOneAway(t *testing.T) {

	for _, tt := range IsOneAwayTests {
		actual := IsOneAway(tt.s1, tt.s2)
		if actual != tt.expected {
			t.Errorf("s1: %v, s2: %v, expected: %v, actual: %v", tt.s1, tt.s2, tt.expected, actual)
		}
	}
}

// 1.6 String Compression
var StringCompressionTests = []struct {
	s1       string
	expected string
}{
	{
		"aabcccccaaa",
		"a2b1c5a3",
	},
	{
		"",
		"",
	},
}

func TestStringCompression(t *testing.T) {

	for _, tt := range StringCompressionTests {
		actual := StringCompression(tt.s1)
		if actual != tt.expected {
			t.Errorf("str: %v, expected: %v, actual: %v", tt.s1, tt.expected, actual)
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	/**
	Given:

	abc
	def
	ghi
	*/
	m := [3][3]string{}
	m[0][0] = "a"
	m[0][1] = "b"
	m[0][2] = "c"
	m[1][0] = "d"
	m[1][1] = "e"
	m[1][2] = "f"
	m[2][0] = "g"
	m[2][1] = "h"
	m[2][2] = "i"

	/**
	Expected:

	gda
	heb
	ifc
	*/

	actual := RotateMatrix(m)

	fmt.Print("\n\nRotated matrix:\n")

	if actual[0][0] != "g" ||
		actual[0][1] != "d" ||
		actual[0][2] != "a" ||
		actual[1][0] != "h" ||
		actual[1][1] != "e" ||
		actual[1][2] != "b" ||
		actual[2][0] != "i" ||
		actual[2][1] != "f" ||
		actual[2][2] != "c" {
		t.Error("1.7 Rotate Matrix | Error")
	}

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {

			fmt.Print(actual[i][j])
		}
		fmt.Println("")
	}
}
