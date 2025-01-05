/*
BFS or DFS for Trees and Graphs
- Trees and graphs are often solved using BFS for shortest paths or DFS for traversals.
- BFS is best for level-order traversal, while DFS is useful for exploring paths.
- Example: Find the shortest path in a graph.
*/

package main

import (
	"fmt"
	"math"
)

type Graph map[int][]int

func bfs(graph Graph, start int) map[int]int {
	distances := make(map[int]int)
	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	queue := []int{start}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[currentNode] {
			if distances[neighbor] == math.MaxInt32 {
				distances[neighbor] = distances[currentNode] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return distances
}

func main() {
	// Example graph
	graph := Graph{
		0: {1, 2},
		1: {2},
		2: {},
	}

	distances := bfs(graph, 0)
	fmt.Println("Distances from node 0:", distances) // Output: Distances from node 0: map[0:0 1:1 2:2]
}
