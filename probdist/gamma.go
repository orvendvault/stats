package stats

import (
	"math"
)

type gamma struct {
	k     float64
	theta float64
}

func (g gamma) Mean() float64 {
	return g.k * g.theta
}

func (g gamma) StdDev() float64 {
	return math.Sqrt(g.Variance())
}

func (g gamma) Variance() float64 {
	return g.k * g.theta * g.theta
}
