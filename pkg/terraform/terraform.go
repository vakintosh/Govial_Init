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

	isV0_13OrLater, err := GetTerraformVersion()
	if err != nil {
		return fmt.Errorf("error checking Terraform version: %v", err)
	}

	data := TemplateData{
		ProjectName:             projectName,
		Author:                  author,
		CreatedAt:               time.Now().Format("2006-01-02"),
		IsTerraformV0_13OrLater: isV0_13OrLater,
	}

	templateFile := getTemplateFile(templateType)

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template file: %v", err)
	}

	filesToGenerate, err := GetDefinedFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to extract defined files from template: %v", err)
	}

	return generateFiles(outputDir, tmpl, filesToGenerate, data)
}

func getTemplateFile(templateType string) string {
	if templateType == "module" {
		return filepath.Join("pkg", "templates", "terraform", "tf-aws-module.tpl")
	}
	return filepath.Join("pkg", "templates", "terraform", "tf-aws.tpl")
}

func generateFiles(outputDir string, tmpl *template.Template, filesToGenerate []string, data TemplateData) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(filesToGenerate))

	for _, file := range filesToGenerate {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			outputFile := filepath.Join(outputDir, file)

			renderedContent, err := RenderTemplate(tmpl, file, data)
			if err != nil {
				errChan <- err
				return
			}

			if err := CreateFile(outputFile, renderedContent); err != nil {
				errChan <- err
			}
		}(file)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
