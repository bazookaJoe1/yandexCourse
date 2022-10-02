package httpclient

import "net/http"

type (
	second uint64
	agent  struct {
		client       *http.Client
		request      []*http.Request
		pollInterval second //timeout between sending to server
	}
)
