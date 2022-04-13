package money

import "math"

func round(val, n float64) float64 {
	decimals64 := math.Pow(10, n)
	multipliedRoundedNumber := math.Round(val * decimals64)
	return multipliedRoundedNumber / decimals64
}

// CalculateRate calculates the ratio between two money values
func CalculateRate(val1, val2 Money) Rate {
	if val2 == 0 {
		val2 = 1
	}

	return RateFromFloat(val1.Float() / val2.Float())
}
