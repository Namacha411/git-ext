package cmd

import (
	"os/exec"
	"regexp"
	"strings"
)

func git(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	return string(out), err
}

func getBranches() ([]string, error) {
	option := []string{"--no-pager", "branch", "-vv"}
	out, err := git(option...)
	if err != nil {
		return nil, err
	}
	branches := strings.Split(out, "\n")
	return branches, nil
}

func getRemoteBranches() ([]string, error) {
	option := []string{"--no-pager", "branch", "-vv", "-r"}
	out, err := git(option...)
	if err != nil {
		return nil, err
	}
	branches := strings.Split(out, "\n")
	return branches, nil
}

func getBranchName(branchInfo string) string {
	regex := regexp.MustCompile(`^\*`)
	branch := strings.TrimSpace(branchInfo)
	branch = regex.ReplaceAllString(branch, "")
	branch = strings.TrimSpace(branch)
	branch = strings.Split(branch, " ")[0]
	return branch
}
