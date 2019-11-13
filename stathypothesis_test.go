package stats

import (
	"math"
	"testing"

	pd "github.com/orvend/stats/probdist"
)

func TestOneSampleZTest(t *testing.T) {
	type args struct {
		sample []float64
		pop    pd.Normal
		alpha  float64
		tails  TailDirection
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 float64
	}{
		{"Normal case", args{[]float64{0.1, 0.02, -0.3, 0.47, 0.015, 0.21, -0.32, -0.05, -0.1, 0.15, 0.17, 0.08, -0.125}, pd.Normal{Mu: 0, Sigma: 1}, 0.05, 1}, true, 0.464639},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OneSampleZTest(tt.args.sample, tt.args.pop, tt.args.alpha, tt.args.tails)
			if got != tt.want {
				t.Errorf("OneSampleZTest() got = %v, want %v", got, tt.want)
			}
			if math.Abs(got1-tt.want1) > 1e-3 {
				t.Errorf("OneSampleZTest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOneSampleTTest(t *testing.T) {
	type args struct {
		sample  []float64
		popmean float64
		alpha   float64
		tails   TailDirection
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 float64
	}{
		{"Normal case", args{[]float64{0.1, 0.02, -0.3, 0.47, 0.015, 0.21, -0.32, -0.05, -0.1, 0.15, 0.17, 0.08, -0.125}, 0.0, 0.05, 0}, true, 1.7823},
		{"Normal case true", args{[]float64{0.1, 0.02, -0.3, 0.47, 0.015, 0.21, -0.32, -0.05, -0.1, 0.15, 0.17, 0.08, -0.125}, 10.0, 0.05, 0}, true, 1.7823},
		{"Normal case false", args{[]float64{0.1, 0.02, -0.3, 0.47, 0.015, 0.21, -0.32, -0.05, -0.1, 0.15, 0.17, 0.08, -0.125}, -10.0, 0.05, 0}, false, 1.7823},
		{"Zero mean case", args{[]float64{0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 12.5, 0, 0, 0, 12.5, 0, 12.5, 12.5, 12.5, 0, 0, 0, 0, 0, 0, 0, 12.5, 0, 12.5, 0, 0, 0, 12.5, 0, 0, 12.5, 0, 12.5, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 0, 0, 0, 12.5, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 0.05, 0}, false, 1.662},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OneSampleTTest(tt.args.sample, tt.args.popmean, tt.args.alpha, tt.args.tails)
			if got != tt.want {
				t.Errorf("OneSampleTTest() got = %v, want %v", got, tt.want)
			}
			if math.Abs(got1-tt.want1) > 1e-3 {
				t.Errorf("OneSampleTTest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPairedTTest(t *testing.T) {
	type args struct {
		presample  []float64
		postsample []float64
		alpha      float64
		tails      TailDirection
	}
	tests := []struct {
		name string
		args args
		want bool
		// want1 float64
	}{
		{"Case H0 false", args{[]float64{0, 0, 25, 0, 12.5, 25, 25, 12.5, 25, 0, 0, 0, 0, 12.5, 0, 0, 12.5, 12.5, 0, 25, 25, 12.5, 0, 12.5, 0, 25, 12.5, 25, 12.5, 25, 25, 12.5, 0, 0, 12.5, 25, 12.5, 12.5, 25, 12.5, 25, 12.5, 12.5, 0, 0, 25, 0, 12.5,
			37.5, 37.5, 25, 12.5, 0, 12.5, 0, 0, 12.5, 12.5, 12.5, 0, 0, 12.5, 0, 12.5, 12.5, 12.5, 0, 25, 12.5, 0, 0, 0, 0, 0, 0, 0, 12.5, 0, 0, 0, 12.5, 37.5, 25, 12.5, 0, 12.5, 0, 12.5, 12.5, 12.5, 12.5, 0}, []float64{0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 12.5, 0, 0, 0, 12.5, 0, 12.5, 12.5, 12.5, 0, 0, 0, 0, 0, 0, 0, 12.5, 0, 12.5, 0, 0, 0, 12.5, 0, 0, 12.5, 0, 12.5, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 0, 0, 0, 12.5, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 0, 12.5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0.05, TailRight}, false},
		{"Case H0 true", args{[]float64{25, 12}, []float64{12, 0}, 0.05, TailRight}, true},
		{"Case H0 true", args{[]float64{25, 12}, []float64{0, 0}, 0.05, TailRight}, true},
		{"Case H0 false", args{[]float64{80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80}, []float64{12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12}, 0.05, TailRight}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := PairedTTest(tt.args.presample, tt.args.postsample, tt.args.alpha, tt.args.tails)
			if got != tt.want {
				t.Errorf("PairedTTest() got = %v, want %v", got, tt.want)
			}
			// if got1 != tt.want1 {
			// 	t.Errorf("PairedTTest() got1 = %v, want %v", got1, tt.want1)
			// }
		})
	}
}
