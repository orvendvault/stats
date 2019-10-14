package stats

import (
	"errors"
	"math"
)

// Exponential is used to represent the exponential distribution parameters.
// The parameter Lambda must be > 0.
type Exponential struct {
	Lambda float64
}

// NewExponential is used to initialize exponential parameters. What is different from
// Exponential type is that here the parameters are validated.
// Lambda must be a real number > 0.
func NewExponential(lambda float64) (Exponential, error) {
	if lambda <= 0 {
		return Exponential{1.0}, errors.New("stats: invalid Exponential parameters. Check Lambda > 0")
	}
	return Exponential{Lambda: lambda}, nil
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
