package main

import (
	"fmt"
	"ginit/cmd/ginit/commands"
	"os"

	// Import the commands package

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
	rootCmd.AddCommand(commands.TfCommand())
	rootCmd.AddCommand(commands.K8sCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
