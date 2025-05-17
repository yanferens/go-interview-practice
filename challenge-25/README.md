[View the Scoreboard](SCOREBOARD.md)

# Challenge 25: Graph Algorithms - Shortest Path

## Problem Statement

Implement multiple graph shortest path algorithms to find the shortest path between vertices in different types of graphs. This challenge will test your understanding of graph theory and path-finding algorithms.

You will implement three different shortest path algorithms:

1. `BreadthFirstSearch` - For unweighted graphs to find the shortest path from a source vertex to all other vertices.
2. `Dijkstra` - For weighted graphs with non-negative weights to find the shortest path from a source vertex to all other vertices.
3. `BellmanFord` - For weighted graphs that may contain negative weight edges to find the shortest path from a source vertex to all other vertices, with detection of negative cycles.

## Function Signatures

```go
func BreadthFirstSearch(graph [][]int, source int) ([]int, []int)
func Dijkstra(graph [][]int, weights [][]int, source int) ([]int, []int)
func BellmanFord(graph [][]int, weights [][]int, source int) ([]int, []bool, []int)
```

## Input Format

- `graph` - A 2D adjacency list where `graph[i]` is a slice of vertices that have an edge from vertex `i`.
- `weights` - A 2D adjacency list where `weights[i][j]` is the weight of the edge from vertex `i` to the vertex `graph[i][j]`.
- `source` - The source vertex from which to find shortest paths.

## Output Format

- `BreadthFirstSearch` returns:
  - A slice of distances where `distances[i]` is the shortest distance from the source to vertex `i`.
  - A slice of predecessors where `predecessors[i]` is the vertex that comes before vertex `i` in the shortest path from the source.

- `Dijkstra` returns:
  - A slice of distances where `distances[i]` is the shortest distance from the source to vertex `i`.
  - A slice of predecessors where `predecessors[i]` is the vertex that comes before vertex `i` in the shortest path from the source.

- `BellmanFord` returns:
  - A slice of distances where `distances[i]` is the shortest distance from the source to vertex `i`.
  - A slice of booleans where `hasPath[i]` is true if there is a path from the source to vertex `i` without a negative cycle, and false otherwise.
  - A slice of predecessors where `predecessors[i]` is the vertex that comes before vertex `i` in the shortest path from the source.

## Requirements

1. `BreadthFirstSearch` should implement a breadth-first search algorithm for unweighted graphs.
2. `Dijkstra` should implement Dijkstra's algorithm for weighted graphs with non-negative weights.
3. `BellmanFord` should implement the Bellman-Ford algorithm for weighted graphs that may have negative weights, with detection of negative cycles.
4. All algorithms should correctly handle edge cases, including isolated vertices and disconnected graphs.
5. If a vertex is unreachable from the source, its distance should be set to infinity (represented as `int(1e9)` or `math.MaxInt32` in Go).
6. If a vertex is the source, its distance should be 0 and its predecessor should be -1.

## Sample Input and Output

### Sample Input 1 (BFS)

```go
graph := [][]int{
    {1, 2},    // Vertex 0 has edges to vertices 1 and 2
    {0, 3, 4}, // Vertex 1 has edges to vertices 0, 3, and 4
    {0, 5},    // Vertex 2 has edges to vertices 0 and 5
    {1},       // Vertex 3 has an edge to vertex 1
    {1},       // Vertex 4 has an edge to vertex 1
    {2},       // Vertex 5 has an edge to vertex 2
}
source := 0
```

### Sample Output 1 (BFS)

```go
distances := []int{0, 1, 1, 2, 2, 2}
predecessors := []int{-1, 0, 0, 1, 1, 2}
```

### Sample Input 2 (Dijkstra)

```go
graph := [][]int{
    {1, 2},    // Vertex 0 has edges to vertices 1 and 2
    {0, 3, 4}, // Vertex 1 has edges to vertices 0, 3, and 4
    {0, 5},    // Vertex 2 has edges to vertices 0 and 5
    {1},       // Vertex 3 has an edge to vertex 1
    {1},       // Vertex 4 has an edge to vertex 1
    {2},       // Vertex 5 has an edge to vertex 2
}
weights := [][]int{
    {5, 10},   // Edge from 0 to 1 has weight 5, edge from 0 to 2 has weight 10
    {5, 3, 2}, // Edge weights from vertex 1
    {10, 2},   // Edge weights from vertex 2
    {3},       // Edge weights from vertex 3
    {2},       // Edge weights from vertex 4
    {2},       // Edge weights from vertex 5
}
source := 0
```

### Sample Output 2 (Dijkstra)

```go
distances := []int{0, 5, 10, 8, 7, 12}
predecessors := []int{-1, 0, 0, 1, 1, 2}
```

### Sample Input 3 (Bellman-Ford)

```go
graph := [][]int{
    {1, 2},
    {3},
    {1, 3},
    {4},
    {},
}
weights := [][]int{
    {6, 7},   // Edge weights from vertex 0
    {5},      // Edge weights from vertex 1
    {-2, 4},  // Edge weights from vertex 2 (note the negative weight)
    {2},      // Edge weights from vertex 3
    {},       // Edge weights from vertex 4
}
source := 0
```

### Sample Output 3 (Bellman-Ford)

```go
distances := []int{0, 6, 7, 11, 13}
hasPath := []bool{true, true, true, true, true}
predecessors := []int{-1, 0, 0, 2, 3}
```

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-25/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required functions.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-25/` directory:

```bash
go test -v
```

## Performance Expectations

- **BreadthFirstSearch**: O(V + E) time complexity, where V is the number of vertices and E is the number of edges.
- **Dijkstra**: O((V + E) log V) time complexity using a priority queue.
- **BellmanFord**: O(V * E) time complexity. 