package binary_tree

import (
	"reflect"
	"testing"
	//"github.com/mpmlj/go-algorithms/util"
)

func TestNewNode(t *testing.T) {

	root := NewNode(0)
	expected := "*binary_tree.Node"
	actual := reflect.TypeOf(root).String()

	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}

	if root.Val != 0 {
		t.Errorf("Expected %+v, got %+v", 0, root.Val)
	}
}

func TestNode_Add(t *testing.T) {
	root := NewNode(8)

	n4 := NewNode(4)
	n10 := NewNode(10)

	n2 := NewNode(2)
	n6 := NewNode(6)

	n20 := NewNode(20)

	root.Add(n4)
	root.Add(n10)

	root.Add(n2)
	root.Add(n6)

	root.Add(n20)

	if !reflect.DeepEqual(root.Left, n4) {
		t.Errorf("Error adding root.Left, got %+v\n", root.Left)
	}
	if !reflect.DeepEqual(root.Right, n10) {
		t.Errorf("Error adding root.Right, got %+v\n", root.Right)
	}
	if !reflect.DeepEqual(root.Left.Left, n2) {
		t.Errorf("Error adding root.Left.Left, got %+v\n", root.Left.Left)
	}
	if !reflect.DeepEqual(root.Left.Right, n6) {
		t.Errorf("Error adding root.Left.Right, got %+v\n", root.Left.Right)
	}

	if !reflect.DeepEqual(root.Right.Right, n20) {
		t.Errorf("Error adding root.Right.Right, got %+v\n", root.Right.Right)
	}
}
