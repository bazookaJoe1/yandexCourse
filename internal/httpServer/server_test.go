package httpserver

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
)

func Test_parseURL(t *testing.T) {
	type args struct {
		context echo.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		{
			name:    "Correct Gauge",
			args:    args{createContextByURL("/update/gauge/DummyMetric/11.11")},
			want:    "gauge",
			want1:   "DummyMetric",
			want2:   "11.11",
			wantErr: false,
		},
		{
			name:    "Correct Counter",
			args:    args{createContextByURL("/update/counter/DummyCounter/11")},
			want:    "counter",
			want1:   "DummyCounter",
			want2:   "11",
			wantErr: false,
		},
		{
			name:    "Incorrect oversize",
			args:    args{createContextByURL("/update/gauge/DummyMetric/11.11/12")},
			want:    "",
			want1:   "",
			want2:   "",
			wantErr: true,
		},
		{
			name:    "Incorrect less size",
			args:    args{createContextByURL("/update/gauge/DummyMetric")},
			want:    "",
			want1:   "",
			want2:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := parseURL(tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseURL() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseURL() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
func createContextByURL(urlS string) echo.Context {
	return echo.New().NewContext(&http.Request{URL: &url.URL{Path: urlS}}, httptest.NewRecorder())
}
