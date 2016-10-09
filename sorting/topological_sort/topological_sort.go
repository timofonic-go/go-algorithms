package main

// TopologicalSort receives a description of a search tree and returns an array with the elements sorted.
// Source: https://github.com/otaviokr/topological-sort/blob/master/toposort.go
func TopologicalSort(tree map[string][]string) []string {
	sorted := []string{}
	inDegree := map[string]int{}

	// 01. Calculate this.indegree of all vertices by going through every edge of the graph;
	// Each child gets indegree++ during breadth-first run.
	// Edge direction is element -> child.
	for element, children := range tree {

		// Set value 0 to inexisting elements.
		if inDegree[element] == 0 {
			// In golang inexisting element of type int will be zero.
			// Same as a check via _, ok := inDegree[element]
			inDegree[element] = 0
		}
		for _, child := range children {
			inDegree[child]++
		}
	}

	// 02. Collect all vertices with indegree == 0 onto a stack;
	stack := []string{}
	for node, value := range inDegree {
		if value == 0 {
			stack = append(stack, node)
			inDegree[node] = -1 // set indegree negative, to skip this node from further processing
		}
	}

	// 03. While zero-degree-stack is not empty:
	for len(stack) > 0 {
		// 03.01. Pop element from zero-degree-stack and append it to topological order;
		var node string
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 03.02. Find all children of element and decrease indegree. If indegree becomes 0, add to zero-degree-stack;
		for _, child := range tree[node] {
			inDegree[child]--
			if inDegree[child] == 0 {
				stack = append(stack, child)
				inDegree[child] = -1 // ensure we skip the node from furher processing
			}
		}

		// 03.03. Append to the sorted list.
		sorted = append(sorted, node)
	}
	return sorted
}
