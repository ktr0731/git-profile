package main

import "github.com/mitchellh/cli"

type AddCommand struct {
	ui cli.Ui
}

func (c *AddCommand) Synopsis() string {
	return "Add new profile"
}

func (c *AddCommand) Help() string {
	return "Usage: git profile add <title> <name> <email>"
}

func (c *AddCommand) Run(args []string) int {
	return 0
}
