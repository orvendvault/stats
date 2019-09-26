package stats

import (
	"math"
)

type exponential struct {
	lambda float64
}

func (exp exponential) Mean() float64 {
	return 1 / exp.lambda
}

func (exp exponential) StdDev() float64 {
	return math.Sqrt(exp.Variance())
}

func (exp exponential) Variance() float64 {
	return 1 / (exp.lambda * exp.lambda)
}
