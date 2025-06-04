# Learning Materials for Rate Limiter Implementation

## Introduction to Rate Limiting

Rate limiting is a crucial technique used in software systems to control the rate of incoming requests or operations. It protects services from being overwhelmed by too many requests and ensures fair resource allocation among users.

## Why Rate Limiting Matters

### 1. **System Protection**
- Prevents system overload and crashes
- Maintains service availability during traffic spikes
- Protects against denial-of-service (DoS) attacks

### 2. **Resource Management**
- Ensures fair usage of computational resources
- Prevents any single user from monopolizing the system
- Maintains consistent performance for all users

### 3. **Cost Control**
- Limits resource consumption and associated costs
- Prevents runaway processes from causing expensive operations
- Enables predictable infrastructure scaling

### 4. **Service Level Agreements (SLAs)**
- Enforces agreed-upon usage limits
- Enables different service tiers with varying limits
- Provides measurable quality of service

## Rate Limiting Algorithms

### 1. Token Bucket Algorithm

The token bucket algorithm is one of the most popular and flexible rate limiting techniques.

#### How It Works

```
┌─────────────────┐
│   Token Bucket  │  ← Tokens added at fixed rate
│  [🪙][🪙][🪙]   │
│  [🪙][🪙][ ]    │  ← Current tokens
│  [ ][ ][ ]      │
└─────────────────┘
        ↓
   Request consumes token
```

1. **Token Generation**: Tokens are added to a bucket at a fixed rate
2. **Bucket Capacity**: The bucket has a maximum capacity (burst limit)
3. **Request Processing**: Each request consumes one or more tokens
4. **Rate Limiting**: If no tokens are available, the request is denied

#### Implementation Key Points

```go
type TokenBucket struct {
    rate       int       // tokens per second
    burst      int       // maximum bucket capacity
    tokens     float64   // current token count
    lastRefill time.Time // last token refill time
    mutex      sync.Mutex
}

func (tb *TokenBucket) Allow() bool {
    tb.mutex.Lock()
    defer tb.mutex.Unlock()
    
    // Calculate elapsed time and add tokens
    now := time.Now()
    elapsed := now.Sub(tb.lastRefill).Seconds()
    tb.tokens += elapsed * float64(tb.rate)
    
    // Cap at burst capacity
    if tb.tokens > float64(tb.burst) {
        tb.tokens = float64(tb.burst)
    }
    
    tb.lastRefill = now
    
    // Check if request can be allowed
    if tb.tokens >= 1.0 {
        tb.tokens -= 1.0
        return true
    }
    
    return false
}
```

#### Advantages
- **Burst Handling**: Allows temporary traffic spikes up to burst capacity
- **Smooth Rate**: Provides consistent long-term rate limiting
- **Flexibility**: Configurable rate and burst parameters
- **Efficiency**: O(1) time complexity for operations

#### Disadvantages
- **Memory Usage**: Requires floating-point arithmetic for precise timing
- **Complexity**: More complex than simpler algorithms

### 2. Sliding Window Algorithm

The sliding window algorithm maintains a more accurate rate limit by tracking requests within a moving time window.

#### How It Works

```
Time: --------|--------|--------|--------|--------
      10:00   10:01   10:02   10:03   10:04
              
Current Time: 10:03:30
Window Size: 1 minute
Window: [10:02:30 - 10:03:30]

Requests in window: ✓✓✓✗✗ (3 requests, limit 5)
```

1. **Window Management**: Maintains a sliding time window of fixed size
2. **Request Tracking**: Records timestamps of all requests
3. **Window Sliding**: Continuously removes old requests outside the window
4. **Rate Checking**: Allows requests if count within window is below limit

#### Implementation Key Points

```go
type SlidingWindow struct {
    rate       int
    windowSize time.Duration
    requests   []time.Time
    mutex      sync.Mutex
}

func (sw *SlidingWindow) Allow() bool {
    sw.mutex.Lock()
    defer sw.mutex.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-sw.windowSize)
    
    // Remove old requests
    validRequests := make([]time.Time, 0)
    for _, req := range sw.requests {
        if req.After(cutoff) {
            validRequests = append(validRequests, req)
        }
    }
    sw.requests = validRequests
    
    // Check if we can allow the request
    if len(sw.requests) < sw.rate {
        sw.requests = append(sw.requests, now)
        return true
    }
    
    return false
}
```

