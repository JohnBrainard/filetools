package tools

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"path"
	"strings"
)

type FileContext struct {
	Path       string
	TargetPath string
	FileInfo   os.FileInfo
}

type EditContext struct {
	Path  string
	Files []FileContext
}

func EditContextNew(path string, recursive bool) EditContext {
	files := listFiles(path, recursive)

	return EditContext{
		Path:  path,
		Files: files,
	}
}

func (context *EditContext) GetPaths(absolute bool) []string {
	paths := make([]string, len(context.Files))

	for i, fileContext := range context.Files {
		if absolute {
			paths[i] = fileContext.Path
		} else {
			paths[i], _ = filepath.Rel(context.Path, fileContext.Path)
		}
	}

	return paths
}

func (context *EditContext) SetTargetPaths(targetPaths []string) error {
	for i, targetPath := range targetPaths {
		fileContext := &context.Files[i]

		if path.IsAbs(targetPath) {
			fileContext.TargetPath = targetPath
		} else {
			targetPath = filepath.Join(context.Path, targetPath)
			targetPath = filepath.Clean(targetPath)
			fileContext.TargetPath = targetPath
		}
	}

	return nil
}

func (context *EditContext) SetTargetPathTemplate(template string) {
	fmt.Println("This will be awesome when it's implemented!")
	for i := range context.Files {
		fileContext := &context.Files[i]
		fileContext.TargetPath = template
	}
}

func (context *EditContext) RenameFiles() error {
	for _, file := range context.Files {
		if strings.Compare(file.Path, file.TargetPath) == 0 {
			fmt.Printf("Skpping %s. No changes.\n", file.Path)
		} else {
			fmt.Printf("REN: %s to \n └── %s\n", file.Path, file.TargetPath)
		}

	}

	return nil
}

func listFiles(path string, recursive bool) []FileContext {
	var files []FileContext

	handleFile := func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsDir() {
			return nil
		}

		files = append(files, FileContextNew(path, info))
		return nil
	}

	if recursive {
		filepath.Walk(path, handleFile)
	} else {
		fileInfos, _ := ioutil.ReadDir(path)
		for _, info := range fileInfos {
			handleFile(filepath.Join(path, info.Name()), info, nil)
		}
	}

	return files
}

func FileContextNew(path string, info os.FileInfo) FileContext {
	return FileContext{
		Path:     path,
		FileInfo: info,
	}
}
