package main

import (
	"fmt"
	"os"

	"github.com/Prayag2003/tig/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tig <command> [<args>]")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		commands.Init()

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tig add <file>")
			return
		}
		commands.Add(os.Args[2])

	case "commit":
		if len(os.Args) < 4 || os.Args[2] != "-m" {
			fmt.Println("Usage: tig commit -m <message>")
			return
		}
		commands.Commit(os.Args[3])

	case "log":
		commands.Log()

	case "clone":
		repoURL := os.Args[2]
		if len(os.Args) < 3 {
			fmt.Println("Usage: tig clone <repo_url> [<destination_path>]")
			return
		}
		var destinationPath string
		if len(os.Args) > 3 {
			destinationPath = os.Args[3]
		} else {
			destinationPath, _ = os.Getwd()
		}
		commands.Clone(repoURL, destinationPath)

	case "branch":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tig branch <list|create|switch> [branch_name]")
			return
		}
		branchCommand := os.Args[2]
		switch branchCommand {
		case "list":
			commands.ListBranches()
		case "create":
			if len(os.Args) < 4 {
				fmt.Println("Usage: tig branch create <branch_name>")
				return
			}
			branchName := os.Args[3]
			commands.CreateBranch(branchName)
		case "switch":
			if len(os.Args) < 4 {
				fmt.Println("Usage: tig branch switch <branch_name>")
				return
			}
			branchName := os.Args[3]
			commands.SwitchBranch(branchName)
		default:
			fmt.Println("Unknown branch command:", branchCommand)
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}
