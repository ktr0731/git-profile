package main

import "github.com/mitchellh/cli"

type CurrentCommand struct {
	ui cli.Ui
}

func (c *CurrentCommand) Synopsis() string {
	return "Show current profile"
}

func (c *CurrentCommand) Help() string {
	return "Usage: git profile current"
}

func (c *CurrentCommand) Run(args []string) int {
	return 0
}
