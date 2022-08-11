package Metricscollector

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type (
	Counter uint64
	second  uint64
	Metrics struct {
		Mu           sync.Mutex
		CanMetrics   *runtime.MemStats
		Pollcount    Counter
		RandomValue  uint64
		pollInterval second
	}
)

func (m *Metrics) ReadMetrics() error {
	var err = errors.New("error")

	tempMet := runtime.MemStats{}

	runtime.ReadMemStats(&tempMet)
	if &tempMet == nil {
		return err
	}
	m.Mu.Lock()
	*m.CanMetrics = tempMet
	m.RandomValue = rand.Uint64()
	m.countTick()
	m.Mu.Unlock()

	return nil
}

func (m *Metrics) countTick() {
	if m.Pollcount == 18446744073709551615 {
		m.Pollcount = 0
	} else {
		m.Pollcount++
	}
}

func (m *Metrics) pollTickerInit() *time.Ticker {
	return time.NewTicker(time.Duration(m.pollInterval * second(time.Second)))
}

func MetricsCollectorMain(mCtx context.Context, m *Metrics, retChan chan struct{}) {
	//mCtx, cancel := context.WithCancel(ctx)
	defer func() { close(retChan) }()

	ticker := m.pollTickerInit()

	for {
		select {
		case <-mCtx.Done():
			log.Printf("---> cancelling MetricsCollectorMain context")
			return
		case <-ticker.C:
			//MetricsStru.Mu.Lock()
			m.ReadMetrics()
			//MetricsStru.Mu.Unlock()
			log.Printf("---> Metrics collected, Pollcount=%d", m.Pollcount)
		default:
			continue
		}
	}
}

func MetricsInit() *Metrics {
	return &Metrics{CanMetrics: &runtime.MemStats{}, pollInterval: 2, Pollcount: 0, RandomValue: rand.Uint64(), Mu: sync.Mutex{}}
}
