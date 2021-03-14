package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// check is the cobra.Command defining the "terraguard check" action.
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Check a Terraform plan for changes to specific resources",
		Long:  "Check a Terraform plan for changes to specific resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func init() {
	checkCmd.Flags().StringSliceP("plan-json", "j", []string{}, "The path to a Terraform plan output JSON file")
	checkCmd.MarkFlagRequired("plan-json")

	checkCmd.Flags().StringSliceP("guarded-resources", "r", []string{}, "A comma-separated list of guarded resource addresses")
	checkCmd.MarkFlagRequired("guarded-resources")

	rootCmd.AddCommand(checkCmd)
}

func pFailure(url string) {
	fmt.Println(fmt.Sprintf(" \u2718 %s", url))
}
