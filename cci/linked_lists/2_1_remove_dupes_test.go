package linked_lists

import (
	"reflect"
	"testing"
)

func TestRemoveDupes(t *testing.T) {

	// 75344474622
	l := &SinglyLinkedList{}
	l.Insert(2)
	l.Insert(2)
	l.Insert(6)
	l.Insert(4)
	l.Insert(7)
	l.Insert(4)
	l.Insert(4)
	l.Insert(4)
	l.Insert(3)
	l.Insert(5)
	l.Insert(7)

	RemoveDupes(l)

	actual := l.ToString()
	expected := "537462"

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestGetKNode(t *testing.T) {

	// 274926
	l := &SinglyLinkedList{}
	l.Insert(6)
	l.Insert(2)
	l.Insert(9)
	n := l.Insert(4)
	l.Insert(7)
	l.Insert(2)

	k := 3
	expected := n
	actual := GetKNode(l, k)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
