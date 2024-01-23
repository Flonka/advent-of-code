package main

import (
	"testing"
)

func TestHandEvaluation(t *testing.T) {
	tests := []struct {
		input    string
		expected HandType
	}{
		{"38JKA 0", HighCard},
		{"A9JK9 0", OnePair},
		{"A9AK9 0", TwoPair},
		{"333KA 0", ThreeOfAKind},
		{"333KK 0", FullHouse},
		{"22224 0", FourOfAKind},
		{"KKKKK 0", FiveOfAKind},
	}

	for _, test := range tests {
		h := HandFromLine(test.input)

		if h.Type != test.expected {
			t.Error("incorrect type", h, test)

		}
	}

}
