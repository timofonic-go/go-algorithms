package main

import (
	"fmt"
)

type Value struct {
	Name      string
	MilesAway int
}

type Node struct {
	Value
	next, prev *Node
}

type List struct {
	head, tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

// Create a new node with a given value.
func (l *List) Push(v Value) *List {

	n := &Node{
		Value: v,
	}

	// No nodes yet.
	if l.head == nil {
		// Create a first node.
		l.head = n
	} else {
		// Connect "previous last" node to the new node we've just created.
		l.tail.next = n

		// Link the new node back to the "previous last" node.
		n.prev = l.tail
	}

	l.tail = n

	return l
}

// Find traverses the Linked List searching for a node with a matching name.
// Returns nil if not found.
func (l *List) Find(name string) *Node {

	found := false
	var ret *Node = nil

	for n := l.First(); n != nil && !found; n = n.Next() {

		if n.Name == name {
			found = true
			ret = n
		}
	}

	return ret
}

func main() {

	l := List{}

	l.Push(Value{Name: "Atlanta", MilesAway: 0})
	l.Push(Value{Name: "Las Vegas", MilesAway: 1961})

	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Printf("Name: %v \n", l.head.Name)
	//fmt.Printf("prev: %v \n", l.head.prev)
	//fmt.Printf("next: %v \n", l.head.next)
	//
	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Printf("Name: %v \n", l.tail.Name)
	//fmt.Printf("prev: %v \n", l.tail.prev)
	//fmt.Printf("next: %v \n", l.tail.next)

	l.Push(Value{Name: "New York", MilesAway: 881})

	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Printf("Name: %v \n", l.head.Name)
	//fmt.Printf("prev: %v \n", l.head.prev)
	//fmt.Printf("next: %v \n", l.head.next)
	//
	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Println(" ")
	//fmt.Printf("Name: %v \n", l.tail.Name)
	//fmt.Printf("prev: %v \n", l.tail.prev)
	//fmt.Printf("next: %v \n", l.tail.next)

	fmt.Println("Search...")
	n := l.Find("Moscow")
	fmt.Printf("Searching for Moscow: %v \n", n)

	n = l.Find("New York")
	fmt.Printf("Searching for New York: %v \n", n)
}
