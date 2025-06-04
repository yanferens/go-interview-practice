#!/bin/bash

# Script to run tests for a participant's submission

# Function to display usage
usage() {
    echo "Usage: $0"
    exit 1
}

# Verify that we are in a challenge directory
if [ ! -f "solution-template_test.go" ]; then
    echo "Error: solution-template_test.go not found. Please run this script from a challenge directory."
    exit 1
fi

# Prompt for GitHub username
read -p "Enter your GitHub username: " USERNAME

SUBMISSION_DIR="submissions/$USERNAME"
SUBMISSION_FILE="$SUBMISSION_DIR/solution-template.go"

# Check if the submission file exists
if [ ! -f "$SUBMISSION_FILE" ]; then
    echo "Error: Solution file '$SUBMISSION_FILE' not found."
    exit 1
fi

# Create a temporary directory to avoid modifying the original files
TEMP_DIR=$(mktemp -d)

# Copy the participant's solution and the test file to the temporary directory
cp "$SUBMISSION_FILE" "solution-template_test.go" "$TEMP_DIR/"

echo "🔥 Running Cache Implementation Tests for user '$USERNAME'..."
echo "============================================================"

# Navigate to the temporary directory
pushd "$TEMP_DIR" > /dev/null

# Initialize a new Go module in the temporary directory
go mod init "challenge" || {
  echo "Failed to initialize Go module."
  popd > /dev/null
  rm -rf "$TEMP_DIR"
  exit 1
}

echo ""
echo "📊 Running Basic Tests..."
go test -v
TEST_EXIT_CODE=$?

if [ $TEST_EXIT_CODE -eq 0 ]; then
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
fi

# Return to the original directory
popd > /dev/null

# Clean up the temporary directory
rm -rf "$TEMP_DIR"

exit $TEST_EXIT_CODE 