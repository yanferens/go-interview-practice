package regex

// ExtractEmails extracts all valid email addresses from a text
func ExtractEmails(text string) []string {
	// TODO: Implement this function
	// 1. Create a regular expression to match email addresses
	// 2. Find all matches in the input text
	// 3. Return the matched emails as a slice of strings

	return nil
}

// ValidatePhone checks if a string is a valid phone number in format (XXX) XXX-XXXX
func ValidatePhone(phone string) bool {
	// TODO: Implement this function
	// 1. Create a regular expression to match the specified phone format
	// 2. Check if the input string matches the pattern
	// 3. Return true if it's a match, false otherwise

	return false
}

// MaskCreditCard replaces all but the last 4 digits of a credit card number with "X"
// Example: "1234-5678-9012-3456" -> "XXXX-XXXX-XXXX-3456"
func MaskCreditCard(cardNumber string) string {
	// TODO: Implement this function
	// 1. Create a regular expression to identify the parts of the card number to mask
	// 2. Use ReplaceAllString or similar method to perform the replacement
	// 3. Return the masked card number

	return ""
}

// ParseLogEntry parses a log entry with format:
// "YYYY-MM-DD HH:MM:SS LEVEL Message"
// Returns a map with keys: "date", "time", "level", "message"
func ParseLogEntry(logLine string) map[string]string {
	// TODO: Implement this function
	// 1. Create a regular expression with capture groups for each component
	// 2. Use FindStringSubmatch to extract the components
	// 3. Populate a map with the extracted values
	// 4. Return the populated map

	return nil
}

// ExtractURLs extracts all valid URLs from a text
func ExtractURLs(text string) []string {
	// TODO: Implement this function
	// 1. Create a regular expression to match URLs (both http and https)
	// 2. Find all matches in the input text
	// 3. Return the matched URLs as a slice of strings

	return nil
}
