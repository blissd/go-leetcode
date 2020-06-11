package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example 1", args{"()"}, true},
		{"example 2", args{"()[]{}"}, true},
		{"example 3", args{"(]"}, false},
		{"example 4", args{"([)]"}, false},
		{"example 5", args{"{[]}"}, true},
		{"failure 1", args{"]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
