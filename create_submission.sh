#!/bin/bash

# Script to create a submission directory and copy the solution template

# Function to display usage
usage() {
    echo "Usage: $0 [challenge-number]"
    exit 1
}

# Check if challenge number is provided
if [ -z "$1" ]; then
    echo "Error: No challenge number provided."
    usage
fi

CHALLENGE="challenge-$1"

# Check if challenge directory exists
if [ ! -d "$CHALLENGE" ]; then
    echo "Error: Challenge directory '$CHALLENGE' does not exist."
    exit 1
fi

# Prompt for GitHub username
read -p "Enter your GitHub username: " USERNAME

# Create the submission directory
SUBMISSION_DIR="$CHALLENGE/submissions/$USERNAME"
if [ -d "$SUBMISSION_DIR" ]; then
    echo "Submission directory '$SUBMISSION_DIR' already exists."
else
    mkdir -p "$SUBMISSION_DIR"
    echo "Created submission directory '$SUBMISSION_DIR'."
fi

# Copy the solution template
cp "$CHALLENGE/solution-template.go" "$SUBMISSION_DIR/"

echo "Copied solution template to your submission directory."

echo "Your submission directory is ready at '$SUBMISSION_DIR'."
echo "You can edit your solution in '$SUBMISSION_DIR/solution-template.go'."

# Check if learning materials exist and inform the user
if [ -f "$CHALLENGE/learning.md" ]; then
    echo "Learning materials for this challenge are available at '$CHALLENGE/learning.md'."
    echo "You can also view them through the web UI on the challenge page."
fi

# Optional: Initialize a Go module in the challenge directory if go.mod doesn't exist
cd "$CHALLENGE"
if [ ! -f "go.mod" ]; then
    echo "Initializing Go module in '$CHALLENGE' directory."
    go mod init "$CHALLENGE" || echo "Failed to initialize Go module."
else
    echo "'go.mod' file already exists in '$CHALLENGE' directory."
fi

echo "Setup complete."