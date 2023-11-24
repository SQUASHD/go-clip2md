package clippy

import (
	"github.com/atotto/clipboard"
)

func WriteToClipboard(text string) error {
	return clipboard.WriteAll(text)
}

func ReadFromClipboard() (string, error) {
	return clipboard.ReadAll()
}
