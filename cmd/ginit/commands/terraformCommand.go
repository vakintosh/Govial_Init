package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vakintosh/Govial_Init/pkg/terraform"

	"github.com/spf13/cobra"
)

// TfCommand initializes the `tf` command for Terraform project setup
func TfCommand() *cobra.Command {
	var isModule bool

	cmd := &cobra.Command{
		Use:   "tf [directory]",
		Short: "Initialize a new Terraform project",
		Long:  `Initialize a new Terraform project with optional directory and module flag.`,
		Args:  cobra.RangeArgs(1, 1), // Allow 1 argument (directory)
		Run: func(cmd *cobra.Command, args []string) {
			// Handle the directory argument
			var outputDir string
			if len(args) == 1 {
				outputDir = args[0] // Use the provided directory name
			} else {
				outputDir = "." // Use the current directory if no directory is specified
			}

			// Get the project name (use the directory name or current directory)
			projectName := filepath.Base(outputDir)

			// Validate directory and create if necessary
			if _, err := os.Stat(outputDir); os.IsNotExist(err) {
				err := os.MkdirAll(outputDir, os.ModePerm)
				if err != nil {
					fmt.Printf("Error: failed to create output directory %s: %v\n", outputDir, err)
					return
				}
			}

			// Pass the template type flag to the GenerateTerraformProject function
			templateType := ""
			if isModule {
				templateType = "module" // Module-specific template
			}

			// Initialize a Terraform project
			err := terraform.GenerateTerraformProject(outputDir, projectName, "Unknown Author", templateType)
			if err != nil {
				fmt.Printf("Error initializing Terraform project: %v\n", err)
			} else {
				fmt.Printf("Terraform project '%s' initialized in '%s'\n", projectName, outputDir)
			}
		},
	}

	// Add the module flag to specify the template for module configuration (only for terraform)
	cmd.Flags().BoolVarP(&isModule, "module", "m", false, "Generate a module-specific Terraform project (only for terraform)")

	return cmd
}
