package main

import "github.com/mitchellh/cli"

type UseCommand struct {
	ui cli.Ui
}

func (c *UseCommand) Synopsis() string {
	return "Use selected profile"
}

func (c *UseCommand) Help() string {
	return "Usage: git profile use <profile-name>"
}

func (c *UseCommand) Run(args []string) int {
	return 0
}
