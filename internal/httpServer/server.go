package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"yandexCourse/internal/storage"
)

func (h handlerUpdate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler: got request <%v>", r.Header)
	switch r.Method {
	case "POST":
		path, err := ParseURL(r.URL)
		if err != nil {
			log.Printf("%s", err)
			http.Error(w, "WrongUrl", http.StatusNotFound)
			return
		}

		typeM := path[2]
		nameM := path[3]
		valueM := path[4]

		err = h.servM.Save(typeM, nameM, valueM)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})

	default:
		http.Error(w, "WrongMethod", http.StatusBadRequest)
	}

}
func ParseURL(url *url.URL) ([]string, error) {
	splitted := strings.Split(url.Path, "/")
	if len(splitted) != 5 {
		return nil, fmt.Errorf("error url request: %s/%s", url.Host, url.Path)
	}
	return splitted, nil
}

func ServerInit() *server {
	return &server{
		Mu:      sync.Mutex{},
		server:  &http.Server{Addr: "localhost:8080"},
		handler: &handlerUpdate{servM: storage.StorageInit()},
	}
}

func (s *server) RegisterHadnlers() {
	http.Handle("/update/", s.handler)
}

func (s *server) Run() {
	s.server.ListenAndServe()
}
