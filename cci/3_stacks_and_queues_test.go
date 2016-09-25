package cci

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := Stack{}
	s.Push(3)
	s.Push(2)
	s.Push(1)

	if s.Values[0] != 3 || s.Values[1] != 2 || s.Values[2] != 1 {
		t.Error("TestStack_Push error!")
	}
}

func TestStack_Pop(t *testing.T) {
	s := Stack{}
	s.Push(3)
	s.Push(2)
	s.Push(1)

	s.Pop()
	s.Pop()

	if len(s.Values) != 1 || s.Values[0] != 3 {
		t.Error("TestStack_Pop error!")
	}
}

func TestSList_Size(t *testing.T) {
	s := Stack{}
	s.Push(3)
	s.Push(2)
	s.Push(1)

	expected := 3
	actual := len(s.Values)

	if expected != actual {
		t.Errorf("TestSList_Size error! Expected: %v, actual: %v", expected, actual)
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack{}
	s.Push(3)
	s.Push(2)
	s.Push(1)

	expected := 1
	actual := s.Peek()

	if expected != actual {
		t.Errorf("TestStack_Peek error! Expected: %v, actual: %v", expected, actual)
	}
}

func TestStack_IsEmptyFail(t *testing.T) {
	s := Stack{}
	s.Push(3)
	s.Push(2)

	expected := false
	actual := s.IsEmpty()

	if expected != actual {
		t.Errorf("TestStack_IsEmptyFail error! Expected: %v, actual: %v", expected, actual)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := Stack{}

	expected := true
	actual := s.IsEmpty()

	if expected != actual {
		t.Errorf("TestStack_IsEmpty error! Expected: %v, actual: %v", expected, actual)
	}
}
