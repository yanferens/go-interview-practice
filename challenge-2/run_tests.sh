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
SOLUTION_FILE="$SUBMISSION_DIR/solution-template.go"

# Check if the submission file exists
if [ ! -f "$SOLUTION_FILE" ]; then
    echo "Error: Solution file '$SOLUTION_FILE' not found."
    exit 1
fi

# Ensure go.mod exists
if [ ! -f "go.mod" ]; then
    echo "Initializing Go module in the current directory."
    go mod init "challenge" || echo "Failed to initialize Go module."
fi

# Copy the participant's solution to the current directory
cp "$SOLUTION_FILE" .

echo "Running tests for user '$USERNAME'..."

# Update the test file to point to the participant's solution
TEST_FILE="solution-template_test.go"

# Backup the original test file
cp "$TEST_FILE" "${TEST_FILE}.bak"

# Replace the cmd execution line to point to the participant's solution
sed -i.bak "s|go run .*|go run solution-template.go|" "$TEST_FILE"

# Run the tests
go test -v

TEST_EXIT_CODE=$?

# Restore the original test file
mv "${TEST_FILE}.bak" "$TEST_FILE"
rm -f "${TEST_FILE}.bak"

# Clean up the copied solution file
rm -f "solution-template.go"

exit $TEST_EXIT_CODE