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
