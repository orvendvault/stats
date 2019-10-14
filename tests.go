package stats

import (
	"math"

	pd "github.com/orvend/stats/probdist"
)

// // OneSampleZTest performs a Z Test.
// // This test can be performed when the population is normally distrivuted and the population variance is known.
// // It returns true if the null hypothesis is accepted and false otherwise.
// func OneSampleZTest(sample []float64, pop pd.Normal, alpha float64) bool {
// 	smean := Mean(sample)
// 	zscore := (smean - pop.Mu) / (pop.Sigma / math.Sqrt(float64(len(sample))))
// 	pvalue := 1 - pop.CDF(zscore)

// 	if pvalue < alpha {
// 		return false
// 	}
// 	return true
// }

//OneSampleTTest performs a One Sample T Test
//This test can be performed when we don't know the populations std dev.
//It returns true if the null hypoyhesis is accepted and false otherwise
//alpha is the significance level of one tail
//Right tail (1):
//Null hypothesis Ho :  sample mean = pop mean
//Alternative hypothesis H1 : sample mean > pop mean
//Two tails (2):
//Null hypothesis Ho :  sample mean = pop mean
//Alternative hypothesis H1 : sample mean != pop mean
//Left tail (-1):
//Null hypothesis Ho :  sample mean = pop mean
//Alternative hypothesis H1 : sample mean < pop mean
func OneSampleTTest(sample []float64, popmean float64, alpha float64, tails int) (bool, float64) {
	mean := Mean(sample)
	s := StdDev(sample)
	n := float64(len(sample))
	tscore := (mean - popmean) / (s / math.Sqrt(n))
	switch tails {
	case 1:
		tcritical := pd.GetTStatistic(n-1, alpha)
		if tscore > tcritical {
			return false, tcritical
		}
		return true, tcritical
	case 2:
		utcritical := pd.GetTStatistic(n-1, 2*alpha)
		ltcritical := -utcritical
		if tscore > utcritical || tscore < ltcritical {
			return false, utcritical
		}
		return true, utcritical
	case -1:
		tcritical := -pd.GetTStatistic(n-1, alpha)
		if tscore < tcritical {
			return false, tcritical
		}
		return true, tcritical
	default:
		panic("stats: incorrect tails input value. Try 1(right), 2 or -1(left)")
	}
}
