package storage

import (
	"fmt"
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
		values map[string][]fmt.Stringer
	}
)

func (g gauge) String() string {
	return fmt.Sprintf("%v", float64(g))
}

func (c Counter) String() string {
	return fmt.Sprintf("%v", uint64(c))
}
