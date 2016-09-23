package cci

type Value struct {
	Name string
}

type Node struct {
	Value
	prev *Node
	next *Node
}

type List struct {
	head, tail *Node
}

func (l *List) Add(v Value) {

	n := &Node{
		Value: v,
	}

	// If I am working with a new list, need to set both a head and a tail.
	if l.head == nil {
		l.head = n
	}

	// If no tail, then set to the current value
	if l.tail == nil {
		l.tail = n
	} else {
		// If non-empty linked list, add a tail. For this:

		// - point old tail's next property to the new node
		l.tail.next = n

		// - point new tail's prev property to the old tail
		n.prev = l.tail

		// - replace tail property of a list with a new node
		l.tail = n
	}

	l.tail = n
}

// 2.1 RemoveDupes
// Write code to remove duplicates from an unsorted linked list.
// Follow up: how would you solve this problem if a temporary buffer is not allowed?
/**

NOTES
1. The only way to know a value of node in the linked list is to traverse it. So we need a way to traverse a list.
2. We need to receive a starting point.

EXAMPLE
Dennis 	-> 	Tania -> 	Tania -> 	Jack
nil,Tania	Dennis,Tania	Tania,Jack	Tania,nil

If we are at first node (head) - "Dennis".
n.Name == "Dennis"
n.next.Name == "Tania"
Names are different, move to the next node.

We are at the second node - "Tania".
n.Name == "Tania"
n.next.Name == "Tania"
Names are the same! Delete a duplicate!
- Point "next" property of the previous node to the next node: n.prev.next = n.next
- Point "prev" property of the next node to the prev node: n.next.prev = n.prev

NOTES:
- check for nil pointers
- how to iterate through a linked list in Go?
*/
func RemoveDupes(l *List) *List {

	if l == nil {
		return &List{}
	}

	// Start with a list head.
	// Continue if node is not nil.
	// If not nil, set cursor to the next node.
	for n := l.head; n != nil; n = n.next {
		if n.next != nil && n.Name == n.next.Name {
			n.Delete()
		}
	}

	return l
}

func (n *Node) Delete() {
	n.prev.next = n.next
	n.next.prev = n.prev
}
