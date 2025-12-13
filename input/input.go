// Package input contains reusable functions and types for parsing the
// advent calendar's different problem inputs.
package input

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"log/slog"
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

// OpenFileBuffered returns a bufio.Scanner of the file path given.
// Default scanner is used, splitting input on lines.
func OpenFileBuffered(fileName string) *bufio.Scanner {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	return bufio.NewScanner(f)
}

func ReadLinesInFile(fileName string) []string {
	f := OpenFile(fileName)
	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("error closing file", "error", err)
		}
	}()
	s := bufio.NewScanner(f)
	outputLines := make([]string, 0, 200)
	for s.Scan() {
		outputLines = append(outputLines, s.Text())
	}

	return outputLines
}

// splitComma is a bufio.SplitFunc, for splitting values on comma signs.
func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
}

// ReadCommaSeparatedInts assumes the first line in a file has all the integers comma separated
func ReadCommaSeparatedInts(filePath string) []int {
	r := OpenFile(filePath)
	defer func() {
		if err := r.Close(); err != nil {
			slog.Error("error closing file", "error", err)
		}
	}()
	s := bufio.NewScanner(r)

	s.Split(splitComma)

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

func ReadCommaSeparated[T ~string](filePath string) []T {
	r := OpenFile(filePath)
	defer func() {
		if err := r.Close(); err != nil {
			slog.Error("error closing file", "error", err)
		}
	}()

	s := bufio.NewScanner(r)

	s.Split(splitComma)

	var data []T
	for s.Scan() {
		data = append(data, T(s.Text()))
	}

	return data
}

func StringsToInts(s []string) []int {
	ints := make([]int, len(s))

	for i := range len(s) {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = n
	}
	return ints
}
