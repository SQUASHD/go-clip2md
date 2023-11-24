package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SQUASHD/go-clip2md/internal/clippy"
	"github.com/SQUASHD/go-clip2md/internal/converter"
	gen "github.com/SQUASHD/go-clip2md/internal/generator"
	"github.com/SQUASHD/go-clip2md/internal/postprocess"
	"github.com/SQUASHD/go-clip2md/internal/sanitize"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch clipboard for changes",
	Long:  `Watch clipboard for changes and generate markdown files`,
	Run:   runWatchCmd,
}

func runWatchCmd(cmd *cobra.Command, args []string) {
	mode, file, count := setupWatch(cmd)
	var lastContent string

	lastContent, _ = clippy.ReadFromClipboard()
	for {
		contentChanged, currentContent := checkClipboardChange(&lastContent)
		if contentChanged {
			var err error
			file, err = processClipboardContent(currentContent, &count, mode, file, cmd)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		time.Sleep(time.Duration(getInterval(cmd)) * time.Second)
	}

	if file != nil {
		file.Close()
	}
}

func setupWatch(cmd *cobra.Command) (string, *os.File, int) {
	out, _ := cmd.Flags().GetString("out")
	mode, _ := cmd.Flags().GetString("mode")

	if mode != "new" && mode != "append" {
		fmt.Println("Invalid mode:", mode)
		os.Exit(1)
	}

	if err := gen.SetWorkingDirectory(out); err != nil {
		fmt.Println("Error setting working directory:", err)
		os.Exit(1)
	}

	return mode, nil, 0
}

func checkClipboardChange(lastContent *string) (bool, string) {
	currentContent, err := clippy.ReadFromClipboard()
	if err != nil {
		fmt.Println("Error reading clipboard:", err)
		return false, *lastContent
	}
	if currentContent != *lastContent {
		*lastContent = currentContent
		return true, currentContent
	}
	return false, *lastContent
}

func processClipboardContent(content string, count *int, mode string, file *os.File, cmd *cobra.Command) (*os.File, error) {
	*count++
	md, err := convertContentToMarkdown(content, cmd)
	if err != nil {
		return file, fmt.Errorf("error converting HTML: %w", err)
	}

	filepath := generateFilePath(md, *count, cmd)
	if mode == "new" || *count == 1 {
		if file != nil {
			file.Close()
		}
		file, err = openFileForWrite(filepath, mode)
		if err != nil {
			return nil, err
		}
	}

	if err := writeFile(file, md); err != nil {
		return file, err
	}

	if mode == "new" || *count == 1 {
		fmt.Println("New file:", filepath)
	} else {
		fmt.Println("Appended to file")
	}
	return file, nil
}

func convertContentToMarkdown(content string, cmd *cobra.Command) (string, error) {
	lang, _ := cmd.Flags().GetString("lang")
	domain, _ := cmd.Flags().GetString("domain")
	md, err := converter.ConvertHTML2Md(domain, content)
	if err != nil {
		return "", fmt.Errorf("error converting HTML: %w", err)
	}
	return postprocess.PostProcessMarkdown(md, lang), nil
}

func generateFilePath(md string, count int, cmd *cobra.Command) string {
	pattern, _ := cmd.Flags().GetString("pattern")
	h1 := postprocess.ExtractFirstHeader(md)

	if pattern != "" {
		return gen.GenerateFilepath(fmt.Sprintf("%s%d-", pattern, count), h1) + ".md"
	}
	return sanitize.ConvertHeaderToFilePath(h1) + ".md"
}

func openFileForWrite(filepath string, mode string) (*os.File, error) {
	if mode == "new" {
		return os.Create(filepath)
	}
	return os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func writeFile(file *os.File, content string) error {
	if _, err := file.WriteString(content + "\n"); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

func getInterval(cmd *cobra.Command) int {
	interval, _ := cmd.Flags().GetInt("interval")
	return interval
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().StringP("out", "o", ".", "Output directory")
	watchCmd.Flags().IntP("interval", "i", 1, "Interval to check clipboard for changes")
	watchCmd.Flags().StringP("pattern", "p", "", "Pattern to use for file naming, e.g. doc, which becomes doc1.md, doc2.md, etc.")
	watchCmd.Flags().StringP("mode", "m", "new", "mode for file writing, new or append")
}
