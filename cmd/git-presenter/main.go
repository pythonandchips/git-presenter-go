package main

import (
	"fmt"
	"git-presenter/internal"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: git-presenter <command> [options]")
		fmt.Println("Commands: init, start, update, next, back, end, start, list, help")
		os.Exit(1)
	}

	command := os.Args[1]
	interactive := true

	if len(os.Args) >= 3 && os.Args[2] == "-c" {
		interactive = false
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	gitPresenter := internal.NewGitPresenter(currentDir, interactive)

	err = gitPresenter.Execute(command)
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
