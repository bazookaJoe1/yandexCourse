package httpserver

import (
	"net/http"
	"yandexCourse/internal/storage"
)

type (
	serverM struct {
		serverHTTP *http.Server
		Storage    storage.Repositories
	}
)
