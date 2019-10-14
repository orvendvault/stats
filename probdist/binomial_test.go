package stats

import (
	"math"
	"math/rand"
	"testing"
)

func Test_binomial_Mean(t *testing.T) {
	tests := []struct {
		name string
		bin  Binomial
		want float64
	}{
		{"NaN case n", Binomial{math.NaN(), 0.25}, math.NaN()},
		{"NaN case p", Binomial{25, math.NaN()}, math.NaN()},
		{"Normal case", Binomial{20, 0.25}, 5.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bin.Mean()
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Mean() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkBinomialMean(seed1 int64, seed2 int64, b *testing.B) {
	bin := Binomial{}
	rand.Seed(seed1)
	bin.N = rand.Float64() * 100
	rand.Seed(seed2)
	bin.P = rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bin.Mean()
	}
}

func Benchmark_binomial_Mean1(b *testing.B) { benchmarkBinomialMean(0, 10, b) }
func Benchmark_binomial_Mean2(b *testing.B) { benchmarkBinomialMean(25, 2, b) }
func Benchmark_binomial_Mean3(b *testing.B) { benchmarkBinomialMean(15, 365, b) }
func Benchmark_binomial_Mean4(b *testing.B) { benchmarkBinomialMean(3, 8, b) }
func Benchmark_binomial_Mean5(b *testing.B) { benchmarkBinomialMean(1e3, 1, b) }
func Benchmark_binomial_Mean6(b *testing.B) { benchmarkBinomialMean(534, 82, b) }

func Test_binomial_StdDev(t *testing.T) {
	tests := []struct {
		name string
		bin  Binomial
		want float64
	}{
		{"NaN case n", Binomial{math.NaN(), 0.25}, math.NaN()},
		{"NaN case p", Binomial{25, math.NaN()}, math.NaN()},
		{"Normal case", Binomial{20, 0.25}, 1.936491},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bin.StdDev()
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

func benchmarkBinomialStdDev(seed1 int64, seed2 int64, b *testing.B) {
	bin := Binomial{}
	rand.Seed(seed1)
	bin.N = rand.Float64() * 100
	rand.Seed(seed2)
	bin.P = rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bin.StdDev()
	}
}

func Benchmark_binomial_StdDev1(b *testing.B) { benchmarkBinomialStdDev(0, 10, b) }
func Benchmark_binomial_StdDev2(b *testing.B) { benchmarkBinomialStdDev(25, 2, b) }
func Benchmark_binomial_StdDev3(b *testing.B) { benchmarkBinomialStdDev(15, 365, b) }
func Benchmark_binomial_StdDev4(b *testing.B) { benchmarkBinomialStdDev(3, 8, b) }
func Benchmark_binomial_StdDev5(b *testing.B) { benchmarkBinomialStdDev(1e3, 1, b) }
func Benchmark_binomial_StdDev6(b *testing.B) { benchmarkBinomialStdDev(534, 82, b) }

func Test_binomial_Variance(t *testing.T) {
	tests := []struct {
		name string
		bin  Binomial
		want float64
	}{
		{"NaN case n", Binomial{math.NaN(), 0.25}, math.NaN()},
		{"NaN case p", Binomial{25, math.NaN()}, math.NaN()},
		{"Normal case", Binomial{20, 0.25}, 3.75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bin.Variance()
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Variance() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkBinomialVariance(seed1 int64, seed2 int64, b *testing.B) {
	bin := Binomial{}
	rand.Seed(seed1)
	bin.N = rand.Float64() * 100
	rand.Seed(seed2)
	bin.P = rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bin.Variance()
	}
}

func Benchmark_binomial_Variance1(b *testing.B) { benchmarkBinomialVariance(0, 10, b) }
func Benchmark_binomial_Variance2(b *testing.B) { benchmarkBinomialVariance(25, 2, b) }
func Benchmark_binomial_Variance3(b *testing.B) { benchmarkBinomialVariance(15, 365, b) }
func Benchmark_binomial_Variance4(b *testing.B) { benchmarkBinomialVariance(3, 8, b) }
func Benchmark_binomial_Variance5(b *testing.B) { benchmarkBinomialVariance(1e3, 1, b) }
func Benchmark_binomial_Variance6(b *testing.B) { benchmarkBinomialVariance(534, 82, b) }
