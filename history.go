package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func historyPath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(
		home,
		".config",
		"project-jump",
		"history.json",
	)
}

func LoadHistory() map[string]int {
	path := historyPath()

	data, err := os.ReadFile(path)
	if err != nil {
		return make(map[string]int)
	}

	var history map[string]int

	if err := json.Unmarshal(data, &history); err != nil {
		return make(map[string]int)
	}

	return history
}

func SaveHistory(history map[string]int) error {
	path := historyPath()

	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(
		history,
		"",
		"  ",
	)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func IncrementProject(path string) {
	history := LoadHistory()

	history[path]++

	err := SaveHistory(history)
	if err != nil {
		panic(err)
	}
}
