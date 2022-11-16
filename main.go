package main

import (
	"benchmark-cli/model"
	"benchmark-cli/utils"
	"fmt"
)

func main() {
	macro := utils.Open("macrobenchmark.json")
	json := utils.Read(macro)

	benchmark := model.Macrobenchmark{}
	utils.Parse(json, &benchmark)

	utils.Close(macro)

	fmt.Printf("Model is %s", benchmark.Context.Build.Model)
}
