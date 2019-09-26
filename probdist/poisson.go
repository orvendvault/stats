package stats

import (
	"math"
)

type poisson struct {
	lambda float64
}

func (p poisson) Mean() float64 {
	return p.lambda
}

func (p poisson) StdDev() float64 {
	return math.Sqrt(p.Variance())
}

func (p poisson) Variance() float64 {
	return p.lambda
}
