package model

type Benchmark struct {
	Name    string  `json:"name"`
	Metrics Metrics `json:"metrics"`
}
