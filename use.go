package main

import (
	"fmt"
	"os/exec"

	"github.com/mitchellh/cli"
)

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
	if len(args) == 0 {
		c.ui.Output(c.Help())
		return 1
	}

	profile, err := profiles.Get(args[0])
	if err != nil {
		c.ui.Error(fmt.Sprint(err))
		return 1
	}

	if err := exec.Command("git", "config", "--local", "user.name", profile.Name).Run(); err != nil {
		c.ui.Error(fmt.Sprintf("execution error: %s", err))
		return 1
	}

	if err := exec.Command("git", "config", "--local", "user.email", profile.Email).Run(); err != nil {
		c.ui.Error(fmt.Sprintf("execution error: %s", err))
		return 1
	}

	return 0
}
