package main

import "github.com/mitchellh/cli"

func commands(ui cli.Ui) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{ui}, nil
		},
		"show": func() (cli.Command, error) {
			return &ShowCommand{ui}, nil
		},
		"list": func() (cli.Command, error) {
			return &ListCommand{ui}, nil
		},
		"rm": func() (cli.Command, error) {
			return &RemoveCommand{ui}, nil
		},
		"use": func() (cli.Command, error) {
			return &UseCommand{ui}, nil
		},
	}
}
