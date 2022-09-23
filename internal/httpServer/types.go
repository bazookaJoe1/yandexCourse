package httpserver

import (
	"net/http"
	"yandexCourse/internal/storage"
)

type (
	serverM struct {
		serverHttp *http.Server
		Storage    storage.Repositories
	}
)
