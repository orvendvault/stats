package stats

import (
	"math"
)

// Exponential is used to represent the exponential distribution parameters.
// The parameter Lambda must be > 0.
type Exponential struct {
	Lambda float64
}

// NewExponential is used to initialize exponential parameters.
func NewExponential() Exponential {
	return Exponential{1.0}
}

// CDF returns the cumulative distribution function output for the exponential distribution.
func (exp Exponential) CDF(x float64) float64 {
	return 1 - math.Pow(math.E, -exp.Lambda*x)
}

// Mean returns the mean of the exponential distirbution.
func (exp Exponential) Mean() float64 {
	return 1 / exp.Lambda
}

// StdDev returns the standard deviation of the exponential distributon.
func (exp Exponential) StdDev() float64 {
	return math.Sqrt(exp.Variance())
}

// Variance returns the standard deviation of the exponential distributon.
func (exp Exponential) Variance() float64 {
	return 1 / (exp.Lambda * exp.Lambda)
}
