package main

import (
	"os/exec"
	"sort"
	"strings"
)

func PickDirectory(dirs []Directory) (Selection, error) {
	history := LoadHistory()

	sort.Slice(
		dirs,
		func(i, j int) bool {
			return history[dirs[i].Path] >
				history[dirs[j].Path]
		},
	)

	displayMap := make(map[string]string)

	var entries []string

	for _, dir := range dirs {

		entry := dir.Name + "\t" + dir.Path

		entries = append(entries, entry)

		displayMap[entry] = dir.Path
	}

	cmd := exec.Command(
		"fzf",
		"--expect=ctrl-n,ctrl-v,ctrl-f",
		"--height=50%",
		"--layout=reverse",
		"--border",
		"--cycle",
		"--info=inline-right",
		"--with-nth=1",
		"--prompt=Jump ❯ ",
		"--header=⏎ Actions • ^N Neovim • ^V VS Code • ^F Finder",
		"--preview=echo {} | cut -f2 | xargs eza -T --level=2 --git-ignore",
		"--preview-window=right:45%",
	)

	cmd.Stdin = strings.NewReader(
		strings.Join(entries, "\n"),
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
