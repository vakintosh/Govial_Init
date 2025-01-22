package terraform

import (
	"fmt"
	"path/filepath"
	"sync"
	"text/template"
	"time"

	"github.com/vakintosh/Govial_Init/pkg/utils"
)

// GenerateTerraformProject initializes a new Terraform project
func GenerateTerraformProject(outputDir, projectName, author, templateType string) error {
	if err := utils.EnsureOutputDirExists(outputDir); err != nil {
		return err
	}

	// Check Terraform version (v0.13 or later)
	isV0_13OrLater, err := GetTerraformVersion()
	if err != nil {
		return fmt.Errorf("error checking Terraform version: %v", err)
	}

	// Define template data
	data := TemplateData{
		ProjectName:             projectName,
		Author:                  author,
		CreatedAt:               time.Now().Format("2006-01-02"),
		IsTerraformV0_13OrLater: isV0_13OrLater,
	}

	// Determine the template file based on the templateType
	var templateFile string
	if templateType == "module" {
		templateFile = filepath.Join("pkg", "templates", "terraform", "tf-aws-module.tpl")
	} else {
		templateFile = filepath.Join("pkg", "templates", "terraform", "tf-aws.tpl")
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template file: %v", err)
	}

	// Get the list of files to generate
	filesToGenerate, err := GetDefinedFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to extract defined files from template: %v", err)
	}

	// Use a WaitGroup to synchronize Go routines
	var wg sync.WaitGroup
	errChan := make(chan error, len(filesToGenerate)) // Capture errors

	// Generate each file in parallel
	for _, file := range filesToGenerate {
		wg.Add(1)
		go func(file string) {
			defer wg.Done() // Signal completion of the Go routine
			outputFile := filepath.Join(outputDir, file)

			// Render the template for the current file
			renderedContent, err := RenderTemplate(tmpl, file, data)
			if err != nil {
				errChan <- err
				return
			}

			// Create the file and write the content
			if err := CreateFile(outputFile, renderedContent); err != nil {
				errChan <- err
			}
		}(file)
	}

	// Wait for all Go routines to complete
	wg.Wait()
	close(errChan)

	// Check for errors during file generation
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
