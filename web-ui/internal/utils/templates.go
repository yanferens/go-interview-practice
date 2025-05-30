package utils

import (
	"html/template"
	"strings"
)

// GetTemplateFuncs returns the template functions used across the application
func GetTemplateFuncs() template.FuncMap {
	return template.FuncMap{
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
		"js": func(s string) template.JS {
			// Safely escape backticks and other special characters for JavaScript
			// Replace backticks with HTML entity
			s = strings.Replace(s, "`", "\\`", -1)
			// Replace dollar signs that might interfere with template literals
			s = strings.Replace(s, "${", "\\${", -1)
			return template.JS(s)
		},
	}
}
