package stats

import (
	"math"
)

// Poisson is used to represent the poisson distribution parameters.
// Lambda is the average number of events per interval and must belong to the set of real positive numbers.
type Poisson struct {
	Lambda float64
}

// NewPoisson is used to initialize poisson parameters.
func NewPoisson() Poisson {
	return Poisson{1.0}
}

// Mean returns the mean of the poisson distribution.
func (p Poisson) Mean() float64 {
	return p.Lambda
}

// StdDev returns the standard deviation of the poisson distribution.
func (p Poisson) StdDev() float64 {
	return math.Sqrt(p.Variance())
}

// Variance returns the variance of the poisson distribution.
func (p Poisson) Variance() float64 {
	return p.Lambda
}
