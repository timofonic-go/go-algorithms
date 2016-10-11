package dijkstra

import (
	"testing"
)

type tcVertex string

type tcEdge struct {
	from   string
	to     string
	length float64
}

type tcPath struct {
	from   string
	to     string
	length float64
	err    bool
}

func TestGraph_ShortestPath(t *testing.T) {

	graph := New()

	vertices := []tcVertex{"1", "2", "3", "4", "5", "6"}
	for _, vertex := range vertices {
		graph.AddVertex(string(vertex))
	}

	edges := []tcEdge{
		{from: "1", to: "6", length: 14},
		{from: "1", to: "3", length: 9},
		{from: "1", to: "2", length: 7},

		{from: "6", to: "5", length: 9},

		{from: "3", to: "6", length: 2},
		{from: "3", to: "4", length: 11},

		{from: "2", to: "3", length: 10},
		{from: "2", to: "4", length: 15},

		{from: "4", to: "5", length: 6},
	}
	for _, edge := range edges {
		graph.AddEdge(edge.from, edge.to, edge.length)
	}

	actual, err := graph.ShortestPath("1", "4")
	expected := 20.0

	if err != nil {
		t.Errorf("Returned error: %s", err.Error())
	}

	if expected != actual {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

}
