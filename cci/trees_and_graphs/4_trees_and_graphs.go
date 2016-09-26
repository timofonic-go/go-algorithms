package trees_and_graphs

import "fmt"

type Node struct {
	Val      int
	Visited  bool
	Children []*Node
	Left     *Node // For trees
	Right    *Node // For trees
	Next     *Node
}

func (n *Node) Add(i int) *Node {

	newNode := &Node{
		Val: i,
	}
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *Node) AddNode(newNode *Node) {
	n.Children = append(n.Children, newNode)
}

type Queue struct {
	Nodes []*Node
}

func (q *Queue) Enqueue(n *Node) {
	q.Nodes = append(q.Nodes, n)
}

func (q *Queue) Dequeue() *Node {

	var out *Node

	if len(q.Nodes) > 0 {
		out = q.Nodes[0]
		q.Nodes = q.Nodes[1:len(q.Nodes)]
	}
	return out
}

func (q *Queue) IsEmpty() bool {

	if len(q.Nodes) == 0 {
		return true
	}
	return false
}

func (start *Node) RouteCheck(end *Node) bool {

	if start == end {
		return true
	}

	// Open a new queue for Breadth First Search (BFS)
	q := &Queue{}

	// Add out starting node there to have smth. to start a search with.
	q.Enqueue(start)

	// While queue is being filled with non-visited nodes...
	for !q.IsEmpty() {

		// Get a first queue element.
		u := q.Dequeue()

		if u != nil {
			// Loop through its children
			for _, v := range u.Children {

				if v.Visited != true {

					// If matches the node we need, return true!
					if v == end {
						return true
					} else {
						// If not, mark the node as visited and
						// add it to the queue to explore its children.
						v.Visited = true
						q.Enqueue(v)
					}

				}
			}

			u.Visited = true

		}
	}

	return false
}

// 4.2 Minimal Tree.
// Given a sorted (increasing order) array with unique integer elements,
// write an algorithm to create a binary search tree with minimal height.
/**

To have a lowest height, we need to make left and right parts as equal as possible.

Algorithm:
0. Make sure *Node has Left and Right properties to connect to subtrees/branches.
1. Insert into the tree the middle element of the array.
2. Insert into the left subtree the left subarray elements.
3. Insert into the right subtree the right subarray elements.
4. Recurse.

1,2,3,4,5,6,7

	     4
	123	567
	recursive for each part
	 2	6
	1 3	5 7

*/
func MinimalTree(arr []int) *Node {
	return CreateMinimalBST(arr, 0, len(arr)-1)
}

func CreateMinimalBST(arr []int, start, end int) *Node {
	var n *Node

	if start > end {
		return n
	}

	var mid int = (start + end) / 2

	n = &Node{
		Val: arr[mid],
	}

	n.Left = CreateMinimalBST(arr, start, mid-1)
	n.Right = CreateMinimalBST(arr, mid+1, end)

	return n
}

// 4.3 List Of Depths
// Given a binary tree, design an algorithm which creates a linked list of all nodes at each depth
// (e.g., if you have a tree with depth D, you'll have D linked lists).
/**
					    5
			2		        7
		  1	      3	 	   6	     8
			     	 4		       9

List 1: 5
List 2: 2->7
List 3: 1->3->6->8
List 4: 4->9


NOTES
We start with a top level, i.e. - a root node of an already prepared binary tree.
Our result - an array of linked lists.

*/

type List struct {
	head *Node
	tail *Node
	size int
}

func (l *List) AddNode(n *Node) {
	l.size++

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
		l.tail.Next = n // 	Tania.next = *Jack
		l.tail = n      //	l.tail = *Jack
	}
}

func (l *List) Size() int {
	return l.size
}

func ListOfDepths(root *Node) []*List {

	result := []*List{}
	current := &List{}

	if root != nil {
		current.AddNode(root)
	}

	for current.Size() > 0 {
		result = append(result, current) // Add previous level

		parents := current // Go to next level

		current = &List{}

		for parent := parents.head; parent != nil; parent = parent.Next {
			if parent.Left != nil {
				current.AddNode(parent.Left)
			}
			if parent.Right != nil {
				current.AddNode(parent.Right)
			}
		}

	}

	return result
}
