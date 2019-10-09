package stats

import (
	"math"
)

// Poisson is used to represent the poisson distribution parameters
type Poisson struct {
	Lambda float64
}

// NewPoisson is used to initialize poisson parameters
func NewPoisson() Poisson {
	return Poisson{1.0}
}

// Mean returns the mean of the poisson distribution
func (p Poisson) Mean() float64 {
	return p.Lambda
}

// StdDev returns the standard deviation of the poisson distribution
func (p Poisson) StdDev() float64 {
	return math.Sqrt(p.Variance())
}

// Variance returns the variance of the poisson distribution
func (p Poisson) Variance() float64 {
	return p.Lambda
}
