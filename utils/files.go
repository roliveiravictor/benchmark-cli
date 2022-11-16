package utils

import "os"

func Open(name string) *os.File {
	jsonFile, err := os.Open(name)
	nilCheck(err, "Failed to open json file.")
	return jsonFile
}

func Close(file *os.File) {
	defer file.Close()
}
