# Learning Materials for Generic Data Structures

## Generics in Go

Go 1.18 introduced generics, allowing for type-parametric programming, which enables writing functions and data structures that work with different types while maintaining type safety. This challenge focuses on implementing generic data structures.

### Introduction to Generics

Generics allow you to write code that operates on values of many types while preserving type safety:

```go
// Before generics: separate functions for each type
func SumInts(numbers []int) int {
    sum := 0
    for _, n := range numbers {
        sum += n
    }
    return sum
}

func SumFloats(numbers []float64) float64 {
    sum := 0.0
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// With generics: one function for multiple types
func Sum[T constraints.Ordered](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// Usage
intSum := Sum([]int{1, 2, 3})               // 6
floatSum := Sum([]float64{1.1, 2.2, 3.3})   // 6.6
```

### Type Parameters and Constraints

Type parameters allow functions and types to work with different types:

```go
// T is a type parameter
// constraints.Ordered is a constraint that T must satisfy
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

Constraints specify what operations can be performed on type parameters:

```go
// Custom constraint: types that support addition
type Addable interface {
    int | int64 | float64 | string
}

// Function that uses the custom constraint
func Add[T Addable](a, b T) T {
    return a + b
}
```

### The `constraints` Package

The Go standard library provides common constraints in the `constraints` package:

```go
import "golang.org/x/exp/constraints"

// Examples of predefined constraints
// constraints.Ordered: types that support < <= >= >
// constraints.Integer: integer types
// constraints.Float: floating-point types
// constraints.Complex: complex number types
```

### Generic Data Structures

Generics enable creating reusable data structures:

#### Generic Stack

```go
// Generic Stack
type Stack[T any] struct {
    elements []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{elements: make([]T, 0)}
}

func (s *Stack[T]) Push(element T) {
    s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.elements) == 0 {
        return zero, false
    }
    
    lastIndex := len(s.elements) - 1
    element := s.elements[lastIndex]
    s.elements = s.elements[:lastIndex]
    return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
    var zero T
    if len(s.elements) == 0 {
        return zero, false
    }
    
    return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
    return len(s.elements)
}
```

#### Generic Queue

```go
// Generic Queue
type Queue[T any] struct {
    elements []T
}

