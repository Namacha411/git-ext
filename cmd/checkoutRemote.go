/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

// checkoutRemoteCmd represents the checkoutRemote command
var checkoutRemoteCmd = &cobra.Command{
	Use:   "cor",
	Short: "Checkout a remote branch locally",
	Long: `Checkout remote branch locally allows you to select and checkout branches from your remote repositories.
It lists all branches available in your remote Git repository and lets you choose one to checkout as a new local branch. This is useful for reviewing and working on branches that are not yet present in your local workspace.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		branches, err := getRemoteBranches()
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
		if _, err := git("checkout", "-b", branch); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkoutRemoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutRemoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutRemoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
