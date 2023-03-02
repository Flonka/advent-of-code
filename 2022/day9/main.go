package main

import (
	"fmt"
	"math"
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

type multiKnotState struct {
	knots []vector.Vec2
}

func main() {
	s := input.OpenFileBuffered("input")

	ropeState := state{}

	multiState := multiKnotState{
		knots: make([]vector.Vec2, 10),
	}

	visited := make(map[vector.Vec2]int)
	visited2 := make(map[vector.Vec2]int)

	// 0,0 visited (startpos)
	visited[vector.Vec2{}]++
	visited2[vector.Vec2{}]++

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

		t2Pos := updateMultiState(&multiState, direction, amount)
		for _, p := range t2Pos {
			visited2[p]++
		}

	}

	fmt.Println("Part1", len(visited))
	fmt.Println("Part2", len(visited2))

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

func updateMultiState(s *multiKnotState, direction vector.Vec2, amount int) []vector.Vec2 {
	tails := make([]vector.Vec2, 0, amount)
	d := vector.Vec2{}

	for i := 0; i < amount; i++ {
		// Move actual head with direction , then perform the "follow head logic"
		head := s.knots[0]
		head.Add(&direction)
		s.knots[0] = head

		// check diff and update
		for knotIndex := 0; knotIndex < len(s.knots)-1; knotIndex++ {
			h := s.knots[knotIndex]
			t := s.knots[knotIndex+1]


			// Calculate difference d
			d = h
			d.Subtract(&t)

			// If any axis distance greater than 1 , tail needs to move
			if d.X > 1 || d.X < -1 || d.Y > 1 || d.Y < -1 {
				fixD(&d)
				t.Add(&d)
			}

			// Update state
			// s.knots[knotIndex] = h
			s.knots[knotIndex+1] = t

			if knotIndex+1 == len(s.knots)-1 {
				tails = append(tails, t)
			}
		}
	}

	return tails
}

// updateState updates the state and returns new visited tail positions
func updateState(s *state, direction vector.Vec2, amount int) []vector.Vec2 {

	tails := make([]vector.Vec2, 0, amount)

	d := vector.Vec2{}

	for i := 0; i < amount; i++ {
		// Progress head
		h := s.head
		t := s.tail

		// Move head with direction
		h.Add(&direction)

		// Calculate difference d
		d = h
		d.Subtract(&t)

		// If any axis distance greater than 1 , tail needs to move
		if d.X > 1 || d.X < -1 || d.Y > 1 || d.Y < -1 {
			fixD(&d)
			t.Add(&d)
		}

		// Update state
		s.head = h
		s.tail = t

		tails = append(tails, t)
	}

	return tails
}

// Reduce v to move one step , or diagonal one step per axis
func fixD(v *vector.Vec2) {
	xNeg := math.Signbit(float64(v.X))
	yNeg := math.Signbit(float64(v.Y))

	if v.X != 0 {
		v.X /= v.X
	}
	if v.Y != 0 {
		v.Y /= v.Y
	}

	if xNeg {
		v.X = -v.X
	}
	if yNeg {
		v.Y = -v.Y
	}
}
