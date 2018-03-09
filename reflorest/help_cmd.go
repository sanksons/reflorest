package main

import "flag"
import "fmt"

func BuildHelpCommand() *Command {

	flagSet := flag.NewFlagSet("help", flag.ExitOnError)
	return &Command{
		Name:         "help",
		FlagSet:      flagSet,
		UsageCommand: "reflorest help",
		Usage: []string{
			"Show Menu options",
		},
		Command: func(args []string, additionalArgs []string) {
			generateHelp()
		},
	}
}

func generateHelp() {
	fmt.Println("I am help")
}
