package main

import "testing"

func Test_playGame(t *testing.T) {
	type args struct {
		nth   int
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nth: 2020, input: "1,3,2"}, want: 1},
		{args: args{nth: 2020, input: "2,1,3"}, want: 10},
		{args: args{nth: 2020, input: "1,2,3"}, want: 27},
		{args: args{nth: 2020, input: "2,3,1"}, want: 78},
		{args: args{nth: 2020, input: "3,2,1"}, want: 438},
		{args: args{nth: 2020, input: "3,1,2"}, want: 1836},
		{args: args{nth: 30000000, input: "0,3,6"}, want: 175594},
		{args: args{nth: 30000000, input: "1,3,2"}, want: 2578},
		{args: args{nth: 30000000, input: "2,1,3"}, want: 3544142},
		{args: args{nth: 30000000, input: "1,2,3"}, want: 261214},
		{args: args{nth: 30000000, input: "2,3,1"}, want: 6895259},
		{args: args{nth: 30000000, input: "3,2,1"}, want: 18},
		{args: args{nth: 30000000, input: "3,1,2"}, want: 362},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := playGame(tt.args.nth, tt.args.input); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
