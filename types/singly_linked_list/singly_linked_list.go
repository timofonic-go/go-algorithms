package singly_linked_list

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

func (l *SinglyLinkedList) Search(d int) *node {
	for n := l.head; n != nil; n = n.next {
		if n.data == d {
			return n
		}
	}

	return nil
}
