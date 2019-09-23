package stats

func mean(input []float64) float64 {
	sum := 0.0
	for _, in := range input {
		sum += in
	}
	return sum / float64(len(input))
}
