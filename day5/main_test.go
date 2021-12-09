package main

import (
	"fmt"
	"testing"
)

func TestPoints(t *testing.T) {
	l := Line{
		start: NewPosFromString("419,110"),
		end:   NewPosFromString("419,109"),
	}

	p := l.StraightPoints()

	if len(p) != 2 {
		t.Fail()
	}
}

func TestDiagPoints(t *testing.T) {

	l := Line{
		start: NewPosFromString("9,7"),
		end:   NewPosFromString("7,9"),
	}
	fmt.Println(l)

	p := l.DiagonalPoints()
	fmt.Println(p)
}
