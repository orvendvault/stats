package stats

import (
	"math"
	"math/rand"
	"testing"
)

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

func benchmarkMean(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mean(s)
	}
}

func BenchmarkMean1(b *testing.B) { benchmarkMean(1, b) }
func BenchmarkMean2(b *testing.B) { benchmarkMean(2, b) }
func BenchmarkMean3(b *testing.B) { benchmarkMean(10, b) }
func BenchmarkMean4(b *testing.B) { benchmarkMean(1e3, b) }
func BenchmarkMean5(b *testing.B) { benchmarkMean(1e6, b) }

func TestMax(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty case", args{[]float64{}}, math.NaN()},
		{"equal case", args{[]float64{5.0, 5.0, 5.0, 5.0}}, 5.0},
		{"normal case", args{[]float64{1.0, 5.0, 2.0, 6.0, 1.0, 2.0, 3.0}}, 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Max() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty case", args{[]float64{}}, math.NaN()},
		{"equal case", args{[]float64{5.0, 5.0, 5.0, 5.0}}, 5.0},
		{"normal case", args{[]float64{1.0, 5.0, 2.0, 6.0, 1.0, 2.0, 3.0}}, 1.0},
	}
	for _, tt := range tests {
		got := Min(tt.args.input)
		if math.IsNaN(got) || math.IsNaN(tt.want) {
			if !math.IsNaN(got) || !math.IsNaN(tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		} else if got != tt.want {
			t.Errorf("Min() = %v, want %v", got, tt.want)
		}
	}
}

func TestRange(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty case", args{[]float64{}}, math.NaN()},
		{"equal case", args{[]float64{5.0, 5.0, 5.0, 5.0}}, 0.0},
		{"normal case", args{[]float64{1.0, 5.0, 2.0, 6.0, 1.0, 2.0, 3.0}}, 5.0},
	}
	for _, tt := range tests {
		got := Range(tt.args.input)
		if math.IsNaN(got) || math.IsNaN(tt.want) {
			if !math.IsNaN(got) || !math.IsNaN(tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		} else if got != tt.want {
			t.Errorf("Range() = %v, want %v", got, tt.want)
		}
	}
}

func TestStandardDeviationSample(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty case", args{[]float64{}}, math.NaN()},
		{"single value case", args{[]float64{1.0}}, math.NaN()},
		{"normal case", args{[]float64{1.0, 2.0, 3.0, 2.3, 1.4, 1.7, 1.5, 1.5, 1.8, 2.6, 2.3, 2.0, 2.2}}, 0.542548},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDeviationSample(tt.args.input); math.Abs(got-tt.want) > 1e6 {
				t.Errorf("StandardDeviationSample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarianceSample(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty case", args{[]float64{}}, math.NaN()},
		{"single value case", args{[]float64{1.0}}, math.NaN()},
		{"normal case", args{[]float64{1.0, 2.0, 3.0, 2.3, 1.4, 1.7, 1.5, 1.5, 1.8, 2.6, 2.3, 2.0, 2.2}}, 0.294358},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VarianceSample(tt.args.input); math.Abs(got-tt.want) > 1e6 {
				t.Errorf("VarianceSample() = %v, want %v", got, tt.want)
			}
		})
	}
}
