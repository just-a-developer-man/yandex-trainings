package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		tRoom    int
		tCond    int
		condMode string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "yandex_test1",
			args: args{
				tRoom:    10,
				tCond:    20,
				condMode: "heat",
			},
			want: 20,
		},
		{
			name: "yandex_test2",
			args: args{
				tRoom:    10,
				tCond:    20,
				condMode: "freeze",
			},
			want: 10,
		},
		{
			name: "simple_heat",
			args: args{
				tRoom:    -200,
				tCond:    200,
				condMode: "heat",
			},
			want: 200,
		},
		{
			name: "simple_freeze",
			args: args{
				tRoom:    200,
				tCond:    -200,
				condMode: "freeze",
			},
			want: -200,
		},
		{
			name: "null_heat",
			args: args{
				tRoom:    0,
				tCond:    0,
				condMode: "heat",
			},
			want: 0,
		},
		{
			name: "no_heat",
			args: args{
				tRoom:    200,
				tCond:    100,
				condMode: "heat",
			},
			want: 200,
		},
		{
			name: "no_freeze",
			args: args{
				tRoom:    -200,
				tCond:    -100,
				condMode: "freeze",
			},
			want: -200,
		},
		{
			name: "auto_heat",
			args: args{
				tRoom:    100,
				tCond:    200,
				condMode: "auto",
			},
			want: 200,
		},
		{
			name: "auto_freeze",
			args: args{
				tRoom: 200,
				tCond: 100, condMode: "auto",
			},
			want: 100,
		},
		{
			name: "fan_test",
			args: args{
				tRoom:    1000,
				tCond:    100,
				condMode: "fan",
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.tRoom, tt.args.tCond, tt.args.condMode); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
