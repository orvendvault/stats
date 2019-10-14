package stats

import (
	"errors"
	"math"
)

// Gamma is used to represent the gamma distribution parameters.
// K is the shape parameter and must be > 0.
// Theta is the scale parameter and it must be > 0.
type Gamma struct {
	K     float64
	Theta float64
}

// NewGamma is used to initialize gamma parameters. What is different from
// Gamma type is that here the parameters are validated.
// K > 0 and Theta > 0 , both real numbers.
func NewGamma(k float64, theta float64) (Gamma, error) {
	if k <= 0 || theta <= 0 {
		return Gamma{1.0, 2.0}, errors.New("stats: invalid Gamma parameters. Check K > 0 and Theta > 0")
	}
	return Gamma{K: k, Theta: theta}, nil
}

// Mean returns the mean of the gamma distribution.
func (g Gamma) Mean() float64 {
	return g.K * g.Theta
}

// StdDev returns the standard deviation of the gamma distribution.
func (g Gamma) StdDev() float64 {
	return math.Sqrt(g.Variance())
}

// Variance returns the variance of the gamma distribution.
func (g Gamma) Variance() float64 {
	return g.K * g.Theta * g.Theta
}
