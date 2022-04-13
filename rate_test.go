package money

import (
	"testing"
)

func TestRateFromFloat(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		args args
		want Rate
	}{
		{"test1", args{3.989999999}, RateFromFloat(3.99)},
		{"test2", args{3.99999}, RateFromFloat(4.0)},
		{"test3", args{1.0000000099999}, RateFromFloat(1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RateFromFloat(tt.args.val); got != tt.want {
				t.Errorf("RateFromFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRate_Equal(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		s    Rate
		args args
		want bool
	}{
		{"test1", RateFromFloat(1), args{1}, true},
		{"test2", RateFromFloat(1.0004), args{1.00039}, true},
		{"test3", RateFromFloat(1.0), args{1.0000499}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.val); got != tt.want {
				t.Errorf("Rate.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
