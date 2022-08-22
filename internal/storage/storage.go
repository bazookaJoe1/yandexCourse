package storage

import (
	"errors"
	"log"
	"strconv"
	"sync"
)

func StorageInit() *Storage {
	return &Storage{Mu: sync.Mutex{}, values: &servMetrics{}}
}

func (s *Storage) Save(nM string, vM string) error {
	var retErr error = nil
	s.Mu.Lock()
	switch nM {
	case "Alloc":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case Alloc): cannot parse metric value <%v>", vM)
			break
		}
		s.values.Alloc = append(s.values.Alloc, gauge(val))
	case "BuckHashSys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case BuckHashSys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.BuckHashSys = append(s.values.BuckHashSys, gauge(val))
	case "Frees":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case Frees): cannot parse metric value <%v>", vM)
			break
		}
		s.values.Frees = append(s.values.Frees, gauge(val))
	case "GCCPUFraction":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case GCCPUFraction): cannot parse metric value <%v>", vM)
			break
		}
		s.values.GCCPUFraction = append(s.values.GCCPUFraction, gauge(val))
	case "GCSys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case GCSys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.GCSys = append(s.values.GCSys, gauge(val))
	case "HeapAlloc":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case HeapAlloc): cannot parse metric value <%v>", vM)
			break
		}
		s.values.HeapAlloc = append(s.values.HeapAlloc, gauge(val))
	case "HeapIdle":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case HeapIdle): cannot parse metric value <%v>", vM)
			break
		}
		s.values.HeapIdle = append(s.values.HeapIdle, gauge(val))
	case "HeapInuse":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case HeapInuse): cannot parse metric value <%v>", vM)
			break
		}
		s.values.HeapInuse = append(s.values.HeapInuse, gauge(val))
	case "HeapObjects":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case HeapObjects): cannot parse metric value <%v>", vM)
			break
		}
		s.values.HeapObjects = append(s.values.HeapObjects, gauge(val))
	case "HeapReleased":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case HeapReleased): cannot parse metric value <%v>", vM)
			break
		}
		s.values.HeapReleased = append(s.values.HeapReleased, gauge(val))
	case "HeapSys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case HeapSys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.HeapSys = append(s.values.HeapSys, gauge(val))
	case "LastGC":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case LastGC): cannot parse metric value <%v>", vM)
			break
		}
		s.values.LastGC = append(s.values.LastGC, gauge(val))
	case "Lookups":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case Lookups): cannot parse metric value <%v>", vM)
			break
		}
		s.values.Lookups = append(s.values.Lookups, gauge(val))
	case "MCacheInuse":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case MCacheInuse): cannot parse metric value <%v>", vM)
			break
		}
		s.values.MCacheInuse = append(s.values.MCacheInuse, gauge(val))
	case "MCacheSys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case MCacheSys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.MCacheSys = append(s.values.MCacheSys, gauge(val))
	case "MSpanInuse":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case MSpanInuse): cannot parse metric value <%v>", vM)
			break
		}
		s.values.MSpanInuse = append(s.values.MSpanInuse, gauge(val))
	case "MSpanSys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case MSpanSys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.MSpanSys = append(s.values.MSpanSys, gauge(val))
	case "Mallocs":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case Mallocs): cannot parse metric value <%v>", vM)
			break
		}
		s.values.Mallocs = append(s.values.Mallocs, gauge(val))
	case "NextGC":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case NextGC): cannot parse metric value <%v>", vM)
			break
		}
		s.values.NextGC = append(s.values.NextGC, gauge(val))
	case "NumForcedGC":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case NumForcedGC): cannot parse metric value <%v>", vM)
			break
		}
		s.values.NumForcedGC = append(s.values.NumForcedGC, gauge(val))
	case "NumGC":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case NumGC): cannot parse metric value <%v>", vM)
			break
		}
		s.values.NumGC = append(s.values.NumGC, gauge(val))
	case "OtherSys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case OtherSys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.OtherSys = append(s.values.OtherSys, gauge(val))
	case "PauseTotalNs":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case PauseTotalNs): cannot parse metric value <%v>", vM)
			break
		}
		s.values.PauseTotalNs = append(s.values.PauseTotalNs, gauge(val))
	case "StackInuse":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case StackInuse): cannot parse metric value <%v>", vM)
			break
		}
		s.values.StackInuse = append(s.values.StackInuse, gauge(val))
	case "StackSys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case StackSys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.StackSys = append(s.values.StackSys, gauge(val))
	case "Sys":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case Sys): cannot parse metric value <%v>", vM)
			break
		}
		s.values.Sys = append(s.values.Sys, gauge(val))
	case "TotalAlloc":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case TotalAlloc): cannot parse metric value <%v>", vM)
			break
		}
		s.values.TotalAlloc = append(s.values.TotalAlloc, gauge(val))
	case "PollCount":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case PollCount): cannot parse metric value <%v>", vM)
			break
		}
		s.values.PollCount = append(s.values.PollCount, Counter(val))
	case "RandomValue":
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("Bad value")
			log.Printf("storage.go (case RandomValue): cannot parse metric value <%v>", vM)
			break
		}
		s.values.RandomValue = append(s.values.RandomValue, gauge(val))
	}
	s.Mu.Unlock()
	return retErr
}
func (s *Storage) Get() {}
