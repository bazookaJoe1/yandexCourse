package httpclient

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	metricscollector "yandexCourse/internal/metricsCollector"
)

// return client with timeout
func clientInit(s second) *http.Client {
	return &http.Client{Timeout: time.Duration(s) * time.Second}
}

// create *http.Request with type <counter> or <gauge>
func (a *agent) createReqs(m *metricscollector.Metrics) {
	m.Mu.Lock()
	for key, value := range *m.ReturnMap() {
		switch _, ok := value.(metricscollector.Counter); ok {
		case true:
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:8080/update/counter/%v/%d", key, value), nil)
			a.request = append(a.request, req)
		case false:
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:8080/update/gauge/%v/%.0f", key, value), nil)
			a.request = append(a.request, req)
		}
	}
	m.Mu.Unlock()
}

func (a *agent) sendReq(m *metricscollector.Metrics) {
	a.createReqs(m)

	for _, req := range a.request {
		var body []byte
		resp, err := a.client.Do(req)
		if err != nil {
			log.Printf("sendReq: error <%v>", err)
		}
		if resp != nil {
			_, _ = resp.Body.Read(body)
			log.Printf("sendReq: Got Response. Status: %v, Header: %v, Body: %v", resp.Status, resp.Header.Get("Content-Type"), body)
			resp.Body.Close()
		}
	}
}
func AgentInit(s second) *agent {
	return &agent{pollInterval: s}
}

func (a *agent) Run(httpCtx context.Context, retChan chan struct{}, m *metricscollector.Metrics) {
	defer func() { close(retChan) }()

	a.client = clientInit(1)

	ticker := time.NewTicker(time.Duration(a.pollInterval) * time.Second)

	for {
		select {
		case <-httpCtx.Done():
			log.Printf("---> cancelling httpClient context")
			return
		case <-ticker.C:
			log.Printf("---> trying to send metrics")
			a.sendReq(m)
		}
	}
}
