package main

import (
	"benchmark-cli/model"
	"benchmark-cli/utils"
	"fmt"
	"os"
	"path/filepath"
)

// adb uninstall <package_name>

func main() {
	buffer := make(map[string]os.DirEntry)
	originBenchmark := model.Macrobenchmark{}
	headBenchmark := model.Macrobenchmark{}

	utils.Checkout(model.Root, *model.Origin)
	utils.GradleStop(model.Root)
	utils.GradleConnectedAndroidTest(model.Root, model.Module)

	loadBenchmark(buffer, &originBenchmark)

	utils.Checkout(model.Root, *model.Head)
	utils.GradleStop(model.Root)
	utils.GradleConnectedAndroidTest(model.Root, model.Module)

	loadBenchmark(buffer, &headBenchmark)

	printBenchmarkMedian(*model.Origin, &originBenchmark)
	printBenchmarkMedian(*model.Head, &headBenchmark)
	evaluateBenchmarks(&originBenchmark, &headBenchmark)
}

func evaluateBenchmarks(
	originBenchmark *model.Macrobenchmark,
	headBenchmark *model.Macrobenchmark,
) {
	for index, origin := range originBenchmark.Benchmarks {
		head := headBenchmark.Benchmarks[index]
		if origin.Name == head.Name {
			median, _ := origin.Metrics.TTID.Median.Float64()
			medianV2, _ := head.Metrics.TTID.Median.Float64()
			delta := (1 - median/medianV2) * 100
			fmt.Println(" ")
			fmt.Printf("%s median delta is %.2f%%", origin.Name, delta)
			fmt.Println(" ")
		}
	}
}

func printBenchmarkMedian(branch string, benchmark *model.Macrobenchmark) {
	for _, data := range benchmark.Benchmarks {
		fmt.Println(" ")
		fmt.Printf("Median %s is %s", branch, data.Metrics.TTID.Median.String())
		fmt.Println(" ")
	}
}

func loadBenchmark(
	buffer map[string]os.DirEntry,
	benchmark *model.Macrobenchmark,
) {
	path := filepath.Join(
		model.Root,
		model.Module,
		model.JSON_PATH,
	)
	originBenchmarks := utils.Find(model.JSON_NAME, path, buffer)
	for key, element := range originBenchmarks {
		joined := filepath.Join(path, key, element.Name())
		file := utils.Open(joined)
		loadModel(file.Name(), benchmark)
	}
}

func loadModel(name string, benchmark *model.Macrobenchmark) {
	macro := utils.Open(name)
	json := utils.Read(macro)
	utils.Parse(json, &benchmark)
	utils.Close(macro)
}
