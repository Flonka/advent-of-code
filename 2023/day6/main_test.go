package main

import "testing"

func Test(t *testing.T) {

	r := bruteForcePart1(7, 9)
	if r != 4 {
		t.Error("result is not 4:", r)
	}

}
