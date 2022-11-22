package model

type Macrobenchmark struct {
	Context    Context  `json:"context"`
	Benchmarks []Suites `json:"benchmarks"`
}

const JSON_NAME = "benchmarkData"

const JSON_PATH = "/build/outputs/connected_android_test_additional_output/benchmark/connected"
