package httpserver

import (
	"net/http"
	"sync"
	"yandexCourse/internal/storage"
)

type (
	handlerUpdate struct {
		resp  []byte
		servM *storage.Storage
	}
	server struct {
		Mu      sync.Mutex
		server  *http.Server
		handler *handlerUpdate
	}
	Middleware func(http.Handler) http.Handler
)
