package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input  []int
		inputN int
		want   []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5, 0, 1}},
	}

	for _, test := range tests {
		tmp := make([]int, len(test.input))
		tmp = rotate(test.input, test.inputN)
		if !reflect.DeepEqual(tmp, test.want) {
			t.Errorf("Rotate( %v ) = %v", tmp, test.want)
		}
	}

}
