package fibo

import "testing"

func TestFib(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	var tests = []test{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		// her kan vi bare legge på nye test cases
	}

	// tc - test case
	for _, tc := range tests {
		got := Fib(tc.input)
		if got != tc.want {
			t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
		}
	}
}
