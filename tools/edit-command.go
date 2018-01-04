package tools

import (
	"flag"
	"os"
	"fmt"
	"strings"
	"github.com/JohnBrainard/filetools/utils"
)

type EditCommand struct {
	fileTools *FileTools
	flagSet   *flag.FlagSet

	path      string
	recursive bool
}

func EditCommandNew(fileTools *FileTools) *EditCommand {
	currentPath, _ := os.Getwd()

	editCommand := EditCommand{
		fileTools: fileTools,
	}

	flagSet := flag.NewFlagSet("edit", flag.ExitOnError)
	flagSet.StringVar(&editCommand.path, "path", currentPath, "Path")
	flagSet.BoolVar(&editCommand.recursive, "recursive", false, "Recursive")
	flagSet.BoolVar(&editCommand.recursive, "r", false, "Recursive")

	editCommand.flagSet = flagSet

	return &editCommand
}

func (command *EditCommand) Init() {
	err := command.flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (command *EditCommand) Validate() bool {
	return true
}

func (command *EditCommand) Usage() {
	command.flagSet.Usage()
}

func (command *EditCommand) Execute() {
	editContext := EditContextNew(command.path, command.recursive)
	filePaths := editContext.GetPaths(false)

	editContent := []byte(strings.Join(filePaths, "\n"))
	editContent = utils.EditTempFile(editContent)
	fmt.Printf("Edited File Names: \n%s\n", editContent)

	targetPaths := getTargetPaths(string(editContent))

	if len(targetPaths) == 1 {
		fmt.Printf("Filename templates aren't currently supported")
		editContext.SetTargetPathTemplate(targetPaths[0])
	} else if len(targetPaths) == len(filePaths) {
		fmt.Printf("Renaming files to:\n%s\n", strings.Join(targetPaths, "\n"))
		editContext.SetTargetPaths(targetPaths)
	} else {
		fmt.Printf("Can't rename files")
		return
	}

	editContext.RenameFiles()
}

func getTargetPaths(text string) []string {
	text = strings.TrimSpace(text)

	return strings.Split(text, "\n")
}
