package main

import (
	"bufio"
	"fmt"
	"unicode"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
)

func main() {

	cli.Default()

	// s := input.OpenFileBuffered("input.txt")
	fileReader := input.OpenFile("input.txt")
	defer fileReader.Close()
	r := bufio.NewReader(fileReader)

	// Pass 1 - find all symbol positions
	for {
		line, err := r.ReadString('\n')
		// Symbol is not digit, and not . and not newline, Readstring includes delimiter

		for _, s := range line {
			if s == '.' || unicode.IsDigit(s) || s == '\n' {
				continue
			}
			// Print symbol
			fmt.Println(fmt.Sprintf("'%v'", string(s)))
		}

		if err != nil {
			break
		}
	}

}
