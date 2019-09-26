package stats

import (
	"math"
	"math/rand"
	"sort"
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

func BenchmarkMean1(b *testing.B)   { benchmarkMean(1, b) }
func BenchmarkMean2(b *testing.B)   { benchmarkMean(2, b) }
func BenchmarkMean10(b *testing.B)  { benchmarkMean(10, b) }
func BenchmarkMean1e3(b *testing.B) { benchmarkMean(1e3, b) }
func BenchmarkMean1e6(b *testing.B) { benchmarkMean(1e6, b) }

func TestMedian(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name      string
		args      args
		want      float64
		wantPanic bool
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 2.5, false},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 3.0, false},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 2.5, true},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 3.0, true},
		{"empty case", args{[]float64{}}, math.NaN(), false},
		{"single value case", args{[]float64{1.0}}, math.NaN(), false},
		{"two values case", args{[]float64{1.0, 9.0}}, 5.0, false},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Error("Median() want panic")
					}
				}
			}()
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

func benchmarkMedian(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	sort.Float64s(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Median(s)
	}
}

func BenchmarkMedian0(b *testing.B)   { benchmarkMedian(0, b) }
func BenchmarkMedian1(b *testing.B)   { benchmarkMedian(1, b) }
func BenchmarkMedian2(b *testing.B)   { benchmarkMedian(2, b) }
func BenchmarkMedian10(b *testing.B)  { benchmarkMedian(10, b) }
func BenchmarkMedian1e3(b *testing.B) { benchmarkMedian(1e3, b) }
func BenchmarkMedian1e6(b *testing.B) { benchmarkMedian(1e6, b) }

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
		{"NaN case", args{[]float64{1.0, 5.0, math.NaN(), 6.0, 1.0, 2.0, 3.0}}, math.NaN()},
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

func benchmarkMax(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(s)
	}
}

func BenchmarkMax0(b *testing.B)   { benchmarkMax(0, b) }
func BenchmarkMax1(b *testing.B)   { benchmarkMax(1, b) }
func BenchmarkMax2(b *testing.B)   { benchmarkMax(2, b) }
func BenchmarkMax10(b *testing.B)  { benchmarkMax(10, b) }
func BenchmarkMax1e3(b *testing.B) { benchmarkMax(1e3, b) }
func BenchmarkMax1e6(b *testing.B) { benchmarkMax(1e6, b) }

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
		{"NaN case", args{[]float64{1.0, 5.0, math.NaN(), 6.0, 1.0, 2.0, 3.0}}, math.NaN()},
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

func benchmarkMin(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(s)
	}
}

func BenchmarkMin0(b *testing.B)   { benchmarkMin(0, b) }
func BenchmarkMin1(b *testing.B)   { benchmarkMin(1, b) }
func BenchmarkMin2(b *testing.B)   { benchmarkMin(2, b) }
func BenchmarkMin10(b *testing.B)  { benchmarkMin(10, b) }
func BenchmarkMin1e3(b *testing.B) { benchmarkMin(1e3, b) }
func BenchmarkMin1e6(b *testing.B) { benchmarkMin(1e6, b) }

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
		{"NaN case", args{[]float64{1.0, 5.0, math.NaN(), 6.0, 1.0, 2.0, 3.0}}, math.NaN()},
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

func benchmarkRange(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Range(s)
	}
}

func BenchmarkRange0(b *testing.B)   { benchmarkRange(0, b) }
func BenchmarkRange1(b *testing.B)   { benchmarkRange(1, b) }
func BenchmarkRange2(b *testing.B)   { benchmarkRange(2, b) }
func BenchmarkRange10(b *testing.B)  { benchmarkRange(10, b) }
func BenchmarkRange1e3(b *testing.B) { benchmarkRange(1e3, b) }
func BenchmarkRange1e6(b *testing.B) { benchmarkRange(1e6, b) }

