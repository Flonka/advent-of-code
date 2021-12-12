package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	fishes := readFishes()

}

type Fish struct {
	timer int
}

func readFishes() []Fish {
	r := input.OpenFile("input")
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

	var f []Fish
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		f = append(f, Fish{timer: i})
	}

	return f
}
