package main

import (
	"fmt"
	"sort"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	d := input.ReadCommaSeparatedInts("input")

	ds := sort.IntSlice(d)

	ds.Sort()

	min := ds[0]
	max := ds[len(ds)-1]
	consumption := calculateConsumption(min, ds)
	for i := min + 1; i <= max; i++ {

		c := calculateConsumption(i, ds)

		if c <= consumption {
			consumption = c
		} else {
			// if trend goes other way, it should never come back
			fmt.Println("Min cons:", consumption, i)
			break
		}
	}

	fmt.Println("Fuel Consumption:", consumption)
}

func calculateConsumption(pos int, data []int) int {

	c := 0
	for i := 0; i < len(data); i++ {
		diff := data[i] - pos
		if diff < 0 {
			diff = -diff
		}

		posConsumption := (diff * (diff + 1)) / 2
		c += posConsumption
	}

	return c
}
