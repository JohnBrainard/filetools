package main

import (
	"flag"
	"github.com/JohnBrainard/filetools/tools"
	"os"
)

type Command interface {
	Init()
	Validate() bool
	Usage()
	Execute()
}

func main() {
	fileTools := tools.FileToolsCreate()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	commands := map[string]Command{
		"edit": tools.EditCommandNew(fileTools),
	}

	command, exists := commands[os.Args[1]]
	if exists {
		command.Init()
		if !command.Validate() {
			command.Usage()
			os.Exit(1)
		}

		command.Execute()
	} else {
		flag.Usage()
	}
}
