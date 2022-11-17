package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Read(file *os.File) []byte {
	data, err := ioutil.ReadAll(file)
	stringNilCheck(err, "Failed to read Macrobenchmark file.")
	return data
}

func Parse[T any](bytes []byte, model *T) {
	err := json.Unmarshal(bytes, model)
	stringNilCheck(err, "Failed to parse Macrobenchmark json.")
}
