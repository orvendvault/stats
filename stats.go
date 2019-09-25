package stats

import (
	"math"
	"sort"
)

// Mean returns the mean of the slice.
func Mean(input []float64) float64 {
	sum := 0.0
	for _, in := range input {
		sum += in
	}
	return sum / float64(len(input))
}

// Median returns the median of the slice. Panics if the input is not sorted.
func Median(input []float64) (output float64) {
	if len(input) < 2 {
		return math.NaN()
	}

	if !sort.Float64sAreSorted(input) {
		panic("stats: input is not sorted.")
	}

	l := len(input)

	if l%2 != 0 {
		output = input[l/2]
	} else {
		output = (input[l/2] + input[l/2-1]) / 2
	}

	return output
}
