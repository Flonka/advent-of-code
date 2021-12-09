package main

import (
	"fmt"
	"testing"
)

func TestPoints(t *testing.T) {
	l := Line{
		start: NewPosFromString("419,112"),
		end: NewPosFromString("419,109"),
	}

	p := l.Points()
	fmt.Println(p)
}