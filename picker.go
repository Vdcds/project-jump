package main

import (
	"os/exec"
	"strings"
)

func PickDirectory(dirs []Directory) (Directory, error) {
	displayMap := make(map[string]Directory)

	var names []string

	for _, dir := range dirs {
		names = append(names, dir.Name)
		displayMap[dir.Name] = dir
	}

	input := strings.Join(names, "\n")

	cmd := exec.Command(
		"fzf",
		"--height=40%",
		"--layout=reverse",
		"--border",
		"--prompt=󰉋  ",
	)

	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.Output()
	if err != nil {
		return Directory{}, err
	}

	selected := strings.TrimSpace(string(out))

	return displayMap[selected], nil
}
