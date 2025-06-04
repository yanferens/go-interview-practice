package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Test Token Bucket Rate Limiter
func TestTokenBucketLimiter_Basic(t *testing.T) {
	limiter := NewTokenBucketLimiter(10, 5) // 10 req/sec, burst of 5

	// Test burst capacity
	allowed := 0
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			allowed++
		}
	}

	if allowed != 5 {
		t.Errorf("Expected 5 requests to be allowed initially (burst), got %d", allowed)
	}

	// Test rate limiting
	time.Sleep(100 * time.Millisecond) // Allow some token refill
	if !limiter.Allow() {
		t.Error("Expected at least one request to be allowed after waiting")
	}
}

func TestTokenBucketLimiter_AllowN(t *testing.T) {
	limiter := NewTokenBucketLimiter(10, 5)

	// Test allowing multiple tokens at once
	if !limiter.AllowN(3) {
		t.Error("Expected AllowN(3) to succeed with burst of 5")
	}

	if limiter.AllowN(3) {
		t.Error("Expected AllowN(3) to fail after consuming 3 tokens (2 remaining)")
	}

	if !limiter.AllowN(2) {
		t.Error("Expected AllowN(2) to succeed with 2 tokens remaining")
	}
}

func TestTokenBucketLimiter_Wait(t *testing.T) {
	limiter := NewTokenBucketLimiter(10, 1) // 10 req/sec, burst of 1

	// Exhaust initial token
	if !limiter.Allow() {
		t.Fatal("Expected first request to be allowed")
	}

	// Test waiting for next token
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := limiter.Wait(ctx)
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected Wait to succeed, got error: %v", err)
	}

	// Should wait approximately 100ms for next token (1/10 second)
	if elapsed < 50*time.Millisecond || elapsed > 200*time.Millisecond {
		t.Errorf("Expected wait time around 100ms, got %v", elapsed)
	}
}

func TestTokenBucketLimiter_WaitTimeout(t *testing.T) {
	limiter := NewTokenBucketLimiter(1, 1) // 1 req/sec, burst of 1

	// Exhaust token
	limiter.Allow()

	// Test timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := limiter.Wait(ctx)
	if err == nil {
		t.Error("Expected Wait to timeout, but it succeeded")
	}
}

func TestTokenBucketLimiter_Concurrent(t *testing.T) {
	limiter := NewTokenBucketLimiter(100, 10)
	var allowed, denied int64

	var wg sync.WaitGroup
	numGoroutines := 50
	requestsPerGoroutine := 20

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerGoroutine; j++ {
				if limiter.Allow() {
					atomic.AddInt64(&allowed, 1)
				} else {
					atomic.AddInt64(&denied, 1)
				}
			}
		}()
	}

	wg.Wait()

	total := allowed + denied
	expectedTotal := int64(numGoroutines * requestsPerGoroutine)

	if total != expectedTotal {
		t.Errorf("Expected total requests %d, got %d", expectedTotal, total)
	}

	// Should have some allowed and some denied requests
	if allowed == 0 {
		t.Error("Expected some requests to be allowed")
	}
	if denied == 0 {
		t.Error("Expected some requests to be denied")
	}
}

// Test Sliding Window Rate Limiter
func TestSlidingWindowLimiter_Basic(t *testing.T) {
	limiter := NewSlidingWindowLimiter(5, time.Second) // 5 req/sec

	// Test initial requests
	allowed := 0
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			allowed++
		}
	}

	if allowed != 5 {
		t.Errorf("Expected 5 requests to be allowed initially, got %d", allowed)
	}

	// Wait for window to slide and test again
	time.Sleep(1100 * time.Millisecond)
	if !limiter.Allow() {
		t.Error("Expected request to be allowed after window slide")
	}
}

func TestSlidingWindowLimiter_WindowSliding(t *testing.T) {
	limiter := NewSlidingWindowLimiter(2, 100*time.Millisecond) // 2 req per 100ms

	// Use up the limit
	if !limiter.Allow() || !limiter.Allow() {
		t.Fatal("Expected first two requests to be allowed")
	}

	// Should be rate limited now
	if limiter.Allow() {
		t.Error("Expected third request to be rate limited")
	}

	// Wait for half the window to pass
	time.Sleep(60 * time.Millisecond)

	// Still should be rate limited
	if limiter.Allow() {
		t.Error("Expected request to still be rate limited")
	}

	// Wait for full window to pass
	time.Sleep(60 * time.Millisecond)

	// Should be allowed now
	if !limiter.Allow() {
		t.Error("Expected request to be allowed after window slide")
	}
}

