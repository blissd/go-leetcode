package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"failed 1", args{"bpfbhmipx"}, 7},
		{"one", args{" "}, 1},
		{"two", args{"ab"}, 2},
		{"three", args{"abc"}, 3},
		{"example 1", args{"abcabcbb"}, 3},
		{"example 2", args{"bbbbb"}, 1},
		{"example 3", args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
