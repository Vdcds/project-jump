package main

import (
	"os/exec"
	"sort"
	"strings"
	"time"
)

func PickDirectory(dirs []Directory) (Selection, error) {
	history := LoadHistory()

	score := func(path string) float64 {
		entry, ok := history[path]
		if !ok {
			return 0
		}

		hoursAgo :=
			(time.Now().Unix() - entry.LastOpened) / 3600

		recencyBonus :=
			100.0 / float64(hoursAgo+1)

		return float64(entry.Count) + recencyBonus
	}

	sort.Slice(
		dirs,
		func(i, j int) bool {
			return score(dirs[i].Path) >
				score(dirs[j].Path)
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
		"--expect=ctrl-n,ctrl-v,ctrl-f,ctrl-g",
		"--height=50%",
		"--layout=reverse",
		"--border",
		"--cycle",
		"--info=inline-right",
		"--with-nth=1",
		"--prompt=Jump ❯ ",
		"--header=⏎ Actions • ^N Neovim • ^V VS Code • ^F Finder • ^G GitHub",
		"--preview=echo {} | cut -f2 | xargs ./scripts/preview.sh",
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
