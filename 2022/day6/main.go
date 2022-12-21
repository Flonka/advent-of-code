package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	s := input.OpenFileBuffered("input")
	s.Split(bufio.ScanRunes)

	var scanCount int
	var last4 string
	// Fill chars
	for i := 0; i < 4; i++ {
		if !s.Scan() {
			os.Exit(1)
		}
		scanCount++

		last4 = last4 + s.Text()
	}

	// Compare
	if isMarker(last4) {
		fmt.Println(scanCount)
	}

	for s.Scan() {
		scanCount++
		c := s.Text()
		// Update last4
		last4 = last4[1:] + c

		if isMarker(last4) {
			fmt.Println("Found it")
			fmt.Println(last4)
			break
		}
	}

	fmt.Println("scan count:", scanCount)
}

func isMarker(chars string) bool {

	unique := 0
	for i := 0; i < len(chars); i++ {
		s := chars[i : i+1]
		count := strings.Count(chars, s)
		if count == 1 {
			unique++
		} else {
			return false
		}
	}

	return unique == len(chars)
}
