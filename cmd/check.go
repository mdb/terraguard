package cmd

import (
	"fmt"
	"strings"

	"github.com/mdb/terraguard"
	"github.com/spf13/cobra"
)

var (
	// CheckShortDescription is the check command's &cobra.Command.Short.
	CheckShortDescription = "Check if a Terraform plan seeks to modify the specified resources"

	// CheckLongDescription is the check command's &cobra.Command.Long.
	CheckLongDescription = CheckShortDescription

	// check is the cobra.Command defining the "terraguard check" action.
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: CheckShortDescription,
		Long:  CheckLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			plan, err := cmd.Flags().GetString("plan")
			if err != nil {
				return err
			}

			resources, err := cmd.Flags().GetStringSlice("guard")
			if err != nil {
				return err
			}

			g, err := terraguard.NewGuard(plan, resources)
			if err != nil {
				return err
			}

			_, violations := g.Check()
			if len(violations) > 0 {
				violatedAddresses := []string{}

				for _, v := range violations {
					violatedAddresses = append(violatedAddresses, v.Address)
				}

				return fmt.Errorf("%s indicates changes to guarded resources:\n\n%s", plan, strings.Join(violatedAddresses, "\n"))
			}

			return nil
		},
	}
)

func init() {
	checkCmd.Flags().StringP("plan", "p", "", "The path to a Terraform plan output JSON file")
	checkCmd.MarkFlagRequired("plan")

	checkCmd.Flags().StringSliceP("guard", "g", []string{}, "A comma-separated list of guarded resource addresses")
	checkCmd.MarkFlagRequired("guard")

	rootCmd.AddCommand(checkCmd)
}

func pFailure(url string) {
	fmt.Println(fmt.Sprintf(" \u2718 %s", url))
}
