package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type AddCommand struct {
	ui cli.Ui
}

func (c *AddCommand) Synopsis() string {
	return "Add new profile"
}

func (c *AddCommand) Help() string {
	return "Usage: git profile add [options] [<title> <name> <email>]"
}

func (c *AddCommand) Run(args []string) int {
	if len(args) != 3 {
		c.ui.Output(c.Help())
		return 1
	}

	if err := NewProfile(args[0], args[1], args[2]).Save(); err != nil {
		c.ui.Output(fmt.Sprintf("error in Profile.Save(): %s", err))
		return 1
	}

	return 0
}
