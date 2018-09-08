package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func filterDirs(files []os.FileInfo) []os.FileInfo {
	y := files[:0]

	for _, file := range files {
		if file.IsDir() {
			y = append(y, file)
		}
	}

	return y
}

func files(path string, withFiles bool) ([]os.FileInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	files, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	if !withFiles {
		files = filterDirs(files)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files, nil
}

func fileSize(file os.FileInfo) string {
	if file.IsDir() {
		return ""
	}

	if file.Size() == 0 {
		return " (empty)"
	}

	return " (" + strconv.FormatInt(file.Size(), 10) + "b)"
}

func separators(fileIndex int, total int) (toFile string, branch string) {
	toFile = "├───"
	branch = "│"

	if total == fileIndex {
		toFile = "└───"
		branch = ""

	}
	return
}

func createTree(output io.Writer, prefix string, path string, printFiles bool) error {

	files, err := files(path, printFiles)
	if err != nil {
		return err
	}

	totalFiles := len(files) - 1
	for i, file := range files {
		toFile, branch := separators(i, totalFiles)
		fileSize := fileSize(file)
		fmt.Fprintln(output, prefix+toFile+file.Name()+fileSize)

		if file.IsDir() {
			createTree(output, prefix+branch+"\t", path+string(os.PathSeparator)+file.Name(), printFiles)
		}
	}

	return nil
}

func dirTree(output io.Writer, path string, printFiles bool) error {
	return createTree(output, "", path, printFiles)
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
