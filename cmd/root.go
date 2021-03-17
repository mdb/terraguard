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

	// rootCmd is the terraguard *cobra.Command
	rootCmd = &cobra.Command{
		Use:          "terraguard",
		Short:        ShortDescription,
		Long:         LongDescription,
		SilenceUsage: true,
	}
)

// Execute executes the terraguard command.
func Execute(version string) {
	rootCmd.Version = version

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
