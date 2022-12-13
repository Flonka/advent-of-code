package day1

import (
	"log"
	"strconv"

	"github.com/Flonka/advent-of-code/input"
)

func GetDepthData(filePath string) []int {

	lines := input.ReadLinesInFile(filePath)

	depthData := make([]int, 0, len(lines))
	for _, v := range lines {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		depthData = append(depthData, i)
	}

	return depthData
}
