package android

import (
	"benchmark-cli/model"
	"benchmark-cli/utils"
	"fmt"
	"os"
)

func Benchmark(
	buffer map[string]os.DirEntry,
	branch string,
	benchmark *model.Macrobenchmark,
) {
	utils.Checkout(utils.Root, branch)
	utils.GradleStop(utils.Root)
	utils.GradleConnectedAndroidTest(utils.Root, utils.Module)
	utils.Charge(buffer, benchmark)
}

func ReportMedian(branch string, benchmark *model.Macrobenchmark) {
	for _, data := range benchmark.Benchmarks {
		fmt.Println(" ")
		fmt.Printf("The %s branch median at %s is %s", branch, data.Name, data.Metrics.TTID.Median.String())
		fmt.Println(" ")
	}
}

func Compare(
	originBenchmark *model.Macrobenchmark,
	headBenchmark *model.Macrobenchmark,
) {
	for index, origin := range originBenchmark.Benchmarks {
		head := headBenchmark.Benchmarks[index]
		if origin.Name == head.Name {
			originMedian, _ := origin.Metrics.TTID.Median.Float64()
			headMedian, _ := head.Metrics.TTID.Median.Float64()
			delta := (1 - originMedian/headMedian) * 100
			fmt.Println(" ")
			if delta >= 0 {
				fmt.Printf("%s median increased is %.2f%%", origin.Name, delta)
			} else {
				fmt.Printf("%s median reduced is %.2f%%", origin.Name, delta)
			}
			fmt.Println(" ")
		}
	}
}
