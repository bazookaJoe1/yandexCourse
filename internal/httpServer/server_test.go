package httpserver

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestParseURL(t *testing.T) {
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Invalid value Counter",
			args: args{
				url: &url.URL{
					Path: "/update/counter/testCounter/none",
				},
			},
			want:    []string{"", "update", "counter", "testCounter", "none"},
			wantErr: false,
		},
		{
			name: "Invalid value gauge",
			args: args{
				url: &url.URL{
					Path: "/update/gauge/testGauge/none",
				},
			},
			want:    []string{"", "update", "gauge", "testGauge", "none"},
			wantErr: false,
		},
		{
			name: "Invalid Path",
			args: args{
				url: &url.URL{
					Path: "/update/gauge/testGauge",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serverREST(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Invalid path counter",
			args: args{
				url: "http://localhost:8080/update/counter/testCounter/none",
			},
			want:    &http.Response{StatusCode: 404},
			wantErr: false,
		},
		{
			name: "Invalid path gauge",
			args: args{
				url: "http://localhost:8080/update/counter/testGauge/none",
			},
			want:    &http.Response{StatusCode: 404},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//POST requests to server
			response, err := http.Post(tt.args.url, "text/plain", nil)

			//we never expect error while server is working
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//Check status code returned by server
			if !reflect.DeepEqual(response.StatusCode, tt.want.StatusCode) {
				t.Errorf("ParseURL() = %v, want %v", response.StatusCode, tt.want.StatusCode)
			}
		})
	}
}
