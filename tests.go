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

//OneSampleTTest performs a One Sample T Test one upper tailed
//This test can be performed when we don't know the populations std dev.
//It returns true if the null hypoyhesis (no difference) is accepted and false otherwise
func OneSampleTTest(sample []float64, popmean float64, alpha float64) (bool, float64) {
	mean := Mean(sample)
	s := StdDev(sample)
	n := float64(len(sample))
	tscore := (mean - popmean) / (s / math.Sqrt(n))
	tcritical := pd.GetTStatistic(n-1, alpha)

	if tscore > tcritical {
		return false, tscore
	}
	return true, tcritical
}
