package commands

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vakintosh/Govial_Init/pkg/terraform"
)

// TfCommand initializes the `tf` command for Terraform project setup
func TfCommand() *cobra.Command {
	var isModule bool
	var dryRun bool
	var author string

	cmd := &cobra.Command{
		Use:   "tf [directory]",
		Short: "Initialize a new Terraform project",
		Long:  `Initialize a new Terraform project in the specified directory with optional flags for modules and dry-run.`,
		Args:  cobra.RangeArgs(1, 1), // Allow exactly one argument (directory)
		Run: func(cmd *cobra.Command, args []string) {
			// Handle the directory argument
			outputDir := args[0]
			projectName := filepath.Base(outputDir)

			// Dry-run preview
			if dryRun {
				fmt.Printf("Dry-run: Terraform project '%s' would be initialized in '%s'\n", projectName, outputDir)
				if isModule {
					fmt.Println("Module-specific template would be used.")
				}
				return
			}

			// Pass the template type flag to the GenerateTerraformProject function
			templateType := ""
			if isModule {
				templateType = "module" // Module-specific template
			}

			// Initialize a Terraform project
			err := terraform.GenerateTerraformProject(outputDir, projectName, author, templateType)
			if err != nil {
				fmt.Printf("Error initializing Terraform project: %v\n", err)
			} else {
				fmt.Printf("Terraform project '%s' initialized in '%s'\n", projectName, outputDir)
			}
		},
	}

	// Flags
	cmd.Flags().BoolVarP(&isModule, "module", "m", false, "Generate a module-specific Terraform project (only for Terraform)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview the project structure without creating files")
	cmd.Flags().StringVar(&author, "author", "Unknown Author", "Specify the author name for the project")

	return cmd
}
