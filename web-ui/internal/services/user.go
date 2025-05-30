package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"web-ui/internal/models"
)

// UserService handles user-related operations
type UserService struct {
	userAttempts models.UserAttemptsMap
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		userAttempts: make(models.UserAttemptsMap),
	}
}

// LoadUserAttempts checks the filesystem for submission directories
func (us *UserService) LoadUserAttempts(username string, challenges models.ChallengeMap) *models.UserAttemptedChallenges {
	// If we already loaded this user's attempts, return from cache
	if attempts, ok := us.userAttempts[username]; ok {
		return attempts
	}

	// Create new tracking structure
	userAttempt := &models.UserAttemptedChallenges{
		Username:     username,
		AttemptedIDs: make(map[int]bool),
	}

	// Scan all challenge directories for this user's submissions
	for id := range challenges {
		if us.hasUserSubmission(username, id) {
			userAttempt.AttemptedIDs[id] = true
		}
	}

	// Cache the results
	us.userAttempts[username] = userAttempt
	return userAttempt
}

// hasUserSubmission checks if a user has a submission for a challenge
func (us *UserService) hasUserSubmission(username string, challengeID int) bool {
	// Try different path formats to handle potential path issues
	// Absolute path
	submissionDir := filepath.Join("..", fmt.Sprintf("challenge-%d", challengeID), "submissions", username)
	submissionFile := filepath.Join(submissionDir, "solution-template.go")

	// Check if the file exists
	if _, err := os.Stat(submissionFile); err == nil {
		return true
	}

	// Alternative path (direct from workspace root)
	altSubmissionFile := filepath.Join(fmt.Sprintf("challenge-%d", challengeID), "submissions", username, "solution-template.go")

	if _, err := os.Stat(altSubmissionFile); err == nil {
		return true
	}

	return false
}

// GetExistingSolution returns the content of an existing solution file if it exists
func (us *UserService) GetExistingSolution(username string, challengeID int) string {
	if username == "" {
		return ""
	}

	// Try different path formats
	// First try the relative path from web-ui
	submissionFile := filepath.Join("..", fmt.Sprintf("challenge-%d", challengeID), "submissions", username, "solution-template.go")
	content, err := ioutil.ReadFile(submissionFile)
	if err == nil {
		return string(content)
	}

	// Try alternative path from root directory
	altSubmissionFile := filepath.Join(fmt.Sprintf("challenge-%d", challengeID), "submissions", username, "solution-template.go")
	content, err = ioutil.ReadFile(altSubmissionFile)
	if err == nil {
		return string(content)
	}

	return ""
}

// RefreshUserAttempts clears the cache for a user and reloads their attempts
func (us *UserService) RefreshUserAttempts(username string, challenges models.ChallengeMap) *models.UserAttemptedChallenges {
	// Clear cache
	delete(us.userAttempts, username)
	// Reload and return
	return us.LoadUserAttempts(username, challenges)
}

// GetUserAttempts returns the cached user attempts or loads them if not cached
func (us *UserService) GetUserAttempts(username string, challenges models.ChallengeMap) *models.UserAttemptedChallenges {
	if attempts, ok := us.userAttempts[username]; ok {
		return attempts
	}
	return us.LoadUserAttempts(username, challenges)
}
