package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	dirs, err := FindDirectories()
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Since(start))
	fmt.Println("Found", len(dirs), "directories")

	selected, err := PickDirectory(dirs)
	if err != nil {
		return
	}

	fmt.Println(selected.Path)
}
