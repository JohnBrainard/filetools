// +build !windows

package utils

import (
	"os/exec"
	"os"
)

const NewLineSeparator = "\n"

func GetSystemEditor(fileName string) *exec.Cmd {
	editor := os.Getenv("EDITOR")

	if len(editor) == 0 {
		editor = "vi"
	}

	command := exec.Command(editor, fileName)
	return command
}
