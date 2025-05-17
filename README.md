# Go Interview Practice

Welcome to the **Go Interview Practice** repository!

This repository contains a series of Go (Golang) coding challenges designed to help you prepare for technical interviews. Each challenge includes detailed instructions, automated tests, learning materials, and per-challenge scoreboards to track your progress. Access challenges through our interactive Web UI or work directly with code files.

## **Key Features**

- 🌐 **Interactive Web UI**: Modern web interface with in-browser code editor and real-time testing
- 📚 **Learning Materials**: Comprehensive Go concepts explanations for each challenge
- 🧪 **Automated Testing**: Immediate feedback on your solution's correctness
- 📊 **Scoreboards**: Track your progress and compare with others
- 🔄 **Continuous Integration**: GitHub Actions workflow to validate submissions
- 🎯 **Progressive Difficulty**: Challenges ranging from beginner to advanced levels

## **Table of Contents**

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Forking the Repository](#forking-the-repository)
  - [Cloning Your Fork](#cloning-your-fork)
- [Web UI](#web-ui)
  - [Features](#features)
  - [Running the Web UI](#running-the-web-ui)
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

## **Web UI**

The Go Interview Practice includes a modern, interactive web interface for a more engaging learning experience.

![Go Interview Practice Web UI](./images/challenges.png)

### **Features**

- **Interactive Challenge Browser**: Browse all coding challenges with difficulty filters and sorting options
- **In-browser Code Editor**: Write and edit Go code directly in your browser
- **Integrated Test Runner**: Run tests against your solutions in real-time
- **Learning Materials**: Access detailed learning content for each challenge
- **Scoreboards**: Track your progress and compare with other users
- **Modern UI**: Clean, responsive design with intuitive navigation
- **Direct Filesystem Submission**: Submit your solution directly to your local repository with one click

### **Running the Web UI**

1. Navigate to the web-ui directory:
   ```bash
   cd web-ui
   ```

2. Run the web server:
   ```bash
   go run main.go
   ```

3. Open your browser and visit:
   ```
   http://localhost:8080
   ```

For more details about the Web UI, check the [web-ui/README.md](web-ui/README.md) file.

---

## **Working on Challenges**

### **Setting Up Your Submission**

You can set up your submission in two ways:

#### **Method 1: Using the Web UI (Recommended for Development)**

1. Run the web UI as described in the [Web UI](#web-ui) section.
2. Browse challenges and select one to work on.
3. Implement your solution in the in-browser editor.
4. Click "Run Tests" to test your solution.
5. Once all tests pass, click "Submit Solution" to submit your solution to the scoreboard.
6. You will be given the option to either:
   - Save your solution directly to the filesystem with one click
   - Copy the commands to manually create the submission files

#### **Method 2: Using the Command Line**

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

The challenges are organized in increasing difficulty from beginner to advanced levels. Start from the easier ones if you're new to Go, and progress to the more complex ones as you build your skills. You can access these challenges through the Web UI or directly via code files.

Each challenge includes:
- Detailed problem statement
- Function signature to implement
- Test cases to validate your solution
- Learning materials related to the concepts required for the challenge

| # | Challenge | Difficulty | Focus |
|---|-----------|------------|-------|
| [1](./challenge-1) | Sum of Two Numbers | Beginner | Basic Go syntax, function implementation |
| [2](./challenge-2) | Reverse a String | Beginner | String manipulation, runes, and character handling |
| [3](./challenge-3) | Employee Data Management | Beginner | Structs, methods, slice operations |
| [4](./challenge-4) | Concurrent Graph BFS Queries | Intermediate | Goroutines, channels, graph algorithms |
| [5](./challenge-5) | HTTP Authentication Middleware | Intermediate | HTTP handling, middleware pattern |
| [6](./challenge-6) | Word Frequency Counter | Beginner | Maps, string handling, data processing |
| [7](./challenge-7) | Bank Account with Error Handling | Intermediate | Custom errors, error handling, thread safety |
| [8](./challenge-8) | Chat Server with Channels | Advanced | Concurrency patterns, message passing, goroutine management |
| [9](./challenge-9) | RESTful Book Management API | Advanced | Web services, HTTP routing, JSON handling, database interactions |
| [10](./challenge-10) | Polymorphic Shape Calculator | Intermediate | Interfaces, polymorphism, Go's interface system |
| [11](./challenge-11) | Concurrent Web Content Aggregator | Advanced | Advanced concurrency patterns, context handling, rate limiting |
| [12](./challenge-12) | File Processing Pipeline with Advanced Error Handling | Advanced | Error wrapping, custom error types, error handling in concurrent code |
| [13](./challenge-13) | Trie Data Structure | Intermediate | Data structures, tree implementation, string processing |
| [14](./challenge-14) | Rate Limiter Implementation | Intermediate | Concurrency control, time handling, distributed systems concepts |
| [15](./challenge-15) | Consistent Hashing | Advanced | Distributed systems, hashing algorithms, load balancing |
| [16](./challenge-16) | Builder Pattern Implementation | Intermediate | Design patterns, fluent interfaces, object creation |
| [17](./challenge-17) | Worker Pool Pattern | Intermediate | Concurrency patterns, resource management, task distribution |
| [18](./challenge-18) | Distributed Lock | Advanced | Distributed systems, synchronization, consensus |
| [19](./challenge-19) | Functional Options Pattern | Intermediate | API design, function closures, configuration patterns |
| [20](./challenge-20) | Circuit Breaker Pattern | Advanced | Resilience patterns, fault tolerance, state management |
| [21](./challenge-21) | Binary Search Implementation | Easy | Search algorithms, array operations, recursion vs. iteration |
| [22](./challenge-22) | Greedy Coin Change | Medium-Easy | Greedy algorithms, optimization problems, currency systems |
| [23](./challenge-23) | String Pattern Matching | Medium | String algorithms, pattern matching, time complexity optimization |
| [24](./challenge-24) | Dynamic Programming - Longest Increasing Subsequence | Medium-Hard | Dynamic programming, optimization, algorithm complexity |
| [25](./challenge-25) | Graph Algorithms - Shortest Path | Hard | Graph theory, path-finding algorithms, BFS, Dijkstra, Bellman-Ford |

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
