package main

import (
	"log/slog"
	"strings"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
)

func main() {
	cli.Default()

	var ops []string
	var numbers [][]int

	input.HandleInputLines(func(lineCount int, lineIndex int, line string) {
		// last line.
		if lineIndex+1 == lineCount {
			ops = strings.Fields(line)
		} else {
			numbers = append(numbers, make([]int, 0))
			numbers[lineIndex] = input.StringsToInts(strings.Fields(line))
		}
	})

	slog.Info("Read data", "ops", len(ops), "numbers", len(numbers))

	resultRow := make([]int, len(ops))

	for _, row := range numbers {
		for columnIdx, value := range row {
			operation := ops[columnIdx]
			switch operation {
			case "+":
				resultRow[columnIdx] += value
			case "*":
				if resultRow[columnIdx] == 0 {
					resultRow[columnIdx] = 1
				}
				resultRow[columnIdx] *= value
			default:
				slog.Error("not handling operation", "operation", operation)
			}
		}
	}

	sum := 0
	for _, v := range resultRow {
		sum += v
	}
	slog.Info("Part1", "sum", sum)
}
