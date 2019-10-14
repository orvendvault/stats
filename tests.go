package stats

import (
	"math"

	pd "github.com/orvend/stats/probdist"
)

// TailDirection represents the direction of the tails that you consider from a distribution to perform statistical tests
type TailDirection uint8

const (
	// TailRight represents the right tail of the distribution
	TailRight TailDirection = iota
	// TailLeft represents the left tail of the distribution
	TailLeft
	// TailBoth represents both left and right tails
	TailBoth
)

// OneSampleZTest performs a Z Test.
// Z is the Standard normal distribution N(0,1)
// This test can be performed when the population is normally distributed and the population variance is known.
// It returns true if the null hypothesis is accepted and false otherwise.
// alpha is the significance level of one tail
// Right tail (1):
// Null hypothesis Ho :  sample mean = pop mean
// Alternative hypothesis H1 : sample mean > pop mean
// Two tails (2):
// Null hypothesis Ho :  sample mean = pop mean
// Alternative hypothesis H1 : sample mean != pop mean
// Left tail (-1):
// Null hypothesis Ho :  sample mean = pop mean
// Alternative hypothesis H1 : sample mean < pop mean
func OneSampleZTest(sample []float64, pop pd.Normal, alpha float64, tails TailDirection) (bool, float64) {
	smean := Mean(sample)
	zscore := (smean - pop.Mu) / (pop.Sigma / math.Sqrt(float64(len(sample))))

	switch tails {
	case TailRight:
		pvalue := 1 - pop.CDF(zscore)
		if pvalue < alpha {
			return false, pvalue
		}
		return true, pvalue
	case TailBoth:
		//symmetric, we can use just one tail
		pvalue := 1 - pop.CDF(zscore)
		if pvalue < alpha {
			return false, pvalue
		}
		return true, pvalue
	case TailLeft:
		pvalue := pop.CDF(-zscore)
		if pvalue < alpha {
			return false, pvalue
		}
		return true, pvalue
	default:
		panic("stats: incorrect tails input value. Try 1(right), 2 or -1(left)")
	}
}

// OneSampleTTest performs a One Sample T Test
// This test can be performed when we don't know the populations std dev.
// It returns true if the null hypoyhesis is accepted and false otherwise
// alpha is the significance level of one tail
// Right tail :
// Null hypothesis Ho :  sample mean = pop mean
// Alternative hypothesis H1 : sample mean > pop mean
// Two tails :
// Null hypothesis Ho :  sample mean = pop mean
// Alternative hypothesis H1 : sample mean != pop mean
// Left tail :
// Null hypothesis Ho :  sample mean = pop mean
// Alternative hypothesis H1 : sample mean < pop mean
func OneSampleTTest(sample []float64, popmean float64, alpha float64, tails TailDirection) (bool, float64) {
	mean := Mean(sample)
	s := StdDev(sample)
	n := float64(len(sample))
	tscore := (mean - popmean) / (s / math.Sqrt(n))
	switch tails {
	case TailRight:
		tcritical := pd.GetTStatistic(n-1, alpha)
		if tscore > tcritical {
			return false, tcritical
		}
		return true, tcritical
	case TailBoth:
		utcritical := pd.GetTStatistic(n-1, 2*alpha)
		ltcritical := -utcritical
		if tscore > utcritical || tscore < ltcritical {
			return false, utcritical
		}
		return true, utcritical
	case TailLeft:
		tcritical := -pd.GetTStatistic(n-1, alpha)
		if tscore < tcritical {
			return false, tcritical
		}
		return true, tcritical
	default:
		panic("stats: incorrect tails input value. Try 1(right), 2 or -1(left)")
	}
}
