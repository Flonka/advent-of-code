package main

import (
	"bufio"
	"fmt"
	"unicode"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

func main() {

	cli.Default()

	// s := input.OpenFileBuffered("input.txt")
	fileReader := input.OpenFile("input.txt")
	defer fileReader.Close()
	r := bufio.NewReader(fileReader)

	symbols := make(map[spatial.DiscretePos2D]string, 50)

	var row, col int
	// Pass 1 - find all symbol positions
	for {
		line, err := r.ReadString('\n')
		// Symbol is not digit, and not . and not newline, Readstring includes delimiter
		col = 0
		for _, s := range line {

			if s == '.' || unicode.IsDigit(s) || s == '\n' {
				continue
			}
			// Print symbol
			// fmt.Println(fmt.Sprintf("'%v'", string(s)))
			symbols[spatial.DiscretePos2D{X: col, Y: row}] = string(s)

			col++
		}

		if err != nil {
			break
		}

		row++
	}

	fmt.Println(symbols)

}
