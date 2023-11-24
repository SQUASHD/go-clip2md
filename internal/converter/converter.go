package converter

import (
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/SQUASHD/go-clip2md/internal/sanitize"
)

func ConvertHTML2Md(domain, content string) (string, error) {
	converter := md.NewConverter(domain, true, nil)
	customRules := []md.Rule{
		{
			Filter: []string{"table"},
			Replacement: func(content string, selection *goquery.Selection, opt *md.Options) *string {
				var markdownTable strings.Builder
				isFirstRow := true

				selection.Find("tr").Each(func(i int, s *goquery.Selection) {
					s.Find("th, td").Each(func(j int, ss *goquery.Selection) {
						// the html-to-markdown library copies cell content verbatim, we need to escape any markdown characters
						cellContent := sanitize.SanitizeCellContent(strings.TrimSpace(ss.Text()))
						markdownTable.WriteString("| ")
						markdownTable.WriteString(cellContent)
					})
					markdownTable.WriteString(" |\n")

					// tables were not being renderered correctly
					if isFirstRow {
						s.Find("th, td").Each(func(j int, ss *goquery.Selection) {
							markdownTable.WriteString("|-")
						})
						markdownTable.WriteString("|\n")
						isFirstRow = false
					}
				})

				result := markdownTable.String()
				return &result
			},
		},
	}

	converter.AddRules(customRules...)

	htmlContent := string(content)
	markdown, err := converter.ConvertString(htmlContent)
	if err != nil {
		return "", err
	}
	return markdown, nil
}
