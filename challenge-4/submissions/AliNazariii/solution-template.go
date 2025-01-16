package main

import (
	"fmt"
	"sync"
)

// ConcurrentBFSQueries concurrently processes BFS queries on the provided graph.
// - graph: adjacency list, e.g., graph[u] = []int{v1, v2, ...}
// - queries: a list of starting nodes for BFS.
// - numWorkers: how many goroutines can process BFS queries simultaneously.
//
// Return a map from the query (starting node) to the BFS order as a slice of nodes.
// YOU MUST use concurrency (goroutines + channels) to pass the performance tests.
func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	if numWorkers <= 0 {
		return nil
	}

	semaphores := make(chan struct{}, numWorkers)

	mu := sync.Mutex{}
	bfs := make(map[int][]int)

	for _, query := range queries {
		semaphores <- struct{}{}
		go func(query int) {
			defer func() { <-semaphores }()
			
			frontier := []int{query}
			visited := make(map[int]bool)
			order := []int{}

			for len(frontier) > 0 {
				node := frontier[0]
				frontier = frontier[1:]

				if visited[node] {
					continue
				}

				visited[node] = true
				order = append(order, node)

				for _, neighbor := range graph[node] {
					if !visited[neighbor] {
						frontier = append(frontier, neighbor)
					}
				}
			}

			mu.Lock()
			bfs[query] = order
			mu.Unlock()

		}(query)
	}

	for i := 0; i < numWorkers; i++ {
		semaphores <- struct{}{}
	}

	return bfs
}

func main() {
    graph := map[int][]int{
        0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {4},
		4: {},
		5: {2},
    }

	fmt.Println(graph, len(graph))
    queries := []int{0}
    numWorkers := 1

    results := ConcurrentBFSQueries(graph, queries, numWorkers)
    /*
       Possible output:
       results[0] = [0 1 2 3 4]
       results[1] = [1 2 3 4]
       results[2] = [2 3 4]
    */

	fmt.Println(results)
}
