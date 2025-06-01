package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"web-ui/internal/models"
	"web-ui/internal/services"
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
