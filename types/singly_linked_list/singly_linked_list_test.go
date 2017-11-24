package singly_linked_list

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {

	l := &SinglyLinkedList{}
	l.Insert(1)
	l.Insert(2)
	n := l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	InsertAfter(n, 8)

	l.Delete(2)
	l.Delete(5)

	l.Display() // 4381

	expected := n
	actual := l.Search(3)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
