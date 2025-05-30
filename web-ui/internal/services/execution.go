package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"web-ui/internal/models"
)

// ExecutionService handles code execution and testing
type ExecutionService struct{}

// NewExecutionService creates a new execution service
func NewExecutionService() *ExecutionService {
	return &ExecutionService{}
}

// ExecutionResult represents the result of code execution
type ExecutionResult struct {
	Passed      bool   `json:"passed"`
	Output      string `json:"output"`
	ExecutionMs int64  `json:"executionMs"`
}

// RunCode executes the provided code against a challenge's tests
func (es *ExecutionService) RunCode(code string, challenge *models.Challenge) ExecutionResult {
	start := time.Now()

	// Create temporary directory for execution
	tempDir, err := ioutil.TempDir("", "challenge-exec")
	if err != nil {
		return ExecutionResult{
			Passed: false,
			Output: fmt.Sprintf("Failed to create temporary directory: %v", err),
		}
	}
	defer os.RemoveAll(tempDir)

	// Write the submitted code to temporary file
	codePath := filepath.Join(tempDir, "solution-template.go")
	err = ioutil.WriteFile(codePath, []byte(code), 0644)
	if err != nil {
		return ExecutionResult{
			Passed: false,
			Output: fmt.Sprintf("Failed to write code file: %v", err),
		}
	}

	// Write the test file to temporary directory
	testPath := filepath.Join(tempDir, "solution_test.go")
	err = ioutil.WriteFile(testPath, []byte(challenge.TestFile), 0644)
	if err != nil {
		return ExecutionResult{
			Passed: false,
			Output: fmt.Sprintf("Failed to write test file: %v", err),
		}
	}

	// Initialize Go module
	err = es.initGoModule(tempDir, challenge.ID)
	if err != nil {
		return ExecutionResult{
			Passed: false,
			Output: fmt.Sprintf("Failed to initialize Go module: %v", err),
		}
	}

	// Install dependencies for challenge 14 (gRPC)
	if challenge.ID == 14 {
		es.installGRPCDependencies(tempDir)
	}

	// Run tests
	cmd := exec.Command("go", "test", "-v")
	cmd.Dir = tempDir

	output, err := cmd.CombinedOutput()
	executionTime := time.Since(start).Milliseconds()
	outputStr := string(output)

	result := ExecutionResult{
		Output:      outputStr,
		ExecutionMs: executionTime,
	}

	if err == nil {
		result.Passed = true
	} else {
		// Check if tests ran but failed (this is the key logic!)
		if _, ok := err.(*exec.ExitError); ok {
			// Test ran but failed - this means tests executed but some failed
			result.Passed = false // Tests failed, so Passed = false
		} else {
			// Command couldn't be run - this is a real error
			result.Passed = false
			result.Output = fmt.Sprintf("Failed to run tests: %v\n%s", err, outputStr)
		}
	}

	return result
}

// initGoModule initializes a Go module in the temporary directory
func (es *ExecutionService) initGoModule(tempDir string, challengeID int) error {
	// Initialize go.mod
	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("challenge-%d", challengeID))
	cmd.Dir = tempDir
	return cmd.Run()
}

// installGRPCDependencies installs gRPC dependencies for challenge 14
func (es *ExecutionService) installGRPCDependencies(tempDir string) {
	// Install gRPC packages
	grpcCmd := exec.Command("go", "get", "google.golang.org/grpc")
	grpcCmd.Dir = tempDir
	grpcCmd.Run() // Ignore errors, just try

	codesCmd := exec.Command("go", "get", "google.golang.org/grpc/codes")
	codesCmd.Dir = tempDir
	codesCmd.Run() // Ignore errors, just try

	statusCmd := exec.Command("go", "get", "google.golang.org/grpc/status")
	statusCmd.Dir = tempDir
	statusCmd.Run() // Ignore errors, just try
}

// SaveSubmissionRequest represents a request to save a submission to filesystem
type SaveSubmissionRequest struct {
	Username    string `json:"username"`
	ChallengeID int    `json:"challengeId"`
	Code        string `json:"code"`
}

// SaveSubmissionResponse represents the response from saving a submission
type SaveSubmissionResponse struct {
	Success     bool     `json:"success"`
	Message     string   `json:"message"`
	FilePath    string   `json:"filePath"`
	GitCommands []string `json:"gitCommands"`
}

// SaveSubmissionToFilesystem saves a user's submission to the filesystem
func (es *ExecutionService) SaveSubmissionToFilesystem(request SaveSubmissionRequest) SaveSubmissionResponse {
	// Get working directory for correct relative paths
	workDir, _ := os.Getwd()

	// Try different path approaches to handle potential path issues
	var submissionDir string
	var fileSaved bool

	// Try multiple path options to ensure it works in different environments
	pathOptions := []string{
		// Option 1: From web-ui directory (standard case)
		filepath.Join("..", fmt.Sprintf("challenge-%d", request.ChallengeID), "submissions", request.Username),
		// Option 2: From root workspace
		filepath.Join(fmt.Sprintf("challenge-%d", request.ChallengeID), "submissions", request.Username),
		// Option 3: Absolute path from detected workspace root
		filepath.Join(workDir, "..", fmt.Sprintf("challenge-%d", request.ChallengeID), "submissions", request.Username),
	}

	for _, dirPath := range pathOptions {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			continue
		}

		solutionFile := filepath.Join(dirPath, "solution-template.go")
		err = ioutil.WriteFile(solutionFile, []byte(request.Code), 0644)
		if err != nil {
			continue
		}

		submissionDir = dirPath
		fileSaved = true
		break
	}

	if !fileSaved {
		return SaveSubmissionResponse{
			Success: false,
			Message: "Failed to save solution to any available path",
		}
	}

	// Return success response with git commands
	return SaveSubmissionResponse{
		Success:  true,
		Message:  "Solution saved to filesystem",
		FilePath: filepath.Join(submissionDir, "solution-template.go"),
		GitCommands: []string{
			"cd " + filepath.Join(workDir, ".."),
			fmt.Sprintf("git add %s", filepath.Join(fmt.Sprintf("challenge-%d", request.ChallengeID), "submissions", request.Username, "solution-template.go")),
			fmt.Sprintf("git commit -m \"Add solution for Challenge %d\"", request.ChallengeID),
			"git push origin main",
		},
	}
}
