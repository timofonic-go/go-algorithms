package cci

import (
	"testing"
)

func TestAddNode(t *testing.T) {

	l := List{}

	v1 := "Dennis"
	v2 := "Tania"
	v3 := "Jack"

	l.Add(Value{
		Name: v1,
	})
	l.Add(Value{
		Name: v2,
	})
	l.Add(Value{
		Name: v3,
	})

	if l.head.Name != v1 {
		t.Errorf("head | expected: %v, got: %v", v1, l.head.Name)
	}
	if l.tail.Name != v3 {
		t.Errorf("head | expected: %v, got: %v", v3, l.head.Name)
	}

	if l.head.next.Name != v2 {
		t.Errorf("head | expected: %v, got: %v", v2, l.head.Name)
	}
	if l.tail.prev.Name != v2 {
		t.Errorf("head | expected: %v, got: %v", v2, l.head.Name)
	}
}

// 2.1 Remove Dupes
func TestRemoveDupes(t *testing.T) {

	l := List{}

	l.Add(Value{
		Name: "Dennis",
	})
	l.Add(Value{
		Name: "Tania",
	})
	l.Add(Value{
		Name: "Tania",
	})
	l.Add(Value{
		Name: "Jack",
	})
	l.Add(Value{
		Name: "Jack",
	})
	l.Add(Value{
		Name: "Jack",
	})
	l.Add(Value{
		Name: "Lola",
	})

	actual := RemoveDupes(&l)
	if actual == nil {
		t.Error("Returned empty list")
	} else {
		if actual.head.Name != "Dennis" ||
			actual.head.next.Name != "Tania" ||
			actual.head.next.next.Name != "Jack" ||
			actual.head.next.next.next.Name != "Lola" {

			t.Log("\n\nExpected:\n")
			for n := l.head; n != nil; n = n.next {
				t.Logf("%v\n", n.Name)
			}

			t.Log("\n\nActual:\n")
			for n := actual.head; n != nil; n = n.next {
				t.Logf("%v\n", n.Name)
			}

		}
	}
}

// 2.2
func TestReturnKthToLast(t *testing.T) {
	l := List{}

	l.Add(Value{
		Name: "a",
	})
	l.Add(Value{
		Name: "b",
	})
	l.Add(Value{
		Name: "c",
	})
	l.Add(Value{
		Name: "d",
	})
	l.Add(Value{
		Name: "e",
	})
	l.Add(Value{
		Name: "f",
	})

	k := 2
	expected := "d"
	actual := ReturnKthToLast(l, k).Name

	if actual != expected {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

// 2.3 Delete Middle Node
func TestDeleteMiddleNode(t *testing.T) {
	l := &SList{}

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)

	n := &SNode{
		Value: 3,
	}

	DeleteMiddleNode(l, n)

	if l == nil {
		t.Error("Returned empty list")
	} else {
		if l.head.Value != 1 ||
			l.head.next.Value != 2 ||
			l.head.next.next.Value != 4 ||
			l.head.next.next.next.Value != 5 ||
			l.head.next.next.next.next.Value != 6 {

			t.Log("\n\nExpected:\n")
			for n := l.head; n != nil; n = n.next {
				t.Errorf("%v\n", n.Value)
			}

			t.Log("\n\nActual:\n")
			for n := l.head; n != nil; n = n.next {
				t.Logf("%v\n", n.Value)
			}

		}
	}
}

// 2.4 Partition
func TestPartitionList(t *testing.T) {

	// given: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1
	// expected: 3 -> 2 -> 1 -> 5 -> 8 -> 5 -> 10
	l := &SList{}
	l.Add(3)
	l.Add(5)
	l.Add(8)
	l.Add(5)
	l.Add(10)
	l.Add(2)
	l.Add(1)

	exp := &SList{}
	exp.Add(3)
	exp.Add(2)
	exp.Add(1)
	exp.Add(5)
	exp.Add(8)
	exp.Add(5)
	exp.Add(10)

	partition := 5
	actual := PartitionSList(l, partition)

	if actual == nil {
		t.Error("Returned empty list")
	} else {

		a1 := SListToIntMap(actual)
		a2 := SListToIntMap(exp)

		if !CompareMaps(a1, a2) {

			t.Log("\n\nExpected:\n")
			for n := exp.head; n != nil; n = n.next {
				t.Errorf("%v\n", n.Value)
			}

			t.Log("\n\nActual:\n")
			for n := actual.head; n != nil; n = n.next {
				t.Logf("%v\n", n.Value)
			}

		}
	}
}

func SListToIntMap(l *SList) map[int]int {

	out := map[int]int{}

	i := 0
	for n := l.head; n != nil; n = n.next {
		out[i] = n.Value
	}

	return out
}

func CompareMaps(m1, m2 map[int]int) bool {

	if len(m1) != len(m2) {
		return false
	}

	for i := 0; i < len(m1); i++ {
		if m1[i] != m2[i] {
			return false
		}
	}

	return true
}

// 2.6 Sum Lists
func TestSumLists(t *testing.T) {

	l1 := &SList{}
	l1.Add(7)
	l1.Add(1)
	l1.Add(6)

	l2 := &SList{}
	l2.Add(5)
	l2.Add(9)
	l2.Add(2)

	expected := 617 + 295
	actual := SumLists(l1, l2)

	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

func TestSumListsFails(t *testing.T) {

	l1 := &SList{}

	l2 := &SList{}
	l2.Add(5)

	expected := -1
	actual := SumLists(l1, l2)

	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

func TestSumListsFwd(t *testing.T) {

	l1 := &SList{}
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)

	l2 := &SList{}
	l2.Add(4)
	l2.Add(5)
	l2.Add(6)

	expected := 123 + 456
	actual := SumListsFwd(l1, l2)

	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}