func TestStdDev(t *testing.T) {
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
		{"NaN case", args{[]float64{1.0, 2.0, 3.0, 2.3, 1.4, math.NaN(), 1.5, 1.5, 1.8, 2.6, 2.3, 2.0, 2.2}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StdDev(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Range() = %v, want %v", got, tt.want)
				}
			} else if math.Abs(got-tt.want) > 1e6 {
				t.Errorf("StdDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkStdDev(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StdDev(s)
	}
}

func BenchmarkStdDev0(b *testing.B)   { benchmarkStdDev(0, b) }
func BenchmarkStdDev1(b *testing.B)   { benchmarkStdDev(1, b) }
func BenchmarkStdDev2(b *testing.B)   { benchmarkStdDev(2, b) }
func BenchmarkStdDev10(b *testing.B)  { benchmarkStdDev(10, b) }
func BenchmarkStdDev1e3(b *testing.B) { benchmarkStdDev(1e3, b) }
func BenchmarkStdDev1e6(b *testing.B) { benchmarkStdDev(1e6, b) }

func TestVariance(t *testing.T) {
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
		{"NaN case", args{[]float64{1.0, 2.0, 3.0, 2.3, 1.4, math.NaN(), 1.5, 1.5, 1.8, 2.6, 2.3, 2.0, 2.2}}, math.NaN()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Variance(tt.args.input)
			if math.IsNaN(got) || math.IsNaN(tt.want) {
				if !math.IsNaN(got) || !math.IsNaN(tt.want) {
					t.Errorf("Range() = %v, want %v", got, tt.want)
				}
			} else if math.Abs(got-tt.want) > 1e6 {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkVariance(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Variance(s)
	}
}

func BenchmarkVariance0(b *testing.B)   { benchmarkVariance(0, b) }
func BenchmarkVariance1(b *testing.B)   { benchmarkVariance(1, b) }
func BenchmarkVariance2(b *testing.B)   { benchmarkVariance(2, b) }
func BenchmarkVariance10(b *testing.B)  { benchmarkVariance(10, b) }
func BenchmarkVariance1e3(b *testing.B) { benchmarkVariance(1e3, b) }
func BenchmarkVariance1e6(b *testing.B) { benchmarkVariance(1e6, b) }

func TestQuartile1(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name      string
		args      args
		want      float64
		wantPanic bool
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 1.5, false},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 1.5, false},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 1.5, true},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 1.5, true},
		{"empty case", args{[]float64{}}, math.NaN(), false},
		{"single value case", args{[]float64{1.0}}, math.NaN(), false},
		{"two values case", args{[]float64{1.0, 9.0}}, math.NaN(), false},
		{"three values case", args{[]float64{1.0, 9.0, 11.0}}, math.NaN(), false},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Error("Quartile2() want panic")
					}
				}
			}()
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

func benchmarkQuartile1(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	sort.Float64s(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Quartile1(s)
	}
}

func BenchmarkQuartile10(b *testing.B)   { benchmarkQuartile1(0, b) }
func BenchmarkQuartile11(b *testing.B)   { benchmarkQuartile1(1, b) }
func BenchmarkQuartile12(b *testing.B)   { benchmarkQuartile1(2, b) }
func BenchmarkQuartile110(b *testing.B)  { benchmarkQuartile1(10, b) }
func BenchmarkQuartile11e3(b *testing.B) { benchmarkQuartile1(1e3, b) }
func BenchmarkQuartile11e6(b *testing.B) { benchmarkQuartile1(1e6, b) }

func TestQuartile2(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name      string
		args      args
		want      float64
		wantPanic bool
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 2.5, false},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 3.0, false},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 2.5, true},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 3.0, true},
		{"empty case", args{[]float64{}}, math.NaN(), false},
		{"single value case", args{[]float64{1.0}}, math.NaN(), false},
		{"two values case", args{[]float64{1.0, 9.0}}, 5.0, false},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Error("Quartile2() want panic")
					}
				}
			}()
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

