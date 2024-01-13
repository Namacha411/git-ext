/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

// branchRemoteCmd represents the branchRemote command
var branchRemoteCmd = &cobra.Command{
	Use:   "brr",
	Short: "List and select remote Git branches",
	Long: `BranchRemote command lists all branches in your remote Git repository and provides an interface to select a branch.
This command is useful for quickly browsing and selecting from the branches that exist in your remote repositories without the need to manually fetch or pull the updates.`,
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

		fmt.Print(branch)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(branchRemoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchRemoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchRemoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
