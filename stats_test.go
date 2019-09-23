package stats

import "testing"

func TestMean(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty", args{[]float64{}}, 0},
		{"equal", args{[]float64{5.0, 5.0, 5.0}}, 5.0},
		{"different", args{[]float64{5.0, 10.0, 0.0}}, 5.0},
		{"different2", args{[]float64{5.0, 10.0}}, 7.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.input); got != tt.want {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}
