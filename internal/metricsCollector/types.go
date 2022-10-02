package metricscollector

import (
	"fmt"
	"sync"
)

type (
	gauge     float64
	Counter   uint64
	second    uint64
	randomVal uint64
	Metrics   struct {
		Mu           sync.Mutex
		values       map[string]fmt.Stringer
		pollCount    Counter
		randomVal    randomVal
		pollInterval second
	}
)

func (g gauge) String() string {
	return fmt.Sprintf("%v", float64(g))
}

func (c Counter) String() string {
	return fmt.Sprintf("%v", uint64(c))
}
