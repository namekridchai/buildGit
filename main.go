// main.go
package main

import (
	"fmt"
	"os"

	command_pkg "github.com/namekridchai/buildGit/command"
	"github.com/namekridchai/buildGit/enum"
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
		command_pkg.Hash(os.Args[2],enum.Blob)
	case "cat":
		command_pkg.Cat(os.Args[2],"blob")
	case "write-tree":
		command_pkg.WriteTree(os.Args[2])
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
