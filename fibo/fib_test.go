package fibo_test

import (
	"testing"

	"github.com/mathhaug/fib/fibo"
)

func TestFib(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		// Legg til flere testtilfeller her
	}

	for _, tc := range tests {
		got := fibo.Fib(tc.input)
		if got != tc.want {
			t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
		}
	}
}
