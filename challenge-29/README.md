[View the Scoreboard](SCOREBOARD.md)

# Challenge 29: Rate Limiter Implementation

## Problem Statement

Implement a comprehensive rate limiter system that can control the rate of requests or operations. This challenge focuses on understanding rate limiting algorithms, concurrency control, and implementing robust systems that can handle high-throughput scenarios.

## Requirements

1. Implement a `RateLimiter` interface with the following methods:
   - `Allow() bool`: Returns true if the request is allowed, false if rate limited
   - `AllowN(n int) bool`: Returns true if n requests are allowed, false if rate limited
   - `Wait(ctx context.Context) error`: Blocks until the request can be processed (or context expires)
   - `WaitN(ctx context.Context, n int) error`: Blocks until n requests can be processed
   - `Limit() int`: Returns the current rate limit (requests per second)
   - `Burst() int`: Returns the current burst capacity
   - `Reset()`: Resets the rate limiter state

2. Implement the following rate limiting algorithms:
   - **Token Bucket**: Classic algorithm with configurable rate and burst capacity
   - **Sliding Window**: More accurate rate limiting with configurable window size
   - **Fixed Window**: Simple counter-based rate limiting with fixed time windows

3. Implement a `RateLimiterFactory` that can create different types of rate limiters:
   - Support for different algorithms (token bucket, sliding window, fixed window)
   - Configurable parameters (rate, burst, window size)
   - Thread-safe implementation for concurrent usage

4. Implement advanced features:
   - **Distributed Rate Limiting**: Support for rate limiting across multiple instances
   - **Adaptive Rate Limiting**: Ability to adjust rate limits based on system load
   - **Rate Limiter Middleware**: HTTP middleware for web applications
   - **Metrics Collection**: Track rate limiter statistics and performance

## Function Signatures

```go
// Core Rate Limiter Interface
type RateLimiter interface {
    Allow() bool
    AllowN(n int) bool
    Wait(ctx context.Context) error
    WaitN(ctx context.Context, n int) error
    Limit() int
    Burst() int
    Reset()
}

// Rate Limiter Types
type TokenBucketLimiter struct {
    // Implementation fields
}

type SlidingWindowLimiter struct {
    // Implementation fields
}

type FixedWindowLimiter struct {
    // Implementation fields
}

// Factory for creating rate limiters
type RateLimiterFactory struct{}

type RateLimiterConfig struct {
    Algorithm    string // "token_bucket", "sliding_window", "fixed_window"
    Rate         int    // requests per second
    Burst        int    // maximum burst capacity
    WindowSize   time.Duration // for sliding window
}

// Constructor functions
func NewTokenBucketLimiter(rate int, burst int) RateLimiter
func NewSlidingWindowLimiter(rate int, windowSize time.Duration) RateLimiter
func NewFixedWindowLimiter(rate int, windowSize time.Duration) RateLimiter
func NewRateLimiterFactory() *RateLimiterFactory
func (f *RateLimiterFactory) CreateLimiter(config RateLimiterConfig) (RateLimiter, error)

// Advanced Features
type DistributedRateLimiter struct {
    // Implementation for distributed scenarios
}

type AdaptiveRateLimiter struct {
    // Implementation for adaptive rate limiting
}

// HTTP Middleware
func RateLimitMiddleware(limiter RateLimiter) func(http.Handler) http.Handler

// Metrics
type RateLimiterMetrics struct {
    TotalRequests   int64
    AllowedRequests int64
    DeniedRequests  int64
    AverageWaitTime time.Duration
}

func (rl RateLimiter) GetMetrics() RateLimiterMetrics
```

## Algorithm Specifications

### Token Bucket Algorithm
- Tokens are added to a bucket at a fixed rate
- Each request consumes one or more tokens
- If insufficient tokens are available, the request is rate limited
- Burst capacity allows for temporary spikes in traffic

### Sliding Window Algorithm
- Maintains a sliding time window of recent requests
- More accurate than fixed window as it doesn't suffer from boundary effects
- Smooths out traffic spikes across window boundaries

