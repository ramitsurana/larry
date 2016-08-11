package main

import (
	"fmt"
	"github.com/mitchellh/cli"	
	"os"
	//"github.com/ramitsurana/larry"

)

func main() {
	
	args := os.Args[1:]

	cli := &cli.CLI{
		Args:     args,
		Commands: commands,
		HelpFunc: cli.BasicHelpFunc("larry"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}
	return exitCode
}
