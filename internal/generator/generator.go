package generator

import (
	"os"
	"path/filepath"

	"github.com/SQUASHD/go-clip2md/internal/sanitize"
)

func SetWorkingDirectory(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if err := os.Chdir(path); err != nil {
		return err
	}
	return nil
}

func WriteContentToFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return err
	}
	return nil
}

func GenerateFilepath(pattern string, h1 string) string {
	h1 = sanitize.ConvertHeaderToFilePath(h1)
	return pattern + h1
}
