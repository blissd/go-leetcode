package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"42"}, 42},
		{"example 2", args{"   -42"}, -42},
		{"example 3", args{"4193 with words"}, 4193},
		{"example 4", args{"words and 987"}, 0},
		{"failure 1", args{"-91283472332"}, -2147483648},
		{"failure 2", args{"3.14159"}, 3},
		{"failure 3", args{"+1"}, 1},
		{"failure 4", args{"-2147483648"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
