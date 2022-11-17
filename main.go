package main

import (
	"benchmark-cli/model"
	"benchmark-cli/utils"
)

func main() {
	utils.GradleStop()
	utils.GradleRun()
	//originBenchmark := model.Macrobenchmark{}
	//loadModel("macrobenchmark.json", &originBenchmark)
	//
	//headBenchmark := model.Macrobenchmark{}
	//loadModel("macrobenchmarkV2.json", &headBenchmark)
	//
	//for index, origin := range originBenchmark.Benchmarks {
	//	head := headBenchmark.Benchmarks[index]
	//	if origin.Name == head.Name {
	//		median, _ := origin.Metrics.TTID.Median.Float64()
	//		medianV2, _ := head.Metrics.TTID.Median.Float64()
	//		delta := (1 - median/medianV2) * 100
	//		fmt.Println(" ")
	//		fmt.Printf("Benchmark %s median delta is %.2f%%", origin.Name, delta)
	//		fmt.Println(" ")
	//	}
	//}
}

func loadModel(name string, benchmark *model.Macrobenchmark) {
	macro := utils.Open(name)
	json := utils.Read(macro)

	utils.Parse(json, &benchmark)

	utils.Close(macro)
}
