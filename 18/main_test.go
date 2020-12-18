package main

import "testing"

func Test_parse(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{expression: "1 + 2 * 3 + 4 * 5 + 6"}, want: 71},
		{args: args{expression: "1 + (2 * 3) + (4 * (5 + 6))"}, want: 51},
		{args: args{expression: "2 * 3 + (4 * 5)"}, want: 26},
		{args: args{expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)"}, want: 437},
		{args: args{expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, want: 12240},
		{args: args{expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, want: 13632},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.expression); got != tt.want {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseAdvanced(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{expression: "1 + 2 * 3 + 4 * 5 + 6"}, want: 231},
		{args: args{expression: "1 + (2 * 3) + (4 * (5 + 6))"}, want: 51},
		{args: args{expression: "2 * 3 + (4 * 5)"}, want: 46},
		{args: args{expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)"}, want: 1445},
		{args: args{expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, want: 669060},
		{args: args{expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, want: 23340},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseAdvanced(tt.args.expression); got != tt.want {
				t.Errorf("parseAdvanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
