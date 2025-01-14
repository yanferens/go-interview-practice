# **Contributing to Go Interview Practice**

Thank you for your interest in contributing to the **Go Interview Practice** repository! We welcome contributions from the community to help improve this project. Whether you want to submit solutions, add new challenges, or improve documentation, your efforts are appreciated.

## **Table of Contents**

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
  - [Submitting a Solution](#submitting-a-solution)
  - [Adding a New Challenge](#adding-a-new-challenge)
- [Style Guidelines](#style-guidelines)
- [Pull Request Process](#pull-request-process)
- [Reporting Issues](#reporting-issues)
- [Contact](#contact)

---

## **Code of Conduct**

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project, you agree to abide by its terms.

---

## **How to Contribute**

### **Submitting a Solution**

If you want to submit a solution to an existing challenge:

1. **Fork the Repository:**

   - Click the "Fork" button on the repository page.

2. **Clone Your Fork:**

   ```bash
   git clone https://github.com/yourusername/go-interview-practice.git
   ```

3. **Create a New Branch:**

   ```bash
   git checkout -b challenge-[number]-solution
   ```

4. **Set Up Your Submission:**

   - Use the provided script to set up your submission:

     ```bash
     ./create_submission.sh [challenge-number]
     ```

5. **Implement Your Solution:**

   - Edit the `solution-template.go` file in your submission directory.
   - Ensure your code passes all the tests.

6. **Run Tests Locally:**

   - Navigate to the challenge directory and use `run_tests.sh`:

     ```bash
     cd challenge-[number]
     ./run_tests.sh
     ```

7. **Commit and Push:**

   ```bash
   git add challenge-[number]/submissions/yourusername/
   git commit -m "Add Challenge [number] solution by [yourusername]"
   git push origin challenge-[number]-solution
   ```

8. **Create a Pull Request:**

   - Go to your fork on GitHub and open a pull request to the `main` branch.

9. **Receive Feedback:**

   - The automated tests will run on your pull request.
   - Address any comments or requested changes.

### **Adding a New Challenge**

If you have an idea for a new challenge:

1. **Create a New Issue:**

   - Open an issue to discuss the new challenge idea.
   - Provide details such as the problem statement and its relevance.

2. **Wait for Approval:**

   - Wait for maintainers or community members to provide feedback.

3. **Create a New Branch:**

   ```bash
   git checkout -b add-challenge-[number]
   ```

4. **Set Up the Challenge Directory:**

   ```
   challenge-[number]/
   ├── README.md
   ├── solution-template.go
   ├── solution-template_test.go
   └── submissions/
   ```

5. **Write the Challenge Description:**

   - Include problem statement, function signature, input/output format, constraints, and sample inputs/outputs in `README.md`.

6. **Create the Solution Template:**

   - Provide a skeleton code in `solution-template.go` with appropriate comments.

7. **Write Comprehensive Tests:**

   - Create `solution-template_test.go` with detailed test cases covering various scenarios, including edge cases.

8. **Update Documentation:**

   - Add the new challenge to the main `README.md`.

9. **Commit and Push:**

   ```bash
   git add challenge-[number]/
   git commit -m "Add Challenge [number]: [Challenge Title]"
   git push origin add-challenge-[number]
   ```

10. **Create a Pull Request:**

    - Submit the pull request for review.
    - Ensure all tests pass in the CI workflow.

---

## **Style Guidelines**

- **Code Formatting:**

  - Use `gofmt` to format your Go code.
  - Maintain consistent indentation and spacing.

- **Naming Conventions:**

  - Use descriptive variable and function names.
  - Follow Go naming conventions (e.g., `camelCase` for variables and functions).

- **Comments:**

  - Include comments to explain complex logic.
  - Use doc comments (`//`) for exported functions and types.

- **Test Writing:**

  - Write thorough tests covering various input cases.
  - Use subtests (`t.Run()`) to organize test cases.

---

## **Pull Request Process**

1. **Ensure All Tests Pass:**

   - Run tests locally before submitting your pull request.
   - Check that your code does not break existing functionality.

2. **Provide a Clear Description:**

   - Explain what changes you have made and why.
   - Reference any related issues.

3. **One Pull Request per Feature:**

   - Keep your pull requests focused on a single feature or fix.

4. **Wait for Review:**

   - A maintainer will review your pull request.
   - Be responsive to feedback and make necessary changes.

---

## **Reporting Issues**

If you encounter any problems or have suggestions:

- **Open an Issue:**

  - Go to the [Issues](https://github.com/RezaSi/go-interview-practice/issues) tab.
  - Provide a detailed description of the issue or suggestion.
  - Include steps to reproduce the issue if applicable.

---

## **Contact**

For any questions or additional support:

- **Email:** [rezashiri88@gmail.com](mailto:rezashiri88@gmail.com)
- **GitHub:** [RezaSi](https://github.com/RezaSi)

---

Thank you for contributing to the Go Interview Practice repository!
