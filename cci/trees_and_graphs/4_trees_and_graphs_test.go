package trees_and_graphs

import (
	//"fmt"
	"fmt"
	"testing"
)

// 4.1
func TestRouteCheck(t *testing.T) {

	n1 := &Node{}
	_ = n1.Add(1)
	_ = n1.Add(2)
	n13 := n1.Add(3)
	_ = n1.Add(4)

	_ = n13.Add(11)
	_ = n13.Add(12)
	n13_3 := n13.Add(13)

	_ = n13_3.Add(21)
	_ = n13_3.Add(22)
	n3_3 := n13_3.Add(23)

	expected := true
	actual := n1.RouteCheck(n3_3)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

// 4.2 Minimal Tree
func TestMinimalTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	/**
	Expected tree:

				    5
			2		        7
		  1	      3	 	   6	     8
			     	 4		       9

	*/

	n := MinimalTree(arr)

	if n.Val != 5 {
		t.Errorf("Lvl 1. Expected %v, got %v ", 5, n.Val)
	}

	if n.Left.Val != 2 {
		t.Errorf("Lvl 2 - Left. Expected %v, got %v ", 2, n.Left.Val)
	}
	if n.Right.Val != 7 {
		t.Errorf("Lvl 2 - Right. Expected %v, got %v ", 7, n.Right.Val)
	}

	if n.Left.Left.Val != 1 {
		t.Errorf("Lvl 3.1 - Left. Expected %v, got %v ", 1, n.Left.Left.Val)
	}
	if n.Left.Right.Val != 3 {
		t.Errorf("Lvl 3.1 - Right. Expected %v, got %v ", 3, n.Left.Right.Val)
	}

	if n.Right.Left.Val != 6 {
		t.Errorf("Lvl 3.2 - Left. Expected %v, got %v ", 6, n.Right.Left.Val)
	}
	if n.Right.Right.Val != 8 {
		t.Errorf("Lvl 3.2 - Right. Expected %v, got %v ", 8, n.Right.Right.Val)
	}

	if n.Left.Right.Right.Val != 4 {
		t.Errorf("Lvl 4.1 - Left. Expected %v, got %v ", 4, n.Left.Right.Right.Val)
	}
	if n.Right.Right.Right.Val != 9 {
		t.Errorf("Lvl 4.1 - Right. Expected %v, got %v ", 9, n.Right.Right.Right.Val)
	}
}

// 4.3 List Of Depths
func TestListOfDepths(t *testing.T) {

	root := &Node{
		Val: 1,
	}

	root.Left = &Node{
		Val: 2,
	}
	root.Left.Left = &Node{
		Val: 4,
	}
	root.Left.Right = &Node{
		Val: 5,
	}

	root.Right = &Node{
		Val: 3,
	}
	root.Right.Left = &Node{
		Val: 6,
	}
	root.Right.Right = &Node{
		Val: 7,
	}

	m := ListOfDepths(root)
	fmt.Printf("%+v \n\n", m[0].Val) // 1

	fmt.Printf("%+v -> ", m[1].Val)       // 2
	fmt.Printf("%+v \n\n", m[1].Next.Val) // 3

	fmt.Printf("%+v -> ", m[2].Val)                 // 4
	fmt.Printf("%+v -> ", m[2].Next.Val)            // 5
	fmt.Printf("%+v -> ", m[2].Next.Next.Val)       // 6
	fmt.Printf("%+v \n\n", m[2].Next.Next.Next.Val) // 7

	if m[0].Val != 1 || m[1].Val != 2 || m[1].Next.Val != 3 || m[2].Val != 4 ||
		m[2].Next.Val != 5 || m[2].Next.Next.Val != 6 || m[2].Next.Next.Next.Val != 7 {
		t.Error("Incorrect ListOfDepths implementation")
	}

}

