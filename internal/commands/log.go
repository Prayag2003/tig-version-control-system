package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Prayag2003/tig/internal/objects"
)

func Log() {
	repoPath := ".tig"
	headPath := filepath.Join(repoPath, "HEAD")

	headData, err := os.ReadFile(headPath)
	if err != nil {
		fmt.Println("Error reading HEAD:", err)
		return
	}

	commitHash := string(headData)
	for commitHash != "" {
		commitPath := filepath.Join(repoPath, "objects", commitHash[:2], commitHash[2:])
		commitData, err := ioutil.ReadFile(commitPath)
		if err != nil {
			fmt.Println("Error reading commit object:", err)
			return
		}

		commit := objects.Commit{}
		err = commit.Load(commitData)
		if err != nil {
			fmt.Println("Error loading commit object:", err)
			return
		}

		fmt.Printf("Commit: %s\nMessage: %s\n\n", commit.Hash, commit.Message)

		// If there is no parent commit, break the loop
		if commit.Parent == "" {
			break
		}

		// Set the next commit hash
		commitHash = commit.Parent
	}
}
