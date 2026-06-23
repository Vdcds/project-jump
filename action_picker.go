package main

import (
	"os/exec"
	"strings"
)

func PickAction() (string, error) {
	actions := []string{
		"nvim",
		"code",
		"finder",
		"cd",
	}

	cmd := exec.Command(
		"fzf",
		"--height=20%",
		"--border",
	)

	cmd.Stdin = strings.NewReader(
		strings.Join(actions, "\n"),
	)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
