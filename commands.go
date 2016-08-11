package main

import (
	"os"
	"os/signal"
	"syscall"
	
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {	

	Commands = map[string]cli.CommandFactory{
		"fetch": func() (cli.Command, error) {
			return &command.FetchCommand
		},

		"convert": func() (cli.Command, error) {
			return &command.convertCommand
		},		
	}
}
