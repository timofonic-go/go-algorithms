package cci

import (
	"strconv"
)

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

type SList struct {
	head *SNode
	tail *SNode
	size int
}

type SNode struct {
	Value int
	next  *SNode
}

func (l *SList) Add(v int) {

	l.size++

	n := &SNode{
		Value: v,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		// Dennis		Tania
		// ->Tania		->nil
		// 			tail

		// Dennis		Tania		Jack
		// ->Tania		->Jack		->nil
		// 					tail

		l.tail.next = n // 	Tania.next = *Jack
		l.tail = n      //	l.tail = *Jack
	}
}

func (l *SList) Size() int {
	return l.size
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

// 2.2 Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
/**
1. Singly linked list
2. kth element from the last element
3. Input: list, k
4. Return: node

EXAMPLE

a -> b -> c -> d
k := 2
expected: b (1: d -> c, 2: c -> b)


ALGORITHM
1. Iterate through list
2. Fill in an array with nodes
0:a
1:b
2:c
3:d

3. Get array length
len(arr): 4

4. Subtract k from length to get an array index

idx := 4 - 1 - 2 = 1, added -1 because index of last element is 1 less then array length


5. Get a node from the array by its index
arr[1] = b

*/
func ReturnKthToLast(l List, k int) *Node {

	arr := []*Node{}

	i := 0
	for n := l.head; n != nil; n = n.next {
		arr = append(arr, n)
		i++
	}

	idx := len(arr) - 1 - k

	return arr[idx]
}

// 2.3 Delete Middle Node
// Implement an algorithm to delete a node in the middle
// (i.e., any node but the first and last node, not necessarily the exact middle)
// of a singly linked list, given only access to that node.
//
// EXAMPLE
// Input: the node c from the linked list a->b->c->d->e->f
// Result: nothing is returned, but the new linked list looks like a->b->d->e->f
/**


if n.next = c { // n == b
	n.next = n.next.next
}

*/

func DeleteMiddleNode(l *SList, sn *SNode) {

	for n := l.head; n != nil; n = n.next {
		if n.next != nil && n.next.Value == sn.Value {
			n.next = n.next.next
		}
	}
}

// 2.4 Partition
// Write code to partition a linked list around a value x,
// such that all nodes less than x come before all nodes greater than or equal to x.
// If x is contained within the list, the values of x only need to be after the elements less than x (see below).
//
// The partition element x can appear anywhere in the "right partition";
// it does not need to appear between the left and right partitions.
/**

EXAMPLE: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1
partition: 5


BRUTE FORCE:
1. Traverse the array
2. Compare current node to the partitioning value
3. If current value < partitioning value: add to list 1
4. if current value >= partitioning value, add to list 2
5. Merge lists

*/
func PartitionSList(l *SList, k int) *SList {

	// Error checking in case an empty list was supplied
	if l == nil {
		return &SList{}
	}

	left := &SList{}
	right := &SList{}

	// Traverse all node starting from the head
	// running while referenced node is not nil
	// in the end of the loop, current pointer is swapped with a pointer of the "next' attribute
	for n := l.head; n != nil; n = n.next {

		if n.Value < k {
			left.Add(n.Value)
		} else {
			right.Add(n.Value)
		}
	}

	// Join the lists
	left.tail.next = right.head

	return left
}

// 2.6 Sum Lists:
// You have two numbers represented by a linked list, where each node contains a single digit.
// The digits are stored in reverse order, such that the 1's digit is at the head of the list.
// Write a function that adds the two numbers and returns the sum as a linked list.
//
// EXAMPLE
// Input: (7-> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
// Output: 2 -> 1 -> 9. That is, 912.
//
// FOLLOW UP
// Suppose the digits are stored in forward order. Repeat the above problem.
//
// EXAMPLE
// Input: (6 -> 1 -> 7) + (2 -> 9 -> 5).That is,617 + 295.
// Output: 9 - > 1 -> 2. That is, 912.
/**

BRUTE FORCE
1. Traverse each of the lists
2. Convert each number to a string character
4. Add each string character to the beginning of a corresponding string.
5. Convert numbers to integer type
6. Return their sum.
7. For variation B, forward order, we will need

TEST:
str := ""
7 -> 1 -> 6
str = "7" + str -> "7"
str = "1" + str -> "17"
str = "6" + str -> "617"
convertion: "617" -> 617

str = ""
5 -> 9 -> 2
str = "5" + str -> "5"
str = "9" + str -> "95"
str = "2" + str -> "295"
convertion: "295" -> 295

sum = 617 + 295 = 912


CASES:
empty lists, return -1

*/

func SumLists(l1, l2 *SList) int {

	if l1.Size() == 0 || l2.Size() == 0 {
		return -1
	}

	var s1, s2 string

	for n := l1.head; n != nil; n = n.next {
		s1 = strconv.Itoa(n.Value) + s1
	}

	for n := l2.head; n != nil; n = n.next {
		s2 = strconv.Itoa(n.Value) + s2
	}

	n1, err := strconv.Atoi(s1)
	if err != nil {
		return -1
	}

	n2, err := strconv.Atoi(s2)
	if err != nil {
		return -1
	}

	return n1 + n2
}

// 2.6 Sum Lists - with forward numbers.
// Suppose the digits are stored in forward order. Repeat the above problem.
func SumListsFwd(l1, l2 *SList) int {

	if l1.Size() == 0 || l2.Size() == 0 {
		return -1
	}

	var s1, s2 string

	for n := l1.head; n != nil; n = n.next {
		s1 += strconv.Itoa(n.Value)
	}

	for n := l2.head; n != nil; n = n.next {
		s2 += strconv.Itoa(n.Value)
	}

	n1, err := strconv.Atoi(s1)
	if err != nil {
		return -1
	}

	n2, err := strconv.Atoi(s2)
	if err != nil {
		return -1
	}

	return n1 + n2
}
