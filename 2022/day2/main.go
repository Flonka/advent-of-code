package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	r := input.OpenFile("input.txt")
	defer r.Close()

	s := bufio.NewScanner(r)
	
	var points int
	var points2 int
	for s.Scan() {
		line := s.Text()
		round := strings.Split(line, " ")
		points += playRound(round[0], round[1])

		var move string
		switch round[1] {
		case "X":
			// lose
			move = loseLUT[round[0]]
		case "Y":
			// draw
			switch round[0] {
			case "A":
				move = "X"
			case "B":
				move = "Y"
			case "C":
				move = "Z"
			}
		case "Z":
			// win
			move = winLUT[round[0]]
		}
		points2 += playRound(round[0], move)
	}

	fmt.Println("Part1:", points)
	fmt.Println("Part2:", points2)
}

// your move ponint
var youMovePointLUT = map[string]int{
	// Rock
	"X": 1,
	// Paper
	"Y": 2,
	// Scissors
	"Z": 3,
}

// Get move to win against opponent, opponent move is the key
var winLUT = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

// Get move to lose against opponent, opponent move is the key
var loseLUT = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

// Return points
func playRound(opponent string, you string) int {

	// add movepoint
	points := youMovePointLUT[you]

	//  check win / draw / lose

	if winLUT[opponent] == you {
		// Win
		points += 6
	} else if loseLUT[opponent] == you {
		// Lose
		points += 0
	} else {
		// Draw
		points += 3
	}
	return points
}
