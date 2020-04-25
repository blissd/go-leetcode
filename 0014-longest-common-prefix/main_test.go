package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example 1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"example 2", args{[]string{"dog", "racecar", "car"}}, ""},
		{"all same", args{[]string{"dog", "dog", "dog"}}, "dog"},
		{"last different", args{[]string{"doga", "doga", "dog"}}, "dog"},
		{"last different 2", args{[]string{"doga", "doga", "dogc"}}, "dog"},
		{"blank", args{[]string{"doga", "", "dogc"}}, ""},
		{"empty", args{[]string{}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
