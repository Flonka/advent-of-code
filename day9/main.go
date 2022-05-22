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
				if neighbourValue < val {
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
	for i, _ := range d {
		d[i] = make([]int, 0, length)
	}

	return heightmap{
		data: d,
	}
}

func main() {

	hmap := readData("input_example")
	fmt.Println(hmap)

	lows := hmap.findLowPositions()
	fmt.Println(lows)



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
