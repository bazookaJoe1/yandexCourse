package httpserver

import (
	"net/http"
	"testing"
	"yandexCourse/internal/storage"

	"github.com/labstack/echo"
)

func TestServerM_processUpdate(t *testing.T) {
	type fields struct {
		serverHTTP *http.Server
		Storage    storage.Repositories
	}
	type args struct {
		context echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ServerM := &ServerM{
				serverHTTP: tt.fields.serverHTTP,
				Storage:    tt.fields.Storage,
			}
			if err := ServerM.processUpdate(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("ServerM.processUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
