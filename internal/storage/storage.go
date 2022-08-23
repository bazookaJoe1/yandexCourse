package storage

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func StorageInit() *Storage {
	return &Storage{Mu: sync.Mutex{}, values: make(map[string][]fmt.Stringer)}
}

func (s *Storage) Save(tM string, nM string, vM string) error {
	var retErr error = nil
	s.Mu.Lock()

	if tM == "gauge" {
		val, err := strconv.ParseFloat(vM, 64)
		if err != nil {
			retErr = errors.New("bad value")
			log.Printf("storage.go (metrics %s): cannot parse metric value <%v>", nM, vM)
		}
		s.values[nM] = append(s.values[nM], gauge(val))
	} else if tM == "counter" {
		val, err := strconv.ParseUint(vM, 10, 64)
		if err != nil {
			retErr = errors.New("bad value")
			log.Printf("storage.go (metrics %s): cannot parse metric value <%v>", nM, vM)
		}
		s.values[nM] = append(s.values[nM], Counter(val))
	} else {
		retErr = errors.New("not implemented")
	}
	s.Mu.Unlock()
	return retErr
}
func (s *Storage) Get() {}
