package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"I"}, 1},
		{"2", args{"II"}, 2},
		{"3", args{"III"}, 3},
		{"4", args{"IV"}, 4},
		{"5", args{"V"}, 5},
		{"6", args{"VI"}, 6},
		{"7", args{"VII"}, 7},
		{"8", args{"VIII"}, 8},
		{"9", args{"IX"}, 9},
		{"10", args{"X"}, 10},
		{"example 4", args{"LVIII"}, 58},
		{"example 5", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
