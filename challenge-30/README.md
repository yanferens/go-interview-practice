# Challenge 30: Context Management Implementation

## Overview

Learn Go's essential `context` package by implementing basic context management patterns. The `context` package is fundamental to Go programming, providing a way to carry cancellation signals, timeouts, and request-scoped values across function calls.

## Learning Objectives

- Understand context cancellation and timeouts
- Learn to pass values through context safely
- Implement context-aware functions
- Apply contexts in real-world scenarios

## Your Task

Implement a context manager with **6 core methods** and **2 helper functions**:

### ContextManager Interface (6 methods)

```go
type ContextManager interface {
    // Create a cancellable context from a parent context
    CreateCancellableContext(parent context.Context) (context.Context, context.CancelFunc)
    
    // Create a context with timeout
    CreateTimeoutContext(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
    
    // Add a value to context
    AddValue(parent context.Context, key, value interface{}) context.Context
    
    // Get a value from context
    GetValue(ctx context.Context, key interface{}) (interface{}, bool)
    
    // Execute a task with context cancellation support
    ExecuteWithContext(ctx context.Context, task func() error) error
    
    // Wait for a duration or until context is cancelled
    WaitForCompletion(ctx context.Context, duration time.Duration) error
}
```

### Helper Functions (2 functions)

```go
// Simulate work that can be cancelled
func SimulateWork(ctx context.Context, workDuration time.Duration, description string) error

// Process multiple items with context awareness
func ProcessItems(ctx context.Context, items []string) ([]string, error)
```

## Context Types to Learn

### 1. **Cancellation Context**
```go
ctx, cancel := context.WithCancel(parent)
defer cancel() // Always call cancel to avoid memory leaks
```

### 2. **Timeout Context**
```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel() // Always call cancel
```

### 3. **Value Context**
```go
ctx := context.WithValue(parent, "userID", "12345")
```

## Real-World Examples

### HTTP Request with Timeout
```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    // Create request context with timeout
    ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
    defer cancel()
    
    // Add request ID for tracing
    ctx = context.WithValue(ctx, "requestID", generateID())
    
    // Process with context
    result, err := processRequest(ctx, r)
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            http.Error(w, "Request timeout", http.StatusRequestTimeout)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(result)
}
```

### Worker with Graceful Shutdown
```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            log.Println("Worker shutting down:", ctx.Err())
            return
        case job := <-jobs:
            processJob(ctx, job)
        }
    }
}
```

## Key Context Patterns

### 1. **Always Check ctx.Done()**
```go
func longRunningWork(ctx context.Context) error {
    for i := 0; i < 1000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err() // Return cancellation error
        default:
            // Continue work
            doWorkStep(i)
        }
    }
    return nil
}
```

### 2. **Use Select for Context Operations**
```go
func executeWithTimeout(ctx context.Context, task func() error) error {
    done := make(chan error, 1)
    
    go func() {
        done <- task()
    }()
    
    select {
    case err := <-done:
        return err // Task completed
    case <-ctx.Done():
        return ctx.Err() // Context cancelled/timeout
    }
}
```

### 3. **Pass Context as First Parameter**
```go
// ✅ Good
func ProcessData(ctx context.Context, data []byte) error

// ❌ Bad - context should be first parameter
func ProcessData(data []byte, ctx context.Context) error
```

## Context Best Practices

### ✅ Do:
- Always pass context as the first parameter
- Call `cancel()` to avoid memory leaks (use `defer cancel()`)
- Check `ctx.Done()` in loops and long operations
- Use context for request-scoped values (user ID, request ID)
- Derive child contexts from parent contexts

### ❌ Don't:
- Store contexts in structs
- Pass nil context (use `context.Background()` instead)
- Use context for optional parameters
- Ignore context cancellation
- Create contexts in tight loops

## Implementation Hints

### CreateCancellableContext
```go
// Use the standard library function
return context.WithCancel(parent)
```

### CreateTimeoutContext
```go
// Use the standard library function
return context.WithTimeout(parent, timeout)
```

### AddValue
```go
// Use the standard library function
return context.WithValue(parent, key, value)
```

### GetValue
```go
// Check if value exists and return appropriately
value := ctx.Value(key)
if value == nil {
    return nil, false
}
return value, true
```

### ExecuteWithContext
```go
// Run task in goroutine and race against context cancellation
done := make(chan error, 1)
go func() {
    done <- task()
}()

select {
case err := <-done:
    return err
case <-ctx.Done():
    return ctx.Err()
}
```

## Common Context Errors

- `context.Canceled` - Context was manually cancelled
- `context.DeadlineExceeded` - Context timeout was reached

## Testing Your Implementation

The test suite includes **13 focused test functions**:

1. **TestCreateCancellableContext** - Basic cancellation
2. **TestCreateTimeoutContext** - Timeout handling
3. **TestAddAndGetValue** - Value storage and retrieval
4. **TestExecuteWithContext_Success** - Successful task execution
5. **TestExecuteWithContext_TaskError** - Task error handling
6. **TestExecuteWithContext_Cancellation** - Cancellation during execution
7. **TestWaitForCompletion_Success** - Successful waiting
8. **TestWaitForCompletion_Cancellation** - Cancelled waiting
9. **TestSimulateWork_Success** - Work simulation
10. **TestSimulateWork_Cancellation** - Cancelled work
11. **TestProcessItems_Success** - Batch processing
12. **TestProcessItems_Cancellation** - Cancelled batch processing
13. **TestContextIntegration** - Combined functionality

## Success Criteria

Your implementation should:

- ✅ Handle context cancellation properly
- ✅ Implement timeout functionality correctly
- ✅ Store and retrieve context values safely
- ✅ Execute tasks with cancellation support
- ✅ Pass all 13 test cases
- ✅ Follow Go context best practices
- ✅ Handle errors appropriately

## Getting Started

1. Look at the solution template with TODO sections
2. Implement one method at a time
3. Run tests frequently: `go test -v`
4. Start with the simple methods (AddValue, GetValue)
5. Then tackle the more complex ones (ExecuteWithContext)

Good luck! Context management is essential for building robust Go applications. Master these patterns and you'll be well-equipped for real-world Go development. 