package stats

import (
	"math"
	"math/rand"
	"testing"
)

func Test_exponential_Mean(t *testing.T) {
	tests := []struct {
		name string
		exp  Exponential
		want float64
	}{
		{"NaN case lambda", Exponential{math.NaN()}, math.NaN()},
		{"Normal case", Exponential{0.5}, 2.0},
		{"Normal case 2", Exponential{1.5}, 0.6666667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.exp.Mean()
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Mean() = %v, want %v", got, tt.want)
				}
			} else if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkExponentialMean(seed1 int64, b *testing.B) {
	exp := Exponential{}
	rand.Seed(seed1)
	exp.Lambda = rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exp.Mean()
	}
}

func Benchmark_exponential_Mean1(b *testing.B) { benchmarkExponentialMean(0, b) }
func Benchmark_exponential_Mean2(b *testing.B) { benchmarkExponentialMean(25, b) }
func Benchmark_exponential_Mean3(b *testing.B) { benchmarkExponentialMean(15, b) }
func Benchmark_exponential_Mean4(b *testing.B) { benchmarkExponentialMean(3, b) }
func Benchmark_exponential_Mean5(b *testing.B) { benchmarkExponentialMean(1e3, b) }
func Benchmark_exponential_Mean6(b *testing.B) { benchmarkExponentialMean(534, b) }

func Test_exponential_StdDev(t *testing.T) {
	tests := []struct {
		name string
		exp  Exponential
		want float64
	}{
		{"NaN case lambda", Exponential{math.NaN()}, math.NaN()},
		{"Normal case", Exponential{0.5}, 2.0},
		{"Normal case 2", Exponential{1.5}, 0.6666667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.exp.StdDev()
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("StdDev() = %v, want %v", got, tt.want)
				}
			} else if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("StdDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkExponentialStdDev(seed1 int64, b *testing.B) {
	exp := Exponential{}
	rand.Seed(seed1)
	exp.Lambda = rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exp.StdDev()
	}
}

func Benchmark_exponential_StdDev1(b *testing.B) { benchmarkExponentialStdDev(0, b) }
func Benchmark_exponential_StdDev2(b *testing.B) { benchmarkExponentialStdDev(25, b) }
func Benchmark_exponential_StdDev3(b *testing.B) { benchmarkExponentialStdDev(15, b) }
func Benchmark_exponential_StdDev4(b *testing.B) { benchmarkExponentialStdDev(3, b) }
func Benchmark_exponential_StdDev5(b *testing.B) { benchmarkExponentialStdDev(1e3, b) }
func Benchmark_exponential_StdDev6(b *testing.B) { benchmarkExponentialStdDev(534, b) }

func Test_exponential_Variance(t *testing.T) {
	tests := []struct {
		name string
		exp  Exponential
		want float64
	}{
		{"NaN case lambda", Exponential{math.NaN()}, math.NaN()},
		{"Normal case", Exponential{0.5}, 4.0},
		{"Normal case 2", Exponential{1.5}, 0.444444},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.exp.Variance()
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Variance() = %v, want %v", got, tt.want)
				}
			} else if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkExponentialVariance(seed1 int64, b *testing.B) {
	exp := Exponential{}
	rand.Seed(seed1)
	exp.Lambda = rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exp.Variance()
	}
}

func Benchmark_exponential_Variance1(b *testing.B) { benchmarkExponentialVariance(0, b) }
func Benchmark_exponential_Variance2(b *testing.B) { benchmarkExponentialVariance(25, b) }
func Benchmark_exponential_Variance3(b *testing.B) { benchmarkExponentialVariance(15, b) }
func Benchmark_exponential_Variance4(b *testing.B) { benchmarkExponentialVariance(3, b) }
func Benchmark_exponential_Variance5(b *testing.B) { benchmarkExponentialVariance(1e3, b) }
func Benchmark_exponential_Variance6(b *testing.B) { benchmarkExponentialVariance(534, b) }
