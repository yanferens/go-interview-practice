# Go Interview Practice Web UI

This is a web-based user interface for the Go Interview Practice project, providing an interactive environment for solving Go programming challenges.

## Features

- **Challenge Browser**: View all available coding challenges with difficulty indicators.
- **In-browser Code Editor**: Edit and run Go code directly in your browser with syntax highlighting.
- **Test Runner**: Run tests against your solution and see results in real-time.
- **Scoreboard**: Track your progress and see how you compare to others.
- **Markdown Support**: Challenge descriptions rendered with full Markdown support.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Web browser (Chrome, Firefox, Safari, Edge)

### Running the Web UI

1. Navigate to the web-ui directory:
   ```
   cd web-ui
   ```

2. Run the web server:
   ```
   go run main.go
   ```

3. Open your browser and visit:
   ```
   http://localhost:8080
   ```

## Project Structure

```
web-ui/
├── main.go                  # Main server entry point
├── static/                  # Static assets
│   ├── css/                 # CSS stylesheets
│   │   └── style.css        # Custom CSS for the UI
│   └── js/                  # JavaScript files
│       └── main.js          # Common JavaScript utilities
├── templates/               # HTML templates
│   ├── base.html            # Base template with common layout
│   ├── challenge.html       # Challenge page with code editor
│   ├── home.html            # Home page with challenge list
│   └── scoreboard.html      # Scoreboard page for results
└── README.md                # This file
```

## Technical Details

### Templates and HTML Rendering

The web UI uses Go's `html/template` package for server-side rendering, with a base template that defines the common layout and individual content templates for each page type.

### JavaScript Libraries

- **Bootstrap**: For responsive UI components
- **Ace Editor**: For the in-browser code editor
- **Marked**: For Markdown parsing
- **Highlight.js**: For syntax highlighting

### API Endpoints

The web UI exposes the following API endpoints:

- `GET /api/challenges`: Get all challenges
- `GET /api/challenges/{id}`: Get a specific challenge
- `POST /api/run`: Run code for a specific challenge
- `POST /api/submissions`: Submit a solution
- `GET /api/scoreboard/{id}`: Get scoreboard for a challenge

## Development

### Adding New Features

1. If adding new pages, create a new template in the `templates` directory.
2. Add any new API handlers in `main.go`.
3. Add CSS styles to `static/css/style.css`.
4. Add JavaScript utilities to `static/js/main.js`.

### Running in Development Mode

To enable hot-reloading during development, you can use tools like [Air](https://github.com/cosmtrek/air):

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with hot reloading
air
```

## Contributing

Contributions to improve the web UI are welcome! Please feel free to submit pull requests or open issues for new features or bug fixes.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 