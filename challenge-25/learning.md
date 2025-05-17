# Graph Algorithms: Shortest Path

## Overview

Finding the shortest path in a graph is a fundamental problem in computer science with numerous real-world applications, from routing in computer networks to GPS navigation systems. This document covers three essential shortest path algorithms:

1. **Breadth-First Search (BFS)** - For unweighted graphs
2. **Dijkstra's Algorithm** - For weighted graphs with non-negative weights
3. **Bellman-Ford Algorithm** - For weighted graphs that may contain negative weights

## Shortest Path Algorithms

### 1. Breadth-First Search (BFS)

BFS is used to find the shortest path in an **unweighted graph**. It works by exploring all neighbor vertices at the present depth prior to moving on to vertices at the next depth level.

#### How BFS Works:

1. Start at the source vertex and mark it as visited
2. Enqueue the source vertex into a queue
3. While the queue is not empty:
   - Dequeue a vertex from the queue
   - For each unvisited adjacent vertex:
     - Mark it as visited
     - Set its distance as current vertex's distance + 1
     - Set its predecessor as the current vertex
     - Enqueue it into the queue

#### Time and Space Complexity:

- **Time Complexity**: O(V + E) where V is the number of vertices and E is the number of edges
- **Space Complexity**: O(V) for the queue and visited array

#### Example:

```go
func BreadthFirstSearch(graph [][]int, source int) ([]int, []int) {
    n := len(graph)
    distances := make([]int, n)
    predecessors := make([]int, n)
    visited := make([]bool, n)
    
    // Initialize distances and predecessors
    for i := 0; i < n; i++ {
        distances[i] = 1e9 // Infinity
        predecessors[i] = -1
    }
    
    distances[source] = 0
    visited[source] = true
    
    queue := []int{source}
    
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        
        for _, neighbor := range graph[curr] {
            if !visited[neighbor] {
                visited[neighbor] = true
                distances[neighbor] = distances[curr] + 1
                predecessors[neighbor] = curr
                queue = append(queue, neighbor)
            }
        }
    }
    
    return distances, predecessors
}
```

#### Applications:

- Social network friend recommendations (shortest connection between users)
- Web crawling
- Finding the shortest route in maze solving
- Network broadcasting

### 2. Dijkstra's Algorithm

Dijkstra's algorithm is used to find the shortest path in a **weighted graph with non-negative weights**. It works by greedily selecting the vertex with the minimum distance and relaxing all its outgoing edges.

#### How Dijkstra Works:

1. Initialize distances to all vertices as infinite and distance to the source as 0
2. Create a priority queue and insert source with distance 0
3. While the priority queue is not empty:
   - Extract the vertex with minimum distance
   - For each adjacent vertex:
     - If the distance can be improved by going through the current vertex, update the distance and predecessor

#### Time and Space Complexity:

- **Time Complexity**: O((V + E) log V) using a binary heap-based priority queue
- **Space Complexity**: O(V) for the distance array, predecessor array, and priority queue

#### Example:

```go
func Dijkstra(graph [][]int, weights [][]int, source int) ([]int, []int) {
    n := len(graph)
    distances := make([]int, n)
    predecessors := make([]int, n)
    visited := make([]bool, n)
    
    // Initialize distances and predecessors
    for i := 0; i < n; i++ {
        distances[i] = 1e9 // Infinity
        predecessors[i] = -1
    }
    
    distances[source] = 0
    
    // Priority queue implementation with a simple loop
    // In practice, use a proper min-heap priority queue
    for i := 0; i < n; i++ {
        // Find vertex with minimum distance
        u := -1
        for v := 0; v < n; v++ {
            if !visited[v] && (u == -1 || distances[v] < distances[u]) {
                u = v
            }
        }
        
        if u == -1 || distances[u] == 1e9 {
            break // All remaining vertices are unreachable
        }
        
        visited[u] = true
        
        // Update distances to adjacent vertices
        for i, v := range graph[u] {
            weight := weights[u][i]
            if !visited[v] && distances[u] + weight < distances[v] {
                distances[v] = distances[u] + weight
                predecessors[v] = u
            }
        }
    }
    
    return distances, predecessors
}
```

#### Applications:

- GPS navigation systems
- Network routing protocols (e.g., OSPF)
- Flight scheduling
- Telecommunications networks

### 3. Bellman-Ford Algorithm

The Bellman-Ford algorithm is used to find the shortest path in a **weighted graph that may contain negative weights**. It can also detect negative weight cycles.

#### How Bellman-Ford Works:

1. Initialize distances to all vertices as infinite and distance to the source as 0
2. Relax all edges V-1 times (where V is the number of vertices):
   - For each edge (u, v) with weight w:
     - If distance[u] + w < distance[v], then update distance[v] = distance[u] + w and predecessor[v] = u
