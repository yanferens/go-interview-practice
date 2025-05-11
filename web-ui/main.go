package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed templates static
var content embed.FS

// Challenge represents a coding challenge
type Challenge struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Template    string `json:"template"`
	TestFile    string `json:"testFile"`
}

// Submission represents a user's submitted solution
type Submission struct {
	Username     string    `json:"username"`
	ChallengeID  int       `json:"challengeId"`
	Code         string    `json:"code"`
	SubmittedAt  time.Time `json:"submittedAt"`
	Passed       bool      `json:"passed"`
	TestOutput   string    `json:"testOutput"`
	ExecutionMs  int64     `json:"executionMs"`
}

// ScoreboardEntry represents an entry in the scoreboard
type ScoreboardEntry struct {
	Username    string    `json:"username"`
	ChallengeID int       `json:"challengeId"`
	SubmittedAt time.Time `json:"submittedAt"`
}

// Global variables
var challenges map[int]*Challenge
var submissions []Submission
var scoreboards map[int][]ScoreboardEntry

// Template functions
var templateFuncs = template.FuncMap{
	"lower": strings.ToLower,
	"truncateDescription": func(s string) string {
		// Extract first paragraph that is not a heading or link
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "!") || strings.HasPrefix(line, "[") {
				continue
			}
			// Found an actual paragraph
			if len(line) > 150 {
				return line[:150] + "..."
			}
			return line
		}
		
		// Fallback to simple truncation
		if len(s) > 150 {
			return s[:150] + "..."
		}
		return s
	},
	"add": func(a, b int) int {
		return a + b
	},
	"extractTitle": func(description string) string {
		// Extract title from markdown content
		lines := strings.Split(description, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "# ") {
				return strings.TrimPrefix(line, "# ")
			}
		}
		return ""
	},
}

func main() {
	// Initialize data
	challenges = make(map[int]*Challenge)
	scoreboards = make(map[int][]ScoreboardEntry)
	
	// Load challenges
	loadChallenges()
	
	// Set up HTTP server
	mux := http.NewServeMux()
	
	// Handle static files
	fsys, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}
	
	staticHandler := http.FileServer(http.FS(fsys))
	mux.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add debug logging
		log.Printf("Serving static file: %s", r.URL.Path)
		
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		}
		staticHandler.ServeHTTP(w, r)
	})))
	
	// API routes
	mux.HandleFunc("/api/challenges", getAllChallenges)
	mux.HandleFunc("/api/challenges/", getChallengeByID)
	mux.HandleFunc("/api/submissions", handleSubmissions)
	mux.HandleFunc("/api/scoreboard/", getScoreboard)
	mux.HandleFunc("/api/run", runCode)
	
	// Web routes
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/challenge/", challengePage)
	mux.HandleFunc("/scoreboard", scoreboardPage)
	mux.HandleFunc("/scoreboard/", scoreChallengeHandler)
	
	// Start server
	port := 8080
	log.Printf("Server starting on http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

// loadChallenges loads all challenges from the filesystem
func loadChallenges() {
	// Find challenge directories (challenge-1, challenge-2, etc.)
	challengeDirs, err := filepath.Glob("../challenge-*")
	if err != nil {
		log.Fatalf("Failed to find challenge directories: %v", err)
	}
	
	for _, dir := range challengeDirs {
		// Extract challenge number
		re := regexp.MustCompile(`challenge-(\d+)`)
		match := re.FindStringSubmatch(dir)
		if len(match) < 2 {
			continue
		}
		
		id, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}
		
		// Read README.md for title and description
		readmePath := filepath.Join(dir, "README.md")
		readmeContent, err := ioutil.ReadFile(readmePath)
		if err != nil {
			log.Printf("Warning: Could not read README for challenge %d: %v", id, err)
			continue
		}
		
		// Extract title from README (first heading)
		titleRe := regexp.MustCompile(`#\s+(.+)`)
		titleMatch := titleRe.FindSubmatch(readmeContent)
		title := ""
		if len(titleMatch) >= 2 {
			title = string(titleMatch[1])
			// Clean up the title - remove "Challenge X: " prefix if present
			cleanTitle := regexp.MustCompile(`^Challenge\s+\d+:\s+`).ReplaceAllString(title, "")
			title = cleanTitle
		} else {
			title = fmt.Sprintf("Challenge %d", id)
		}
		
		// Determine difficulty level based on the challenge number or content
		var difficulty string
		switch {
		case id <= 3 || id == 6:
			difficulty = "Beginner"
		case id == 4 || id == 5 || id == 7:
			difficulty = "Intermediate"
		default:
			difficulty = "Advanced"
		}
		
		// Read solution template
		templatePath := filepath.Join(dir, "solution-template.go")
		templateContent, err := ioutil.ReadFile(templatePath)
		if err != nil {
			log.Printf("Warning: Could not read solution template for challenge %d: %v", id, err)
			continue
		}
		
		// Read test file
		testPath := filepath.Join(dir, "solution-template_test.go")
		testContent, err := ioutil.ReadFile(testPath)
		if err != nil {
			log.Printf("Warning: Could not read test file for challenge %d: %v", id, err)
		}
		
		// Create challenge
		challenge := &Challenge{
			ID:          id,
			Title:       title,
			Description: string(readmeContent),
			Difficulty:  difficulty,
			Template:    string(templateContent),
			TestFile:    string(testContent),
		}
		
		challenges[id] = challenge
		
		// Load scoreboard
		loadScoreboardForChallenge(id, dir)
	}
	
	log.Printf("Loaded %d challenges", len(challenges))
}

