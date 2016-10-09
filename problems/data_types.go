package problems

type Node struct {
	Val      int
	Children []*Node
	visited  bool
}

func NewNode(v int) *Node {
	return &Node{
		Val: v,
	}
}

func (n *Node) Add(nn *Node) {
	if n != nil {
		n.Children = append(n.Children, nn)
	}
}

func (n *Node) Visit() {
	n.visited = true
}

func (n *Node) IsVisited() bool {
	return n.visited
}

// BFS-based node search by value.
func (root *Node) Find(v int) *Node {

	if root == nil {
		return root
	}

	if root.Val == v {
		return root
	}

	q := Queue{}
	root.Visit()
	q.Enqueue(root)

	for !q.IsEmpty() {

		r := q.Dequeue()
		r.Visit()

		for _, n := range r.Children {

			if n.IsVisited() {

				if n.Val == v {
					return n
				}

				n.Visit()
				q.Enqueue(n)
			}
		}

	}

	return nil
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
