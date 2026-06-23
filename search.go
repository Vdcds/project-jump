package main

import (
	"os"
	"path/filepath"
	"strings"
)

func IsGitRepo(path string) bool {
	_, err := os.Stat(filepath.Join(path, ".git"))
	return err == nil
}

func FindDirectories() ([]Directory, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	var dirs []Directory

	// -------------------------
	// DEV -> Git repos only
	// -------------------------

	devRoot := filepath.Join(home, "dev")

	if _, err := os.Stat(devRoot); err == nil {

		err = filepath.Walk(
			devRoot,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}

				if !info.IsDir() {
					return nil
				}

				// Skip app bundles
				if strings.HasSuffix(info.Name(), ".app") {
					return filepath.SkipDir
				}

				// Found a repo
				if IsGitRepo(path) {

					dirs = append(dirs, Directory{
						Name: filepath.Base(path),
						Path: path,
					})

					return filepath.SkipDir
				}

				return nil
			},
		)
		if err != nil {
			return nil, err
		}
	}

	// -------------------------
	// Desktop
	// -------------------------

	if err := addTopLevelDirs(
		filepath.Join(home, "Desktop"),
		&dirs,
	); err != nil {
		return nil, err
	}

	// -------------------------
	// Downloads
	// -------------------------

	if err := addTopLevelDirs(
		filepath.Join(home, "Downloads"),
		&dirs,
	); err != nil {
		return nil, err
	}

	// -------------------------
	// .config
	// -------------------------

	if err := addTopLevelDirs(
		filepath.Join(home, ".config"),
		&dirs,
	); err != nil {
		return nil, err
	}

	return dirs, nil
}

func addTopLevelDirs(root string, dirs *[]Directory) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil
	}

	for _, entry := range entries {

		if !entry.IsDir() {
			continue
		}

		if strings.HasSuffix(entry.Name(), ".app") {
			continue
		}

		*dirs = append(*dirs, Directory{
			Name: entry.Name(),
			Path: filepath.Join(root, entry.Name()),
		})
	}

	return nil
}
