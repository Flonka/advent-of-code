package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

type Digit struct {
	Value int
	End   spatial.DiscretePos2D
}

func main() {

	cli.Default()

	// s := input.OpenFileBuffered("input.txt")
	fileReader := input.OpenFile("input.txt")
	defer fileReader.Close()
	r := bufio.NewReader(fileReader)

	symbols := make(map[spatial.DiscretePos2D]string, 50)
	digits := make([]Digit, 0, 100)

	var row, col int
	// Pass 1 - find all symbol positions
	for {
		line, err := r.ReadString('\n')
		// Symbol is not digit, and not . and not newline, Readstring includes delimiter
		col = 0

		lastWasDigit := false
		var digitAccumulation string
		for _, s := range line {
			if unicode.IsDigit(s) {
				lastWasDigit = true
				digitAccumulation += string(s)

			} else {
				// Not a digit now
				if lastWasDigit {

					digitVal, err := strconv.Atoi(digitAccumulation)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					digits = append(digits, Digit{Value: digitVal, End: spatial.DiscretePos2D{X: col - 1, Y: row}})
					lastWasDigit = false
					digitAccumulation = ""
				}

				if s != '.' && s != '\n' {
					// Must be symbol
					// fmt.Println(fmt.Sprintf("'%v'", string(s)))
					symbols[spatial.DiscretePos2D{X: col, Y: row}] = string(s)

				}
			}

			col++
		}

		if err != nil {
			break
		}

		row++
	}

	fmt.Println("part1", part1(digits, symbols))

}

func part1(digits []Digit, symbols map[spatial.DiscretePos2D]string) int {

	sum := 0

	for _, d := range digits {

	}

	return sum
}
