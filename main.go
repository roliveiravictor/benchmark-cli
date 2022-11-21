package main

import (
	"benchmark-cli/model"
	"benchmark-cli/utils"
	"fmt"
	"os"
)

func main() {
	project := utils.Home() + "/StudioProjects/play-store-buy-me-a-coffee"
	tests := "/benchmark/build/outputs/connected_android_test_additional_output/benchmark/connected"

	utils.Checkout(project, "master")
	utils.GradleStop(project)
	utils.GradleRun(project)

	path := project + tests

	buffer := make(map[string]os.DirEntry)

	fileName := "benchmarkData"
	originBenchmark := model.Macrobenchmark{}

	originBenchmarks := utils.Find(fileName, path, buffer)
	for key, element := range originBenchmarks {
		file := utils.Open(path + "/" + key + "/" + element.Name())
		loadModel(file.Name(), &originBenchmark)
	}

	utils.Checkout(project, "fix/fb-verify-dynamic-link)")
	utils.GradleStop(project)
	utils.GradleRun(project)

	headBenchmark := model.Macrobenchmark{}

	headBenchmarks := utils.Find(fileName, path, buffer)
	for key, element := range headBenchmarks {
		file := utils.Open(path + "/" + key + "/" + element.Name())
		loadModel(file.Name(), &headBenchmark)
	}

	for _, data := range originBenchmark.Benchmarks {
		fmt.Println("Median master is ", data.Metrics.TTID.Median)
		fmt.Println(" ")
	}

	for _, data := range headBenchmark.Benchmarks {
		fmt.Println("Median release is ", data.Metrics.TTID.Median)
		fmt.Println(" ")
	}

	for index, origin := range originBenchmark.Benchmarks {
		head := headBenchmark.Benchmarks[index]
		if origin.Name == head.Name {
			median, _ := origin.Metrics.TTID.Median.Float64()
			medianV2, _ := head.Metrics.TTID.Median.Float64()
			delta := (1 - median/medianV2) * 100
			fmt.Printf("Benchmark %s median delta is %.2f%%", origin.Name, delta)
			fmt.Println(" ")
		}
	}
}

func loadModel(name string, benchmark *model.Macrobenchmark) {
	macro := utils.Open(name)
	json := utils.Read(macro)

	utils.Parse(json, &benchmark)

	utils.Close(macro)
}
