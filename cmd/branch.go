/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "br",
	Short: "List and manage Git branches",
	Long: `Branch command lists, creates, or deletes branches in a Git repository.
This command displays the list of branches available in the repository,
with an option to create or delete specific branches.`,
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

		fmt.Print(branch)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
