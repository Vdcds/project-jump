package main

import (
	"fmt"
)

func main() {
	dirs, err := FindDirectories()
	if err != nil {
		panic(err)
	}

	selection, err := PickDirectory(dirs)
	if err != nil {
		return
	}

	action := ""

	switch selection.Key {

	case "ctrl-n":
		action = "nvim"

	case "ctrl-v":
		action = "code"

	case "ctrl-f":
		action = "finder"
	}

	if action == "" {

		action, err = PickAction()
		if err != nil {
			return
		}
	}
	IncrementProject(selection.Path)
	fmt.Printf(
		"%s::%s\n",
		action,
		selection.Path,
	)
}
