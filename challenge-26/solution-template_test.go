package regex

import (
	"reflect"
	"testing"
)

func TestExtractEmails(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Basic email extraction",
			input:    "Contact us at support@example.com or sales@company.co.uk for more info.",
			expected: []string{"support@example.com", "sales@company.co.uk"},
		},
		{
			name:     "No emails in text",
			input:    "There are no email addresses in this text.",
			expected: []string{},
		},
		{
			name:     "Mixed valid and invalid emails",
			input:    "Valid: user@domain.com, invalid: user@, also invalid: @domain.com, another valid: name.surname+tag@domain-name.co",
			expected: []string{"user@domain.com", "name.surname+tag@domain-name.co"},
		},
		{
			name:     "Emails with subdomains",
			input:    "Contact admin@server.department.company.com or user@subdomain.example.org",
			expected: []string{"admin@server.department.company.com", "user@subdomain.example.org"},
		},
		{
			name:     "Email within HTML-like text",
			input:    "<p>Please contact <a href='mailto:info@example.com'>info@example.com</a> for support.</p>",
			expected: []string{"info@example.com", "info@example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractEmails(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ExtractEmails() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid phone number with parentheses",
			input:    "(555) 123-4567",
			expected: true,
		},
		{
			name:     "Invalid format - missing parentheses",
			input:    "555 123-4567",
			expected: false,
		},
		{
			name:     "Invalid format - hyphen instead of parentheses",
			input:    "555-123-4567",
			expected: false,
		},
		{
			name:     "Invalid format - extra digits",
			input:    "(555) 123-45678",
			expected: false,
		},
		{
			name:     "Invalid format - letters",
			input:    "(555) ABC-DEFG",
			expected: false,
		},
		{
			name:     "Invalid format - missing space",
			input:    "(555)123-4567",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePhone(tt.input)
			if got != tt.expected {
				t.Errorf("ValidatePhone() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMaskCreditCard(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Card with hyphens",
			input:    "1234-5678-9012-3456",
			expected: "XXXX-XXXX-XXXX-3456",
		},
		{
			name:     "Card without hyphens",
			input:    "1234567890123456",
			expected: "XXXXXXXXXXXX3456",
		},
		{
			name:     "Already partially masked card",
			input:    "XXXX-XXXX-XXXX-3456",
			expected: "XXXX-XXXX-XXXX-3456",
		},
		{
			name:     "Short card number",
			input:    "1234-5678",
			expected: "XXXX-5678",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaskCreditCard(tt.input)
			if got != tt.expected {
				t.Errorf("MaskCreditCard() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParseLogEntry(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:  "Valid INFO log entry",
			input: "2023-11-15 14:23:45 INFO Server started on port 8080",
			expected: map[string]string{
				"date":    "2023-11-15",
				"time":    "14:23:45",
				"level":   "INFO",
				"message": "Server started on port 8080",
			},
		},
		{
			name:  "Valid ERROR log entry",
			input: "2023-11-15 14:24:12 ERROR Failed to connect to database: timeout",
			expected: map[string]string{
				"date":    "2023-11-15",
				"time":    "14:24:12",
				"level":   "ERROR",
				"message": "Failed to connect to database: timeout",
			},
		},
		{
			name:  "Valid WARNING log entry",
			input: "2023-11-15 14:25:01 WARNING High memory usage: 85%",
			expected: map[string]string{
				"date":    "2023-11-15",
				"time":    "14:25:01",
				"level":   "WARNING",
				"message": "High memory usage: 85%",
			},
		},
		{
			name:     "Invalid log format",
			input:    "Invalid log format",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseLogEntry(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseLogEntry() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestExtractURLs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Basic URL extraction",
			input:    "Visit https://golang.org and http://example.com/page?q=123 for more information.",
			expected: []string{"https://golang.org", "http://example.com/page?q=123"},
		},
		{
			name:     "No URLs in text",
			input:    "There are no URLs in this text.",
			expected: []string{},
		},
		{
			name:     "URLs with various components",
			input:    "Check out https://user:pass@example.com:8080/path/to/page?query=value#section and http://localhost:3000",
			expected: []string{"https://user:pass@example.com:8080/path/to/page?query=value#section", "http://localhost:3000"},
		},
		{
			name:     "URLs in brackets",
			input:    "URLs can be in brackets (https://example.org) or [http://domain.com]",
			expected: []string{"https://example.org", "http://domain.com"},
		},
		{
			name:     "URLs with various TLDs",
			input:    "Visit https://example.com, https://example.org, https://example.net, and https://example.co.uk",
			expected: []string{"https://example.com", "https://example.org", "https://example.net", "https://example.co.uk"},
		},
		{
			name:     "URLs in HTML-like text",
			input:    "<a href='https://example.com'>Link</a> and <img src='http://example.org/image.jpg'>",
			expected: []string{"https://example.com", "http://example.org/image.jpg"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractURLs(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ExtractURLs() = %v, want %v", got, tt.expected)
			}
		})
	}
}
