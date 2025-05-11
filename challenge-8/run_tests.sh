#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Get username
if [ $# -eq 1 ]; then
    USERNAME=$1
else
    read -p "Enter your GitHub username: " USERNAME
fi

# Check if submission directory exists
if [ ! -d "submissions/$USERNAME" ]; then
    echo -e "${RED}Directory submissions/$USERNAME does not exist.${NC}"
    echo -e "${YELLOW}Use the create_submission.sh script from the root directory to set up your submission.${NC}"
    exit 1
fi

# Check if solution file exists
if [ ! -f "submissions/$USERNAME/solution-template.go" ]; then
    echo -e "${RED}Solution file submissions/$USERNAME/solution-template.go does not exist.${NC}"
    exit 1
fi

# Create a temporary file with a unique name for testing
TEMP_FILE="temp_solution_$(date +%s).go"

# Copy the solution file to the temporary file
cp "submissions/$USERNAME/solution-template.go" "$TEMP_FILE"

# Hide the original solution-template.go file by renaming it temporarily
mv solution-template.go solution-template.go.bak

# Run tests
echo -e "${YELLOW}Running tests...${NC}"
go test -v
TEST_RESULT=$?

# Restore the original solution-template.go file
mv solution-template.go.bak solution-template.go

# Clean up
rm -f "$TEMP_FILE"

# Check if tests passed and show appropriate message
if [ $TEST_RESULT -eq 0 ]; then
    echo -e "${GREEN}All tests passed! You can submit your solution.${NC}"
else
    echo -e "${RED}Tests failed. Please fix the issues and try again.${NC}"
    exit $TEST_RESULT
fi 