// loadScoreboardForChallenge loads the scoreboard for a challenge
func loadScoreboardForChallenge(id int, dir string) {
	scoreboardPath := filepath.Join(dir, "SCOREBOARD.md")
	scoreboardContent, err := ioutil.ReadFile(scoreboardPath)
	if err != nil {
		log.Printf("No scoreboard found for challenge %d", id)
		return
	}
	
	// Parse scoreboard markdown table
	// There are two formats:
	// Format 1: | Username | Passed Tests | Total Tests |
	// Format 2: | Rank | Username | Solution | Date Submitted |
	lines := strings.Split(string(scoreboardContent), "\n")
	entries := []ScoreboardEntry{}
	
	// Determine format by looking at header line
	var format int = 1 // Default to format 1
	headerLine := ""
	for i, line := range lines {
		if i > 0 && strings.Contains(line, "|") {
			headerLine = line
			break
		}
	}
	
	if strings.Contains(headerLine, "Rank") && strings.Contains(headerLine, "Username") {
		format = 2
	}
	
	for i, line := range lines {
		// Skip header and separator lines
		if i < 3 {
			continue
		}
		
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}
		
		// Skip separator lines
		if strings.Contains(line, "---") {
			continue
		}
		
		parts := strings.Split(line, "|")
		if len(parts) < 3 {
			continue
		}
		
		var username string
		
		if format == 1 {
			// Format is: | Username | Passed Tests | Total Tests |
			username = strings.TrimSpace(parts[1])
		} else {
			// Format is: | Rank | Username | Solution | Date Submitted |
			username = strings.TrimSpace(parts[2])
		}
		
		// Skip empty usernames or placeholders
		if username == "" || username == "------" || IsNumeric(username) {
			continue
		}
		
		log.Printf("Parsed username from scoreboard: %s for challenge %d", username, id)
		
		// Use current time for existing entries
		entry := ScoreboardEntry{
			Username:    username,
			ChallengeID: id,
			SubmittedAt: time.Now(),
		}
		
		entries = append(entries, entry)
	}
	
	scoreboards[id] = entries
	log.Printf("Loaded %d scoreboard entries for challenge %d", len(entries), id)
}

// IsNumeric checks if a string contains only digits
func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// --- API Handlers ---

