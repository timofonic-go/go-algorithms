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

func TestArr_Push(t *testing.T) {
	s0 := StackFixed{
		Capacity: 3,
		Idx:      0,
	}
	s1 := StackFixed{
		Capacity: 3,
		Idx:      3,
	}
	s2 := StackFixed{
		Capacity: 3,
		Idx:      6,
	}

	a := Arr{
		Sizes:  []int{0, 0, 0},
		Stacks: []StackFixed{s0, s1, s2},
		Values: make([]int, 9),
	}

	a.Push(0, 1)
	a.Push(1, 3)
	a.Push(1, 4)
	a.Push(2, 4)
	a.Push(2, 5)
	a.Push(2, 6)

	// Fill different stacks with different number of elements
	// Expected: [1, 0, 0, 3, 4, 0, 4, 5, 6]

	expected := []int{1, 0, 0, 3, 4, 0, 4, 5, 6}
	actual := a.Values

	for i := 0; i < len(expected); i++ {

		if expected[i] != actual[i] {
			t.Errorf("Bad TestArr_Push! Expected: %v, actual: %v", expected[i], actual[i])
		}

	}

}

func TestArr_Pop(t *testing.T) {
	s0 := StackFixed{
		Capacity: 3,
		Idx:      0,
	}
	s1 := StackFixed{
		Capacity: 3,
		Idx:      3,
	}
	s2 := StackFixed{
		Capacity: 3,
		Idx:      6,
	}

	a := Arr{
		Sizes:  []int{0, 0, 0},
		Stacks: []StackFixed{s0, s1, s2},
		Values: make([]int, 9),
	}

	a.Push(0, 1)
	a.Push(1, 3)
	a.Push(1, 4)
	a.Push(2, 4)
	a.Push(2, 5)
	a.Push(2, 6)
	a.Pop(0)
	a.Pop(1)
	a.Pop(2)

	// Fill different stacks with different number of elements
	// Expected: [1, 0, 0, 3, 4, 0, 4, 5, 6]

	expected := []int{0, 0, 0, 3, 0, 0, 4, 5, 0}
	actual := a.Values

	for i := 0; i < len(expected); i++ {

		if expected[i] != actual[i] {
			t.Errorf("Bad TestArr_Pop! Expected: %v, actual: %v", expected[i], actual[i])
		}

	}
}

func TestArr_Peek(t *testing.T) {
	s0 := StackFixed{
		Capacity: 3,
		Idx:      0,
	}
	s1 := StackFixed{
		Capacity: 3,
		Idx:      3,
	}
	s2 := StackFixed{
		Capacity: 3,
		Idx:      6,
	}

	a := Arr{
		Sizes:  []int{0, 0, 0},
		Stacks: []StackFixed{s0, s1, s2},
		Values: make([]int, 9),
	}

	a.Push(0, 1)
	a.Push(1, 3)
	a.Push(1, 4)
	a.Push(2, 4)
	a.Push(2, 5)
	a.Push(2, 6)

	expected := []int{1, 4, 6}

	for i := 0; i < len(expected); i++ {

		actual := a.Peek(i)

		if expected[i] != actual {
			t.Errorf("Bad TestArr_Peek! Expected: %v, actual: %v", expected[i], actual)
		}

	}
}

// 3.2 Stack Min
func TestStack_Min(t *testing.T) {
	s := Stack{}
	s.Push(3)
	s.Push(2)
	s.Push(6)
	s.Push(4)
	s.Push(1)
	s.Pop()

	expected := 2
	actual := s.GetMin()

	if expected != actual {
		t.Errorf("TestStack_Min error! Expected: %v, actual: %v", expected, actual)
	}
}

// 3.3 Stack Of Plates
func TestStacksOfPlates(t *testing.T) {
	s := StacksOfPlates{
		Capacity: 3,
	}

	s.Push(1)
	s.Push(5)
	s.Push(8)

	s.Push(9)
	s.Push(2)
	s.Push(4)

	s.Push(7)
	s.Pop()

	expectedStackIdx := 1
	expectedStackEl := 4

	actualStackIdx := s.GetCurrStackIdx()
	actualStackEl := s.Peek()

	if expectedStackIdx != actualStackIdx {
		t.Errorf("TestStacksOfPlates error! Expected idx: %v, actual idx: %v", expectedStackIdx, actualStackIdx)
	}

	if expectedStackEl != actualStackEl {
		t.Errorf("TestStacksOfPlates error! Expected el: %v, actual el: %v", expectedStackEl, actualStackEl)
	}
}

