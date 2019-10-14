package stats

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func Test_normal_CDF(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		norm Normal
		args args
		want float64
	}{
		{"NaN case mu", Normal{math.NaN(), 0.25}, args{0.2}, math.NaN()},
		{"NaN case sigma", Normal{0.0, math.NaN()}, args{0.2}, math.NaN()},
		{"Normal case", Normal{0.0, 1.0}, args{0.2}, 0.579259},
		{"Normal case", Normal{0.0, 0.1}, args{0.174}, 0.959070},
		{"Median case", Normal{0.0, 1.0}, args{0.0}, 0.5},
		{"Negative input case", Normal{0.0, 1.0}, args{-0.2}, 0.420741},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.norm.CDF(tt.args.x)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("CDF() = %v, want %v", got, tt.want)
				}
			} else if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("CDF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkNormalCDF(seed1 int64, seed2 int64, b *testing.B) {
	norm := Normal{}
	rand.Seed(seed1)
	norm.Mu = rand.Float64() * 10
	rand.Seed(seed2)
	norm.Sigma = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		norm.CDF(norm.Mu + rand.Float64()*3)
	}
}

func Benchmark_normal_CDF1(b *testing.B) { benchmarkNormalCDF(0, 10, b) }
func Benchmark_normal_CDF2(b *testing.B) { benchmarkNormalCDF(25, 2, b) }
func Benchmark_normal_CDF3(b *testing.B) { benchmarkNormalCDF(15, 365, b) }
func Benchmark_normal_CDF4(b *testing.B) { benchmarkNormalCDF(3, 8, b) }
func Benchmark_normal_CDF5(b *testing.B) { benchmarkNormalCDF(1e3, 1, b) }
func Benchmark_normal_CDF6(b *testing.B) { benchmarkNormalCDF(534, 82, b) }

func Test_normal_Mean(t *testing.T) {
	tests := []struct {
		name string
		norm Normal
		want float64
	}{
		{"NaN case mu", Normal{math.NaN(), 1.0}, math.NaN()},
		{"NaN case sigma", Normal{0.0, math.NaN()}, 0.0},
		{"Normal case", Normal{0.0, 1.0}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.norm.Mean()
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

func benchmarkNormalMean(seed1 int64, seed2 int64, b *testing.B) {
	norm := Normal{}
	rand.Seed(seed1)
	norm.Mu = rand.Float64() * 10
	rand.Seed(seed2)
	norm.Sigma = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		norm.Mean()
	}
}

func Benchmark_normal_Mean1(b *testing.B) { benchmarkNormalMean(0, 10, b) }
func Benchmark_normal_Mean2(b *testing.B) { benchmarkNormalMean(25, 2, b) }
func Benchmark_normal_Mean3(b *testing.B) { benchmarkNormalMean(15, 365, b) }
func Benchmark_normal_Mean4(b *testing.B) { benchmarkNormalMean(3, 8, b) }
func Benchmark_normal_Mean5(b *testing.B) { benchmarkNormalMean(1e3, 1, b) }
func Benchmark_normal_Mean6(b *testing.B) { benchmarkNormalMean(534, 82, b) }

func Test_normal_Median(t *testing.T) {
	tests := []struct {
		name string
		norm Normal
		want float64
	}{
		{"NaN case mu", Normal{math.NaN(), 1.0}, math.NaN()},
		{"NaN case sigma", Normal{0.0, math.NaN()}, 0.0},
		{"Normal case", Normal{0.0, 1.0}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.norm.Median()
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Median() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkNormalMedian(seed1 int64, seed2 int64, b *testing.B) {
	norm := Normal{}
	rand.Seed(seed1)
	norm.Mu = rand.Float64() * 10
	rand.Seed(seed2)
	norm.Sigma = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		norm.Median()
	}
}

func Benchmark_normal_Median1(b *testing.B) { benchmarkNormalMedian(0, 10, b) }
func Benchmark_normal_Median2(b *testing.B) { benchmarkNormalMedian(25, 2, b) }
func Benchmark_normal_Median3(b *testing.B) { benchmarkNormalMedian(15, 365, b) }
func Benchmark_normal_Median4(b *testing.B) { benchmarkNormalMedian(3, 8, b) }
func Benchmark_normal_Median5(b *testing.B) { benchmarkNormalMedian(1e3, 1, b) }
func Benchmark_normal_Median6(b *testing.B) { benchmarkNormalMedian(534, 82, b) }

func Test_normal_StdDev(t *testing.T) {
	tests := []struct {
		name string
		norm Normal
		want float64
	}{
		{"NaN case mu", Normal{math.NaN(), 1.0}, 1.0},
		{"NaN case sigma", Normal{0.0, math.NaN()}, math.NaN()},
		{"Normal case", Normal{0.0, 1.0}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.norm.StdDev()
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("StdDev() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("StdDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkNormalStdDev(seed1 int64, seed2 int64, b *testing.B) {
	norm := Normal{}
	rand.Seed(seed1)
	norm.Mu = rand.Float64() * 10
	rand.Seed(seed2)
	norm.Sigma = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		norm.StdDev()
	}
}

func Benchmark_normal_StdDev1(b *testing.B) { benchmarkNormalStdDev(0, 10, b) }
func Benchmark_normal_StdDev2(b *testing.B) { benchmarkNormalStdDev(25, 2, b) }
func Benchmark_normal_StdDev3(b *testing.B) { benchmarkNormalStdDev(15, 365, b) }
func Benchmark_normal_StdDev4(b *testing.B) { benchmarkNormalStdDev(3, 8, b) }
func Benchmark_normal_StdDev5(b *testing.B) { benchmarkNormalStdDev(1e3, 1, b) }
func Benchmark_normal_StdDev6(b *testing.B) { benchmarkNormalStdDev(534, 82, b) }

func Test_normal_Variance(t *testing.T) {
	tests := []struct {
		name string
		norm Normal
		want float64
	}{
		{"NaN case mu", Normal{math.NaN(), 1.0}, 1.0},
		{"NaN case sigma", Normal{0.0, math.NaN()}, math.NaN()},
		{"Normal case", Normal{0.0, 1.0}, 1.0},
		{"Normal case 2", Normal{1.0, 2.0}, 4.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.norm.Variance()
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

func benchmarkNormalVariance(seed1 int64, seed2 int64, b *testing.B) {
	norm := Normal{}
	rand.Seed(seed1)
	norm.Mu = rand.Float64() * 10
	rand.Seed(seed2)
	norm.Sigma = rand.Float64() * 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		norm.Variance()
	}
}

func Benchmark_normal_Variance1(b *testing.B) { benchmarkNormalVariance(0, 10, b) }
func Benchmark_normal_Variance2(b *testing.B) { benchmarkNormalVariance(25, 2, b) }
func Benchmark_normal_Variance3(b *testing.B) { benchmarkNormalVariance(15, 365, b) }
func Benchmark_normal_Variance4(b *testing.B) { benchmarkNormalVariance(3, 8, b) }
func Benchmark_normal_Variance5(b *testing.B) { benchmarkNormalVariance(1e3, 1, b) }
func Benchmark_normal_Variance6(b *testing.B) { benchmarkNormalVariance(534, 82, b) }

func TestNewNormal(t *testing.T) {
	tests := []struct {
		name string
		want Normal
	}{
		{"Normal case", Normal{0.0, 1.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNormal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNormal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewNormal()
	}
}
