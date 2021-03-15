package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// ShortDescription is terraguard's &cobra.Command.Short
	ShortDescription string = "Protect against problematic Terraform actions"

	// LongDescription is terraguard's &cobra.Command.Long
	LongDescription string = "Ensure a Terraform plan does not modify guarded resources"
)

var rootCmd = &cobra.Command{
	Use:          "terraguard",
	Short:        ShortDescription,
	Long:         LongDescription,
	Version:      "0.0.1",
	SilenceUsage: true,
}

// Execute executes the terraguard command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
