package kubernetes

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// GenerateManifest generates standalone Kubernetes manifests from a template.
func GenerateManifest(outputDir string, chart HelmChart, templateFile string) error {
	// Ensure the output directory exists
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory: %v", err)
	}

	// Parse the template file
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Prepare output file path and remove .tpl suffix
	outputFile := filepath.Join(outputDir, filepath.Base(templateFile))
	if filepath.Ext(outputFile) == ".tpl" {
		outputFile = outputFile[:len(outputFile)-4]
	}

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Execute the template with the provided data
	if err := tmpl.Execute(file, chart); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	fmt.Printf("Manifest generated from template %s at %s\n", templateFile, outputFile)
	return nil
}
