# Go Interview Practice

Welcome to the **Go Interview Practice** repository!

This repository contains a series of Go (Golang) coding challenges designed to help you prepare for technical interviews. Each challenge includes detailed instructions, automated tests, and per-challenge scoreboards to track your progress. Bash scripts are provided to streamline solution submission and testing.

## **Table of Contents**

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Forking the Repository](#forking-the-repository)
  - [Cloning Your Fork](#cloning-your-fork)
- [Working on Challenges](#working-on-challenges)
  - [Setting Up Your Submission](#setting-up-your-submission)
  - [Implementing the Solution](#implementing-the-solution)
  - [Running Tests](#running-tests)
- [Submitting Your Solution](#submitting-your-solution)
  - [Creating a Pull Request](#creating-a-pull-request)
- [Scripts and Automation](#scripts-and-automation)
  - [create_submission.sh](#createsubmissionsh)
  - [run_tests.sh](#run_testssh)
- [Challenge List](#challenge-list)
  - [Beginner Challenges](#beginner-challenges)
  - [Intermediate Challenges](#intermediate-challenges)
  - [Advanced Challenges](#advanced-challenges)
- [Contributing](#contributing)
  - [Adding a New Challenge](#adding-a-new-challenge)
- [License](#license)

---

## **Getting Started**

### **Prerequisites**

- **Go (Golang) Installed**: Ensure you have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).
- **Git Installed**: You'll need Git to clone repositories and manage code versions. Download it from [here](https://git-scm.com/downloads).

### **Forking the Repository**

Forking the repository creates a copy under your own GitHub account where you can make changes without affecting the original repository.

1. Go to the [original repository](https://github.com/RezaSi/go-interview-practice).
2. Click the **Fork** button in the top right corner.

### **Cloning Your Fork**

Clone your forked repository to your local machine:

```bash
git clone https://github.com/yourusername/go-interview-practice.git
```

---

## **Working on Challenges**

### **Setting Up Your Submission**

We have provided a bash script called `create_submission.sh` to help you set up your submission directory and copy the solution template.

**Usage:**

```bash
./create_submission.sh [challenge-number]
```

**Example:**

```bash
./create_submission.sh 1
```

**Steps:**

1. Navigate to the root directory of the cloned repository.

2. Run the script with the challenge number you want to work on.

   ```bash
   ./create_submission.sh 1
   ```

3. The script will prompt you to enter your GitHub username. This is used to create a unique submission directory for you.

4. The script will create the directory `challenge-1/submissions/yourusername` and copy the solution template into it.

### **Implementing the Solution**

1. Navigate to your submission directory:

   ```bash
   cd challenge-1/submissions/yourusername/
   ```

2. Open the `solution-template.go` file in your preferred text editor.

3. Implement the required function(s) as specified in the challenge's `README.md`.

### **Running Tests**

Each challenge includes a `run_tests.sh` script to help you run tests against your solution.

**Usage:**

1. Navigate to the challenge directory:

   ```bash
   cd challenge-1
   ```

2. Run the test script:

   ```bash
   ./run_tests.sh
   ```

3. Enter your GitHub username when prompted.

4. The script will run the tests and display the results.

---

## **Submitting Your Solution**

### **Creating a Pull Request**

1. **Commit Your Changes:**

   ```bash
   git add challenge-1/submissions/yourusername/solution-template.go
   git commit -m "Add Challenge 1 solution"
   ```

2. **Push to Your Fork:**

   ```bash
   git push origin challenge-1-solution
   ```

3. **Create a Pull Request:**

   - Go to your forked repository on GitHub.
   - Click on the "Compare & pull request" button.
   - Submit the pull request to the original repository's `main` branch.

4. **Wait for Automated Feedback:**

   - The GitHub Actions workflow will automatically run tests on your submission.
   - Review the test results in the **Checks** tab of your pull request.
   - If tests pass, your submission will be reviewed and merged.
   - If tests fail, fix the issues and push the changes. The workflow will re-run the tests.

---

## **Scripts and Automation**

We've included two bash scripts to streamline your workflow:

### **create_submission.sh**

This script sets up your submission directory and copies the solution template.

**Location:** Root directory of the repository.

**Usage:**

```bash
./create_submission.sh [challenge-number]
```

**What It Does:**

- Prompts you for your GitHub username.
- Creates a submission directory under `challenge-[number]/submissions/yourusername`.
- Copies the `solution-template.go` into your submission directory.
- Initializes the Go module if necessary.

### **run_tests.sh**

This script runs the tests against your solution.

**Location:** Inside each challenge directory (e.g., `challenge-1/run_tests.sh`).

**Usage:**

```bash
./run_tests.sh
```

**What It Does:**

- Prompts you for your GitHub username.
- Copies your solution file into the challenge directory.
- Runs the tests.
- Cleans up any temporary files used during testing.

---

## **Challenge List**

The challenges are organized in increasing difficulty from beginner to advanced levels. Start from the easier ones if you're new to Go, and progress to the more complex ones as you build your skills.

### **Beginner Challenges**

1. **Challenge 1: Sum of Two Numbers**
   - Implement a function to sum two integers.
   - Focus: Basic Go syntax, function implementation.

2. **Challenge 2: Reverse a String**
   - Implement a function to reverse a string.
   - Focus: String manipulation, runes, and character handling.

3. **Challenge 3: Employee Data Management**
   - Implement a struct to manage employee data with basic operations.
   - Focus: Structs, methods, slice operations.

6. **Challenge 6: Word Frequency Counter**
   - Implement a function that counts word occurrences in a text.
   - Focus: Maps, string handling, data processing.

### **Intermediate Challenges**

4. **Challenge 4: Concurrent Graph BFS Queries**
   - Implement a BFS algorithm for graph traversal with concurrency.
   - Focus: Goroutines, channels, graph algorithms.

5. **Challenge 5: HTTP Authentication Middleware**
   - Implement an HTTP middleware for authentication.
   - Focus: HTTP handling, middleware pattern.

7. **Challenge 7: Bank Account with Error Handling**
   - Implement a banking system with proper error handling.
   - Focus: Custom errors, error handling, thread safety.

### **Advanced Challenges**

8. **Challenge 8: Chat Server with Channels**
   - Implement a chat server using channels and goroutines.
   - Focus: Concurrency patterns, message passing, goroutine management.

9. **Challenge 9: RESTful Book Management API**
   - Implement a RESTful API for a book management system.
   - Focus: Web services, HTTP routing, JSON handling, database interactions.

---

## **Contributing**

We welcome contributions! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on our code of conduct, and the process for submitting pull requests.

### **Adding a New Challenge**

If you'd like to add a new challenge to the repository, please follow these steps:

1. **Create a New Branch:**

   ```bash
   git checkout -b add-challenge-[number]
   ```

2. **Create the Challenge Directory Structure:**

   ```
   challenge-[number]/
   ├── README.md
   ├── solution-template.go
   ├── solution-template_test.go
   └── submissions/
   ```

3. **Write the Challenge Description:**

   - In `challenge-[number]/README.md`, provide:
     - Problem statement
     - Function signature
     - Input/output formats
     - Sample inputs/outputs
     - Detailed instructions

4. **Create the Solution Template:**

   - In `challenge-[number]/solution-template.go`, provide the skeleton code with function signatures and comments.

5. **Write the Test File:**

   - In `challenge-[number]/solution-template_test.go`, write comprehensive tests covering various cases.

6. **Update the Main README:**

   - Add the new challenge to the list of available challenges in the main `README.md`.

7. **Commit and Push:**

   ```bash
   git add challenge-[number]
   git commit -m "Add Challenge [number]: [Challenge Title]"
   git push origin add-challenge-[number]
   ```

8. **Create a Pull Request:**

   - Submit a pull request to the original repository.
   - Ensure that all tests pass in the continuous integration workflow.

---

## **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
