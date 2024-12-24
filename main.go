// main.go
package main

import (
	"fmt"
	"os"

	"github.com/namekridchai/buildGit/git"
)

func main() {
	// Check the number of command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main.go [command]")
		os.Exit(1)
	}

	// The first argument is the command
	command := os.Args[1]

	switch command {
	case "init":
		git.Init()
	case "hash":
		git.Hash(os.Args[2])
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
