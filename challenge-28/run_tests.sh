#!/bin/bash

echo "🔥 Running Cache Implementation Tests"
echo "====================================="

# Set up Go module if needed
if [ ! -f "go.mod" ]; then
    echo "Initializing Go module..."
    go mod init cache-challenge
fi

echo ""
echo "📊 Running Basic Tests..."
go test -v

echo ""
echo "🏎️  Running Benchmark Tests..."
go test -bench=. -benchmem

echo ""
echo "🔄 Running Race Detection Tests..."
go test -v -race

echo ""
echo "⚡ Running Coverage Analysis..."
go test -cover -coverprofile=coverage.out
if [ -f coverage.out ]; then
    echo "📈 Coverage Report:"
    go tool cover -func=coverage.out | tail -1
    echo "   (Run 'go tool cover -html=coverage.out' to see detailed coverage)"
fi

echo ""
echo "🧪 Running Stress Tests..."
go test -v -timeout=30s

echo ""
echo "✅ All tests completed!"
echo ""
echo "💡 Quick Performance Check:"
echo "   Expected: O(1) time complexity for Get/Put/Delete operations"
echo "   The benchmark tests above should show consistent performance"
echo "   regardless of cache size (within reasonable limits)." 