// getAllChallenges returns all challenges
func getAllChallenges(w http.ResponseWriter, r *http.Request) {
	// Convert map to slice for JSON response
	var challengeList []*Challenge
	for _, challenge := range challenges {
		challengeList = append(challengeList, challenge)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(challengeList)
}

// getChallengeByID returns a specific challenge by ID
func getChallengeByID(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`/api/challenges/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)
	
	if len(matches) < 2 {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	challenge, ok := challenges[id]
	if !ok {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(challenge)
}

// handleSubmissions handles submission creation and listing
func handleSubmissions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse the request body
		var submission Submission
		err := json.NewDecoder(r.Body).Decode(&submission)
		if err != nil {
			http.Error(w, "Invalid submission data", http.StatusBadRequest)
			return
		}
		
		submission.SubmittedAt = time.Now()
		
		// Ensure username is set
		if submission.Username == "" {
			submission.Username = "anonymous"
		}
		
		// Run tests on the submission
		results, err := testSubmission(submission)
		if err != nil {
			http.Error(w, "Failed to test submission: "+err.Error(), http.StatusInternalServerError)
			return
		}
		
		submission.Passed = results.Passed
		submission.TestOutput = results.Output
		submission.ExecutionMs = results.ExecutionMs
		
		// Add submission to the list
		submissions = append(submissions, submission)
		
		// If passed, update scoreboard
		if submission.Passed {
			updateScoreboard(submission)
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(submission)
		return
	}
	
	// For GET requests, return all submissions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submissions)
}

// getScoreboard returns the scoreboard for a specific challenge
func getScoreboard(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`/api/scoreboard/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)
	
	if len(matches) < 2 {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	entries, ok := scoreboards[id]
	if !ok {
		// Return empty array if no scoreboard exists
		entries = []ScoreboardEntry{}
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

// TestResult represents the result of running tests on a submission
type TestResult struct {
	Passed      bool   `json:"passed"`
	Output      string `json:"output"`
	ExecutionMs int64  `json:"executionMs"`
}

// runCode runs user code without saving a submission
func runCode(w http.ResponseWriter, r *http.Request) {
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
	
	// Create a temporary submission
	submission := Submission{
		Username:    "temp",
		ChallengeID: request.ChallengeID,
		Code:        request.Code,
	}
	
	// Run tests
	results, err := testSubmission(submission)
	if err != nil {
		http.Error(w, "Failed to run code: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// testSubmission tests a submission and returns the results
func testSubmission(submission Submission) (TestResult, error) {
	result := TestResult{
		Passed: false,
	}
	
	challenge, ok := challenges[submission.ChallengeID]
	if !ok {
		return result, fmt.Errorf("challenge not found")
	}
	
	// Create a unique ID for this submission
	submissionID := fmt.Sprintf("%s-%d-%d", submission.Username, submission.ChallengeID, time.Now().UnixNano())
	
	// Create a temporary directory for testing
	tempDir, err := ioutil.TempDir("", fmt.Sprintf("challenge-%d-", submission.ChallengeID))
	if err != nil {
		return result, fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Write the submission code to files
	solutionFile := filepath.Join(tempDir, "solution-template.go")
	if err := ioutil.WriteFile(solutionFile, []byte(submission.Code), 0644); err != nil {
		return result, fmt.Errorf("failed to write solution file: %v", err)
	}
	
	// Write the test file
	testFile := filepath.Join(tempDir, "solution_test.go")
	if err := ioutil.WriteFile(testFile, []byte(challenge.TestFile), 0644); err != nil {
		return result, fmt.Errorf("failed to write test file: %v", err)
	}
	
	// Initialize a Go module in the temp directory
	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("challenge%d_%s", submission.ChallengeID, submissionID))
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		return result, fmt.Errorf("failed to initialize Go module: %v", err)
	}
	
	// Run tests
	startTime := time.Now()
	cmd = exec.Command("go", "test", "-v")
	cmd.Dir = tempDir
	
	output, err := cmd.CombinedOutput()
	result.ExecutionMs = time.Since(startTime).Milliseconds()
	result.Output = string(output)
	
	if err == nil {
		result.Passed = true
	} else {
		// Check if tests ran but failed
		if exitErr, ok := err.(*exec.ExitError); ok {
			// Test ran but failed
			result.Passed = exitErr.ExitCode() == 0
		} else {
			// Command couldn't be run
			return result, fmt.Errorf("failed to run tests: %v", err)
		}
	}
	
	return result, nil
}

// updateScoreboard adds a new entry to the scoreboard for a passed submission
func updateScoreboard(submission Submission) {
	// Ensure username isn't empty, use "anonymous" as fallback
	username := submission.Username
	if username == "" {
		username = "anonymous"
	}
	
	entry := ScoreboardEntry{
		Username:    username,
		ChallengeID: submission.ChallengeID,
		SubmittedAt: submission.SubmittedAt,
	}
	
	// Log the entry details for debugging
	log.Printf("Adding scoreboard entry: Challenge %d, Username: %s, Time: %v", 
		entry.ChallengeID, entry.Username, entry.SubmittedAt)
	
	// Add to memory scoreboard
	scoreboards[submission.ChallengeID] = append(scoreboards[submission.ChallengeID], entry)
}

// --- Web Handlers ---

// homePage renders the home page with a list of challenges
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	tmpl, err := template.New("").Funcs(templateFuncs).ParseFS(content, "templates/base.html", "templates/home.html")
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Convert map to slice for template
	var challengeList []*Challenge
	for _, challenge := range challenges {
		challengeList = append(challengeList, challenge)
	}
	
	data := struct {
		Challenges []*Challenge
	}{
		Challenges: challengeList,
	}
	
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
	}
}

// challengePage renders a specific challenge page
func challengePage(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`/challenge/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)
	
	if len(matches) < 2 {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	challenge, ok := challenges[id]
	if !ok {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}
	
	tmpl, err := template.New("").Funcs(templateFuncs).ParseFS(content, "templates/base.html", "templates/challenge.html")
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := struct {
		Challenge *Challenge
	}{
		Challenge: challenge,
	}
	
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
	}
}

// scoreboardPage renders the overall scoreboard page
func scoreboardPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").Funcs(templateFuncs).ParseFS(content, "templates/base.html", "templates/scoreboard.html")
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := struct {
		Challenges  map[int]*Challenge
		Scoreboards map[int][]ScoreboardEntry
	}{
		Challenges:  challenges,
		Scoreboards: scoreboards,
	}
	
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
	}
}

// scoreChallengeHandler handles displaying the scoreboard for a specific challenge
func scoreChallengeHandler(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`/scoreboard/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)
	
	if len(matches) < 2 {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}
	
	challenge, ok := challenges[id]
	if !ok {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}
	
	entries, ok := scoreboards[id]
	if !ok {
		// Empty array if no scoreboard exists
		entries = []ScoreboardEntry{}
	}
	
	// Use the dedicated challenge scoreboard template
	tmpl, err := template.New("").Funcs(templateFuncs).ParseFS(content, "templates/base.html", "templates/challenge_scoreboard.html")
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := struct {
		Challenge *Challenge
		Entries   []ScoreboardEntry
	}{
		Challenge: challenge,
		Entries:   entries,
	}
	
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
	}
} 