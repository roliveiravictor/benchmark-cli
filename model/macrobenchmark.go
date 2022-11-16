package model

type Macrobenchmark struct {
	Context    Context     `json:"context"`
	Benchmarks []Benchmark `json:"benchmarks"`
}
