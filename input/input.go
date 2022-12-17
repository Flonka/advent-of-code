package input

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) io.ReadCloser {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	return f
}

func OpenFileBuffered(fileName string) *bufio.Scanner {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	return bufio.NewScanner(f)
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

// ReadCommaSeparatedInts assumes the first line in a file has all the integers comma separated
func ReadCommaSeparatedInts(fiePath string) []int {
	r := OpenFile(fiePath)
	defer r.Close()

	s := bufio.NewScanner(r)

	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		commaIndex := bytes.IndexByte(data, ',')
		newLineIndex := bytes.IndexByte(data, '\n')
		if atEOF {
			return len(data), nil, nil
		}
		i := commaIndex

		if commaIndex == -1 && newLineIndex != -1 {
			i = newLineIndex
		}

		return i + 1, data[:i], nil

	})

	var data []int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, i)
	}

	return data
}
