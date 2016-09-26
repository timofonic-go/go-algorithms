package trees_and_graphs

import "testing"

func TestRouteCheck(t *testing.T) {

	n1 := &Node{}
	_ = n1.Add(1)
	_ = n1.Add(2)
	n13 := n1.Add(3)
	_ = n1.Add(4)

	_ = n13.Add(11)
	n13_2 := n13.Add(12)
	_ = n13.Add(13)

	n2 := &Node{}
	_ = n1.Add(5)
	n22 := n1.Add(6)
	_ = n1.Add(7)
	_ = n1.Add(8)

	n22.AddNode(n13_2) // Intersection point
	_ = n22.Add(21)
	_ = n22.Add(21)

	expected := true
	actual := n1.RouteCheck(n2)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}
