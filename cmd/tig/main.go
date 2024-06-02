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
	default:
		fmt.Println("Unknown command:", command)
	}
}
