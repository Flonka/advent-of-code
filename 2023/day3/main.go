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
	endCol := d.End.X + 1

	if row >= 0 {
		// above row
		for i := startCol; i <= endCol; i++ {
			p = append(p, spatial.DiscretePos2D{X: i, Y: row})
		}
	}

	// left , right
	p = append(p, spatial.DiscretePos2D{X: startCol, Y: d.End.Y})
	p = append(p, spatial.DiscretePos2D{X: endCol, Y: d.End.Y})

	row = d.End.Y + 1
	if row < size {
		// Below row
		for i := startCol; i <= endCol; i++ {
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
				// Special case on last character
				if col == size-1 {
					digits = append(digits, createDigit(digitAccumulation, row, col))
					// Reset
					lastWasDigit = false
					digitAccumulation = ""
				}

			} else { // Not a digit now
				if lastWasDigit {
					digits = append(digits, createDigit(digitAccumulation, row, col-1))
					// Reset
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

func createDigit(acc string, row int, col int) Digit {

	digitVal, err := strconv.Atoi(acc)
	pos := spatial.DiscretePos2D{X: col, Y: row}
	if err != nil {
		fmt.Println(err, pos)
		os.Exit(1)
	}

	d := Digit{Value: digitVal, End: pos}
	return d
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
