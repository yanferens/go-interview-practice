package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"web-ui/internal/models"
	"web-ui/internal/services"
	"web-ui/internal/utils"
)

// APIHandler handles all API endpoints
type APIHandler struct {
	challengeService  *services.ChallengeService
	scoreboardService *services.ScoreboardService
	userService       *services.UserService
	executionService  *services.ExecutionService
	submissions       []models.Submission
}

// NewAPIHandler creates a new API handler
func NewAPIHandler(
	challengeService *services.ChallengeService,
	scoreboardService *services.ScoreboardService,
	userService *services.UserService,
	executionService *services.ExecutionService,
) *APIHandler {
	return &APIHandler{
		challengeService:  challengeService,
		scoreboardService: scoreboardService,
		userService:       userService,
		executionService:  executionService,
		submissions:       make([]models.Submission, 0),
	}
}

// GetAllChallenges returns all challenges
func (h *APIHandler) GetAllChallenges(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	challenges := h.challengeService.GetChallenges()

	// Convert map to slice for JSON response
	var challengeList []*models.Challenge
	for _, challenge := range challenges {
		challengeList = append(challengeList, challenge)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(challengeList)
}

// GetChallengeByID returns a specific challenge by ID
func (h *APIHandler) GetChallengeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract challenge ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/api/challenges/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}

	challenge, exists := h.challengeService.GetChallenge(id)
	if !exists {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(challenge)
}

// HandleSubmissions handles submission operations
func (h *APIHandler) HandleSubmissions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createSubmission(w, r)
	case "GET":
		h.getSubmissions(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// createSubmission creates a new submission
func (h *APIHandler) createSubmission(w http.ResponseWriter, r *http.Request) {
	var submission models.Submission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Set submission timestamp
	submission.SubmittedAt = time.Now()

	// Validate challenge exists
	challenge, exists := h.challengeService.GetChallenge(submission.ChallengeID)
	if !exists {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}

	// Run the code
	result := h.executionService.RunCode(submission.Code, challenge)
	submission.Passed = result.Passed
	submission.TestOutput = result.Output
	submission.ExecutionMs = result.ExecutionMs

	// Store submission
	h.submissions = append(h.submissions, submission)

	// Add to scoreboard if passed
	if submission.Passed {
		h.scoreboardService.AddSubmission(submission)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submission)
}

// getSubmissions returns all submissions
func (h *APIHandler) getSubmissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.submissions)
}

// GetScoreboard returns the scoreboard for a challenge
func (h *APIHandler) GetScoreboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract challenge ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/api/scoreboard/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}

	scoreboard, exists := h.scoreboardService.GetScoreboard(id)
	if !exists {
		scoreboard = []models.ScoreboardEntry{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scoreboard)
}

