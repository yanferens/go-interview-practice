package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"web-ui/internal/handlers"
	"web-ui/internal/services"
)

// Server represents the web server with all its dependencies
type Server struct {
	content           embed.FS
	challengeService  *services.ChallengeService
	scoreboardService *services.ScoreboardService
	userService       *services.UserService
	executionService  *services.ExecutionService
}

// NewServer creates a new server instance
func NewServer(
	content embed.FS,
	challengeService *services.ChallengeService,
	scoreboardService *services.ScoreboardService,
	userService *services.UserService,
	executionService *services.ExecutionService,
) *Server {
	return &Server{
		content:           content,
		challengeService:  challengeService,
		scoreboardService: scoreboardService,
		userService:       userService,
		executionService:  executionService,
	}
}

// SetupRoutes configures all HTTP routes
func (s *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Setup static file handling
	s.setupStaticFiles(mux)

	// Initialize handlers
	apiHandler := handlers.NewAPIHandler(
		s.challengeService,
		s.scoreboardService,
		s.userService,
		s.executionService,
	)

	webHandler := handlers.NewWebHandler(
		s.content,
		s.challengeService,
		s.scoreboardService,
		s.userService,
	)

	// API routes
	mux.HandleFunc("/api/challenges", apiHandler.GetAllChallenges)
	mux.HandleFunc("/api/challenges/", apiHandler.GetChallengeByID)
	mux.HandleFunc("/api/submissions", apiHandler.HandleSubmissions)
	mux.HandleFunc("/api/scoreboard/", apiHandler.GetScoreboard)
	mux.HandleFunc("/api/run", apiHandler.RunCode)
	mux.HandleFunc("/api/save-to-filesystem", apiHandler.SaveSubmissionToFilesystem)
	mux.HandleFunc("/api/refresh-attempts", apiHandler.RefreshUserAttempts)
	mux.HandleFunc("/api/git-username", apiHandler.GetGitUsername)
	mux.HandleFunc("/api/main-scoreboard-rank", apiHandler.GetMainScoreboardRank)
	mux.HandleFunc("/api/main-leaderboard", apiHandler.GetMainLeaderboard)

	// Web routes
	mux.HandleFunc("/", webHandler.HomePage)
	mux.HandleFunc("/challenge/", webHandler.ChallengePage)
	mux.HandleFunc("/scoreboard", webHandler.ScoreboardPage)
	mux.HandleFunc("/scoreboard/", webHandler.ScoreChallengeHandler)

	return mux
}

// setupStaticFiles configures static file serving
func (s *Server) setupStaticFiles(mux *http.ServeMux) {
	fsys, err := fs.Sub(s.content, "static")
	if err != nil {
		log.Fatal(err)
	}

	staticHandler := http.FileServer(http.FS(fsys))
	mux.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set appropriate content type headers
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		}
		staticHandler.ServeHTTP(w, r)
	})))
}
