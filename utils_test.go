package money

import "testing"

func TestCalculateRate(t *testing.T) {
	type args struct {
		val1 Money
		val2 Money
	}
	tests := []struct {
		name string
		args args
		want Rate
	}{
		{"", args{val1: MoneyFromFloat(400.00), val2: MoneyFromFloat(5.30)}, RateFromFloat(75.4717)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateRate(tt.args.val1, tt.args.val2); got != tt.want {
				t.Errorf("CalculateRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
