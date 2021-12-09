package input

import (
	"bufio"
	"io"
	"log"
	"os"
)

func OpenFile(fileName string) io.ReadCloser {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	return f
}

func ReadLinesInFile(fileName string) []string {
	f := OpenFile(fileName)
	defer f.Close()
	s := bufio.NewScanner(f)
	outputLines := make([]string, 0, 200)
	for s.Scan() {
		outputLines = append(outputLines, s.Text())
	}

	return outputLines
}