// Test Fixed Window Rate Limiter
func TestFixedWindowLimiter_Basic(t *testing.T) {
	limiter := NewFixedWindowLimiter(3, 100*time.Millisecond) // 3 req per 100ms window

	// Test window capacity
	allowed := 0
	for i := 0; i < 5; i++ {
		if limiter.Allow() {
			allowed++
		}
	}

	if allowed != 3 {
		t.Errorf("Expected 3 requests to be allowed in window, got %d", allowed)
	}

	// Wait for new window
	time.Sleep(110 * time.Millisecond)

	// Should be allowed in new window
	if !limiter.Allow() {
		t.Error("Expected request to be allowed in new window")
	}
}

func TestFixedWindowLimiter_WindowReset(t *testing.T) {
	limiter := NewFixedWindowLimiter(2, 50*time.Millisecond)

	// Use window capacity
	limiter.Allow()
	limiter.Allow()

	// Should be rate limited
	if limiter.Allow() {
		t.Error("Expected request to be rate limited")
	}

	// Wait for window reset
	time.Sleep(60 * time.Millisecond)

	// Should be allowed in new window
	if !limiter.Allow() {
		t.Error("Expected request to be allowed after window reset")
	}
}

// Test Rate Limiter Factory
func TestRateLimiterFactory(t *testing.T) {
	factory := NewRateLimiterFactory()

	tests := []struct {
		name         string
		config       RateLimiterConfig
		expectError  bool
		expectedType string
	}{
		{
			name: "Valid Token Bucket",
			config: RateLimiterConfig{
				Algorithm: "token_bucket",
				Rate:      10,
				Burst:     5,
			},
			expectError:  false,
			expectedType: "*main.TokenBucketLimiter",
		},
		{
			name: "Valid Sliding Window",
			config: RateLimiterConfig{
				Algorithm:  "sliding_window",
				Rate:       5,
				WindowSize: time.Second,
			},
			expectError:  false,
			expectedType: "*main.SlidingWindowLimiter",
		},
		{
			name: "Valid Fixed Window",
			config: RateLimiterConfig{
				Algorithm:  "fixed_window",
				Rate:       3,
				WindowSize: time.Second,
			},
			expectError:  false,
			expectedType: "*main.FixedWindowLimiter",
		},
		{
			name: "Invalid Algorithm",
			config: RateLimiterConfig{
				Algorithm: "invalid",
				Rate:      10,
			},
			expectError: true,
		},
		{
			name: "Invalid Token Bucket Rate",
			config: RateLimiterConfig{
				Algorithm: "token_bucket",
				Rate:      0,
				Burst:     5,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter, err := factory.CreateLimiter(tt.config)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Expected no error but got: %v", err)
				return
			}

			if limiter == nil {
				t.Error("Expected limiter to be created")
				return
			}

			// Test basic functionality
			limiter.Allow()

			if limiter.Limit() != tt.config.Rate {
				t.Errorf("Expected rate %d, got %d", tt.config.Rate, limiter.Limit())
			}
		})
	}
}

// Test HTTP Middleware
func TestRateLimitMiddleware(t *testing.T) {
	limiter := NewTokenBucketLimiter(2, 2) // 2 req/sec, burst of 2

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	middleware := RateLimitMiddleware(limiter)
	wrappedHandler := middleware(handler)

	// First two requests should succeed
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		wrappedHandler.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Request %d: expected status 200, got %d", i+1, w.Code)
		}
	}

	// Third request should be rate limited
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(w, req)

	if w.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got %d", w.Code)
	}

	// Check rate limit headers
	if w.Header().Get("X-RateLimit-Limit") != "2" {
		t.Errorf("Expected X-RateLimit-Limit header to be '2', got '%s'",
			w.Header().Get("X-RateLimit-Limit"))
	}
}

// Test Metrics
func TestRateLimiterMetrics(t *testing.T) {
	limiter := NewTokenBucketLimiter(5, 5)

	// Make some requests
	for i := 0; i < 10; i++ {
		limiter.Allow()
	}

	metrics := limiter.GetMetrics()

	if metrics.TotalRequests != 10 {
		t.Errorf("Expected total requests to be 10, got %d", metrics.TotalRequests)
	}

	if metrics.AllowedRequests != 5 {
		t.Errorf("Expected allowed requests to be 5, got %d", metrics.AllowedRequests)
	}

	if metrics.DeniedRequests != 5 {
		t.Errorf("Expected denied requests to be 5, got %d", metrics.DeniedRequests)
	}
}

