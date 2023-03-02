package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input")

	X := 1
	cycleCheck := 20
	cycleCount := 1
	runCycles := 0

	signalSum := 0

	for s.Scan() {
		line := s.Text()
		splits := strings.Split(line, " ")
		cmd := splits[0]

		var V int
		switch cmd {
		case "addx":
			am, err := strconv.Atoi(splits[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			runCycles = 2
			V = am

		case "noop":
			runCycles = 1
		default:
			os.Exit(1)
		}

		for runCycles > 0 {
			if cycleCount == cycleCheck {
				signal := cycleCount * X
				// fmt.Println(cycleCount, signal)
				signalSum += signal
				cycleCheck += 40
			}

			// Check if sprite X (middle pos) +-1 is the current row pixel index
			rowPix := cycleCount % 40
			if rowPix == 0 {
				rowPix = 39
			} else {
				rowPix = rowPix -1 
			}

			drawValue := "."
			if (X-1) == rowPix || X == rowPix || (X+1) == rowPix {
				drawValue = "#"
			}

			fmt.Print(drawValue)
			if rowPix == 39 {
				fmt.Print("\n")
			}

			cycleCount++
			runCycles--
		}

		X += V

	}

	fmt.Println("part1", signalSum)

}
