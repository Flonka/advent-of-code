package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type vec2 struct {
	x int
	y int
}

type state struct {
	head vec2
	tail vec2
}

func main() {
	s := input.OpenFileBuffered("input")

	ropeState := state{}

	visited := make(map[vec2]int)

	// 0,0 visited (startpos)
	visited[vec2{}]++

	for s.Scan() {
		line := s.Text()
		splits := strings.Split(line, " ")

		dirString := splits[0]
		amount, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		direction := getDir(dirString)

		newTailPositions := updateState(&ropeState, direction, amount)
		for _, p := range newTailPositions {
			visited[p]++
		}

	}

	for k, v := range visited {
		fmt.Println(k, v)
	}
}

func getDir(s string) vec2 {
	v := vec2{}
	switch s {
	case "U":
		v.y = 1
	case "D":
		v.y = -1
	case "L":
		v.x = -1
	case "R":
		v.x = 1
	default:
		fmt.Println("Unhandled dir", s)
		os.Exit(1)
	}

	return v
}

// updateState updates the state and returns new visited tail positions
func updateState(s *state, direction vec2, amount int) []vec2 {

	tails := make([]vec2, 0, amount)
	// fmt.Println(direction, amount)

	return tails
}
