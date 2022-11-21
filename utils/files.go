package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func Open(name string) *os.File {
	file, err := os.Open(name)
	stringNilCheck(err, "Failed to open json file.")
	return file
}

func Close(file *os.File) {
	defer file.Close()
}

func Find(subName string, path string, buffer map[string]os.DirEntry) map[string]os.DirEntry {
	files, err := os.ReadDir(path)
	stringNilCheck(err, "Failed reading path")

	var pos = 0
	for _, file := range files {
		if file.IsDir() {
			info, err := file.Info()
			stringNilCheck(err, "Failed to find Info")
			Find(subName, path+"/"+info.Name(), buffer)
		}
		if strings.Contains(file.Name(), subName) {
			buffer[filepath.Base(path)] = file
			pos++
		}
	}

	return buffer
}
