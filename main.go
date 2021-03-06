package main

import (
	"fmt"
	"os"
	"gitlab.com/soerenschneider/ssh-login-notification/cmd"
)

func main() {
	command := cmd.FireUp()
	if err := command.Execute(); err != nil {
		fmt.Printf("Error while executing command: %v\n", err)
		os.Exit(1)
	}
}
