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