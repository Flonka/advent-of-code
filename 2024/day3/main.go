package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	d := Day3{lines: input.ReadLinesInFile("input.txt")}

	d.part1()
}

type Day3 struct {
	lines []string
}

func (d *Day3) part1() {

	// filter out valid mult strings
	mults := make([]multPair, 0, 100)

	for _, l := range d.lines {
		mults = append(mults, readMults(l)...)
	}

	var sum int

	for _, m := range mults {
		sum += m.a * m.b

	}

	fmt.Println("Part1 ", sum)
}

type multPair struct {
	a int
	b int
}

var multRex *regexp.Regexp = regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)

func readMults(s string) []multPair {

	mulstrings := multRex.FindAllString(s, -1)

	mults := make([]multPair, 0, len(mulstrings))

	for _, ms := range mulstrings {
		cIndex := strings.Index(ms, ",")
		as := ms[4:cIndex]
		bs := ms[cIndex+1 : len(ms)-1]

		a, _ := strconv.Atoi(as)
		b, _ := strconv.Atoi(bs)

		mults = append(mults, multPair{a: a, b: b})
	}

	return mults
}
