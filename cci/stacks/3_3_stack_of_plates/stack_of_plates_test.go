package stacks

import "testing"

func TestStackOfPlates_NewStack(t *testing.T) {

	s := StackOfPlates{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)

	t.Logf("Last stack index: %d", s.last)
	t.Logf("Stack %d: %+v", 0, s.stacks[0])
	t.Logf("Stack %d: %+v", 1, s.stacks[1])
	t.Logf("Stack %d: %+v", 2, s.stacks[2])

	expected := 2
	actual := s.last

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}

}

func TestStackOfPlates_Pop(t *testing.T) {

	s := StackOfPlates{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)

	s.pop()
	s.pop()
	el := s.pop()

	// Last stack should have index 1
	expected := 1
	actual := s.last

	t.Logf("Last index: %d", s.last)
	if expected != actual {
		t.Fatalf("Stack count | Expected %d, got %d", expected, actual)
	}

	expected = 6
	actual = el

	t.Logf("Last popped element: %d", el)
	if expected != actual {
		t.Errorf("Last el val | Expected %d, got %d", expected, actual)
	}
}

func TestStackOfPlatespopAt(t *testing.T) {

	s := StackOfPlates{}
	s.push(1) // stack 0
	s.push(2) // stack 0
	s.push(3) // stack 0

	s.push(4) // stack 1
	s.push(5) // stack 1
	s.push(6) // stack 1

	s.push(7) // stack 2

	expected := 6
	actual := s.popAt(1) // pop from stack 1, should be 6

	// expected:
	// el = 6
	// stack 2 collapsed, and s.last = 1

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}

	if s.last != 1 {
		t.Errorf("Expected %d, got %d", 1, s.last)
	}

}
