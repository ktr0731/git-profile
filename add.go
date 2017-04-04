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

	if err := profiles.Add(Profile{
		Title: args[0],
		Name:  args[1],
		Email: args[2],
	}); err != nil {
		c.ui.Output(fmt.Sprintf("error in Profile.Add(): %s", err))
		return 1
	}

	if err := profiles.Save(); err != nil {
		c.ui.Output(fmt.Sprintf("error in Profile.Save(): %s", err))
		return 1
	}

	return 0
}
