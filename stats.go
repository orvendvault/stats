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

// Max returns the maximum value of the sample
func Max(input []float64) float64 {
	if len(input) == 0 {
		return math.NaN()
	}
	max := input[0]
	for i := 1; i < len(input); i++ {
		if math.IsNaN(input[i]) {
			return math.NaN()
		}
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

// Min returns the minimum value of the sample
func Min(input []float64) float64 {
	if len(input) == 0 {
		return math.NaN()
	}
	min := input[0]
	for i := 1; i < len(input); i++ {
		if math.IsNaN(input[i]) {
			return math.NaN()
		}
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

// Range returns the difference between the largest and smallest values
func Range(input []float64) float64 {
	if len(input) < 2 {
		return math.NaN()
	}
	return Max(input) - Min(input)
}

// sumOfSquaredDifferences returns the sum of the squared differences of each observation from the mean
func sumOfSquaredDifferences(input []float64) float64 {
	if len(input) < 2 {
		return math.NaN()
	}
	mean := Mean(input)
	ssd := 0.0
	for _, o := range input {
		ssd += math.Pow(o-mean, 2.0)
	}
	return ssd
}

// StdDev returns the standard deviation of the sample
func StdDev(input []float64) float64 {
	return math.Sqrt(Variance(input))
}

// Variance returns the variance of the sample
func Variance(input []float64) float64 {
	return sumOfSquaredDifferences(input) / float64(len(input)-1)
}

// Quartile1 returns the first quartile
func Quartile1(input []float64) (Q1 float64) {
	if len(input) < 4 {
		return math.NaN()
	}
	if !sort.Float64sAreSorted(input) {
		panic("stats: input is not sorted.")
	}
	if len(input)%2 == 0 {
		Q1 = Median(input[:len(input)/2])
	} else {
		Q1 = Median(input[:len(input)/2])
	}
	return Q1
}

// Quartile2 returns the second quartile (equivalent to the median)
func Quartile2(input []float64) float64 {
	return Median(input)
}

// Quartile3 returns the third quartile
func Quartile3(input []float64) (Q3 float64) {
	if len(input) < 4 {
		return math.NaN()
	}
	if !sort.Float64sAreSorted(input) {
		panic("stats: input is not sorted.")
	}
	if len(input)%2 == 0 {
		Q3 = Median(input[len(input)/2:])
	} else {
		Q3 = Median(input[len(input)/2+1:])
	}
	return Q3
}

// InterQuartileRange returns the difference between the third and the first quartiles
func InterQuartileRange(input []float64) float64 {
	return Quartile3(input) - Quartile1(input)
}
