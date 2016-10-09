package binary_tree

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

func NewNode(v int) *Node {
	return &Node{
		Val: v,
	}
}

// Insert a node into a binary search tree
func (n *Node) Add(nn *Node) {
	if n != nil {
		if nn.Val < n.Val {
			// Add to the left side

			if n.Left == nil {
				n.Left = nn
			} else {
				n.Left.Add(nn)
			}
		} else {
			// Add to the right side

			if n.Right == nil {
				n.Right = nn
			} else {
				n.Right.Add(nn)
			}
		}
	}
}
