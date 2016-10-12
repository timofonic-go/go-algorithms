package dijkstra

import (
	"errors"
	"fmt"
	"math"
)

func (g *Graph) ShortestPath(from, to string) (float64, error) {

	From, ok := g.Vertices[from]
	if !ok {
		return math.Inf(+1), errors.New(fmt.Sprintf("\"From\" vertex \"%+v\" does not exists", from))
	}

	To, ok := g.Vertices[to]
	if !ok {
		return math.Inf(+1), errors.New(fmt.Sprintf("\"To\" vertex \"%+v\" does not exists", to))
	}

	weight := g.Dijkstra(From, To)

	if math.IsInf(weight, +1) {
		return weight, errors.New(fmt.Sprintf("Path between \"%v\" and \"%v\" does not exist!", from, to))
	}

	return weight, nil
}

type nodes map[*Vertex]float64

func (g *Graph) Dijkstra(from, to *Vertex) float64 {

	// If source and destination are the same.
	if from == to {
		return 0
	}

	// Nodes with temporary distance values that may be recalculated later.
	todo := make(nodes)

	// Array of shortest distances.
	visited := make(nodes)

	// 0.1 Init
	visited[from] = 0
	current := from

	for {

		// 0.2 Recalculate distances to all neighbors of the current node.
		for _, edge := range current.Neighborhood {

			// Skip already marked as shortest distance.
			_, ok := visited[edge.To]
			if ok {
				continue
			}

			// Determine if this node's distance has already been calculated.
			// If not, calculate and add to the todo list (will determine if shortest later).
			// If yes, recalculate distance if it is longer than previously calculated.
			weight, ok := todo[edge.To]
			if ok {
				// Skip node, if distance to it is already lower than the distance to it
				// via the current node.
				if weight < visited[current]+edge.Length {
					continue
				}
			}

			todo[edge.To] = visited[current] + edge.Length
		}

		// Stop if no vertices left to go to.
		if len(todo) == 0 {
			break
		}

		// Distances to all neighbor nodes have been recalculated.
		// Now choose the shortest one.
		min := math.Inf(+1)

		// Find the shortest path and next node among the remaining nodes.
		for node, distance := range todo {

			// If distance via this node is lower than existing minimum for this node
			if distance <= min {

				// Select as a current node / next step.
				current = node

				// Adjust minimal distance.
				min = distance
			}
		}

		// Distance to this node is the shortest.
		// Store in the list of shortest distances.
		visited[current] = min

		// Delete from the list of unprocessed nodes.
		delete(todo, current)
	}

	// If the node we are looking for is among the visited nodes - return its weight.
	weight, ok := visited[to]
	if ok {
		return weight
	}

	// Otherwise return infinity to show that its not found.
	return math.Inf(+1)
}
