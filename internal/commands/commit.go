package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Prayag2003/tig/internal/objects"
	"github.com/Prayag2003/tig/internal/repository"
)

func Commit(message string) {
	repoPath := ".tig"
	indexPath := filepath.Join(repoPath, "index")

	// Load index
	indexData, err := os.ReadFile(indexPath)
	if err != nil {
		fmt.Println("Error reading index:", err)
		return
	}
	index := repository.NewIndex()
	index.Load(indexData)

	// Create tree object from index
	entries := []objects.TreeEntry{}
	for file, hash := range index.Entries {
		entries = append(entries, objects.TreeEntry{Name: file, Hash: hash})
	}
	tree := objects.NewTree(entries)
	err = tree.Save()
	if err != nil {
		fmt.Println("Error saving tree object:", err)
		return
	}

	// Create commit object
	headPath := filepath.Join(repoPath, "HEAD")
	headData, err := os.ReadFile(headPath)
	if err != nil {
		fmt.Println("Error reading HEAD:", err)
		return
	}
	headRef := string(headData)
	parentCommit := ""
	if headRef != "ref: refs/heads/master\n" {
		parentCommit = headRef
	}
	commit := objects.NewCommit(message, tree.Hash, parentCommit)
	err = commit.Save()
	if err != nil {
		fmt.Println("Error saving commit object:", err)
		return
	}

	// Update HEAD to point to the new commit
	err = os.WriteFile(headPath, []byte(commit.Hash), 0644)
	if err != nil {
		fmt.Println("Error updating HEAD:", err)
		return
	}

	fmt.Println("Committed with message:", message)
}
