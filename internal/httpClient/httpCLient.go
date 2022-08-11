package httpclient

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	metricscollector "yandexCourse/internal/metricsCollector"
)

type (
	second uint64
	agent  struct {
		client       *http.Client
		metrMap      map[string]interface{}
		request      []*http.Request
		pollInterval second
	}
)

func mapInit() map[string]interface{} {
	metrMap := map[string]interface{}{
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
		"RandomValue":   nil}
	return metrMap
}

func clientInit() *http.Client {
	return &http.Client{Timeout: 1 * time.Second}
}

func (a *agent) updateMap(metrics *metricscollector.Metrics) {
	metrics.Mu.Lock()
	a.metrMap["Alloc"] = metrics.CanMetrics.Alloc
	a.metrMap["BuckHashSys"] = metrics.CanMetrics.BuckHashSys
	a.metrMap["Frees"] = metrics.CanMetrics.Frees
	a.metrMap["GCCPUFraction"] = metrics.CanMetrics.GCCPUFraction
	a.metrMap["GCSys"] = metrics.CanMetrics.GCSys
	a.metrMap["HeapAlloc"] = metrics.CanMetrics.HeapAlloc
	a.metrMap["HeapIdle"] = metrics.CanMetrics.HeapIdle
	a.metrMap["HeapInuse"] = metrics.CanMetrics.HeapInuse
	a.metrMap["HeapObjects"] = metrics.CanMetrics.HeapObjects
	a.metrMap["HeapReleased"] = metrics.CanMetrics.HeapReleased
	a.metrMap["HeapSys"] = metrics.CanMetrics.HeapSys
	a.metrMap["LastGC"] = metrics.CanMetrics.LastGC
	a.metrMap["Lookups"] = metrics.CanMetrics.Lookups
	a.metrMap["MCacheInuse"] = metrics.CanMetrics.MCacheInuse
	a.metrMap["MCacheSys"] = metrics.CanMetrics.MCacheSys
	a.metrMap["MSpanInuse"] = metrics.CanMetrics.MSpanInuse
	a.metrMap["MSpanSys"] = metrics.CanMetrics.MSpanSys
	a.metrMap["Mallocs"] = metrics.CanMetrics.Mallocs
	a.metrMap["NextGC"] = metrics.CanMetrics.NextGC
	a.metrMap["NumForcedGC"] = metrics.CanMetrics.NumForcedGC
	a.metrMap["NumGC"] = metrics.CanMetrics.NumGC
	a.metrMap["OtherSys"] = metrics.CanMetrics.OtherSys
	a.metrMap["PauseTotalNs"] = metrics.CanMetrics.PauseTotalNs
	a.metrMap["StackInuse"] = metrics.CanMetrics.StackInuse
	a.metrMap["StackSys"] = metrics.CanMetrics.StackSys
	a.metrMap["Sys"] = metrics.CanMetrics.Sys
	a.metrMap["TotalAlloc"] = metrics.CanMetrics.TotalAlloc
	a.metrMap["PollCount"] = metrics.Pollcount
	a.metrMap["RandomValue"] = metrics.RandomValue
	metrics.Mu.Unlock()

}

func typeCheck(i interface{}) int {
	if _, ok := i.(uint64); ok {
		return 0
	} else if _, ok := i.(float64); ok {
		return 1
	} else if _, ok := i.(metricscollector.Counter); ok {
		return 2
	}
	return -1
}

func (a *agent) createReqs(metrics *metricscollector.Metrics) {
	a.updateMap(metrics)
	for key, value := range a.metrMap {
		flag := typeCheck(value)
		switch flag {
		case 0:
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:8080/update/gauge/%s/%d", key, value), nil)
			a.request = append(a.request, req)
		case 1:
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:8080/update/gauge/%s/%f", key, value), nil)
			a.request = append(a.request, req)
		case 2:
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:8080/update/counter/%s/%d", key, value), nil)
			a.request = append(a.request, req)
		}
	}
}

func (a *agent) sendReq(metrics *metricscollector.Metrics) {
	a.createReqs(metrics)

	for _, req := range a.request {
		a.client.Do(req)
	}
}
func AgentInit() *agent {
	return &agent{}
}

func (a *agent) Run(httpCtx context.Context, retChan chan struct{}, metrics *metricscollector.Metrics) {
	//httpCtx, cancel := context.WithCancel(ctx)
	defer func() { close(retChan) }()

	a.client = clientInit()
	a.metrMap = mapInit()
	a.pollInterval = second(10 * time.Second)

	ticker := time.NewTicker(time.Duration(a.pollInterval))

	for {
		select {
		case <-httpCtx.Done():
			log.Printf("---> cancelling httpClient context")
			return
		case <-ticker.C:
			log.Printf("---> trying to send metrics")
			a.sendReq(metrics)
		}
	}

}
