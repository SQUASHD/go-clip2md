package watcher

import (
	"fmt"
	"github.com/SQUASHD/go-clip2md/internal/clipboard"
	"time"
)

func Watch() {
	var lastContent string

	lastContent, _ = clipboard.ReadFromClipboard()

	for {
		// Read current clipboard content
		currentContent, err := clipboard.ReadFromClipboard()
		if err != nil {
			fmt.Println("Error reading clipboard:", err)
			continue
		}

		// Check if content has changed
		if currentContent != lastContent {
			fmt.Println("Clipboard content changed:", currentContent)
			// Handle new clipboard content here

			// Update lastContent
			lastContent = currentContent
		}

		// Wait for a bit before checking again
		time.Sleep(1 * time.Second)
	}
}
