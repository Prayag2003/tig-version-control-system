package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Init() {
	err := os.Mkdir(".tig", 0755)
	if err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}

	dirs := []string{
		filepath.Join(".tig", "objects"),
		filepath.Join(".tig", "refs", "heads"),
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	headFile := filepath.Join(".tig", "HEAD")
	err = os.WriteFile(headFile, []byte("ref: refs/heads/master\n"), 0644)
	if err != nil {
		fmt.Println("Error creating HEAD file:", err)
		return
	}

	fmt.Println("Initialized empty Tig repository in .tig")
}
