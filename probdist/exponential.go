package stats

import (
	"math"
)

// Exponential is used to represent the exponential distribution parameters
type Exponential struct {
	Lambda float64
}

// Mean returns the mean of the exponential distirbution
func (exp Exponential) Mean() float64 {
	return 1 / exp.Lambda
}

// StdDev returns the standard deviation of the exponential distributon
func (exp Exponential) StdDev() float64 {
	return math.Sqrt(exp.Variance())
}

// Variance returns the standard deviation of the exponential distributon
func (exp Exponential) Variance() float64 {
	return 1 / (exp.Lambda * exp.Lambda)
}
