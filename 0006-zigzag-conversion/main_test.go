package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example 1", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"example 2", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"one row", args{"PAYPALISHIRING", 1}, "PAYPALISHIRING"},
		{"two rows", args{"PAYPALISHIRING", 2}, "PYAIHRNAPLSIIG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
