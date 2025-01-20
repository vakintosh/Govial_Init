package terraform

import (
	"fmt"
	"os"
	"strings"
)

// CreateFile creates a file and writes the trimmed content to it
func CreateFile(filePath, content string) error {
	// Trim the leading whitespace (including empty lines)
	trimmedContent := strings.TrimSpace(content)

	// Create the file
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer f.Close()

	// Write the trimmed content to the file
	_, err = f.WriteString(trimmedContent)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %v", filePath, err)
	}

	return nil
}
