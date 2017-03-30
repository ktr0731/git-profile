package main

import "github.com/mitchellh/cli"

type ListCommand struct {
	ui cli.Ui
}

func (c *ListCommand) Synopsis() string {
	return "List all profiles"
}

func (c *ListCommand) Help() string {
	return "Usage: git profile list"
}

func (c *ListCommand) Run(args []string) int {
	return 0
}
