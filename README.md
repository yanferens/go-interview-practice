## **Instructions for Participants**

### **Step-by-Step Guide**

1. **Fork the Repository**

   - Click the "Fork" button at the top-right corner of the repository page on GitHub.

2. **Clone Your Fork**

   ```bash
   git clone https://github.com/yourusername/interview-practice.git
   ```

3. **Create a New Branch**

   ```bash
   cd interview-practice
   git checkout -b challenge-1-solution
   ```

4. **Create Your Submission Directory**

   ```bash
   mkdir -p challenge-1/submissions/yourusername
   ```

5. **Copy the Solution Template**

   ```bash
   cp challenge-1/solution-template.go challenge-1/submissions/yourusername/
   ```

6. **Implement the `Sum` Function**

   - Open `challenge-1/submissions/yourusername/solution-template.go` in your text editor.
   - Implement the `Sum` function by replacing the `return 0` with the correct logic.

     ```go
     func Sum(a int, b int) int {
         return a + b
     }
     ```

7. **Test Your Solution Locally**

   - Navigate to the `challenge-1/` directory.

     ```bash
     cd challenge-1
     ```

   - Run the tests:

     ```bash
     go test -v
     ```

   - Ensure all tests pass.

8. **Commit and Push Your Changes**

   ```bash
   git add challenge-1/submissions/yourusername/solution-template.go
   git commit -m "Add Challenge 1 solution"
   git push origin challenge-1-solution
   ```

9. **Create a Pull Request**

   - Go to your forked repository on GitHub.
   - Click on "Compare & pull request".
   - Submit the pull request to the original repository's `main` branch.

10. **Wait for Automated Feedback**

    - The GitHub Actions workflow will run tests on your submission.
    - Check the "Checks" tab in your pull request to see the test results.

---

## **Additional Notes**

- **Input Reading:** The `main` function reads input in the format `a, b`. Make sure to enter the inputs separated by a comma and a space.
- **Handling Errors:** If you're extending the functionality, ensure that your program handles invalid inputs gracefully.
- **Code Style:** Follow standard Go code formatting. You can use `go fmt` to format your code.

---

## **Example Implementation**

Here's how your `solution-template.go` might look after implementing the `Sum` function:

```go
package main

import (
	"fmt"
)

func main() {
	var a, b int
	// Read two integers from standard input
	_, err := fmt.Scanf("%d, %d", &a, &b)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Call the Sum function and print the result
	result := Sum(a, b)
	fmt.Println(result)
}

// Sum returns the sum of a and b.
func Sum(a int, b int) int {
	return a + b
}
```

---

## **Understanding the Test File**

- The test file `solution-template_test.go` automates the testing of your `Sum` function.
- It runs your program with predefined inputs and checks if the output matches the expected result.
- Make sure not to modify the test file unless necessary.

---

## **Troubleshooting**

- **Compilation Errors:**
  - Ensure all your code files are in the correct directory.
  - Check for typos or syntax errors.

- **Test Failures:**
  - Read the error messages carefully.
  - Verify that your `Sum` function handles all cases in the test file.

- **Input/Output Issues:**
  - Ensure that you're reading the input and printing the output exactly as specified.

---

## **Frequently Asked Questions**

**Q:** *Can I add additional functions or import packages?*

**A:** Yes, as long as the main functionality and input/output formats remain as specified.

**Q:** *What should I do if I encounter an error when running the tests?*

**A:** Check the error message for clues. It may be due to incorrect input handling or logic errors in your `Sum` function.

**Q:** *Can I see other participants' submissions?*

**A:** Submissions are located in the `submissions/` directory. Please refrain from copying others' work.

---

## **Contact and Support**

If you have any questions or need help, feel free to open an issue on the repository or contact the maintainers.

---

## **Maintainer Notes (For Repository Owner)**

- **Review Submissions:**
  - Check pull requests for correctness and adherence to guidelines.
  - Merge valid submissions into the `main` branch.

- **Workflow Checks:**
  - Ensure that the GitHub Actions workflows are running properly.
  - Monitor test results and troubleshoot any failures in the workflow.

- **Updating Challenges:**
  - Keep the challenge descriptions and test cases up to date.
  - Consider adding more test cases for comprehensive coverage.

---

## **Next Steps**

- Once you've completed Challenge 1, you can proceed to the next challenge.
- Continue practicing to improve your coding skills and prepare for interviews.
