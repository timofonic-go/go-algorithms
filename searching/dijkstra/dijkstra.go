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

	// Nodes with weights to process from the current node.
	distances := make(nodes)

	// Already visited nodes
	visited := make(nodes)

	visited[from] = 0
	current := from

	for {

		// Recalculate distances to all neighbors.
		for _, edge := range current.Neighborhood {

			// Skip nodes we already visited.
			_, ok := visited[edge.To]
			if ok {
				continue
			}

			// If node has already been visited.
			weight, ok := distances[edge.To]
			if ok {
				// Skip node, if distance to it is already lower than the distance
				// to current node + length of edge
				if weight < visited[current]+edge.Length {
					continue
				}
			}

			// At this point, we either have not calculated node's distance yet
			// or its distance is higher than the new one.
			// Set its distance to the current one: distance to the current node + length of current edge.
			distances[edge.To] = visited[current] + edge.Length
		}

		// Stop if nothing left to process.
		if len(distances) == 0 {
			break
		}

		// Distances of all neighbors have been recalculated.
		// Now choose the shortest one.
		min := math.Inf(+1)

		// Process all remaining nodes.
		for node, distance := range distances {

			// If distance via this node is lower than existing minimum for this node
			if distance <= min {

				// Select as a current node / next step.
				current = node

				// Adjust minimal distance.
				min = distance
			}
		}

		// Delete processed node from the list of unprocessed.
		delete(distances, current)

		// Update a list of visited nodes with a current node and its shortest distance.
		visited[current] = min
	}

	// If the node we are looking for is among the visited nodes - return its weight.
	weight, ok := visited[to]
	if ok {
		return weight
	}

	// Otherwise return infinity to show that its not found.
	return math.Inf(+1)
}
