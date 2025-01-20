package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vakintosh/Govial_Init/internal/terraform"

	"github.com/spf13/cobra"
)

// InitCommand initializes the `init` command for project setup
func InitCommand() *cobra.Command {
	var isModule bool

	cmd := &cobra.Command{
		Use:   "init [technology] [directory]",
		Short: "Initialize a new project",
		Long:  `Initialize a new project with the specified technology and optional directory.`,
		Args:  cobra.RangeArgs(1, 2), // Allow 1 or 2 arguments
		Run: func(cmd *cobra.Command, args []string) {
			// Ensure technology is provided
			if len(args) < 1 {
				fmt.Println("Error: Technology is required (e.g., 'ginit init terraform').")
				cmd.Help()
				return
			}

			// Set the technology (e.g., 'terraform', 'python')
			technology := args[0]

			// Validate if the -module flag is set for terraform
			if isModule && technology != "tf" {
				fmt.Println("Error: The --module flag can only be used with terraform.")
				return
			}

			// Handle the directory argument
			var outputDir string
			if len(args) == 2 {
				outputDir = args[1] // Use the provided directory name
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

			// Initialize the project based on the technology
			switch technology {
			case "tf":
				// Pass the template type flag to the GenerateTerraformProject function
				templateType := ""
				if isModule {
					templateType = "module" // Module-specific template
				} else {
					templateType = "" // Default template
				}

				// Initialize a Terraform project
				err := terraform.GenerateTerraformProject(outputDir, projectName, "Unknown Author", templateType)
				if err != nil {
					fmt.Printf("Error initializing Terraform project: %v\n", err)
				} else {
					fmt.Printf("Terraform project '%s' initialized in '%s'\n", projectName, outputDir)
				}
			default:
				fmt.Printf("Unknown technology: %s\n", technology)
			}
		},
	}

	// Add the module flag to specify the template for module configuration (only for terraform)
	cmd.Flags().BoolVarP(&isModule, "module", "m", false, "Generate a module-specific Terraform project (only for terraform)")

	return cmd
}
