/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "co",
	Short: "Switch branches or restore working tree files",
	Long:  `Checkout command is used to switch branches or restore working tree files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		branches, err := getBranches()
		if err != nil {
			return err
		}

		idx, err := fuzzyfinder.Find(
			branches,
			func(i int) string {
				return branches[i]
			},
		)
		if err != nil {
			return err
		}
		branch := getBranchName(branches[idx])

		if _, err := git("checkout", branch); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
