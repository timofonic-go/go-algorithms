package cci

import (
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
}

func TestIsOneAway(t *testing.T) {

	for _, tt := range IsOneAwayTests {
		actual := IsOneAway(tt.s1, tt.s2)
		if actual != tt.expected {
			t.Errorf("s1: %v, s2: %v, expected: %v, actual: %v", tt.s1, tt.s2, tt.expected, actual)
		}
	}
}
