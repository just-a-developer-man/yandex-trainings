package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple triangle",
			args: args{
				a: 3,
				b: 2,
				c: 4,
			},
			want: "YES",
		},
		{
			name: "invalid triangle",
			args: args{
				a: 3,
				b: 2,
				c: 6,
			},
			want: "NO",
		},
		{
			name: "equal sides triangle",
			args: args{
				a: 3,
				b: 3,
				c: 3,
			},
			want: "YES",
		},
		{
			name: "isosceles triangle",
			args: args{
				a: 4,
				b: 4,
				c: 5,
			},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
