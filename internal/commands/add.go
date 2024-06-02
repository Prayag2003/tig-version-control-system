package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Prayag2003/tig/internal/objects"
	"github.com/Prayag2003/tig/internal/repository"
)

func Add(file string) {
	repoPath := ".tig"
	indexPath := filepath.Join(repoPath, "index")

	// Read file content
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Save the file as a blob object
	blob := objects.NewBlob(content)
	err = blob.Save()
	if err != nil {
		fmt.Println("Error saving blob object:", err)
		return
	}

	// Load or create index
	index := repository.NewIndex()
	indexData, err := os.ReadFile(indexPath)
	if err == nil {
		err = index.Load(indexData)
		if err != nil {
			fmt.Println("Error loading index:", err)
			return
		}
	}

	// Update index
	index.Add(file, blob.Hash)
	indexData, err = index.Save()
	if err != nil {
		fmt.Println("Error saving index:", err)
		return
	}

	// Write index to file
	err = os.WriteFile(indexPath, indexData, 0644)
	if err != nil {
		fmt.Println("Error writing index:", err)
		return
	}

	fmt.Println("Added file:", file)
}
