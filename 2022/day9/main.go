package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/vector"
)

type state struct {
	head vector.Vec2
	tail vector.Vec2
}

func main() {
	s := input.OpenFileBuffered("input")

	ropeState := state{}

	visited := make(map[vector.Vec2]int)

	// 0,0 visited (startpos)
	visited[vector.Vec2{}]++

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

func getDir(s string) vector.Vec2 {
	v := vector.Vec2{}
	switch s {
	case "U":
		v.Y = 1
	case "D":
		v.Y = -1
	case "L":
		v.X = -1
	case "R":
		v.X = 1
	default:
		fmt.Println("Unhandled dir", s)
		os.Exit(1)
	}

	return v
}

// updateState updates the state and returns new visited tail positions
func updateState(s *state, direction vector.Vec2, amount int) []vector.Vec2 {

	tails := make([]vector.Vec2, 0, amount)

	fmt.Println(direction, amount)
	fmt.Println("state", s)

	d := vector.Vec2{}

	for i := 0; i < amount; i++ {
		// Progress head
		h := s.head
		t := s.tail

		h.Add(&direction)

		fmt.Println(d, h)

		d = h
		d.Subtract(&t)

		// If distance greater than 1 , tail needs to move
		fmt.Println("diff", d, "l", d.Length())

		

		// Update state
		s.head = h
		s.tail = t
	}
	fmt.Println("state", s)

	os.Exit(1)

	return tails
}
