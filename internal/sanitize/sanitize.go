package sanitize

import (
	"regexp"
	"strings"
)

func SanitizeCellContent(cellContent string) string {
	// Example sanitization logic:
	// If the asterisk is surrounded by spaces, it's more likely to be a formatting character
	// You can extend this logic based on your specific use case
	cellContent = strings.ReplaceAll(cellContent, "*", "\\*")
	cellContent = strings.ReplaceAll(cellContent, "[", "\\[")
	cellContent = strings.ReplaceAll(cellContent, "]", "\\]")
	return cellContent
}

func ConvertHeaderToFilePath(h1 string) string {
	result := strings.ToLower(h1)
	result = strings.ReplaceAll(result, " ", "-")

	// Remove any non-alphanumeric characters except for dashes
	reg := regexp.MustCompile(`[^a-z0-9-]+`)
	result = reg.ReplaceAllString(result, "")

	return result
}
