package money

import (
	"testing"
)

func TestMoneyFromFloat(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		args args
		want Money
	}{
		{"test1", args{3.989999999}, MoneyFromFloat(3.99)},
		{"test2", args{3.99}, MoneyFromFloat(3.99)},
		{"test3", args{3.994}, MoneyFromFloat(3.99)},
		{"test4", args{3.995}, MoneyFromFloat(4.00)},
		{"test5", args{4.044}, MoneyFromFloat(4.04)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoneyFromFloat(tt.args.val); got != tt.want {
				t.Errorf("MoneyFromFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Equal(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		s    Money
		args args
		want bool
	}{
		{"test1", MoneyFromFloat(0.30), args{0.3000000009}, true},
		{"test2", MoneyFromFloat(0.30), args{0.3100000009}, false},
		{"test3", MoneyFromFloat(0.3), args{0.3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.val); got != tt.want {
				t.Errorf("Money.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
