#!/bin/bash

# ANSI color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if a GitHub username was provided as an argument
if [ $# -eq 1 ]; then
    GITHUB_USERNAME=$1
else
    # Prompt for GitHub username if not provided
    read -p "Enter your GitHub username: " GITHUB_USERNAME
fi

# Check if the submission directory exists
SUBMISSION_DIR="submissions/$GITHUB_USERNAME"
if [ ! -d "$SUBMISSION_DIR" ]; then
    echo -e "${RED}Error: Submission directory '$SUBMISSION_DIR' not found.${NC}"
    echo -e "Please run the ${BLUE}create_submission.sh${NC} script first to set up your submission directory."
    exit 1
fi

# Check if the solution file exists
SOLUTION_FILE="$SUBMISSION_DIR/solution-template.go"
if [ ! -f "$SOLUTION_FILE" ]; then
    echo -e "${RED}Error: Solution file '$SOLUTION_FILE' not found.${NC}"
    exit 1
fi

# Create a temporary copy of the solution
echo -e "${BLUE}Preparing to test your solution...${NC}"
cp "$SOLUTION_FILE" "solution-template.go.tmp"

# Run the tests
echo -e "${YELLOW}Running tests...${NC}"
TEST_OUTPUT=$(go test -v 2>&1)
TEST_RESULT=$?

# Display test results
if [ $TEST_RESULT -eq 0 ]; then
    echo -e "${GREEN}All tests passed!${NC}"
    
    # Extract test results for display
    PASSED_TESTS=$(echo "$TEST_OUTPUT" | grep -c "PASS")
    echo -e "${GREEN}Passed: $PASSED_TESTS tests${NC}"
    
    # Update scoreboard if all tests pass
    echo -e "${BLUE}Updating scoreboard...${NC}"
    # Check if the user is already on the scoreboard
    if grep -q "$GITHUB_USERNAME" SCOREBOARD.md; then
        # User already on scoreboard, do nothing
        echo -e "${YELLOW}You are already on the scoreboard.${NC}"
    else
        # Add user to scoreboard
        echo "| $GITHUB_USERNAME | ✅ |" >> SCOREBOARD.md
        echo -e "${GREEN}Added to scoreboard!${NC}"
    fi
else
    echo -e "${RED}Some tests failed. See details below:${NC}"
    echo "$TEST_OUTPUT" | grep -E "FAIL:|--- FAIL:"
fi

# Cleanup
rm "solution-template.go.tmp"

exit $TEST_RESULT 