#### Advantages
- **Accuracy**: More precise rate limiting without boundary effects
- **Fairness**: Smooth distribution of allowed requests
- **Predictability**: Consistent behavior across time boundaries

#### Disadvantages
- **Memory Usage**: Stores timestamps for all requests in the window
- **Complexity**: O(n) time complexity for cleanup operations
- **Scalability**: Memory usage grows with request rate

### 3. Fixed Window Algorithm

The fixed window algorithm is the simplest approach, using a counter that resets at fixed intervals.

#### How It Works

```
Window 1     Window 2     Window 3
[10:00-10:01][10:01-10:02][10:02-10:03]
✓✓✓✓✓✗✗✗    ✓✓✓✓✓✗      ✓✓✓✓✓
(5/5 limit)  (5/5 limit) (5/5 limit)
```

1. **Time Windows**: Divides time into fixed-size windows
2. **Counter Reset**: Request counter resets at window boundaries
3. **Simple Counting**: Increments counter for each request
4. **Limit Enforcement**: Denies requests when counter exceeds limit

#### Implementation Key Points

```go
type FixedWindow struct {
    rate         int
    windowSize   time.Duration
    windowStart  time.Time
    requestCount int
    mutex        sync.Mutex
}

func (fw *FixedWindow) Allow() bool {
    fw.mutex.Lock()
    defer fw.mutex.Unlock()
    
    now := time.Now()
    
    // Check if we're in a new window
    if now.Sub(fw.windowStart) >= fw.windowSize {
        fw.windowStart = now
        fw.requestCount = 0
    }
    
    // Check if request can be allowed
    if fw.requestCount < fw.rate {
        fw.requestCount++
        return true
    }
    
    return false
}
```

#### Advantages
- **Simplicity**: Easy to understand and implement
- **Performance**: O(1) time complexity
- **Memory Efficiency**: Minimal memory usage

#### Disadvantages
- **Boundary Effects**: Allows bursts at window boundaries
- **Unfairness**: Can allow 2x rate limit at window transitions

## Advanced Rate Limiting Concepts

### 1. Distributed Rate Limiting

When running multiple instances of a service, rate limits need to be coordinated across instances.

#### Approaches

1. **Centralized Storage**: Use Redis or similar for shared state
2. **Consistent Hashing**: Distribute rate limiting across nodes
3. **Approximate Algorithms**: Trade accuracy for performance

```go
type DistributedRateLimiter struct {
    redis  *redis.Client
    key    string
    rate   int
    window time.Duration
}

func (drl *DistributedRateLimiter) Allow() bool {
    script := `
        local key = KEYS[1]
        local window = tonumber(ARGV[1])
        local rate = tonumber(ARGV[2])
        local now = tonumber(ARGV[3])
        
        -- Remove old entries
        redis.call('zremrangebyscore', key, '-inf', now - window)
        
        -- Count current entries
        local count = redis.call('zcard', key)
        
        if count < rate then
            -- Add current request
            redis.call('zadd', key, now, now)
            redis.call('expire', key, window)
            return 1
        else
            return 0
        end
    `
    
    now := time.Now().Unix()
    result := drl.redis.Eval(script, []string{drl.key}, 
        drl.window.Seconds(), drl.rate, now)
    
    return result.Val().(int64) == 1
}
```

### 2. Adaptive Rate Limiting

Adjusts rate limits based on system conditions and performance metrics.

#### Strategies

1. **Load-Based**: Adjust limits based on CPU, memory, or latency
2. **Queue-Based**: Use queue length as an indicator
3. **Success Rate**: Reduce limits when error rates increase

```go
type AdaptiveRateLimiter struct {
    baseLimiter  RateLimiter
    baseRate     int
    currentRate  int
    metrics      *SystemMetrics
    mutex        sync.RWMutex
}

func (arl *AdaptiveRateLimiter) adjustRate() {
    arl.mutex.Lock()
    defer arl.mutex.Unlock()
    
    // Get current system metrics
    cpuUsage := arl.metrics.GetCPUUsage()
    errorRate := arl.metrics.GetErrorRate()
    
    // Adjust rate based on conditions
    if cpuUsage > 0.8 || errorRate > 0.1 {
        // Reduce rate when system is stressed
        arl.currentRate = int(float64(arl.baseRate) * 0.5)
    } else if cpuUsage < 0.4 && errorRate < 0.01 {
        // Increase rate when system is healthy
        arl.currentRate = int(float64(arl.baseRate) * 1.2)
    }
    
    // Update the underlying limiter
    // (implementation depends on limiter type)
}
```

