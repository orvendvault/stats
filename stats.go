package stats

// Mean returns the mean of the slice.
func Mean(input []float64) float64 {
	sum := 0.0
	for _, in := range input {
		sum += in
	}
	return sum / float64(len(input))
}