### Fixed Window Algorithm
- Simple counter that resets at fixed intervals
- Fast and memory-efficient
- May allow bursts at window boundaries

## Constraints

- All rate limiters must be thread-safe for concurrent usage
- Implement proper error handling for invalid configurations
- Support for context cancellation in blocking operations
- Efficient memory usage for high-throughput scenarios
- Configurable precision for timing operations

## Sample Usage

### Basic Usage

```go
// Create a token bucket rate limiter (10 requests/second, burst of 5)
limiter := NewTokenBucketLimiter(10, 5)

// Check if request is allowed
if limiter.Allow() {
    fmt.Println("Request allowed")
    // Process request
} else {
    fmt.Println("Request rate limited")
    // Handle rate limiting
}

// Wait for request to be allowed (with timeout)
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

if err := limiter.Wait(ctx); err != nil {
    fmt.Printf("Request timed out: %v\n", err)
} else {
    fmt.Println("Request processed after waiting")
}
```

### Factory Usage

```go
factory := NewRateLimiterFactory()

config := RateLimiterConfig{
    Algorithm:  "sliding_window",
    Rate:       100,
    WindowSize: time.Minute,
}

limiter, err := factory.CreateLimiter(config)
if err != nil {
    log.Fatal(err)
}

// Use the limiter
for i := 0; i < 200; i++ {
    if limiter.Allow() {
        fmt.Printf("Request %d allowed\n", i+1)
    } else {
        fmt.Printf("Request %d rate limited\n", i+1)
    }
    time.Sleep(10 * time.Millisecond)
}
```

### HTTP Middleware Usage

```go
limiter := NewTokenBucketLimiter(100, 10) // 100 req/sec, burst of 10

mux := http.NewServeMux()
mux.HandleFunc("/api/endpoint", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Request processed"))
})

// Apply rate limiting middleware
handler := RateLimitMiddleware(limiter)(mux)

server := &http.Server{
    Addr:    ":8080",
    Handler: handler,
}

log.Println("Server starting on :8080")
server.ListenAndServe()
```

### Advanced Usage with Metrics

```go
limiter := NewTokenBucketLimiter(50, 10)

// Simulate load
for i := 0; i < 1000; i++ {
    limiter.Allow()
    time.Sleep(time.Millisecond)
}

// Get metrics
metrics := limiter.GetMetrics()
fmt.Printf("Total requests: %d\n", metrics.TotalRequests)
fmt.Printf("Allowed requests: %d\n", metrics.AllowedRequests)
fmt.Printf("Denied requests: %d\n", metrics.DeniedRequests)
fmt.Printf("Success rate: %.2f%%\n", 
    float64(metrics.AllowedRequests)/float64(metrics.TotalRequests)*100)
```

## Performance Requirements

- **Token Bucket**: O(1) time complexity for Allow() operations
- **Sliding Window**: O(log n) time complexity where n is the number of requests in window
- **Fixed Window**: O(1) time complexity for Allow() operations
- Memory usage should be bounded and configurable
- Support for at least 10,000 concurrent goroutines

## Testing Requirements

Your implementation should pass tests for:
- Basic functionality of each algorithm
- Concurrent access from multiple goroutines
- Context cancellation in blocking operations
- Rate limit accuracy under various load patterns
- Memory leak detection under sustained load
- Performance benchmarks for high-throughput scenarios

## Instructions

- **Fork** the repository.
- **Clone** your fork to your local machine.
- **Create** a directory named after your GitHub username inside `challenge-29/submissions/`.
- **Copy** the `solution-template.go` file into your submission directory.
- **Implement** the required interfaces, types, and methods.
- **Test** your solution locally by running the test file.
- **Commit** and **push** your code to your fork.
- **Create** a pull request to submit your solution.

## Testing Your Solution Locally

Run the following command in the `challenge-29/` directory:

```bash
go test -v
```

For performance testing:

```bash
go test -v -bench=.
```

For race condition testing:

```bash
go test -v -race
``` 