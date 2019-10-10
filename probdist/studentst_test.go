package stats

import "testing"

func TestGetTStatistic(t *testing.T) {
	type args struct {
		v     float64
		alpha float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case 1", args{1.0, 0.1}, 3.078},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTStatistic(tt.args.v, tt.args.alpha); got != tt.want {
				t.Errorf("GetTStatistic() = %v, want %v", got, tt.want)
			}
		})
	}
}
