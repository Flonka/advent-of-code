package main

import (
	"slices"
	"testing"
)

func TestPart1(t *testing.T) {

	tests := []struct {
		input  []int
		output int
	}{
		{[]int{0, 3, 6, 9, 12, 15}, 18},
		{[]int{1, 3, 6, 10, 15, 21}, 28},
		{[]int{10, 13, 16, 21, 30, 45}, 68},
	}

	for _, test := range tests {
		result := part1(test.input)
		if result != test.output {
			t.Error(result, " is not ", test.output)
		}

	}
}

func TestDiffSlice(t *testing.T) {

	r := createDiffSlice([]int{2, 2, -6})

	if slices.Compare(r, []int{0, -8}) != 0 {
		t.Fail()
	}

}
