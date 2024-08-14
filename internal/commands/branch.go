package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListBranches() {
	repoPath := ".tig"
	branchesPath := filepath.Join(repoPath, "refs", "heads")

	branches, err := os.ReadDir(branchesPath)
	if err != nil {
		fmt.Println("Error reading branches:", err)
		return
	}

	for _, branch := range branches {
		fmt.Println(branch.Name())
	}
}

// CreateBranch creates a new branch.
func CreateBranch(branchName string) {
	repoPath := ".tig"
	headPath := filepath.Join(repoPath, "HEAD")

	// Read the current commit
	headContent, err := os.ReadFile(headPath)
	if err != nil {
		fmt.Println("Error reading HEAD:", err)
		return
	}

	currentCommit := strings.TrimSpace(string(headContent))

	// Create the branch file
	branchPath := filepath.Join(repoPath, "refs", "heads", branchName)
	err = os.WriteFile(branchPath, []byte(currentCommit), 0644)
	if err != nil {
		fmt.Println("Error creating branch:", err)
		return
	}

	fmt.Println("Branch created:", branchName)
}

// SwitchBranch switches to another branch.
func SwitchBranch(branchName string) {
	repoPath := ".tig"
	branchPath := filepath.Join(repoPath, "refs", "heads", branchName)

	// Check if the branch exists
	_, err := os.Stat(branchPath)
	if os.IsNotExist(err) {
		fmt.Println("Branch does not exist:", branchName)
		return
	}

	// Update HEAD to point to the new branch
	headPath := filepath.Join(repoPath, "HEAD")
	err = os.WriteFile(headPath, []byte(branchName), 0644)
	if err != nil {
		fmt.Println("Error switching branch:", err)
		return
	}

	fmt.Println("Switched to branch:", branchName)
}
