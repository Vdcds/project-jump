package main

import (
	"os"
	"path/filepath"
)

type Directory struct {
	Name string
	Path string
}

var ignored = map[string]bool{
	".git":         true,
	"node_modules": true,
	".next":        true,
	"dist":         true,
	"build":        true,
	"coverage":     true,
	".turbo":       true,
	".cache":       true,
	"vendor":       true,
	"Contents":     true,
}

func FindDirectories() ([]Directory, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	roots := []string{
		filepath.Join(home, "dev"),
		filepath.Join(home, "Desktop"),
		filepath.Join(home, "Downloads"),
		filepath.Join(home, ".config"),
	}

	var dirs []Directory

	for _, root := range roots {

		if _, err := os.Stat(root); err != nil {
			continue
		}

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if info.IsDir() && ignored[info.Name()] {
				return filepath.SkipDir
			}

			if info.IsDir() {
				dirs = append(dirs, Directory{
					Name: filepath.Base(path),
					Path: path,
				})
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return dirs, nil
}
