package stats

import (
	"math"
)

// Normal is used to represent the normal distribution parameters
type Normal struct {
	Mu    float64
	Sigma float64
}

// CDF returns the cumulative distribution function output of the normal distribution for a given x
func (norm Normal) CDF(x float64) float64 {
	return 0.5 * (1 + math.Erf((x-norm.Mu)/(math.Sqrt(2)*norm.Sigma)))
}

// Mean returns the mean of the normal distribution
func (norm Normal) Mean() float64 {
	return norm.Mu
}

// Median returns the median of the normal distribution
func (norm Normal) Median() float64 {
	return norm.Mu
}

// StdDev returns the standard deviation of the normal distribution
func (norm Normal) StdDev() float64 {
	return norm.Sigma
}

// Variance returns the Variance of the normal distribution
func (norm Normal) Variance() float64 {
	return math.Pow(norm.Sigma, 2.0)
}
