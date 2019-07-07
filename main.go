package main

import (
	"sshnot/cmd"
	"fmt"
	"os"
)

func main() {
	command := cmd.Get()
	if err := command.Execute(); err != nil {
		fmt.Printf("Error while executing command: %v\n", err)
		os.Exit(1)
	}
}
