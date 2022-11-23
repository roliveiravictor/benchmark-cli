package utils

import (
	"benchmark-cli/model"
	"os"
	"path/filepath"
)

func Charge(
	buffer map[string]os.DirEntry,
	benchmark *model.Macrobenchmark,
) {
	path := filepath.Join(
		Root,
		Module,
		model.JSON_PATH,
	)
	suites := Find(model.JSON_NAME, path, buffer)
	for key, element := range suites {
		joined := filepath.Join(path, key, element.Name())
		file := Open(joined)
		json := Read(file)
		Parse(json, &benchmark)
		Close(file)
	}
}
