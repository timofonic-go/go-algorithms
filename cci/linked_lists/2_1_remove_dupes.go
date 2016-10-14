package linked_lists

import "strconv"

// 2.1 Remove dupes w/ buffer
func RemoveDupes(l *SinglyLinkedList) {

	matches := make(map[int]struct{})

	for n := l.head; n != nil; n = n.next {
		_, ok := matches[n.data]
		if ok {
			// match found, need to remove

			l.Delete(n.data)

		} else {
			// update a map
			matches[n.data] = struct{}{}
		}
	}

}

// 2.2 Remove Kth to Last
/**
To understand if a node is Kth to last, we need to know:
- length of the linked list
- index of a node we check

INPUT:
- singly linked list
- K

OUT:
- *node

ALGO:
- calculate total list length
- index counter during traversal
- measuring distance to the end

EXAMPLE
given sll: 629472
given K: 2
expected: node 4

1. Total qty: 6
last index = 6-1 = 5
2.
at el 6: 	index = 0	dist = 5-0 = 5		is 5 == 2 -> no, continue
at el 2:	index = 1	dist = 5-1 = 4		is 4 == 2 -> no, continue
at el 9:	index = 2	dist = 5-2 = 3		is 3 == 2 -> no, contunue
at el 4:	index = 3	fist = 5-3 = 2		is 2 == 2 -> yes, return node

OPTIMIZING
if we know total length
we know K
we can calculate position of the element we need
we can traverse to it and return

NOTE

if k > length of list, then return an error or nil as not found

*/

func GetKNode(l *SinglyLinkedList, k int) *node {

	// 1. Calculate total length (assuming this is not implemented via Insert() and Delete() functions
	length := 0
	for n := l.head; n != nil; n = n.next {
		length++
	}

	index := length - k - 1

	n := l.head
	for i := 0; i < index; i++ {
		if n.next == nil {
			return nil
		}
		n = n.next
	}

	return n
}

type SinglyLinkedList struct {
	head *node
}

type node struct {
	data int
	next *node
}

// http://www.algolist.net/Data_structures/Singly-linked_list/Insertion
func (l *SinglyLinkedList) Insert(d int) *node {

	n := &node{
		data: d,
	}

	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}

	return n
}

func InsertAfter(src *node, d int) {

	n := &node{
		data: d,
	}

	if src.next == nil {
		src.next = n
	} else {
		n.next = src.next
		src.next = n
	}
}

func (l *SinglyLinkedList) Delete(d int) {

	if l.head.data == d {
		l.head = l.head.next
	} else {
		for n := l.head; n.next != nil; n = n.next {
			if n.next.data == d {
				n.next = n.next.next
				break
			}
		}
	}
}

func (l *SinglyLinkedList) Display() {

	for n := l.head; n != nil; n = n.next {
		print(n.data)
	}
}

func (l *SinglyLinkedList) ToString() string {

	var out string
	for n := l.head; n != nil; n = n.next {
		out += strconv.Itoa(n.data)
	}

	return out
}

func (l *SinglyLinkedList) Search(d int) *node {
	for n := l.head; n != nil; n = n.next {
		if n.data == d {
			return n
		}
	}

	return nil
}
