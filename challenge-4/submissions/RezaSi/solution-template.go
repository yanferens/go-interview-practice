package main

// ConcurrentBFSQueries concurrently processes BFS queries on the provided graph.
// - graph: adjacency list, e.g., graph[u] = []int{v1, v2, ...}
// - queries: a list of starting nodes for BFS.
// - numWorkers: how many goroutines can process BFS queries simultaneously.
//
// Return a map from the query (starting node) to the BFS order as a slice of nodes.
// YOU MUST use concurrency (goroutines + channels) to pass the performance tests.
func ConcurrentBFSQueries(graph map[int][]int, queries []int, numWorkers int) map[int][]int {
	// TODO: Implement concurrency-based BFS for multiple queries.
	// Return an empty map so the code compiles but fails tests if unchanged.
	if numWorkers == 0 {
	    return map[int][]int{}
	}

	result := make(map[int][]int, 0)
    
    for _, query := range(queries) {
        queryRes := BFS(graph, query)
        result[query] = queryRes
    }

	return result
}

func BFS(graph map[int][]int, startNode int) []int {
    queue := make([]int, 0)
    queue = append(queue, startNode)
    
    result := make([]int, 0)
    
    visitedNodes := make(map[int]bool, 0)
    visitedNodes[startNode] = true
    
    for len(queue) > 0 {
        currentNode := queue[0]
        result = append(result, currentNode)
        
        for _, adj := range(graph[currentNode]) {
            _, visited := visitedNodes[adj]
            if !visited {
                visitedNodes[adj] = true
                queue = append(queue, adj)
            }
        }
        
        queue = queue[1:]
    }
    
    return result
}

func main() {
	// You can insert optional local tests here if desired.
}