### 3. Rate Limiting Patterns

#### Per-User Rate Limiting

```go
type PerUserRateLimiter struct {
    limiters map[string]RateLimiter
    factory  *RateLimiterFactory
    config   RateLimiterConfig
    mutex    sync.RWMutex
}

func (purl *PerUserRateLimiter) Allow(userID string) bool {
    purl.mutex.RLock()
    limiter, exists := purl.limiters[userID]
    purl.mutex.RUnlock()
    
    if !exists {
        purl.mutex.Lock()
        // Double-check pattern
        if limiter, exists = purl.limiters[userID]; !exists {
            limiter, _ = purl.factory.CreateLimiter(purl.config)
            purl.limiters[userID] = limiter
        }
        purl.mutex.Unlock()
    }
    
    return limiter.Allow()
}
```

#### Hierarchical Rate Limiting

```go
type HierarchicalRateLimiter struct {
    globalLimiter RateLimiter
    userLimiters  map[string]RateLimiter
}

func (hrl *HierarchicalRateLimiter) Allow(userID string) bool {
    // Check global limit first
    if !hrl.globalLimiter.Allow() {
        return false
    }
    
    // Then check per-user limit
    userLimiter := hrl.getUserLimiter(userID)
    if !userLimiter.Allow() {
        // Return token to global limiter if user limit exceeded
        // (implementation depends on limiter type)
        return false
    }
    
    return true
}
```

## Concurrency and Thread Safety

### Key Considerations

1. **Race Conditions**: Multiple goroutines accessing shared state
2. **Atomic Operations**: Use atomic operations for simple counters
3. **Mutex Protection**: Protect complex state with mutexes
4. **Lock-Free Algorithms**: Consider lock-free approaches for high performance

### Thread-Safe Implementation Patterns

```go
// Using atomic operations for simple counters
type AtomicCounter struct {
    count int64
    limit int64
}

func (ac *AtomicCounter) Allow() bool {
    current := atomic.LoadInt64(&ac.count)
    if current >= ac.limit {
        return false
    }
    
    // Try to increment atomically
    newCount := atomic.AddInt64(&ac.count, 1)
    return newCount <= ac.limit
}

// Using read-write mutexes for better read performance
type RWMutexLimiter struct {
    mu    sync.RWMutex
    count int
    limit int
}

func (rwl *RWMutexLimiter) Allow() bool {
    rwl.mu.Lock()
    defer rwl.mu.Unlock()
    
    if rwl.count >= rwl.limit {
        return false
    }
    
    rwl.count++
    return true
}
```

## Performance Optimization

### 1. Minimize Lock Contention

```go
// Use separate locks for different operations
type OptimizedLimiter struct {
    // Separate mutexes for different concerns
    tokenMu   sync.Mutex
    metricsMu sync.Mutex
    
    tokens  float64
    metrics RateLimiterMetrics
}
```

### 2. Batch Operations

```go
func (tb *TokenBucket) AllowN(n int) bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()
    
    tb.refillTokens()
    
    if tb.tokens >= float64(n) {
        tb.tokens -= float64(n)
        return true
    }
    
    return false
}
```

### 3. Lazy Cleanup

```go
// Only clean up old requests when necessary
func (sw *SlidingWindow) cleanupIfNeeded() {
    if len(sw.requests) > sw.maxSize {
        sw.cleanup()
    }
}
```

## Testing Rate Limiters

### Unit Testing Strategies

1. **Basic Functionality**: Test allow/deny behavior
2. **Timing Tests**: Verify rate limiting over time
3. **Concurrency Tests**: Test thread safety
4. **Edge Cases**: Test boundary conditions

### Integration Testing

