package main

import (
	"reflect"
	"testing"
	"time"
)

func bfsReference(graph map[int][]int, start int) []int {
	// A simple reference BFS for checking correctness (sequential).
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true
	var order []int

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)

		for _, v := range graph[u] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return order
}

func buildSampleGraph() map[int][]int {
	// A sample graph
	// 0 -> 1, 2
	// 1 -> 2, 3
	// 2 -> 3
	// 3 -> 4
	// 4 -> (none)
	// 5 -> 2
	return map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {4},
		4: {},
		5: {2},
	}
}

func TestSmallGraphs(t *testing.T) {
	graph := buildSampleGraph()

	testCases := []struct {
		name       string
		graph      map[int][]int
		queries    []int
		numWorkers int
	}{
		{
			name:       "One query from node 0",
			graph:      graph,
			queries:    []int{0},
			numWorkers: 1,
		},
		{
			name:       "Multiple queries from {0, 1, 5}",
			graph:      graph,
			queries:    []int{0, 1, 5},
			numWorkers: 2,
		},
		{
			name:       "No queries",
			graph:      graph,
			queries:    []int{},
			numWorkers: 2,
		},
		{
			name:       "Many workers, single query",
			graph:      graph,
			queries:    []int{3},
			numWorkers: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results := ConcurrentBFSQueries(tc.graph, tc.queries, tc.numWorkers)

			if len(results) != len(tc.queries) {
				t.Errorf("Expected %d results, got %d", len(tc.queries), len(results))
				return
			}

			for _, start := range tc.queries {
				refOrder := bfsReference(tc.graph, start)
				if !reflect.DeepEqual(results[start], refOrder) {
					t.Errorf("For start %d, expected %v, got %v", start, refOrder, results[start])
				}
			}
		})
	}
}

func TestZeroWorkers(t *testing.T) {
	graph := buildSampleGraph()
	queries := []int{0, 1}
	results := ConcurrentBFSQueries(graph, queries, 0)
	if len(results) != 0 {
		t.Errorf("Expected empty results when numWorkers=0, but got %v", results)
	}
}

func TestLargeGraphPerformance(t *testing.T) {
	// Build a larger graph that simulates concurrency requirements
	// We'll create a chain and some branching to ensure BFS is non-trivial
	graph := make(map[int][]int)
	numNodes := 20000
	for i := 0; i < numNodes-1; i++ {
		graph[i] = []int{i + 1}
	}
	graph[numNodes-1] = []int{}

	// We'll run BFS queries from the first few nodes
	queries := []int{0, 1, 5, 100, 9999, numNodes - 2}
	numWorkers := 10

	done := make(chan struct{})
	go func() {
		_ = ConcurrentBFSQueries(graph, queries, numWorkers)
		close(done)
	}()

	select {
	case <-done:
		// success if completed quickly
	case <-time.After(3 * time.Second):
		t.Errorf("Timed out for concurrency BFS with large graph!")
	}
}

func TestDisconnectedGraph(t *testing.T) {
	// Graph with multiple disconnected components
	// 0 -> 1
	// 1 -> (none)
	// 2 -> 3
	// 3 -> (none)
	graph := map[int][]int{
		0: {1},
		1: {},
		2: {3},
		3: {},
		4: {},
	}

	queries := []int{0, 2, 4}
	numWorkers := 3
	results := ConcurrentBFSQueries(graph, queries, numWorkers)

	// Check BFS reference
	if !reflect.DeepEqual(results[0], []int{0, 1}) {
		t.Errorf("Start=0, expected [0 1], got %v", results[0])
	}
	if !reflect.DeepEqual(results[2], []int{2, 3}) {
		t.Errorf("Start=2, expected [2 3], got %v", results[2])
	}
	if !reflect.DeepEqual(results[4], []int{4}) {
		t.Errorf("Start=4, expected [4], got %v", results[4])
	}
}
