package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mitchellh/cli"
)

type ShowCommand struct {
	ui cli.Ui
}

func (c *ShowCommand) Synopsis() string {
	return "Show current using profile"
}

func (c *ShowCommand) Help() string {
	return "Usage: git profile show"
}

func (c *ShowCommand) Run(args []string) int {
	err := execCmd("git", "config", "user.email")
	if err != nil {
		c.ui.Error(fmt.Sprintf("execution error: %s", err))
		return 1
	}
	err = execCmd("git", "config", "user.name")
	if err != nil {
		c.ui.Error(fmt.Sprintf("execution error: %s", err))
		return 1
	}
	return 0
}

func execCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
