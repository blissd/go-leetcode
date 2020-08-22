package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"example 1", args{3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"failure 1", args{4}, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args.n)
			sort.Strings(got)
			want := tt.want
			sort.Strings(want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
