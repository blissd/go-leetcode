package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{123}, 321},
		{"example 2", args{-123}, -321},
		{"example 3", args{120}, 21},
		{"failure 1", args{1}, 1},
		{"failure 1 (negative)", args{-1}, -1},
		{"failure 2", args{900000}, 9},
		{"failure 3", args{1534236469}, 0},
		{"failure 4", args{-2147483648}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
