package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// EnsureFileExists promises to create the given file inside the given directory if it not exists.
// If the file already exist nothing happens. In both cases the absolute filePath is returned.
func EnsureFileExists(directory, filename string) (string, error) {
	// Ensure the directory exists
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Construct the full file path
	filePath := filepath.Join(directory, filename)

	// Check if the file already exists
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("failed to check file existence: %w", err)
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return filePath, nil
}
