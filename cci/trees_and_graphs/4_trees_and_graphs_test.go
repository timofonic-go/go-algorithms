package trees_and_graphs

import (
	"testing"
)

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
