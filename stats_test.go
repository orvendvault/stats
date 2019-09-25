package stats

import (
	"fmt"
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
		{"empty", args{[]float64{}}, math.NaN()},
		{"equal", args{[]float64{5.0, 5.0, 5.0}}, 5.0},
		{"different", args{[]float64{5.0, 10.0, 0.0}}, 5.0},
		{"different2", args{[]float64{5.0, 10.0}}, 7.5},
		{"nan", args{[]float64{5.0, math.NaN()}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mean(tt.args.input)
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

func TestMedian(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 2.5},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 3.0},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 2.5},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 3.0},
		{"empty case", args{[]float64{}}, math.NaN()},
		{"single value case", args{[]float64{1.0}}, math.NaN()},
		{"two values case", args{[]float64{1.0, 9.0}}, 5.0},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer handlepanic()
			got := Median(tt.args.input)
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

func handlepanic() {
	if r := recover(); r != nil {
		fmt.Println("Recover", r)
	}
}

func benchmarkMedian(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Median(s)
	}
}

func BenchmarkMedian0(b *testing.B)   { benchmarkMean(0, b) }
func BenchmarkMedian1(b *testing.B)   { benchmarkMean(1, b) }
func BenchmarkMedian2(b *testing.B)   { benchmarkMean(2, b) }
func BenchmarkMedian10(b *testing.B)  { benchmarkMean(10, b) }
func BenchmarkMedian1e3(b *testing.B) { benchmarkMean(1e3, b) }
func BenchmarkMedian1e6(b *testing.B) { benchmarkMean(1e6, b) }

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
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Min() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			got := Range(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Range() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
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

func TestQuartile1(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 1.5},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 1.5},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 1.5},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 1.5},
		{"empty case", args{[]float64{}}, math.NaN()},
		{"single value case", args{[]float64{1.0}}, math.NaN()},
		{"two values case", args{[]float64{1.0, 9.0}}, math.NaN()},
		{"three values case", args{[]float64{1.0, 9.0, 11.0}}, math.NaN()},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer handlepanic()
			got := Quartile1(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Quartile1() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Quartile1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartile2(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 2.5},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 3.0},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 2.5},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 3.0},
		{"empty case", args{[]float64{}}, math.NaN()},
		{"single value case", args{[]float64{1.0}}, math.NaN()},
		{"two values case", args{[]float64{1.0, 9.0}}, 5.0},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer handlepanic()
			got := Quartile2(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Quartile2() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Quartile2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartile3(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 3.5},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 4.5},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 3.5},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 4.5},
		{"empty case", args{[]float64{}}, math.NaN()},
		{"single value case", args{[]float64{1.0}}, math.NaN()},
		{"two values case", args{[]float64{1.0, 9.0}}, math.NaN()},
		{"three values case", args{[]float64{1.0, 9.0, 11.0}}, math.NaN()},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer handlepanic()
			got := Quartile3(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Quartile3() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("Quartile3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterQuartileRange(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 2.0},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 3.0},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 2.0},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 3.0},
		{"empty case", args{[]float64{}}, math.NaN()},
		{"single value case", args{[]float64{1.0}}, math.NaN()},
		{"two values case", args{[]float64{1.0, 9.0}}, math.NaN()},
		{"three values case", args{[]float64{1.0, 9.0, 11.0}}, math.NaN()},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer handlepanic()
			got := InterQuartileRange(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("InterQuartileRange() = %v, want %v", got, tt.want)
				}
			} else if got != tt.want {
				t.Errorf("InterQuartileRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
