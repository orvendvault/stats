package stats

import (
	"math"
)

type binomial struct {
	n float64
	p float64
}

func (bin binomial) Mean() float64 {
	return bin.n * bin.p
}

func (bin binomial) StdDev() float64 {
	return math.Sqrt(bin.Variance())
}

func (bin binomial) Variance() float64 {
	return bin.Mean() * (1 - bin.p)
}
