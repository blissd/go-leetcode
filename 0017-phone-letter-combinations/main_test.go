package main

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{"23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"123", args{"123"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"213", args{"213"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"231", args{"231"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"1", args{"1"}, []string{}},
		{"2", args{"2"}, []string{"a", "b", "c"}},
		{"22", args{"22"}, []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
