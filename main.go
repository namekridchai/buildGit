// main.go
package main

import (
	"fmt"
	"os"

	command_pkg "github.com/namekridchai/buildGit/command"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main.go [command]")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "init":
		command_pkg.Init()
	case "hash":
		command_pkg.Hash(os.Args[2])
	case "cat":
		command_pkg.Cat(os.Args[2],"blob")
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
