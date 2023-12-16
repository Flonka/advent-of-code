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

// magic number width and height
const size int = 140

type Digit struct {
	Value int
	End   spatial.DiscretePos2D
}

func (d *Digit) GetPerimiter() []spatial.DiscretePos2D {

	p := make([]spatial.DiscretePos2D, 0, 10)

	l := digitsOfInt(d.Value)

	// Above row , below, 2 ends
	row := d.End.Y - 1

	startCol := d.End.X - l

	if row > 0 {
		// above row
		for i := startCol; i <= d.End.X+1; i++ {
			p = append(p, spatial.DiscretePos2D{X: i, Y: row})
		}
	}

	// left , right
	p = append(p, spatial.DiscretePos2D{X: d.End.X - l - 1, Y: d.End.Y})
	p = append(p, spatial.DiscretePos2D{X: d.End.X + 1, Y: d.End.Y})

	row = d.End.Y + 1
	if row < size {
		// Below row
		for i := startCol; i <= d.End.X+1; i++ {
			p = append(p, spatial.DiscretePos2D{X: i, Y: row})
		}

	}

	return p

}

func digitsOfInt(i int) (count int) {

	for i > 0 {
		i = i / 10
		count++
	}

	return count
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

			} else { // Not a digit now
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
		for _, p := range d.GetPerimiter() {
			_, present := symbols[p]
			if present {
				sum += d.Value
			}
		}
	}

	return sum
}
