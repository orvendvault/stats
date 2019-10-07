package stats

// import (
// 	"math"
// )

// StudentsT is used to represent the students t distribution parameters
type StudentsT struct {
	V float64
}

// func (t studentsT) CDF(x float64, v int) float64 {
// 	if x == 0 {
// 		return 0.5
// 	} else if x > 0 {
// 		return 1 - 0.5*regIncBeta(t.v/(x*x+t.v), t.v/2, 0.5)
// 	} else if x < 0 {
// 		return 1 - (1 - 0.5*regIncBeta(t.v/(x*x+t.v), t.v/2, 0.5))
// 	} else {
// 		return math.NaN()
// 	}
// }

// //TODO: improve this functions.
// func lgamma(input float64) float64 {
// 	res, _ := math.Lgamma(input)
// 	return res
// }
// func regIncBeta(x, a, b float64) float64 {
// 	return math.Exp(lgamma(a) + lgamma(b) - lgamma(a+b))
// }
