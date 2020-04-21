package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example 1", args{"aa", "a"}, false},
		{"example 2", args{"aa", "a*"}, true},
		{"example 3", args{"ab", ".*"}, true},
		{"example 4", args{"aab", "c*a*b"}, true},
		{"example 5", args{"mississippi", "mis*is*p*."}, false},
		{"failure 1a", args{"aaa", "a*a"}, true},
		{"failure 1b", args{"aaa", "aa*"}, true},
		{"failure 1c", args{"aaa", "a*a*"}, true},
		{"failure 1d", args{"aaa", "a*a*a"}, true},
		{"failure 2a", args{"aaa", "aaaa"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
