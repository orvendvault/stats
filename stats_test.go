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