func benchmarkQuartile2(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	sort.Float64s(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Quartile2(s)
	}
}

func BenchmarkQuartile20(b *testing.B)   { benchmarkQuartile2(0, b) }
func BenchmarkQuartile21(b *testing.B)   { benchmarkQuartile2(1, b) }
func BenchmarkQuartile22(b *testing.B)   { benchmarkQuartile2(2, b) }
func BenchmarkQuartile210(b *testing.B)  { benchmarkQuartile2(10, b) }
func BenchmarkQuartile21e3(b *testing.B) { benchmarkQuartile2(1e3, b) }
func BenchmarkQuartile21e6(b *testing.B) { benchmarkQuartile2(1e6, b) }

func TestQuartile3(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name      string
		args      args
		want      float64
		wantPanic bool
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 3.5, false},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 4.5, false},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 3.5, true},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 4.5, true},
		{"empty case", args{[]float64{}}, math.NaN(), false},
		{"single value case", args{[]float64{1.0}}, math.NaN(), false},
		{"two values case", args{[]float64{1.0, 9.0}}, math.NaN(), false},
		{"three values case", args{[]float64{1.0, 9.0, 11.0}}, math.NaN(), false},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Error("Quartile3() want panic")
					}
				}
			}()
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

func benchmarkQuartile3(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	sort.Float64s(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Quartile3(s)
	}
}

func BenchmarkQuartile30(b *testing.B)   { benchmarkQuartile3(0, b) }
func BenchmarkQuartile31(b *testing.B)   { benchmarkQuartile3(1, b) }
func BenchmarkQuartile32(b *testing.B)   { benchmarkQuartile3(2, b) }
func BenchmarkQuartile310(b *testing.B)  { benchmarkQuartile3(10, b) }
func BenchmarkQuartile31e3(b *testing.B) { benchmarkQuartile3(1e3, b) }
func BenchmarkQuartile31e6(b *testing.B) { benchmarkQuartile3(1e6, b) }

func TestInterQuartileRange(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name      string
		args      args
		want      float64
		wantPanic bool
	}{
		{"even case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0}}, 2.0, false},
		{"odd case sorted", args{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}}, 3.0, false},
		{"even case unsorted", args{[]float64{4.0, 2.0, 1.0, 3.0}}, 2.0, true},
		{"odd case unsorted", args{[]float64{4.0, 3.0, 5.0, 1.0, 2.0}}, 3.0, true},
		{"empty case", args{[]float64{}}, math.NaN(), false},
		{"single value case", args{[]float64{1.0}}, math.NaN(), false},
		{"two values case", args{[]float64{1.0, 9.0}}, math.NaN(), false},
		{"three values case", args{[]float64{1.0, 9.0, 11.0}}, math.NaN(), false},
		{"nan", args{[]float64{math.NaN(), 5.0}}, math.NaN(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Error("InterQuartileRange() want panic")
					}
				}
			}()
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

func benchmarkInterQuartileRange(len int, b *testing.B) {
	s := make([]float64, len)
	for e := 0; e <= len-1; e++ {
		s[e] = rand.Float64()
	}
	sort.Float64s(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InterQuartileRange(s)
	}
}

func BenchmarkInterQuartileRange0(b *testing.B)   { benchmarkInterQuartileRange(0, b) }
func BenchmarkInterQuartileRange1(b *testing.B)   { benchmarkInterQuartileRange(1, b) }
func BenchmarkInterQuartileRange2(b *testing.B)   { benchmarkInterQuartileRange(2, b) }
func BenchmarkInterQuartileRange10(b *testing.B)  { benchmarkInterQuartileRange(10, b) }
func BenchmarkInterQuartileRange1e3(b *testing.B) { benchmarkInterQuartileRange(1e3, b) }
func BenchmarkInterQuartileRange1e6(b *testing.B) { benchmarkInterQuartileRange(1e6, b) }
