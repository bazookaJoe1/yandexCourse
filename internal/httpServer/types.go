package httpserver

import (
	"net/http"
	"yandexCourse/internal/storage"
)

type (
	ServerM struct {
		serverHTTP *http.Server
		Storage    storage.Repositories
	}
)
