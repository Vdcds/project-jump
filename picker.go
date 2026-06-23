package main

import (
	"os/exec"
	"strings"
)

func PickDirectory(dirs []Directory) (Selection, error) {
	displayMap := make(map[string]string)

	var names []string

	for _, dir := range dirs {
		names = append(names, dir.Name)
		displayMap[dir.Name] = dir.Path
	}

	cmd := exec.Command(
		"fzf",
		"--expect=ctrl-n,ctrl-v,ctrl-f",
		"--height=40%",
		"--layout=reverse",
		"--border",
		"--prompt=󰉋  ",
	)

	cmd.Stdin = strings.NewReader(
		strings.Join(names, "\n"),
	)

	out, err := cmd.Output()
	if err != nil {
		return Selection{}, err
	}

	lines := strings.Split(
		strings.TrimSpace(string(out)),
		"\n",
	)

	var key string
	var selected string

	if len(lines) == 1 {
		selected = lines[0]
	} else {
		key = lines[0]
		selected = lines[1]
	}

	return Selection{
		Key:  key,
		Path: displayMap[selected],
	}, nil
}
