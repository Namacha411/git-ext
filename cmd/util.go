package cmd

import (
	"os/exec"
)

func git(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	return string(out[:]), err
}