3. Check for negative weight cycles:
   - For each edge (u, v) with weight w:
     - If distance[u] + w < distance[v], then a negative weight cycle exists

#### Time and Space Complexity:

- **Time Complexity**: O(V * E) where V is the number of vertices and E is the number of edges
- **Space Complexity**: O(V) for the distance and predecessor arrays

#### Example:

```go
func BellmanFord(graph [][]int, weights [][]int, source int) ([]int, []bool, []int) {
    n := len(graph)
    distances := make([]int, n)
    hasPath := make([]bool, n)
    predecessors := make([]int, n)
    
    // Initialize distances and predecessors
    for i := 0; i < n; i++ {
        distances[i] = 1e9 // Infinity
        predecessors[i] = -1
        hasPath[i] = false
    }
    
    distances[source] = 0
    hasPath[source] = true
    
    // Relax all edges |V| - 1 times
    for i := 0; i < n-1; i++ {
        for u := 0; u < n; u++ {
            for j, v := range graph[u] {
                weight := weights[u][j]
                if distances[u] != 1e9 && distances[u] + weight < distances[v] {
                    distances[v] = distances[u] + weight
                    predecessors[v] = u
                    hasPath[v] = true
                }
            }
        }
    }
    
    // Check for negative weight cycles
    for u := 0; u < n; u++ {
        for j, v := range graph[u] {
            weight := weights[u][j]
            if distances[u] != 1e9 && distances[u] + weight < distances[v] {
                // Negative weight cycle detected
                hasPath[v] = false
                // Mark all vertices in the cycle as having no valid path
                markCycleVertices(graph, v, &hasPath)
            }
        }
    }
    
    return distances, hasPath, predecessors
}

func markCycleVertices(graph [][]int, start int, hasPath *[]bool) {
    visited := make([]bool, len(graph))
    queue := []int{start}
    
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        
        if visited[curr] {
            continue
        }
        
        visited[curr] = true
        (*hasPath)[curr] = false
        
        for _, neighbor := range graph[curr] {
            queue = append(queue, neighbor)
        }
    }
}
```

#### Applications:

- Arbitrage detection in currency exchange
- Network routing with reliable negative feedback
- Systems with constraints that can be modeled as negative edges

## Comparison of Algorithms

| Algorithm | Suitable For | Time Complexity | Space Complexity | Handles Negative Weights | Detects Negative Cycles |
|-----------|--------------|-----------------|------------------|--------------------------|-------------------------|
| BFS | Unweighted graphs | O(V + E) | O(V) | N/A | N/A |
| Dijkstra | Weighted graphs with non-negative weights | O((V + E) log V) | O(V) | No | No |
| Bellman-Ford | Weighted graphs, including negative weights | O(V * E) | O(V) | Yes | Yes |

## Retrieving the Actual Path

All three algorithms track predecessors, which allows us to reconstruct the actual shortest path from the source to any destination:

```go
func reconstructPath(predecessors []int, destination int) []int {
    if predecessors[destination] == -1 {
        return []int{destination} // Only the destination itself
    }
    
    path := []int{destination}
    for current := predecessors[destination]; current != -1; current = predecessors[current] {
        path = append([]int{current}, path...) // Prepend to maintain order
    }
    
    return path
}
```

## Common Optimization Techniques

1. **Bidirectional Search**: For BFS, start searching from both source and destination simultaneously.
2. **A* Algorithm**: An extension of Dijkstra that uses heuristics to guide the search towards the destination.
3. **Johnson's Algorithm**: For all-pairs shortest paths in sparse graphs with negative edges (runs Bellman-Ford once, then Dijkstra for each vertex).
4. **Early Termination**: If only interested in the shortest path to a specific destination, terminate the algorithm once that destination is processed.

## Real-World Applications

1. **Navigation Systems**: Google Maps, Waze, and other GPS applications use variants of these algorithms.
2. **Network Routing**: OSPF (Open Shortest Path First) protocol uses Dijkstra's algorithm.
3. **Social Networks**: Finding the shortest connection between users ("degrees of separation").
4. **AI and Robotics**: Pathfinding for autonomous vehicles and robots.
5. **Telecommunications**: Routing data through the fastest or most reliable path.

## Common Pitfalls and Considerations

1. **Edge Weights Precision**: Floating-point imprecision can lead to incorrect results. Consider scaling to integers if possible.
2. **Negative Cycles**: Bellman-Ford can detect them, but no shortest path exists in graphs with negative cycles.
3. **Very Large Graphs**: For extremely large graphs, consider approximation algorithms or hierarchical approaches.
4. **Dynamic Graphs**: If the graph changes frequently, incremental updates to the shortest paths might be more efficient than recalculating. 