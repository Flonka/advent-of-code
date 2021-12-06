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

	lastD := depthData[0:3]
	lastSum := sum3(lastD)
	incs := 0
	for i, _ := range depthData[1 : len(depthData)-2] {

		sumD := sum3(depthData[i+1 : i+1+3])
		if sumD > lastSum {
			incs++
		}
		lastSum = sumD
	}

	fmt.Println(incs)
}

func sum3(s []int) int {
	if len(s) != 3 {
		log.Fatal("wrong size", s)
	}
	return s[0] + s[1] + s[2]
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
