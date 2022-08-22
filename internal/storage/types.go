package storage

import (
	"sync"
)

type (
	gauge        float64
	Counter      uint64
	repositories interface {
		Save()
		Get()
	}

	Storage struct {
		Mu     sync.Mutex
		values *servMetrics
	}

	servMetrics struct {
		Alloc         []gauge
		BuckHashSys   []gauge
		Frees         []gauge
		GCCPUFraction []gauge
		GCSys         []gauge
		HeapAlloc     []gauge
		HeapIdle      []gauge
		HeapInuse     []gauge
		HeapObjects   []gauge
		HeapReleased  []gauge
		HeapSys       []gauge
		LastGC        []gauge
		Lookups       []gauge
		MCacheInuse   []gauge
		MCacheSys     []gauge
		MSpanInuse    []gauge
		MSpanSys      []gauge
		Mallocs       []gauge
		NextGC        []gauge
		NumForcedGC   []gauge
		NumGC         []gauge
		OtherSys      []gauge
		PauseTotalNs  []gauge
		StackInuse    []gauge
		StackSys      []gauge
		Sys           []gauge
		TotalAlloc    []gauge
		PollCount     []Counter
		RandomValue   []gauge
	}
)
