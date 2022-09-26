package metricscollector

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestMetrics_ReturnMap(t *testing.T) {
	type fields struct {
		Mu           sync.Mutex
		values       map[string]fmt.Stringer
		pollCount    Counter
		randomVal    randomVal
		pollInterval second
	}
	tests := []struct {
		name   string
		fields fields
		want   *map[string]fmt.Stringer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metrics{
				Mu:           tt.fields.Mu,
				values:       tt.fields.values,
				pollCount:    tt.fields.pollCount,
				randomVal:    tt.fields.randomVal,
				pollInterval: tt.fields.pollInterval,
			}
			if got := m.ReturnMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Metrics.ReturnMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
