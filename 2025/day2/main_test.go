package main

import (
	"fmt"
	"testing"
)

func TestInvalid(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{1212, true},
		{200, false},
		{123123, true},
		{1231234, false},
		{38593859, true},
	}

	for _, testCase := range tests {
		if testCase.expected != isInvalid(testCase.id) {
			fmt.Println("failed: ", testCase)
			t.Fail()
		}
	}
}
