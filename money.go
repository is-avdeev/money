package money

import (
	"fmt"
)

type Money int64

func MoneyFromFloat(val float64) Money {
	return Money(int64(round(val, 2.0) * 100))
}

func MoneyFromInt(val int64) Money {
	return Money(val)
}

func (s Money) Int() int64 {
	return int64(s)
}

func (s Money) Float() float64 {
	return float64(s) / 100
}

func (s Money) String() string {
	return fmt.Sprintf("%.2f", s.Float())
}

func (s Money) ToRate() Rate {
	return Rate(int64(s) * 100)
}

func (s Money) CalculateByRate(rate Rate, unit float64) Money {
	return MoneyFromFloat(s.Float() * rate.Float() * unit)
}

func (s Money) Equal(val float64) bool {
	return s == MoneyFromFloat(val)
}

// If values equals res - true
func (s Money) Compare(val float64) (str1, str2 string, res bool) {
	str1 = s.String()
	str2 = MoneyFromFloat(val).String()
	res = (str1 == str2)
	return
}
