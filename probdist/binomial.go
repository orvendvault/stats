package stats

import (
	"errors"
	"math"
)

// Binomial is used to represent the binomial distribution parameters.
// N is the number of trials {0, 1, 2, 3, ...}.
// P is the succes probability of each trial [0, 1].
type Binomial struct {
	N float64
	P float64
}

// NewBinomial is used to initialize binomial parameters. What is different from
// Binomial type is that here the parameters are validated.
// N must be an integer positive number and P a real between 0 and 1.
func NewBinomial(n float64, p float64) (Binomial, error) {
	if math.Mod(n, 1.0) != 0 || p < 0 || p > 1.0 {
		return Binomial{20.0, 0.5}, errors.New("stats: invalid Binomial parameters")
	}
	return Binomial{N: n, P: p}, nil
}

// Mean returns the mean of the binomial distribution.
func (bin Binomial) Mean() float64 {
	return bin.N * bin.P
}

// StdDev returns the standard deviation of the binomial distribution.
func (bin Binomial) StdDev() float64 {
	return math.Sqrt(bin.Variance())
}

// Variance returns the variance of the binomial distribution.
func (bin Binomial) Variance() float64 {
	return bin.Mean() * (1 - bin.P)
}
