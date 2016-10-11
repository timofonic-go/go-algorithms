package dijkstra

import (
	"errors"
	"fmt"
	"math"
)

type Graph struct {
	Vertices map[string]*Vertex
	Edges    []*Edge
}

type Vertex struct {
	Name         string
	Neighborhood []*Edge // outgoing edges
}

type Edge struct {
	From, To *Vertex
	Length   float64
}

func New() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex, 0),
		Edges:    make([]*Edge, 0),
	}
}

func (g *Graph) AddVertex(name string) error {

	_, ok := g.Vertices[name]
	if ok {
		return errors.New(fmt.Sprintf("Vertex \"%+v\" already exists", name))
	}

	g.Vertices[name] = &Vertex{
		Name:         name,
		Neighborhood: make([]*Edge, 0),
	}

	return nil
}

func (g *Graph) AddEdge(from, to string, length float64) error {
	if length < 0.0 || length == math.Inf(+1) {
		return errors.New("Weight must not be negative number or infinity")
	}

	From, ok := g.Vertices[from]
	if !ok {
		return errors.New(fmt.Sprintf("\"From\" vertex \"%+v\" does not exists", from))
	}

	To, ok := g.Vertices[to]
	if !ok {
		return errors.New(fmt.Sprintf("\"To\" vertex \"%+v\" does not exists", to))
	}

	edge := Edge{
		From:   From,
		To:     To,
		Length: length,
	}

	g.Edges = append(g.Edges, &edge)

	// Fill neighborhood of the From node with outgoing edges.
	From.Neighborhood = append(From.Neighborhood, &edge)

	return nil
}
