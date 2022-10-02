package storage

import (
	"sync"
	"testing"
)

func TestStorage_Save(t *testing.T) {
	type args struct {
		TypeM  string
		NameM  string
		ValueM string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name: "Correct Gauge",
			args: args{
				TypeM:  "gauge",
				NameM:  "DummyGauge",
				ValueM: "11.11",
			},
			wantErr: false,
		},
		{
			name: "Correct Counter",
			args: args{
				TypeM:  "counter",
				NameM:  "DummyCounter",
				ValueM: "11",
			},
			wantErr: false,
		},
		{
			name: "Incorrect value",
			args: args{
				TypeM:  "gauge",
				NameM:  "DummyGauge",
				ValueM: "none",
			},
			wantErr: true,
		},
		{
			name: "Incorrect type",
			args: args{
				TypeM:  "none",
				NameM:  "DummyGauge",
				ValueM: "11.11",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				RWMutex: sync.RWMutex{},
				values:  make(map[string]any),
			}
			if err := s.Save(tt.args.TypeM, tt.args.NameM, tt.args.ValueM); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
