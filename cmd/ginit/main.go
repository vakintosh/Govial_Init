package main

import (
	"fmt"
	"os"

	"github.com/vakintosh/Govial_Init/cmd/ginit/commands" // Import the commands package

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ginit",
		Short: "Govial-init is a CLI tool for project initialization",
		Long:  `Govial-init streamlines the setup of new projects across various technologies.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to G-init! Use 'ginit <technology>' to initialize a project.")
			cmd.Help()
		},
	}

	// Add subcommands here
	rootCmd.AddCommand(commands.InitCommand()) // Use the InitCommand function from the ginit/commands package

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
