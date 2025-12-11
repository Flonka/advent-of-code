package main

import (
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

func main() {
	d := initDay("XMAS")
	d.part1()
}

type Day4 struct {
	data       spatial.DiscreteMap2D[string]
	searchWord string
}

func (d *Day4) part1() {
	// For each string pos, if its X; search for XMAS in all directions
	p := spatial.DiscretePos2D{}
	count := 0

	for x := range d.data.Width {
		p.X = x
		for y := range d.data.Height {
			p.Y = y
			s := d.data.GetValue(0, p)
			if s == string(d.searchWord[0]) {
				// start search in all directions
			}
		}
	}

	fmt.Println("Part1:", count)
}

func initDay(targetWord string) Day4 {
	lines := input.ReadLinesInFile("input.txt")
	dmap := spatial.NewDiscreteMap2D[string](140, 140, 1)

	p := spatial.DiscretePos2D{}

	for y, l := range lines {
		p.Y = y
		for x, r := range l {
			p.X = x

			var val string
			if strings.ContainsRune(targetWord, r) {
				val = string(r)
			} else {
				val = "."
			}

			dmap.SetValue(0, p, val)
		}

	}

	fmt.Println(dmap.Data)

	return Day4{
		data:       dmap,
		searchWord: targetWord,
	}
}
