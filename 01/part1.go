package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFileName string = "input"

func main() {
	depthData := getDepthData()

	lastD := depthData[0]
	incs := 0
	for _, d := range depthData[1:] {

		if d > lastD {
			incs++
		}
		lastD = d
	}

	fmt.Println(incs)
}

func getDepthData() []int {

	lines := readLinesInFile(inputFileName)

	depthData := make([]int, 0, len(lines))
	for _, v := range lines {
		if len(v) == 0 {
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		depthData = append(depthData, i)
	}

	return depthData
}

func readLinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	s := bufio.NewScanner(f)

	outputLines := make([]string, 0, 200)
	for s.Scan() {
		outputLines = append(outputLines, s.Text())
	}

	return outputLines

}
