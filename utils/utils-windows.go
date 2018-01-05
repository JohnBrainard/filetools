// +build windows

package utils

import (
	"os"
	"os/exec"
	"strings"
)

const NewLineSeparator = "\r\n"

func GetSystemEditor(fileName string) *exec.Cmd {
	editor := os.Getenv("EDITOR")
	editorArgs := os.Getenv("EDITOR_ARGS")

	var args []string

	if len(editor) == 0 {
		editor = "notepad.exe"
	}

	if len(editorArgs) > 0 {
		args = append(args, strings.Split(editorArgs, ",")...)
	}

	args = append(args, fileName)

	command := exec.Command(editor, args...)
	return command
}
