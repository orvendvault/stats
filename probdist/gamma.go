package stats

import (
	"math"
)

// Gamma is used to represent the gamma distribution parameters
type Gamma struct {
	K     float64
	Theta float64
}

// Mean returns the mean of the gamma distribution
func (g Gamma) Mean() float64 {
	return g.K * g.Theta
}

// StdDev returns the standard deviation of the gamma distribution
func (g Gamma) StdDev() float64 {
	return math.Sqrt(g.Variance())
}

// Variance returns the variance of the gamma distribution
func (g Gamma) Variance() float64 {
	return g.K * g.Theta * g.Theta
}
