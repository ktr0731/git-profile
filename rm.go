package main

import "github.com/mitchellh/cli"

type RemoveCommand struct {
	ui cli.Ui
}

func (c *RemoveCommand) Synopsis() string {
	return "Remove selected profile"
}

func (c *RemoveCommand) Help() string {
	return "Usage: git profile rm <profile-name>"
}

func (c *RemoveCommand) Run(args []string) int {
	return 0
}
