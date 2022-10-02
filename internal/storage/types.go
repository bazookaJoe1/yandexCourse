package storage

import (
	"sync"
)

const (
	BADCOUNTER = iota
	BADGAUGE
)

type (
	Repositories interface {
		Save(TypeM, NameM, ValueM string) error
		Get() error
	}

	Storage struct {
		sync.RWMutex
		values map[string][]any
	}
)
