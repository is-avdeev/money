package money

import (
	"fmt"
	"strconv"
	"strings"
)

type Rate int64

func RateFromFloat(val float64) Rate {
	return Rate(int64(round(val, 4.0) * 10000))
}

func RateFromString(str string) (Rate, error) {
	str = strings.Replace(str, ",", ".", -1)
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return Rate(0), err
	}

	return RateFromFloat(val), nil
}

func (s Rate) Int() int64 {
	return int64(s)
}

func (s Rate) Float() float64 {
	return float64(s) / 10000
}

func (s Rate) String() string {
	return fmt.Sprintf("%.4f", s.Float())
}

func (s Rate) Equal(val float64) bool {
	return s == RateFromFloat(val)
}

// If values equals res - true
func (s Rate) Compare(val float64) (str1, str2 string, res bool) {
	str1 = s.String()
	str2 = RateFromFloat(val).String()
	res = (str1 == str2)
	return
}
