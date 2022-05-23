package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code-2021/input"
)

type heightmap struct {
	data [][]int
}

func (h *heightmap) calculateRisk(lowPoints []pos) int {
	var risk int
	for _, p := range lowPoints {
		// The risk level of a low point is 1 plus its height.
		risk += h.data[p.y][p.x] + 1
	}

	return risk
}

func (h *heightmap) trim() {
	for i, row := range h.data {
		if len(row) == 0 {
			h.data = h.data[:i]
			return
		}
	}
}

type pos struct {
	x int
	y int
}

type basin struct {
	positions []pos
}

func (h *heightmap) findLowPositions() []pos {
	lowPositions := make([]pos, 0, 20)

	for y, rowData := range h.data {
		for x, val := range rowData {
			// find neighbour positions based on current x y
			neighbours := getNeighbouringPositions(x, y, len(h.data)-1, len(rowData)-1)

			// is the current value low
			lowVal := true

			for _, n := range neighbours {
				neighbourValue := h.data[n.y][n.x]
				if neighbourValue <= val {
					lowVal = false
					break
				}
			}

			if lowVal {
				lowPositions = append(lowPositions, pos{x: x, y: y})
			}

		}
	}

	return lowPositions
}

func getNeighbouringPositions(x int, y int, yMax int, xMax int) []pos {
	neighbours := make([]pos, 0, 4)

	n := pos{x: x, y: y - 1}
	if n.y >= 0 {
		neighbours = append(neighbours, n)
	}
	w := pos{x: x - 1, y: y}
	if w.x >= 0 {
		neighbours = append(neighbours, w)
	}
	e := pos{x: x + 1, y: y}
	if e.x <= xMax {
		neighbours = append(neighbours, e)
	}
	s := pos{x: x, y: y + 1}
	if s.y <= yMax {
		neighbours = append(neighbours, s)
	}

	return neighbours

}

func InitHeightmap(length int) heightmap {

	d := make([][]int, length)
	for i := range d {
		d[i] = make([]int, 0, length)
	}

	return heightmap{
		data: d,
	}
}

func main() {

	hmap := readData("input_example")

	lows := hmap.findLowPositions()
	// fmt.Println(lows)

	risk := hmap.calculateRisk(lows)
	fmt.Println("Risk:", risk)

	basins := findBasins(hmap, lows)
	fmt.Println(basins)

}

func findBasins(hmap heightmap, lowPositions []pos) []basin {

	basins := make([]basin, 0, 10)

	// Basin
	// per low pos , djikstra utåt ? , spara basin och dess pos ( applicera basin-regler), spara LUT för positioner använda i basin
	// reducera alla basin mha LUT.

	// 	A basin is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin, although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations will always be part of exactly one basin.

	// The size of a basin is the number of locations within the basin, including the low point. The example above has four basins.§

	// inBasin := make(map[pos]bool)
	for _, lowPos := range lowPositions {
		hmap.fillBasin(lowPos)
	}

	return basins

}

// Return positions forming a basin
func (h *heightmap) fillBasin(p pos) []pos {

	basinPositions := make([]pos, 0, 10)
	queue := getNeighbouringPositions(p.x, p.y, len(h.data), len(h.data[0]))


	for len(queue) > 0 {
		n = queue[0]
		queue = queue[1:]

		// Check if n is valid 
		// if so, write to basinPosition
		// add neighbours of n to queue


	}

	return basinPositions
}

func readData(p string) heightmap {
	r := input.OpenFile(p)
	defer r.Close()
	s := bufio.NewScanner(r)

	hmap := InitHeightmap(100)
	lineCount := 0
	for s.Scan() {
		l := s.Text()
		for _, strInt := range strings.Split(l, "") {
			i, err := strconv.Atoi(strInt)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			hmap.data[lineCount] = append(hmap.data[lineCount], i)

		}
		lineCount++
	}
	hmap.trim()
	return hmap
}
