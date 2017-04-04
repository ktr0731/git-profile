package main

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

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
	var list string
	for _, profile := range profiles {
		list += fmt.Sprintf("[%s]\n- %s\n- %s\n\n", profile.Title, profile.Name, profile.Email)
	}
	c.ui.Output(strings.TrimSpace(list))

	return 0
}
