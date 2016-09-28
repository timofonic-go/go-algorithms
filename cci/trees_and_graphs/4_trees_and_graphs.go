package trees_and_graphs

import (
	"math"
	"math/rand"
)

type Node struct {
	Val      int
	Visited  bool
	Children []*Node

	Left   *Node
	Right  *Node
	Next   *Node
	Parent *Node
}

func (n *Node) Add(i int) *Node {

	newNode := &Node{
		Val:    i,
		Parent: n,
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

// 4.4 Checked Balanced.
// Implement a function to check if a binary tree is balanced.
// For the purposes of this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any node never differ by more than one.
/**
				    5
			2		        7
		  1	      3	 	   6	     8
			     	 4		       9


CCI SOLUTION (#2)
This improved algorithm works by checking the height of each subtree as we recurse down from the root.
On each node, we recursively get the heights of the left and right subtrees through the checkHeight
method. If the subtree is balanced, then checkHeight will return the actual height of the subtree. If the
subtree is not balanced, then checkHeight will return an error code. We will immediately break and
return an error code from the current call.

This code runs in O(N) time and O(H) space, where H is the height of the tree.
*/

const LongerBranchFound = math.MinInt32

func IsBalanced(n *Node) bool {
	h := checkHeight(n)
	if h == LongerBranchFound {
		return false
	}

	return true
}

func checkHeight(n *Node) int {

	if n == nil {
		return -1
	}

	leftHeight := checkHeight(n.Left)
	if leftHeight == LongerBranchFound {
		return LongerBranchFound // Pass error up.
	}
	rightHeight := checkHeight(n.Right)
	if rightHeight == LongerBranchFound {
		return LongerBranchFound // Pass error up.
	}

	if math.Abs(float64(rightHeight-leftHeight)) > 1 {
		return LongerBranchFound // Found error -> pass it back
	} else {
		return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}
}

// 4.5 Validate BST: Implement a function to check if a binary tree is a _binary search tree_.
/**
We leverage the definition of the binary search tree.

...the condition is that all left nodes must be less than or equal to the current node,
which must be less than all the right nodes.
Using this thought, we can approach the problem by passing down the min and max values.
As we iterate through the tree, we verify against progressively narrower ranges.

				    5
			2		        7
		  1	      3	 	   6	     8
			     	 4		       9
*/
func CheckBST(n *Node) bool {
	return isBST(n, math.MinInt32, math.MaxInt32)
}

func isBST(n *Node, min, max int) bool {

	// Check if node exists. If it is nil, it still matches BST definition.
	if n == nil {
		return true
	}

	// math.MinInt32 && math.MaxInt32 are used as null values.
	//
	// Return false, if:
	//
	// 1. min value is set (not null)
	//
	// AND
	//
	// 2.1 Value of current node is lower or equal than the minimal value of current range,
	// breaking BST definition.
	//
	// OR
	//
	// 2.2 Value of current node is bigger than the highest value of current range,
	// breaking the BST definition.
	//
	if (min != math.MinInt32 && n.Val <= min) || (max != math.MaxInt32 && n.Val > max) {
		return false
	}

	// If left part ( << n ) is not BST OR right part ( > n) is not BST, return false.
	if !isBST(n.Left, min, n.Val) || !isBST(n.Right, n.Val, max) {
		return false
	}

	return true
}

// 4.6 Successor:
//
// Write an algorithm to find the "next" node (i.e., in-order successor) of a given node
// in a binary search tree.
// You may assume that each node has a link to its parent.
//
/**

INFO

In-order traversal means to "visit" (often, print) the left branch, then the current node, and finally, the right
branch.

1 void inOrderTraversal(TreeNode node) {
2 	if (node!= null) {
3		inOrderTraversal(node.left);
4 		visit(node);
5 		inOrderTraversal(node.right);
6 	}
7 }

When performed on a binary search tree, it visits the nodes in ascending order (hence the name "in-order").

LINKS
http://www.java2blog.com/2014/07/binary-tree-inorder-traversal-in-java.html
http://www.geeksforgeeks.org/inorder-successor-in-binary-search-tree/
"In Binary Search Tree, Inorder Successor of an input node can also be defined as the node with the smallest key greater than the key of input node."

NOTES
1. Requires a Parent node property to exist and related modification of Add()

*/

func Successor(n *Node) *Node {

	if n == nil {
		return &Node{}
	}

	// Found right children - return leftmost node of the right subtree.
	if n.Right != nil {
		return leftMostChild(n.Right)
	} else {
		q := n
		x := q.Parent // parent

		// Go up until we're on left instead of right
		for x != nil && x.Left != q {
			q = x
			x = x.Parent
		}

		return x
	}
}

func leftMostChild(n *Node) *Node {

	if n == nil {
		return nil
	}

	for n.Left != nil {
		n = n.Left
	}

	return n
}

// 4.7 Build Order - WIP
// You are given a list of projects and a list of dependencies
// (which is a list of pairs of projects, where the second project is dependent on the first project).
// All of a project's dependencies must be built before the project is.
// Find a build order that will allow the projects to be built.
// If there is no valid build order, return an error.
//
// EXAMPLE
// Input:
// 	projects: a, b, c, d, e, f
// 	dependencies: (a, d), (f, b), (b, d), (f, a), (d, c)
// Output: f, e, a, b, d, c

// ...this problem is called topological sort: linearly ordering the vertices in a graph
// such that for every edge (a, b), a appears before b in the linear order.
/*

func FindBuildOrder(projects []string, dependencies [][]string) *Stack {
	graph := BuildGraph(projects, dependencies)
	return OrderProjects(graph.GetNodes())
}

const (
	Blank = iota
	Partial
	Complete
)

type Project struct {
	state        int
	children     []*Project
	pmap         map[string]*Project
	name         string
	dependencies int
}

func (p *Project) GetState() int {
	return p.state
}

func (p *Project) SetState(s int) {
	p.state = s
}

func (p *Project) GetChildren() []*Project {
	return p.children
}

func (p *Project) GetName() string {
	return p.name
}

func (p *Project) AddNeighbor(node *Project) {
	_, ok := p.pmap[node.GetName()]
	if !ok {
		p.children = append(p.children, node)
		node.IncrementDependencies()
	}
}

func (p *Project) IncrementDependencies() {
	p.dependencies++
}
func (p *Project) DecrementDependencies() {
	p.dependencies--
}

///////////////// #2

type Stack struct {
	projects []*Project
}

func (s *Stack) Push(p *Project) {
	s.projects = append(s.projects, p)
}

func OrderProjects(projects []*Project) *Stack {

	stack := &Stack{}
	for _, project := range projects {
		if project.GetState() == Blank {
			if !doDFS(project, stack) {
				return nil
			}
		}
	}

	return stack
}

func doDFS(project *Project, stack *Stack) bool {

	if project.GetState() == Partial {
		return false // Cycle
	}

	if project.GetState() == Blank {
		project.SetState(Partial)
		children := project.GetChildren()

		for _, child := range children {
			if !doDFS(child, stack) {
				return false
			}
		}
		project.SetState(Complete)
		stack.Push(project)
	}

	return true
}

func BuildGraph(projects []string, dependencies [][]string) *Graph {
	graph := &Graph{}
	for _, project := range projects {
		graph.CreateNode(project)
	}

	for _, dependency := range dependencies {
		first := dependency[0]
		second := dependency[1]
		graph.AddEdge(first, second)
	}

	return graph
}

type Graph struct {
	nodes []*Project
	gmap  map[string]Project
}

func (g *Graph) GetOrCreateNode(name string) *Project {
	_, ok := g.gmap[name]
	if !ok {
		node := &Project{
			name: name,
		}
		g.nodes = append(g.nodes, node)
		g.gmap[name] = node
	}

	return g.gmap[name]
}

func (g *Graph) AddEdge(startName, endName string) {
	start := g.GetOrCreateNode(startName)
	end := g.GetOrCreateNode(endName)
	start.AddNeighbor(end)
}
func (g *Graph) GetNodes() []*Project {
	return g.nodes
}

func (g *Graph) CreateNode(project *Project) {
	g.nodes = append(g.nodes, project)
	g.gmap[project.GetName()] = project
}
*/

// 4.8 First Common Ancestor
// Find first common ancestor of two nodes in a binary tree.
// Avoid storing additional nodes in a data structure.
// NOTE: This is not necessary in a binary search tree.
/**
Solution #3: Without links to parents
If p and q are both on the left of the node, branch left to look for the common ancestor.
If they are both on the right, branch right to look for the common ancestor.
When p and q are no longer on the same side, current node should be the common ancestor.
*/
func CommonAncestor(root, p, q *Node) *Node {

	// Check if node is in the tree.
	if !covers(root, p) || !covers(root, q) {
		return nil
	}

	return ancestorHelper(root, p, q)
}

func ancestorHelper(root, p, q *Node) *Node {
	if root == nil || p == nil || q == nil {
		return nil
	}

	// Check if both p and q are on the same side.
	pIsOnLeft := covers(root.Left, p)
	qIsOnLeft := covers(root.Left, q)

	// If p and q are on different sides, this means current node is an ancestor.
	if pIsOnLeft != qIsOnLeft {
		return root
	}

	// If p and q are both on the same side, determine this side.
	childSide := root.Left
	if qIsOnLeft {
		childSide = root.Right
	}

	// Use this side to continue traversing to find a common ancestor.
	return ancestorHelper(childSide, p, q)
}

// covers checks if node "p" is under the tree of the node "root".
func covers(root, p *Node) bool {
	if root == nil {
		return false
	}

	// If p is the same as root, we consider this a correct case, i.e. "covers".
	if root == p {
		return true
	}

	// Return true if any of the sides of the root node covers p.
	return covers(root.Left, p) || covers(root.Right, p)
}

// 4.10 Check Subtree
//
// T1 and T2 are two very large binary trees, with T1 much bigger than T2.
// Create an algorithm to determine if T2 is a subtree of T1.
// A tree T2 is a subtree of T1 if there exists a node n in T1 such that the subtree of n is identical to T2.
// That is, if you cut off the tree at node n, the two trees would be identical.
func ContainsTree(t1, t2 *Node) bool {
	if t2 == nil {
		return true // The empty tree is always a subtree
	}

	return subTree(t1, t2)
}

// subTree traverses bigger tree t1 to find a matching root node to start comparing their trees.
func subTree(r1, r2 *Node) bool {

	if r1 == nil {
		// Big tree is empty.
		return false
	} else if r1.Val == r2.Val && matchTree(r1, r2) {
		// Matching node found.
		return true
	}

	// Matching node not found. Continue traversing both sides and return true if a smaller tree found.
	return subTree(r1.Left, r2) || subTree(r1.Right, r2)
}

// matchTree checks if trees match.
func matchTree(r1, r2 *Node) bool {

	if r1 == nil && r2 == nil {
		return true // nothing left in subtree
	} else if r1 == nil || r2 == nil {
		return false // one of the trees is empty, therefore trees don't match
	} else if r1.Val != r2.Val {
		return false // data does not match
	} else {
		return matchTree(r1.Left, r2.Left) && matchTree(r1.Right, r2.Right)
	}
}

// 4.11 Random Node:
//
// You are implementing a binary search tree class from scratch, which, in addition
// to insert, find, and delete, has a method getRandomNode() which returns a random node from the tree.
// All nodes should be equally likely to be chosen. Design and implement an algorithm
// for getRandomNode, and explain how you would implement the rest of the methods.
//
// Option #6
// In a balanced tree, this algorithm will be O(log N), where N is the number of nodes.
// Key - evenly distribute probabilities across the options.
// Every node = 1/N probability for N nodes in total.
// "Even in a balanced tree, the number of nodes on each side might not be equal. If we have more nodes on the left than the right, then we need to go left more often."

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
	size  int // size attribute that will be used for random selection of nodes
}

// NewTreeNode is a helper for tree generation.
func NewTreeNode(d int) *TreeNode {
	return &TreeNode{
		data: d,
		size: 1,
	}
}

// getRandomNode fetch a random node based on a probability of a specific side of node.
func (t *TreeNode) getRandomNode() *TreeNode {
	leftSize := 0
	if t.left != nil {
		leftSize = t.left.size
	}

	// get number of children for a current node
	index := rand.Intn(t.size)

	// Determine a size for a next random generation.
	if index < leftSize {

		// Select random node from the left part
		return t.left.getRandomNode()
	} else if index == leftSize {
		return t
	} else {
		// Select random node from the right part
		return t.right.getRandomNode()
	}
}

func (t *TreeNode) insertInOrder(d int) {
	if d <= t.data {
		if t.left == nil {
			t.left = NewTreeNode(d)
		} else {
			t.left.insertInOrder(d)
		}
	} else {
		if t.right == nil {
			t.right = NewTreeNode(d)
		} else {
			t.right.insertInOrder(d)
		}
	}
	t.size++
}

// recursively search for a node with value d and return it or nil if not found
func (t *TreeNode) find(d int) *TreeNode {
	if t.data == d {
		return t
	} else if d <= t.data {
		if t.left != nil {
			return t.left.find(d)
		}
		return nil

	} else if d > t.data {
		if t.right != nil {
			return t.right.find(d)
		}
		return nil
	}

	return nil
}

// 4.12 Paths with Sum:
//
// You are given a binary tree in which each node contains an integer value (which might be positive or negative).
// Design an algorithm to count the number of paths that sum to a given value.
// The path does not need to start or end at the root or a leaf, but it must go downwards
// (traveling only from parent nodes to child nodes).
func PathCounter(root *TreeNode, targetSum int) int {
	m := make(map[int]int)
	return countPathsWithSum(root, targetSum, 0, m)
}

func countPathsWithSum(node *TreeNode, targetSum, runningSum int, pathCount map[int]int) int {

	if node == nil {
		return 0
	}

	// Count paths with sum ending at the current node.
	runningSum += node.data
	sum := runningSum - targetSum

	// Lookup ...
	totalPaths := 0
	val, ok := pathCount[sum]
	if ok {
		totalPaths = val
	}

	// If runningSum equals targetSum, then one additional path starts at root.
	// Add in this path.
	if runningSum == targetSum {
		totalPaths++
	}

	// Increment pathCount, recurse, then decrement pathCount.
	incrementHashTable(pathCount, runningSum, 1) // Increment pathCount
	totalPaths += countPathsWithSum(node.left, targetSum, runningSum, pathCount)
	totalPaths += countPathsWithSum(node.right, targetSum, runningSum, pathCount)
	incrementHashTable(pathCount, runningSum, -1) // Decrement pathCount

	return totalPaths
}

func incrementHashTable(hashtable map[int]int, key, delta int) {
	newCount := 0
	val, ok := hashtable[key]
	if ok {
		newCount = val + delta
	}

	// Remove when zero to reduce space usage
	if newCount == 0 {
		delete(hashtable, key)
	} else {
		hashtable[key] = newCount
	}
}