// 4.4 Checked Balanced.
func TestIsBalanced_Balanced(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := MinimalTree(arr)

	expected := true
	actual := IsBalanced(n)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

func TestIsBalanced_Unbalanced(t *testing.T) {
	n := &Node{}

	n.Left = &Node{
		Val: 1,
	}
	n.Right = &Node{
		Val: 2,
	}
	n.Right.Right = &Node{
		Val: 3,
	}
	n.Right.Right.Right = &Node{
		Val: 4,
	}

	expected := false
	actual := IsBalanced(n)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

// 4.5 CheckBST
func TestCheckBST(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := MinimalTree(arr)

	expected := true
	actual := CheckBST(n)
	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

func TestCheckBST_OnUnbalanced(t *testing.T) {

	n := &Node{}

	n.Left = &Node{
		Val: 1,
	}
	n.Right = &Node{
		Val: 2,
	}
	n.Right.Right = &Node{
		Val: 3,
	}
	n.Right.Right.Right = &Node{
		Val: 4,
	}

	expected := false
	actual := CheckBST(n)
	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

func (n *Node) AddLeft(i int) *Node {
	n.Left = &Node{
		Val:    i,
		Parent: n,
	}

	return n.Left
}

func (n *Node) AddRight(i int) *Node {
	n.Right = &Node{
		Val:    i,
		Parent: n,
	}

	return n.Right
}

// 4.6 Find successor in a binary search tree
func TestSuccessor(t *testing.T) {

	n20 := NewNode(20)

	n8 := n20.AddLeft(8)
	n8.AddLeft(4)
	n12 := n8.AddRight(12)
	n12.AddLeft(10)
	n14 := n12.AddRight(14)

	n20.AddRight(2)

	expected := n20
	actual := Successor(n14)

	if nil == actual {
		t.Fatal("Should return non-nil actual value")
	}

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected.Val, actual.Val)
	}

}

// 4.7 Build Order
/**
PLACEHOLDER, WIP
*/

// 4.8 Common Ancestor
func TestCommonAncestor(t *testing.T) {

	n40 := &Node{
		Val: 40,
	}

	n20 := n40.AddLeft(20)
	n60 := n40.AddRight(60)

	_ = n20.AddLeft(10)
	_ = n20.AddRight(30)

	_ = n60.AddLeft(50)
	n70 := n60.AddRight(70)

	expected := n40
	actual := CommonAncestor(n40, n20, n70)

	if nil == actual {
		t.Fatal("Should return non-nil actual value")
	}

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected.Val, actual.Val)
	}
}

