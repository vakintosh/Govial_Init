package terraform

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

// GenerateTerraformProject initializes a new Terraform project
func GenerateTerraformProject(outputDir, projectName, author, templateType string) error {
	// Ensure the output directory exists
	if err := ensureOutputDirExists(outputDir); err != nil {
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

	// Select the template file based on the --module flag
	templateFile := "pkg/templates/terraform/tf-aws.tpl"
	if templateType == "module" {
		templateFile = "pkg/templates/terraform/tf-aws-module.tpl"
	}

	// Load and apply the selected template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template file: %v", err)
	}

	// Dynamically get the list of files to generate
	filesToGenerate, err := GetDefinedFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to extract defined files from template: %v", err)
	}

	// Generate each file based on the template
	for _, file := range filesToGenerate {
		outputFile := filepath.Join(outputDir, file)

		// Render the template for the current file
		renderedContent, err := RenderTemplate(tmpl, file, data)
		if err != nil {
			return err
		}

		// Create the file and write the content
		if err := CreateFile(outputFile, renderedContent); err != nil {
			return err
		}
	}

	return nil
}

// Ensure the output directory exists, create it if not
func ensureOutputDirExists(outputDir string) error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}
	return nil
}
