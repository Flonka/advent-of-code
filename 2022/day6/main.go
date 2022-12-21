package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	fmt.Println("Part1", scanUnique(4))
	fmt.Println("Part2", scanUnique(14))
}

func scanUnique(n int) int {
	s := input.OpenFileBuffered("input")

	s.Split(bufio.ScanRunes)
	var scanCount int
	var compString string
	// Fill chars
	for i := 0; i < n; i++ {
		if !s.Scan() {
			os.Exit(1)
		}
		scanCount++

		compString = compString + s.Text()
	}

	if len(compString) != n {
		log.Fatal("Not correct length")
	}

	// Compare
	if isMarker(compString) {
		return scanCount
	}

	for s.Scan() {
		scanCount++
		c := s.Text()
		// Update string
		compString = compString[1:] + c

		if isMarker(compString) {
			fmt.Println("Found it", compString)
			break
		}
	}

	return scanCount
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