// RunCode executes submitted code
func (h *APIHandler) RunCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		ChallengeID int    `json:"challengeId"`
		Code        string `json:"code"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	challenge, exists := h.challengeService.GetChallenge(request.ChallengeID)
	if !exists {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}

	result := h.executionService.RunCode(request.Code, challenge)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// SaveSubmissionToFilesystem saves a submission to the filesystem
func (h *APIHandler) SaveSubmissionToFilesystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request services.SaveSubmissionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if request.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Set username cookie
	h.setUsernameCookie(w, request.Username)

	// Validate challenge exists
	_, exists := h.challengeService.GetChallenge(request.ChallengeID)
	if !exists {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}

	response := h.executionService.SaveSubmissionToFilesystem(request)

	// Clear user attempts cache
	h.userService.RefreshUserAttempts(request.Username, h.challengeService.GetChallenges())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RefreshUserAttempts refreshes user's attempt cache
func (h *APIHandler) RefreshUserAttempts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if request.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	attempts := h.userService.RefreshUserAttempts(request.Username, h.challengeService.GetChallenges())

	response := struct {
		Username     string       `json:"username"`
		AttemptedIDs map[int]bool `json:"attemptedIds"`
		Scores       map[int]int  `json:"scores"`
		Success      bool         `json:"success"`
	}{
		Username:     request.Username,
		AttemptedIDs: attempts.AttemptedIDs,
		Scores:       attempts.Scores,
		Success:      true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetGitUsername returns the username extracted from git configuration
func (h *APIHandler) GetGitUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	gitInfo := utils.GetGitUsername()

	response := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Source   string `json:"source"`
		Success  bool   `json:"success"`
	}{
		Username: gitInfo.Username,
		Email:    gitInfo.Email,
		Source:   gitInfo.Source,
		Success:  gitInfo.Username != "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// setUsernameCookie sets the username cookie
func (h *APIHandler) setUsernameCookie(w http.ResponseWriter, username string) {
	// Cookie expires in 30 days
	expiration := time.Now().Add(30 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     "username",
		Value:    username,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: false, // Allow JavaScript to access it
	}
	http.SetCookie(w, &cookie)
}

// GetMainScoreboardRank returns the user's rank in the main scoreboard
func (h *APIHandler) GetMainScoreboardRank(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username parameter required", http.StatusBadRequest)
		return
	}

	// Calculate user's rank in main scoreboard
	rank := h.calculateMainScoreboardRank(username)

	response := struct {
		Username string `json:"username"`
		Rank     int    `json:"rank"`
		Success  bool   `json:"success"`
	}{
		Username: username,
		Rank:     rank,
		Success:  true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// calculateMainScoreboardRank calculates the user's rank based on completed challenges
func (h *APIHandler) calculateMainScoreboardRank(username string) int {
	// Get all users and their completion counts (only count if ALL tests passed)
	challenges := h.challengeService.GetChallenges()
	userCompletions := make(map[string]int)

	// Process all challenge scoreboards to count actual completions
	for challengeID := range challenges {
		// Read scoreboard file directly to check test results
		scoreboardPath := fmt.Sprintf("../challenge-%d/SCOREBOARD.md", challengeID)
		content, err := ioutil.ReadFile(scoreboardPath)
		if err != nil {
			// Try alternative path
			scoreboardPath = fmt.Sprintf("challenge-%d/SCOREBOARD.md", challengeID)
			content, err = ioutil.ReadFile(scoreboardPath)
			if err != nil {
				continue
			}
		}

		// Parse scoreboard to find users who passed ALL tests
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			// Skip header and separator lines
			if !strings.Contains(line, "|") || strings.Contains(line, "Username") || strings.Contains(line, "---") {
				continue
			}

			parts := strings.Split(line, "|")
			if len(parts) < 4 {
				continue
			}

			username := strings.TrimSpace(parts[1])
			passedTestsStr := strings.TrimSpace(parts[2])
			totalTestsStr := strings.TrimSpace(parts[3])

			// Skip empty usernames or placeholders
			if username == "" || username == "------" {
				continue
			}

			// Parse test counts
			passedTests, err1 := strconv.Atoi(passedTestsStr)
			totalTests, err2 := strconv.Atoi(totalTestsStr)

			// Only count as completed if ALL tests passed
			if err1 == nil && err2 == nil && passedTests > 0 && passedTests == totalTests {
				userCompletions[username]++
			}
		}
	}

	// Get the target user's completion count
	targetCompletions := userCompletions[username]
	if targetCompletions == 0 {
		return 0 // User is unranked
	}

	// Count how many users have more completions (following Python script logic)
	rank := 1
	for user, completions := range userCompletions {
		if user != username && completions > targetCompletions {
			rank++
		}
	}

	return rank
}

// GetMainLeaderboard returns the main leaderboard data
func (h *APIHandler) GetMainLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Calculate leaderboard data
	leaderboard := h.calculateMainLeaderboard()

	response := struct {
		Leaderboard []LeaderboardUser `json:"leaderboard"`
		Success     bool              `json:"success"`
	}{
		Leaderboard: leaderboard,
		Success:     true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LeaderboardUser represents a user in the leaderboard
type LeaderboardUser struct {
	Username            string       `json:"username"`
	CompletedCount      int          `json:"completedCount"`
	CompletionRate      float64      `json:"completionRate"`
	CompletedChallenges map[int]bool `json:"completedChallenges"`
	Achievement         string       `json:"achievement"`
	Rank                int          `json:"rank"`
}

// calculateMainLeaderboard calculates the main leaderboard data
func (h *APIHandler) calculateMainLeaderboard() []LeaderboardUser {
	challenges := h.challengeService.GetChallenges()
	totalChallenges := len(challenges)
	userCompletions := make(map[string]map[int]bool)

	// Process all challenge scoreboards to find completions
	for challengeID := range challenges {
		// Read scoreboard file directly to check test results
		scoreboardPath := fmt.Sprintf("../challenge-%d/SCOREBOARD.md", challengeID)
		content, err := ioutil.ReadFile(scoreboardPath)
		if err != nil {
			// Try alternative path
			scoreboardPath = fmt.Sprintf("challenge-%d/SCOREBOARD.md", challengeID)
			content, err = ioutil.ReadFile(scoreboardPath)
			if err != nil {
				continue
			}
		}

		// Parse scoreboard to find users who passed ALL tests
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			// Skip header and separator lines
			if !strings.Contains(line, "|") || strings.Contains(line, "Username") || strings.Contains(line, "---") {
				continue
			}

			parts := strings.Split(line, "|")
			if len(parts) < 4 {
				continue
			}

			username := strings.TrimSpace(parts[1])
			passedTestsStr := strings.TrimSpace(parts[2])
			totalTestsStr := strings.TrimSpace(parts[3])

			// Skip empty usernames or placeholders
			if username == "" || username == "------" {
				continue
			}

			// Parse test counts
			passedTests, err1 := strconv.Atoi(passedTestsStr)
			totalTests, err2 := strconv.Atoi(totalTestsStr)

			// Only count as completed if ALL tests passed
			if err1 == nil && err2 == nil && passedTests > 0 && passedTests == totalTests {
				if userCompletions[username] == nil {
					userCompletions[username] = make(map[int]bool)
				}
				userCompletions[username][challengeID] = true
			}
		}
	}

	// Convert to leaderboard format
	var leaderboard []LeaderboardUser
	for username, completions := range userCompletions {
		completedCount := len(completions)
		completionRate := float64(completedCount) / float64(totalChallenges) * 100

		// Determine achievement
		achievement := "🌱 Beginner"
		if completedCount >= 20 {
			achievement = "🔥 Master"
		} else if completedCount >= 15 {
			achievement = "⭐ Expert"
		} else if completedCount >= 10 {
			achievement = "💪 Advanced"
		} else if completedCount >= 5 {
			achievement = "🚀 Intermediate"
		}

		leaderboard = append(leaderboard, LeaderboardUser{
			Username:            username,
			CompletedCount:      completedCount,
			CompletionRate:      completionRate,
			CompletedChallenges: completions,
			Achievement:         achievement,
		})
	}

	// Sort by completion count (descending), then by username
	for i := 0; i < len(leaderboard); i++ {
		for j := i + 1; j < len(leaderboard); j++ {
			if leaderboard[j].CompletedCount > leaderboard[i].CompletedCount ||
				(leaderboard[j].CompletedCount == leaderboard[i].CompletedCount && leaderboard[j].Username < leaderboard[i].Username) {
				leaderboard[i], leaderboard[j] = leaderboard[j], leaderboard[i]
			}
		}
	}

	// Assign ranks
	for i := range leaderboard {
		leaderboard[i].Rank = i + 1
	}

	return leaderboard
}
