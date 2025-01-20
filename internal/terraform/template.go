package terraform

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template" // Change to text/template
)

// TemplateData holds data to be used in templates
type TemplateData struct {
	ProjectName             string
	Author                  string
	CreatedAt               string
	IsTerraformV0_13OrLater bool
}

// GetDefinedFiles extracts the defined blocks from the template file
func GetDefinedFiles(templateFile string) ([]string, error) {
	// Read the template content
	content, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read template file: %v", err)
	}

	// Use regex to find all `{{define "file-name"}}` blocks
	re := regexp.MustCompile(`{{define "(.*?)"}}`)
	matches := re.FindAllSubmatch(content, -1)

	var fileNames []string
	for _, match := range matches {
		if len(match) > 1 {
			fileNames = append(fileNames, string(match[1]))
		}
	}
	return fileNames, nil
}

// RenderTemplate renders the template with the provided data
func RenderTemplate(tmpl *template.Template, templateName string, data TemplateData) (string, error) {
	var renderedContent strings.Builder
	err := tmpl.ExecuteTemplate(&renderedContent, templateName, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template for %s: %v", templateName, err)
	}

	return renderedContent.String(), nil
}
