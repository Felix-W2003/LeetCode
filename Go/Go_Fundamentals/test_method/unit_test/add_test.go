package main

import (
	"testing"
)

func Test_reduce(t *testing.T) {
}
func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Addtion:Positive numbers",
			args: args{a: 1, b: 2},
			want: 3,
		},
		{
			name: "Addtion:Negative numbers",
			args: args{a: -3, b: -2},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
func setupTest(t *testing.T) (int, int) {
	t.Helper()
	t.Errorf("setupTest error")
	return 10, 20
}
func TestAdd(t *testing.T) {
	a, b := setupTest(t)
	got := add(a, b)
	want := 30
	if got != want {
		t.Errorf("add() = %v, want %v", got, want)
	}

}
