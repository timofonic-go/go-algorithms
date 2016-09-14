package main

import (
	"testing"
)

func TestLinkedListPush(t *testing.T) {

	l := List{}

	l.Push(Value{
		Name:  "Moscow",
		Miles: 10000,
	})
	l.Push(Value{
		Name:  "London",
		Miles: 5000,
	})
	l.Push(Value{
		Name:  "New York",
		Miles: 1,
	})

	if l.tail.prev.Name != "London" {
		t.Errorf("Expected: London, got: %s", l.tail.prev.Name)
	}

}

func TestLinkedListFirst(t *testing.T) {

	expected := "Moscow"

	l := List{}

	l.Push(Value{
		Name:  "Moscow",
		Miles: 10000,
	})
	l.Push(Value{
		Name:  "London",
		Miles: 5000,
	})

	actual := l.First().Name

	if actual != expected {
		t.Errorf("Expected: %s, got: %s", expected, actual)
	}

}

func TestLinkedListFind(t *testing.T) {

	l := List{}

	l.Push(Value{
		Name:  "Moscow",
		Miles: 10000,
	})
	l.Push(Value{
		Name:  "London",
		Miles: 5000,
	})
	l.Push(Value{
		Name:  "New York",
		Miles: 1,
	})

	node := l.Find("London")

	if node == nil {
		t.Error("Node not found")
	}
}
