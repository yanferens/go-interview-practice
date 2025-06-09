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

#### Key Concepts:
- **Greedy approach**: Always select the closest unvisited vertex
- **Relaxation**: Update distance if a shorter path is found
- **Priority queue**: Use to efficiently get the minimum distance vertex
- **Non-negative weights**: Algorithm doesn't work with negative edge weights

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

#### Key Concepts:
- **Edge relaxation**: Core operation that updates shortest distances
- **V-1 iterations**: Maximum number of edges in any shortest path
- **Negative cycle detection**: Additional iteration to detect negative cycles
- **Dynamic programming**: Bottom-up approach to shortest paths

#### Applications:

- Currency arbitrage detection
- Network routing with negative weights
- Finding shortest paths in graphs with negative edge weights
- Distributed systems algorithms

## General Concepts

### Graph Representation

Graphs can be represented in several ways:

1. **Adjacency Matrix**: 2D array where `matrix[i][j]` represents the weight of edge from vertex i to vertex j
2. **Adjacency List**: Array of lists where each list contains the neighbors of a vertex
3. **Edge List**: List of all edges in the graph

### Path Reconstruction

To reconstruct the actual shortest path:
1. Use a predecessor array to track the previous vertex in the shortest path
2. Start from the destination and follow predecessors back to the source
3. Reverse the path to get the correct order

### When to Use Each Algorithm

- **BFS**: Unweighted graphs or when all edges have the same weight
- **Dijkstra**: Weighted graphs with non-negative weights
- **Bellman-Ford**: Weighted graphs that may contain negative weights or when you need to detect negative cycles

## Further Reading

- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/books/introduction-algorithms-third-edition)
- [Graph Algorithms Visualization](https://visualgo.net/en/sssp)
- [Dijkstra's Algorithm Explained](https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/)
- [Bellman-Ford Algorithm](https://www.geeksforgeeks.org/bellman-ford-algorithm-dp-23/) 