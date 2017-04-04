package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

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
	if len(args) == 0 {
		c.ui.Output(c.Help())
		return 1
	}

	if err := profiles.Remove(args[0]); err != nil {
		c.ui.Error(fmt.Sprint(err))
		return 1
	}

	if err := profiles.Save(); err != nil {
		c.ui.Error(fmt.Sprint(err))
		return 1
	}

	return 0
}
