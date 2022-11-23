package main

import (
	"benchmark-cli/android"
	"benchmark-cli/model"
	"benchmark-cli/utils"
	"os"
)

func main() {
	buffer := make(map[string]os.DirEntry)
	origin := model.Macrobenchmark{}
	head := model.Macrobenchmark{}
	android.Benchmark(buffer, *utils.Origin, &origin)
	android.Benchmark(buffer, *utils.Head, &head)
	android.ReportMedian(*utils.Origin, &origin)
	android.ReportMedian(*utils.Head, &head)
	android.Compare(&origin, &head)
}
