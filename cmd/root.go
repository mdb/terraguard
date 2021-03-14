package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "terraguard",
	Short:        "Protect against problematic Terraform actions",
	Long:         "Ensure a Terraform plan does not modify guarded resources",
	Version:      "0.0.1",
	SilenceUsage: true,
}

// Execute executes the terraguard command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
