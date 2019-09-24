package stats

import (
	"sort"
)

// Mean returns the mean of the slice.
func Mean(input []float64) float64 {
	if len(input) == 0 {
		return 0
	}
	sum := 0.0
	for _, in := range input {
		sum += in
	}
	return sum / float64(len(input))
}

// Median returns the median of the slice
func Median(input []float64) (output float64) {
	if len(input) == 0 {
		return 0
	}
	if len(input) == 1 {
		return input[0]
	}

	sort.Float64s(input)
	l := len(input)

	if l%2 != 0 {
		output = input[l/2]
	} else {
		output = (input[l/2] + input[l/2-1]) / 2
	}

	return output
}