// Test Reset functionality
func TestRateLimiterReset(t *testing.T) {
	limiter := NewTokenBucketLimiter(10, 2)

	// Exhaust tokens
	limiter.Allow()
	limiter.Allow()

	// Should be rate limited
	if limiter.Allow() {
		t.Error("Expected request to be rate limited before reset")
	}

	// Reset limiter
	limiter.Reset()

	// Should be allowed after reset
	if !limiter.Allow() {
		t.Error("Expected request to be allowed after reset")
	}

	// Metrics should be reset
	metrics := limiter.GetMetrics()
	if metrics.TotalRequests != 1 { // Only the post-reset request
		t.Errorf("Expected total requests to be 1 after reset, got %d", metrics.TotalRequests)
	}
}

// Benchmark Tests
func BenchmarkTokenBucketLimiter_Allow(b *testing.B) {
	limiter := NewTokenBucketLimiter(1000000, 1000) // High rate to avoid limiting

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow()
		}
	})
}

func BenchmarkSlidingWindowLimiter_Allow(b *testing.B) {
	limiter := NewSlidingWindowLimiter(1000000, time.Second)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow()
		}
	})
}

func BenchmarkFixedWindowLimiter_Allow(b *testing.B) {
	limiter := NewFixedWindowLimiter(1000000, time.Second)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow()
		}
	})
}

// Test for memory leaks and goroutine leaks
func TestRateLimiterMemoryLeaks(t *testing.T) {
	initialGoroutines := runtime.NumGoroutine()

	// Create and use multiple limiters
	for i := 0; i < 100; i++ {
		limiter := NewTokenBucketLimiter(10, 5)

		// Use the limiter
		for j := 0; j < 10; j++ {
			limiter.Allow()
		}

		// Test Wait with quick timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		limiter.Wait(ctx)
		cancel()
	}

	// Force garbage collection
	runtime.GC()
	runtime.GC()

	// Allow some time for cleanup
	time.Sleep(10 * time.Millisecond)

	finalGoroutines := runtime.NumGoroutine()

	// Should not have significant goroutine leaks
	if finalGoroutines > initialGoroutines+5 {
		t.Errorf("Potential goroutine leak: started with %d, ended with %d",
			initialGoroutines, finalGoroutines)
	}
}

// Integration test with realistic scenario
func TestRateLimiterIntegration(t *testing.T) {
	// Simulate a realistic API rate limiting scenario
	limiter := NewTokenBucketLimiter(100, 20) // 100 req/sec, burst of 20

	var successCount, errorCount int64
	var wg sync.WaitGroup

	// Simulate 10 concurrent clients
	numClients := 10
	requestsPerClient := 50

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			for j := 0; j < requestsPerClient; j++ {
				ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

				if err := limiter.Wait(ctx); err == nil {
					atomic.AddInt64(&successCount, 1)
				} else {
					atomic.AddInt64(&errorCount, 1)
				}

				cancel()
				time.Sleep(time.Millisecond) // Small delay between requests
			}
		}(i)
	}

	wg.Wait()

	total := successCount + errorCount
	expectedTotal := int64(numClients * requestsPerClient)

	if total != expectedTotal {
		t.Errorf("Expected total operations %d, got %d", expectedTotal, total)
	}

	// Check metrics
	metrics := limiter.GetMetrics()
	fmt.Printf("Integration test results: %d successful, %d timed out, metrics: %+v\n",
		successCount, errorCount, metrics)

	// Should have processed some requests successfully
	if successCount == 0 {
		t.Error("Expected some requests to succeed")
	}
}

// Edge case tests
func TestRateLimiterEdgeCases(t *testing.T) {
	// Test with very small rates
	limiter := NewTokenBucketLimiter(1, 1)
	if !limiter.Allow() {
		t.Error("Expected first request to be allowed")
	}
	if limiter.Allow() {
		t.Error("Expected second request to be denied")
	}

	// Test with zero burst (should still work with rate)
	// Note: This might not be valid depending on implementation
	// limiter2 := NewTokenBucketLimiter(10, 0)

	// Test factory with edge case configurations
	factory := NewRateLimiterFactory()

	_, err := factory.CreateLimiter(RateLimiterConfig{
		Algorithm: "token_bucket",
		Rate:      1,
		Burst:     1,
	})
	if err != nil {
		t.Errorf("Expected minimal valid config to work, got error: %v", err)
	}
}