var StackPopAtTests = []struct {
	stackIndex int
	expected   int
}{
	{
		0,
		8,
	},
	{
		1,
		3,
	},
	{
		2,
		2,
	},
	{
		3,
		1,
	},
}

func TestStacksOfPlates_PopAt(t *testing.T) {

	s := StacksOfPlates{
		Capacity: 3,
	}

	s.Push(1)
	s.Push(5)
	s.Push(8)

	s.Push(9)
	s.Push(2)
	s.Push(4)

	s.Push(7)
	s.Push(8)
	s.Push(3)

	s.Push(1)
	s.Push(2)

	s.PopAt(1) // Expected to Pop el 4

	// After this change:
	// - Peek() from stack 1 should return Peek from stack 2, i.e. 3
	// - Peek() from stack 2 should return Peek from stack 3, i.e. 2
	// - Peel() from stack 3 should return 1

	for _, v := range StackPopAtTests {
		actual := s.GetTopEl(v.stackIndex)
		if v.expected != actual {
			t.Errorf("Stack %d, expected %v, got %v", v.stackIndex, v.expected, actual)
		}
	}
}

func TestMyQueue(t *testing.T) {

	q := &MyQueue{}
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Remove()

	actual := q.Peek()
	expected := 3

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStackQueue(t *testing.T) {

	q := NewStackQueue()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Add(5)

	q.Pop()
	q.Pop()
	el := q.Peek()

	expected := 3
	actual := el

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStack_Sort(t *testing.T) {

	s := &Stack2{}
	s.Push(4)
	s.Push(3)
	s.Push(1)
	s.Push(5)
	s.Push(6)
	s.Push(2)

	s.Sort()

	expectedArr := []int{6, 5, 4, 3, 2, 1}

	for i, expected := range expectedArr {
		actual := s.Values[i]
		if expected != actual {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

// 3.6 Animal Shelter
func TestAnimalQueue_Enqueue(t *testing.T) {

	dog1 := &Dog{}
	dog1.SetName("Norton")

	cat1 := &Cat{}
	cat1.SetName("Musia")

	dog2 := &Dog{}
	dog2.SetName("Memphis")

	cat2 := &Cat{}
	cat2.SetName("Jessie")

	q := AnimalQueue{}
	q.Enqueue(dog1)
	q.Enqueue(cat1)
	q.Enqueue(dog2)
	q.Enqueue(cat2)

	expected := "Norton"
	actual := q.DequeAny().GetName()
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAnimalQueue_DequeueDogs(t *testing.T) {

	dog1 := &Dog{}
	dog1.SetName("Norton")

	cat1 := &Cat{}
	cat1.SetName("Musia")

	dog2 := &Dog{}
	dog2.SetName("Memphis")

	cat2 := &Cat{}
	cat2.SetName("Jessie")

	q := AnimalQueue{}
	q.Enqueue(dog1)
	q.Enqueue(cat1)
	q.Enqueue(dog2)
	q.Enqueue(cat2)

	_ = q.DequeueDogs().GetName()

	expected := "Memphis"
	actual := q.DequeueDogs().GetName()
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAnimalQueue_DequeueCats(t *testing.T) {

	dog1 := &Dog{}
	dog1.SetName("Norton")

	cat1 := &Cat{}
	cat1.SetName("Musia")

	dog2 := &Dog{}
	dog2.SetName("Memphis")

	cat2 := &Cat{}
	cat2.SetName("Jessie")

	q := AnimalQueue{}
	q.Enqueue(dog1)
	q.Enqueue(cat1)
	q.Enqueue(dog2)
	q.Enqueue(cat2)

	_ = q.DequeueCats().GetName()

	expected := "Jessie"
	actual := q.DequeueCats().GetName()
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
