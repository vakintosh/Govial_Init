package utils

import (
	"fmt"
	"os"
)

// EnsureOutputDirExists ensures the specified output directory exists,
// creating it if it does not.
func EnsureOutputDirExists(outputDir string) error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}
	return nil
}