```go
func TestRateLimiterWithRealTraffic(t *testing.T) {
    limiter := NewTokenBucketLimiter(100, 10)
    
    // Simulate realistic traffic patterns
    var wg sync.WaitGroup
    clients := 50
    duration := 5 * time.Second
    
    for i := 0; i < clients; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            
            end := time.Now().Add(duration)
            for time.Now().Before(end) {
                limiter.Allow()
                time.Sleep(time.Millisecond * 10)
            }
        }()
    }
    
    wg.Wait()
    
    // Verify metrics and behavior
    metrics := limiter.GetMetrics()
    // Assert expected behavior
}
```

### Performance Benchmarking

```go
func BenchmarkRateLimiter(b *testing.B) {
    limiter := NewTokenBucketLimiter(1000000, 1000)
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            limiter.Allow()
        }
    })
}
```

## Real-World Applications

### 1. API Rate Limiting

```go
func APIRateLimitMiddleware(limiter RateLimiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Limit()))
            c.Header("X-RateLimit-Remaining", "0")
            c.Header("Retry-After", "1")
            c.AbortWithStatusJSON(429, gin.H{
                "error": "Rate limit exceeded",
            })
            return
        }
        
        c.Next()
    }
}
```

### 2. Database Connection Limiting

```go
type DBConnectionLimiter struct {
    limiter RateLimiter
    db      *sql.DB
}

func (dcl *DBConnectionLimiter) Query(query string, args ...interface{}) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := dcl.limiter.Wait(ctx); err != nil {
        return fmt.Errorf("database rate limit exceeded: %w", err)
    }
    
    _, err := dcl.db.QueryContext(ctx, query, args...)
    return err
}
```

### 3. Background Job Processing

```go
type JobProcessor struct {
    limiter RateLimiter
    queue   <-chan Job
}

func (jp *JobProcessor) processJobs() {
    for job := range jp.queue {
        if jp.limiter.Allow() {
            go jp.processJob(job)
        } else {
            // Queue job for later or drop it
            jp.requeueJob(job)
        }
    }
}
```

## Best Practices

### 1. Configuration

- **Choose Appropriate Algorithm**: Token bucket for burst, sliding window for accuracy
- **Set Reasonable Limits**: Based on system capacity and SLA requirements
- **Monitor and Adjust**: Continuously monitor and tune rate limits

### 2. Error Handling

- **Graceful Degradation**: Provide meaningful error messages
- **Retry Logic**: Implement exponential backoff for clients
- **Circuit Breaking**: Combine with circuit breaker patterns

### 3. Observability

- **Metrics Collection**: Track allowed/denied requests, wait times
- **Logging**: Log rate limiting events for debugging
- **Alerting**: Alert on unusual rate limiting patterns

### 4. Client-Side Considerations

```go
type RateLimitedClient struct {
    client  *http.Client
    limiter RateLimiter
}

func (rlc *RateLimitedClient) Do(req *http.Request) (*http.Response, error) {
    ctx := req.Context()
    
    // Wait for rate limiter approval
    if err := rlc.limiter.Wait(ctx); err != nil {
        return nil, fmt.Errorf("rate limit wait failed: %w", err)
    }
    
    return rlc.client.Do(req)
}
```

## Common Pitfalls and Solutions

### 1. Clock Skew in Distributed Systems
- **Problem**: Different servers have different times
- **Solution**: Use logical clocks or synchronized time sources

### 2. Memory Leaks
- **Problem**: Storing too much historical data
- **Solution**: Implement cleanup mechanisms and bounded storage

### 3. Thundering Herd
- **Problem**: Many requests hitting at window boundary
- **Solution**: Use jitter or staggered resets

### 4. Precision vs Performance
- **Problem**: High precision requires complex calculations
- **Solution**: Balance precision needs with performance requirements

## Further Reading

- [Go's golang.org/x/time/rate package](https://pkg.go.dev/golang.org/x/time/rate)
- [Rate Limiting Algorithms](https://en.wikipedia.org/wiki/Rate_limiting)
- [The Go Memory Model](https://golang.org/ref/mem)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)
- [Concurrency in Go (Book)](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)

## Conclusion

Rate limiting is a critical component of robust, scalable systems. Understanding different algorithms and their trade-offs allows you to choose the right approach for your specific use case. Remember to always test your rate limiters under realistic conditions and monitor their behavior in production. 