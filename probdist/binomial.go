package stats

import (
	"math"
)

// Binomial is used to represent the binomial distribution parameters
type Binomial struct {
	N float64
	P float64
}

// Mean returns the mean of the binomial distribution
func (bin Binomial) Mean() float64 {
	return bin.N * bin.P
}

// StdDev returns the standard deviation of the binomial distribution
func (bin Binomial) StdDev() float64 {
	return math.Sqrt(bin.Variance())
}

// Variance returns the variance of the binomial distribution
func (bin Binomial) Variance() float64 {
	return bin.Mean() * (1 - bin.P)
}
