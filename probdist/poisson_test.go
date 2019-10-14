package stats

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func Test_poisson_Mean(t *testing.T) {
	tests := []struct {
		name string
		p    Poisson
		want float64
	}{
		{"NaN case", Poisson{math.NaN()}, math.NaN()},
		{"Normal case", Poisson{1.0}, 1.0},
		{"Normal case 2", Poisson{2.0}, 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Mean()
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

func benchmarkPoissonMean(seed1 int64, b *testing.B) {
	poi := Exponential{}
	rand.Seed(seed1)
	poi.Lambda = rand.Float64() * 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		poi.Mean()
	}
}

func Benchmark_poisson_Mean1(b *testing.B) { benchmarkPoissonMean(0, b) }
func Benchmark_poisson_Mean2(b *testing.B) { benchmarkPoissonMean(25, b) }
func Benchmark_poisson_Mean3(b *testing.B) { benchmarkPoissonMean(15, b) }
func Benchmark_poisson_Mean4(b *testing.B) { benchmarkPoissonMean(3, b) }
func Benchmark_poisson_Mean5(b *testing.B) { benchmarkPoissonMean(1e3, b) }
func Benchmark_poisson_Mean6(b *testing.B) { benchmarkPoissonMean(534, b) }

func Test_poisson_StdDev(t *testing.T) {
	tests := []struct {
		name string
		p    Poisson
		want float64
	}{
		{"NaN case lambda", Poisson{math.NaN()}, math.NaN()},
		{"Normal case", Poisson{1.0}, 1.0},
		{"Normal case 2", Poisson{2.0}, 1.414213},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.StdDev()
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

func benchmarkPoissonStdDev(seed1 int64, b *testing.B) {
	poi := Exponential{}
	rand.Seed(seed1)
	poi.Lambda = rand.Float64() * 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		poi.StdDev()
	}
}

func Benchmark_poisson_StdDev1(b *testing.B) { benchmarkPoissonStdDev(0, b) }
func Benchmark_poisson_StdDev2(b *testing.B) { benchmarkPoissonStdDev(25, b) }
func Benchmark_poisson_StdDev3(b *testing.B) { benchmarkPoissonStdDev(15, b) }
func Benchmark_poisson_StdDev4(b *testing.B) { benchmarkPoissonStdDev(3, b) }
func Benchmark_poisson_StdDev5(b *testing.B) { benchmarkPoissonStdDev(1e3, b) }
func Benchmark_poisson_StdDev6(b *testing.B) { benchmarkPoissonStdDev(534, b) }

func Test_poisson_Variance(t *testing.T) {
	tests := []struct {
		name string
		p    Poisson
		want float64
	}{
		{"NaN case", Poisson{math.NaN()}, math.NaN()},
		{"Normal case", Poisson{1.0}, 1.0},
		{"Normal case 2", Poisson{2.0}, 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Variance()
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

func benchmarkPoissonVariance(seed1 int64, b *testing.B) {
	poi := Exponential{}
	rand.Seed(seed1)
	poi.Lambda = rand.Float64() * 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		poi.Variance()
	}
}

func Benchmark_poisson_Variance1(b *testing.B) { benchmarkPoissonVariance(0, b) }
func Benchmark_poisson_Variance2(b *testing.B) { benchmarkPoissonVariance(25, b) }
func Benchmark_poisson_Variance3(b *testing.B) { benchmarkPoissonVariance(15, b) }
func Benchmark_poisson_Variance4(b *testing.B) { benchmarkPoissonVariance(3, b) }
func Benchmark_poisson_Variance5(b *testing.B) { benchmarkPoissonVariance(1e3, b) }
func Benchmark_poisson_Variance6(b *testing.B) { benchmarkPoissonVariance(534, b) }

func TestNewPoisson(t *testing.T) {
	type args struct {
		lambda float64
	}
	tests := []struct {
		name    string
		args    args
		want    Poisson
		wantErr bool
	}{
		{"Normal case", args{1.0}, Poisson{1.0}, false},
		{"Invalid negative case", args{-1.0}, Poisson{1.0}, true},
		{"Invalid zero case", args{0.0}, Poisson{1.0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPoisson(tt.args.lambda)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPoisson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoisson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewPoisson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPoisson(1.0)
	}
}
