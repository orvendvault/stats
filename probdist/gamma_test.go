package stats

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func Test_gamma_Mean(t *testing.T) {
	tests := []struct {
		name string
		g    Gamma
		want float64
	}{
		{"NaN case k", Gamma{math.NaN(), 2.0}, math.NaN()},
		{"NaN case theta", Gamma{2.0, math.NaN()}, math.NaN()},
		{"Normal case", Gamma{5.0, 1.0}, 5.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.Mean()
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

func benchmarkGammaMean(seed1 int64, seed2 int64, b *testing.B) {
	g := Gamma{}
	rand.Seed(seed1)
	g.K = rand.Float64() * 10
	rand.Seed(seed2)
	g.Theta = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Mean()
	}
}

func Benchmark_gamma_Mean1(b *testing.B) { benchmarkGammaMean(0, 10, b) }
func Benchmark_gamma_Mean2(b *testing.B) { benchmarkGammaMean(25, 2, b) }
func Benchmark_gamma_Mean3(b *testing.B) { benchmarkGammaMean(15, 365, b) }
func Benchmark_gamma_Mean4(b *testing.B) { benchmarkGammaMean(3, 8, b) }
func Benchmark_gamma_Mean5(b *testing.B) { benchmarkGammaMean(1e3, 1, b) }
func Benchmark_gamma_Mean6(b *testing.B) { benchmarkGammaMean(534, 82, b) }

func Test_gamma_StdDev(t *testing.T) {
	tests := []struct {
		name string
		g    Gamma
		want float64
	}{
		{"NaN case k", Gamma{math.NaN(), 2.0}, math.NaN()},
		{"NaN case theta", Gamma{2.0, math.NaN()}, math.NaN()},
		{"Normal case", Gamma{5.0, 1.0}, 2.236067},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.StdDev()
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

func benchmarkGammaStdDev(seed1 int64, seed2 int64, b *testing.B) {
	g := Gamma{}
	rand.Seed(seed1)
	g.K = rand.Float64() * 10
	rand.Seed(seed2)
	g.Theta = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.StdDev()
	}
}

func Benchmark_gamma_StdDev1(b *testing.B) { benchmarkGammaStdDev(0, 10, b) }
func Benchmark_gamma_StdDev2(b *testing.B) { benchmarkGammaStdDev(25, 2, b) }
func Benchmark_gamma_StdDev3(b *testing.B) { benchmarkGammaStdDev(15, 365, b) }
func Benchmark_gamma_StdDev4(b *testing.B) { benchmarkGammaStdDev(3, 8, b) }
func Benchmark_gamma_StdDev5(b *testing.B) { benchmarkGammaStdDev(1e3, 1, b) }
func Benchmark_gamma_StdDev6(b *testing.B) { benchmarkGammaStdDev(534, 82, b) }

func Test_gamma_Variance(t *testing.T) {
	tests := []struct {
		name string
		g    Gamma
		want float64
	}{
		{"NaN case k", Gamma{math.NaN(), 2.0}, math.NaN()},
		{"NaN case theta", Gamma{2.0, math.NaN()}, math.NaN()},
		{"Normal case", Gamma{5.0, 1.0}, 5.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.Variance()
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

func benchmarkGammaVariance(seed1 int64, seed2 int64, b *testing.B) {
	g := Gamma{}
	rand.Seed(seed1)
	g.K = rand.Float64() * 10
	rand.Seed(seed2)
	g.Theta = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.StdDev()
	}
}

func Benchmark_gamma_Variance1(b *testing.B) { benchmarkGammaVariance(0, 10, b) }
func Benchmark_gamma_Variance2(b *testing.B) { benchmarkGammaVariance(25, 2, b) }
func Benchmark_gamma_Variance3(b *testing.B) { benchmarkGammaVariance(15, 365, b) }
func Benchmark_gamma_Variance4(b *testing.B) { benchmarkGammaVariance(3, 8, b) }
func Benchmark_gamma_Variance5(b *testing.B) { benchmarkGammaVariance(1e3, 1, b) }
func Benchmark_gamma_Variance6(b *testing.B) { benchmarkGammaVariance(534, 82, b) }

func TestNewGamma(t *testing.T) {
	type args struct {
		k     float64
		theta float64
	}
	tests := []struct {
		name    string
		args    args
		want    Gamma
		wantErr bool
	}{
		{"Normal case", args{2.0, 1.0}, Gamma{2.0, 1.0}, false},
		{"Invalid zero K case", args{0.0, 2.0}, Gamma{1.0, 2.0}, true},
		{"Invalid zero Theta case", args{1.0, 0.0}, Gamma{1.0, 2.0}, true},
		{"Invalid negative K case", args{-1.0, 2.0}, Gamma{1.0, 2.0}, true},
		{"Invalid negative Theta case", args{1.0, -5.0}, Gamma{1.0, 2.0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGamma(tt.args.k, tt.args.theta)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGamma() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGamma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewGamma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGamma(2.0, 1.5)
	}
}
