/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

// branchDeleteCmd represents the branchDelete command
var branchDeleteCmd = &cobra.Command{
	Use:   "brd",
	Short: "Delete local Git branches",
	Long: `BranchDelete command provides an interactive interface to delete one or more local Git branches.
It lists all local branches and allows you to select multiple branches for deletion. This command is particularly useful for cleaning up old or obsolete branches from your local repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		branches, err := getBranches()
		if err != nil {
			return err
		}

		indexes, err := fuzzyfinder.FindMulti(
			branches,
			func(i int) string {
				return branches[i]
			},
		)
		if err != nil {
			return err
		}

		for _, i := range indexes {
			branch := getBranchName(branches[i])
			out, err := git("branch", "-d", branch)
			if err != nil {
				return fmt.Errorf("failed to delete branch '%s': %v", branch, err)
			}
			fmt.Println(out)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(branchDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
