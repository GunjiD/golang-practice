package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input [5]int
		want  [5]int
	}{
		{[5]int{0, 1, 2, 3, 4}, [5]int{4, 3, 2, 1, 0}},
	}

	for _, test := range tests {
		reverse(&test.input)
		if test.input != test.want {
			t.Errorf("Reverse( %v ) = %v", test.input, test.want)
		}
	}

}
