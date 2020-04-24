package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"I", args{1}, "I"},
		{"V", args{5}, "V"},
		{"X", args{10}, "X"},
		{"L", args{50}, "L"},
		{"C", args{100}, "C"},
		{"D", args{500}, "D"},
		{"M", args{1000}, "M"},
		{"example 1", args{3}, "III"},
		{"example 2", args{4}, "IV"},
		{"example 3", args{9}, "IX"},
		{"example 4", args{58}, "LVIII"},
		{"example 5", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
