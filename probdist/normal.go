package stats

import (
	"math"
)

type normal struct {
	mu    float64
	sigma float64
}

func (norm normal) CDF(x float64) float64 {
	return 0.5 * (1 + math.Erfc((x-norm.mu)/math.Sqrt(2)*norm.sigma))
}

func (norm normal) Mean() float64 {
	return norm.mu
}

func (norm normal) Median() float64 {
	return norm.mu
}

func (norm normal) StdDev() float64 {
	return norm.sigma
}

func (norm normal) Variance() float64 {
	return math.Pow(norm.sigma, 2.0)
}
