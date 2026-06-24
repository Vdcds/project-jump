package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
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

func LoadHistory() map[string]ProjectHistory {
	path := historyPath()

	data, err := os.ReadFile(path)
	if err != nil {
		return make(map[string]ProjectHistory)
	}

	var history map[string]ProjectHistory

	if err := json.Unmarshal(data, &history); err != nil {
		return make(map[string]ProjectHistory)
	}

	return history
}

func SaveHistory(history map[string]ProjectHistory) error {
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

	entry := history[path]

	entry.Count++
	entry.LastOpened = time.Now().Unix()

	history[path] = entry

	err := SaveHistory(history)
	if err != nil {
		panic(err)
	}
}
