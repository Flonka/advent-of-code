package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	d := readInts()

	w := len(d[0]) - 1
	h := len(d) - 1
	fmt.Println("width", w)
	fmt.Println("height", h)

	visibleTrees := 0
	maxScenicScore := 0
	for y, row := range d {
		for x, tree := range row {

			atE := atEdge(x, y, w, h)
			// Is tree at edge?-> visible
			// else
			// Traverse data for given position in all directions
			if atE || visibleToOutside(x, y, w, h, tree, d) {
				visibleTrees++
			}

			if !atE {
				s := calculateScenicScore(x, y, w, h, tree, d)

				if s > maxScenicScore {
					maxScenicScore = s
				}
			}

		}
	}

	fmt.Println("Visible trees from outside", visibleTrees)
	fmt.Println("MaxScenicScore", maxScenicScore)

}

func calculateScenicScore(x, y, w, h, treeHeight int, data [][]int) int {

	var leftVis, rightVis, upVis, downVis int
	compareHeight := treeHeight
	for xPos := x - 1; xPos >= 0; xPos-- {
		leftVis++
		t := data[y][xPos]
		if t >= compareHeight {
			break
		}
	}

	for xPos := x + 1; xPos <= w; xPos++ {
		rightVis++
		t := data[y][xPos]
		if t >= compareHeight {
			break
		}
	}

	for yPos := y - 1; yPos >= 0; yPos-- {
		upVis++
		t := data[yPos][x]
		if t >= compareHeight {
			break
		}
	}

	for yPos := y + 1; yPos <= h; yPos++ {
		downVis++
		t := data[yPos][x]
		if t >= compareHeight {
			break
		}
	}

	return leftVis * rightVis * upVis * downVis
}

func visibleToOutside(x, y, w, h, treeHeight int, data [][]int) bool {

	// march left, right, up, down

	blocked := false
	compareHeight := treeHeight
	for xPos := x - 1; xPos >= 0; xPos-- {
		t := data[y][xPos]
		if t >= compareHeight {
			blocked = true
			break
		}
	}

	if !blocked {
		return true
	}

	// Reset
	blocked = false
	for xPos := x + 1; xPos <= w; xPos++ {
		t := data[y][xPos]
		if t >= compareHeight {
			blocked = true
			break
		}
	}

	if !blocked {
		return true
	}

	blocked = false
	for yPos := y - 1; yPos >= 0; yPos-- {
		t := data[yPos][x]
		if t >= compareHeight {
			blocked = true
			break
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for yPos := y + 1; yPos <= h; yPos++ {
		t := data[yPos][x]
		if t >= compareHeight {
			blocked = true
			break
		}
	}

	return !blocked
}

func atEdge(x, y, w, h int) bool {
	if x == 0 || x == w {
		return true
	}

	if y == 0 || y == h {
		return true
	}

	return false
}

func readInts() [][]int {

	rows := make([][]int, 0, 100)
	s := input.OpenFileBuffered("input")

	for s.Scan() {
		ints := make([]int, 0, 100)

		line := s.Text()

		for _, v := range strings.Split(line, "") {
			i, _ := strconv.Atoi(v)
			ints = append(ints, i)
		}
		rows = append(rows, ints)
	}

	return rows
}
