package sanitize

import (
	"regexp"
	"strings"
)

// SanitizeCellContent sanitizes the content of a table cell
// because the html-to-markdown library copies cell content verbatim, we need to escape any markdown characters
func SanitizeCellContent(cellContent string) string {
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