// 4.10 Check Subtree
func TestContainsTree(t *testing.T) {
	n40 := &Node{
		Val: 40,
	}

	n20 := n40.AddLeft(20)
	n60 := n40.AddRight(60)

	_ = n20.AddLeft(10)
	_ = n20.AddRight(30)

	_ = n60.AddLeft(50)
	_ = n60.AddRight(70)

	t1 := n40
	t2 := n60

	expected := true
	actual := ContainsTree(t1, t2)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

// 4.11 Random Node
func TestNewTreeNode(t *testing.T) {

	n := NewTreeNode(40)

	n.insertInOrder(20)
	n.insertInOrder(60)
	n.insertInOrder(10)
	n.insertInOrder(30)
	n.insertInOrder(50)
	n.insertInOrder(70)

	//fmt.Printf("%+v\n", n.data) // 40
	//
	//fmt.Printf("%+v\n", n.left.data)  // 20
	//fmt.Printf("%+v\n", n.right.data) // 60
	//
	//fmt.Printf("%+v\n", n.left.left.data)  // 10
	//fmt.Printf("%+v\n", n.left.right.data) // 30
	//
	//fmt.Printf("%+v\n", n.right.left.data)  // 50
	//fmt.Printf("%+v\n", n.right.right.data) // 70

	if n.data != 40 ||
		n.left.data != 20 ||
		n.right.data != 60 ||
		n.left.left.data != 10 ||
		n.left.right.data != 30 ||
		n.right.left.data != 50 ||
		n.right.right.data != 70 {
		t.Error("TreeNode structure error")
	}

	if n.size != 7 {
		t.Errorf("TreeNode node 0 size error. Expected %v, got %v", 7, n.size)
	}
	if n.left.size != 3 {
		t.Errorf("TreeNode node 1.left size error. Expected %v, got %v", 3, n.left.size)
	}
	if n.right.size != 3 {
		t.Errorf("TreeNode node 1.right size error. Expected %v, got %v", 3, n.right.size)
	}

	if n.left.left.size != 1 {
		t.Errorf("TreeNode node 2.left.left size error. Expected %v, got %v", 1, n.left.left.size)
	}
	if n.left.right.size != 1 {
		t.Errorf("TreeNode node 2.left.right size error. Expected %v, got %v", 1, n.left.right.size)
	}

	if n.right.left.size != 1 {
		t.Errorf("TreeNode node 2.right.left size error. Expected %v, got %v", 1, n.right.left.size)
	}
	if n.right.right.size != 1 {
		t.Errorf("TreeNode node 2.right.right size error. Expected %v, got %v", 1, n.right.right.size)
	}
}

func TestNewTreeNode_RandomNode(t *testing.T) {

	n := NewTreeNode(40)

	n.insertInOrder(20)
	n.insertInOrder(60)
	n.insertInOrder(10)
	n.insertInOrder(30)
	n.insertInOrder(50)
	n.insertInOrder(70)

	fmt.Printf("\nRandom node fetched: %+v\n", n.getRandomNode().size)
}

// 4.12 Paths with Sum
func TestPathCounter(t *testing.T) {

	n := NewTreeNode(40)

	n.left = &TreeNode{
		data: 20,
	}
	n.left.left = &TreeNode{
		data: 10,
	}
	n.left.right = &TreeNode{
		data: 30,
	}
	n.left.left.left = &TreeNode{
		data: 5,
	}
	n.left.left.right = &TreeNode{
		data: 15,
	}

	n.right = &TreeNode{
		data: 45,
	}
	n.right.left = &TreeNode{
		data: 50,
	}
	n.right.right = &TreeNode{
		data: 60,
	}

	expected := 2
	actual := PathCounter(n, 85)

	// 2 paths:
	// 40 -> 45
	// 40 -> 20 -> 10 -> 15

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

func TestPathCounter2(t *testing.T) {

	n := NewTreeNode(40)

	n.left = &TreeNode{
		data: 20,
	}
	n.left.left = &TreeNode{
		data: 10,
	}
	n.left.right = &TreeNode{
		data: 30,
	}
	n.left.left.left = &TreeNode{
		data: 5,
	}
	n.left.left.right = &TreeNode{
		data: 15,
	}

	n.right = &TreeNode{
		data: 45,
	}
	n.right.left = &TreeNode{
		data: 50,
	}
	n.right.right = &TreeNode{
		data: 60,
	}

	expected := 2
	actual := CountPathsWithSum(n, 85)

	// 2 paths:
	// 40 -> 45
	// 40 -> 20 -> 10 -> 15

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

func TestDFS(t *testing.T) {
	n40 := &Node{
		Val: 40,
	}

	n20 := n40.Add(20)
	n60 := n40.Add(60)

	_ = n20.Add(10)
	_ = n20.Add(30)

	n50 := n60.Add(50)
	_ = n60.Add(70)

	expected := n50
	actual := DFS(n40, 50)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected.Val, actual.Val)
	}
}

func TestBFS(t *testing.T) {
	n40 := &Node{
		Val: 40,
	}

	n20 := n40.Add(20)
	n60 := n40.Add(60)

	_ = n20.Add(10)
	_ = n20.Add(30)

	n50 := n60.Add(50)
	_ = n60.Add(70)

	expected := n50
	actual := BFS(n40, 50)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected.Val, actual.Val)
	}
}
