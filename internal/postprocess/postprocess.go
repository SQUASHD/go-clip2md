package postprocess

import (
	"regexp"
	"strings"
)

func ExtractFirstHeader(markdown string) string {
	lines := strings.Split(markdown, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return line[2:]
		} else if strings.HasPrefix(line, "## ") {
			return line[3:]
		}
	}

	return extractFirstWords(markdown, 5)
}

// extractFirstWords extracts the first 'n' words from a string.
func extractFirstWords(s string, n int) string {
	words := strings.Fields(s)
	if len(words) < n {
		n = len(words)
	}
	return strings.Join(words[:n], "-")
}

func PostProcessMarkdown(markdown, flag string) string {
	markdown = escapeSpecialCharsInMarkdown(markdown)
	markdown = compressMultipleNewlines(markdown)
	markdown = escapePercentSigns(markdown)
	markdown = replaceCodeBlocksWithLanguage(markdown, flag)
	return markdown
}

func escapeSpecialCharsInMarkdown(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		// Check if the line starts with > (Markdown blockquote)
		if !strings.HasPrefix(line, ">") {
			// Replace < and > with their escaped versions
			line = strings.ReplaceAll(line, "<", "\\<")
			line = strings.ReplaceAll(line, ">", "\\>")
		}
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

func compressMultipleNewlines(input string) string {
	// Regular expression to match multiple newlines
	re := regexp.MustCompile(`\n{3,}`)
	return re.ReplaceAllString(input, "\n\n")
}

func escapePercentSigns(input string) string {
	return strings.ReplaceAll(input, "%%", "\\%%")
}

func replaceCodeBlocksWithLanguage(content, codeblockLang string) string {
	// codeblockLang is empty if the --lang flag is not set
	if codeblockLang == "" {
		return content
	}

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "```") && line != "```" {
			lines[i] = "```" + codeblockLang
		}
	}

	return strings.Join(lines, "\n")
}
