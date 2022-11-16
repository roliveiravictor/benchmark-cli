package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Read(file *os.File) []byte {
	data, err := ioutil.ReadAll(file)
	nilCheck(err, "Failed to read Macrobenchmark file.")
	return data
}

func Parse[T any](bytes []byte, model *T) {
	err := json.Unmarshal(bytes, model)
	nilCheck(err, "Failed to parse Macrobenchmark json.")
}