func NewQueue[T any]() *Queue[T] {
    return &Queue[T]{elements: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(element T) {
    q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    var zero T
    if len(q.elements) == 0 {
        return zero, false
    }
    
    element := q.elements[0]
    q.elements = q.elements[1:]
    return element, true
}

func (q *Queue[T]) Peek() (T, bool) {
    var zero T
    if len(q.elements) == 0 {
        return zero, false
    }
    
    return q.elements[0], true
}

func (q *Queue[T]) IsEmpty() bool {
    return len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
    return len(q.elements)
}
```

#### Generic Linked List

```go
// Generic Linked List Node
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

// Generic Linked List
type LinkedList[T any] struct {
    head *Node[T]
    tail *Node[T]
    size int
}

func NewLinkedList[T any]() *LinkedList[T] {
    return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(value T) {
    node := &Node[T]{Value: value}
    
    if l.head == nil {
        l.head = node
        l.tail = node
    } else {
        l.tail.Next = node
        l.tail = node
    }
    
    l.size++
}

func (l *LinkedList[T]) Prepend(value T) {
    node := &Node[T]{Value: value, Next: l.head}
    
    if l.head == nil {
        l.tail = node
    }
    
    l.head = node
    l.size++
}

func (l *LinkedList[T]) Remove(value T, equals func(a, b T) bool) bool {
    if l.head == nil {
        return false
    }
    
    // Special case: remove head
    if equals(l.head.Value, value) {
        l.head = l.head.Next
        l.size--
        
        if l.head == nil {
            l.tail = nil
        }
        
        return true
    }
    
    // Find the node before the one to remove
    current := l.head
    for current.Next != nil && !equals(current.Next.Value, value) {
        current = current.Next
    }
    
    // If found, remove it
    if current.Next != nil {
        if current.Next == l.tail {
            l.tail = current
        }
        
        current.Next = current.Next.Next
        l.size--
        return true
    }
    
    return false
}

func (l *LinkedList[T]) Contains(value T, equals func(a, b T) bool) bool {
    current := l.head
    
    for current != nil {
        if equals(current.Value, value) {
            return true
        }
        current = current.Next
    }
    
    return false
}

func (l *LinkedList[T]) Size() int {
    return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
    return l.size == 0
}

func (l *LinkedList[T]) ToSlice() []T {
    result := make([]T, l.size)
    current := l.head
    i := 0
    
    for current != nil {
        result[i] = current.Value
        current = current.Next
        i++
    }
    
    return result
}
```

#### Generic Binary Search Tree

```go
// Generic Binary Search Tree Node
type TreeNode[T constraints.Ordered] struct {
    Value T
    Left  *TreeNode[T]
    Right *TreeNode[T]
}

// Generic Binary Search Tree
type BinarySearchTree[T constraints.Ordered] struct {
    root *TreeNode[T]
    size int
}

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
    return &BinarySearchTree[T]{}
}

func (t *BinarySearchTree[T]) Insert(value T) {
    t.root = t.insertHelper(t.root, value)
    t.size++
}

func (t *BinarySearchTree[T]) insertHelper(node *TreeNode[T], value T) *TreeNode[T] {
    if node == nil {
        return &TreeNode[T]{Value: value}
    }
    
    if value < node.Value {
        node.Left = t.insertHelper(node.Left, value)
    } else {
        node.Right = t.insertHelper(node.Right, value)
    }
    
    return node
}

func (t *BinarySearchTree[T]) Contains(value T) bool {
    return t.containsHelper(t.root, value)
}

func (t *BinarySearchTree[T]) containsHelper(node *TreeNode[T], value T) bool {
    if node == nil {
        return false
    }
    
    if value == node.Value {
        return true
    }
    
    if value < node.Value {
        return t.containsHelper(node.Left, value)
    }
    
    return t.containsHelper(node.Right, value)
}

func (t *BinarySearchTree[T]) InOrderTraversal() []T {
    result := make([]T, 0, t.size)
    t.inOrderHelper(t.root, &result)
    return result
}

func (t *BinarySearchTree[T]) inOrderHelper(node *TreeNode[T], result *[]T) {
    if node == nil {
        return
    }
    
    t.inOrderHelper(node.Left, result)
    *result = append(*result, node.Value)
    t.inOrderHelper(node.Right, result)
}

func (t *BinarySearchTree[T]) Size() int {
    return t.size
}

func (t *BinarySearchTree[T]) IsEmpty() bool {
    return t.size == 0
}
```

#### Generic Map

```go
// Generic Map (requires a hash function for the key)
type Map[K comparable, V any] struct {
    data map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
    return &Map[K, V]{
        data: make(map[K]V),
    }
}

func (m *Map[K, V]) Put(key K, value V) {
    m.data[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
    value, ok := m.data[key]
    return value, ok
}

func (m *Map[K, V]) Remove(key K) {
    delete(m.data, key)
}

func (m *Map[K, V]) Contains(key K) bool {
    _, ok := m.data[key]
    return ok
}

func (m *Map[K, V]) Keys() []K {
    keys := make([]K, 0, len(m.data))
    for k := range m.data {
        keys = append(keys, k)
    }
    return keys
}

func (m *Map[K, V]) Values() []V {
    values := make([]V, 0, len(m.data))
    for _, v := range m.data {
        values = append(values, v)
    }
    return values
}

func (m *Map[K, V]) Size() int {
    return len(m.data)
}

func (m *Map[K, V]) IsEmpty() bool {
    return len(m.data) == 0
}
```

### Generic Algorithms

Generics allow for implementing algorithms that work with multiple types:

#### Generic Binary Search

```go
// Binary search on a sorted slice
func BinarySearch[T constraints.Ordered](slice []T, target T) (int, bool) {
    low, high := 0, len(slice)-1
    
    for low <= high {
        mid := (low + high) / 2
        
        if slice[mid] == target {
            return mid, true
        }
        
        if slice[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    return -1, false
}
```

#### Generic Sorting

```go
// Generic bubble sort
func BubbleSort[T constraints.Ordered](slice []T) {
    n := len(slice)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
}

// With custom comparator
func BubbleSortFunc[T any](slice []T, less func(a, b T) bool) {
    n := len(slice)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if less(slice[j+1], slice[j]) {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
}
```

### Type Parameters with Methods

Methods can also use type parameters, but they must be declared on the struct itself, not added later:

```go
// This works - type parameter on the struct
type Pair[T any] struct {
    First, Second T
}

func (p *Pair[T]) Swap() {
    p.First, p.Second = p.Second, p.First
}

// This doesn't work - can't add methods with type parameters
// func (p Pair) SwapAny[T any](pair Pair[T]) {
//     p.First, p.Second = pair.Second, pair.First
// }
```

### Designing Generic Interfaces

Generic interfaces allow for specifying contracts that work with different types:

```go
// Generic Collection interface
type Collection[T any] interface {
    Add(item T)
    Remove(item T) bool
    Contains(item T) bool
    Size() int
    IsEmpty() bool
    Clear()
    ForEach(func(T))
}

// Implementing the interface
type ArrayList[T any] struct {
    items []T
    equals func(a, b T) bool
}

func NewArrayList[T any](equals func(a, b T) bool) *ArrayList[T] {
    return &ArrayList[T]{
        items: make([]T, 0),
        equals: equals,
    }
}

func (a *ArrayList[T]) Add(item T) {
    a.items = append(a.items, item)
}

func (a *ArrayList[T]) Remove(item T) bool {
    for i, val := range a.items {
        if a.equals(val, item) {
            a.items = append(a.items[:i], a.items[i+1:]...)
            return true
        }
    }
    return false
}

func (a *ArrayList[T]) Contains(item T) bool {
    for _, val := range a.items {
        if a.equals(val, item) {
            return true
        }
    }
    return false
}

func (a *ArrayList[T]) Size() int {
    return len(a.items)
}

func (a *ArrayList[T]) IsEmpty() bool {
    return len(a.items) == 0
}

func (a *ArrayList[T]) Clear() {
    a.items = make([]T, 0)
}

func (a *ArrayList[T]) ForEach(f func(T)) {
    for _, item := range a.items {
        f(item)
    }
}
```

### Generic Function Types

Functions can also be parameterized:

```go
// Generic function type
type Transformer[T, U any] func(T) U

// Map function that applies a transformation to each element
func Map[T, U any](slice []T, transformer Transformer[T, U]) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = transformer(v)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4}
squares := Map(numbers, func(x int) int { return x * x })
// squares is [1, 4, 9, 16]
```

### Best Practices for Generics

1. **Use generics to reduce duplication**: Apply when you have similar functions for different types
2. **Choose appropriate constraints**: Use the most restrictive constraint that works for your needs
3. **Don't overuse generics**: Only use them when the benefits outweigh the added complexity
4. **Consider performance implications**: Generic code can sometimes be slower than type-specific code
5. **Provide concrete type helper functions**: Offer convenience functions for common concrete types

```go
// Helper function for string comparison
func NewStringArrayList() *ArrayList[string] {
    return NewArrayList[string](func(a, b string) bool { return a == b })
}

// Helper function for int comparison
func NewIntArrayList() *ArrayList[int] {
    return NewArrayList[int](func(a, b int) bool { return a == b })
}
```

## Further Reading

- [Go Generics Tutorial](https://go.dev/doc/tutorial/generics)
- [Using Generics in Go](https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220613132600-b0d781184e0d/rand)
- [When To Use Generics](https://go.dev/blog/when-generics)
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md) 