package input

import (
	"bufio"
	"log"
	"os"
)

func ReadLinesInFile(fileName string) []string {
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
