package metricscollector

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func (m *Metrics) updateCounter() {
	m.Mu.Lock()
	m.pollCount = (m.pollCount + 1) % 18446744073709551615
	m.Mu.Unlock()
}

func (m *Metrics) updateMap(rMs *runtime.MemStats) {
	m.Mu.Lock()
	m.values["Alloc"] = gauge(rMs.Alloc)
	m.values["BuckHashSys"] = gauge(rMs.BuckHashSys)
	m.values["Frees"] = gauge(rMs.Frees)
	m.values["GCCPUFraction"] = gauge(rMs.GCCPUFraction)
	m.values["GCSys"] = gauge(rMs.GCSys)
	m.values["HeapAlloc"] = gauge(rMs.HeapAlloc)
	m.values["HeapIdle"] = gauge(rMs.HeapIdle)
	m.values["HeapInuse"] = gauge(rMs.HeapInuse)
	m.values["HeapObjects"] = gauge(rMs.HeapObjects)
	m.values["HeapReleased"] = gauge(rMs.HeapReleased)
	m.values["HeapSys"] = gauge(rMs.HeapSys)
	m.values["LastGC"] = gauge(rMs.LastGC)
	m.values["Lookups"] = gauge(rMs.Lookups)
	m.values["MCacheInuse"] = gauge(rMs.MCacheInuse)
	m.values["MCacheSys"] = gauge(rMs.MCacheSys)
	m.values["MSpanInuse"] = gauge(rMs.MSpanInuse)
	m.values["MSpanSys"] = gauge(rMs.MSpanSys)
	m.values["Mallocs"] = gauge(rMs.Mallocs)
	m.values["NextGC"] = gauge(rMs.NextGC)
	m.values["NumForcedGC"] = gauge(rMs.NumForcedGC)
	m.values["NumGC"] = gauge(rMs.NumGC)
	m.values["OtherSys"] = gauge(rMs.OtherSys)
	m.values["PauseTotalNs"] = gauge(rMs.PauseTotalNs)
	m.values["StackInuse"] = gauge(rMs.StackInuse)
	m.values["StackSys"] = gauge(rMs.StackSys)
	m.values["Sys"] = gauge(rMs.Sys)
	m.values["TotalAlloc"] = gauge(rMs.TotalAlloc)
	m.values["PollCount"] = m.pollCount
	m.values["RandomValue"] = gauge(m.randomVal)
	m.Mu.Unlock()
}

func (m *Metrics) ReturnMap() *map[string]fmt.Stringer {
	return &m.values
}

func (m *Metrics) ReadMetrics() {
	tempMet := runtime.MemStats{}

	runtime.ReadMemStats(&tempMet)

	m.updateMap(&tempMet)
	m.updateCounter()
}

func (m *Metrics) pollTickerInit() *time.Ticker {
	return time.NewTicker(time.Duration(m.pollInterval * second(time.Second)))
}

func MetricsCollectorMain(mCtx context.Context, m *Metrics, retChan chan struct{}) {
	defer func() { close(retChan) }()

	ticker := m.pollTickerInit()

	for {
		select {
		case <-mCtx.Done():
			log.Printf("---> cancelling MetricsCollectorMain context")
			return
		case <-ticker.C:
			m.ReadMetrics()
			log.Printf("---> Metrics collected, Pollcount=%d", m.pollCount)
		default:
			continue
		}
	}
}

func MetricsInit(pollInterval second) *Metrics {
	return &Metrics{Mu: sync.Mutex{},
		values: map[string]fmt.Stringer{
			"Alloc":         nil,
			"BuckHashSys":   nil,
			"Frees":         nil,
			"GCCPUFraction": nil,
			"GCSys":         nil,
			"HeapAlloc":     nil,
			"HeapIdle":      nil,
			"HeapInuse":     nil,
			"HeapObjects":   nil,
			"HeapReleased":  nil,
			"HeapSys":       nil,
			"LastGC":        nil,
			"Lookups":       nil,
			"MCacheInuse":   nil,
			"MCacheSys":     nil,
			"MSpanInuse":    nil,
			"MSpanSys":      nil,
			"Mallocs":       nil,
			"NextGC":        nil,
			"NumForcedGC":   nil,
			"NumGC":         nil,
			"OtherSys":      nil,
			"PauseTotalNs":  nil,
			"StackInuse":    nil,
			"StackSys":      nil,
			"Sys":           nil,
			"TotalAlloc":    nil,
			"PollCount":     nil,
			"RandomValue":   nil},
		pollInterval: pollInterval,
		pollCount:    0,
		randomVal:    randomVal(rand.Uint64()),
	}
}
