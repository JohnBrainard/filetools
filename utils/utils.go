package utils

import (
	"os"
	"io/ioutil"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func EditTempFile(content []byte) []byte {
	tempFile, _ := ioutil.TempFile("", "filetools-edit-")
	tempFile.Write(content)
	tempFile.Close()

	defer os.Remove(tempFile.Name())
	return EditFile(tempFile.Name())
}

func EditFile(fileName string) []byte {
	cmd := GetSystemEditor(fileName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	CheckError(err)

	err = cmd.Wait()
	CheckError(err)

	file, err := os.Open(fileName)
	CheckError(err)
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	CheckError(err)

	return text
}
