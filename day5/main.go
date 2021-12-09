package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	lines := readData()
	result := make(map[Pos]int)
	for i := range lines {
		line := lines[i]

		if (line.start.x == line.end.x) || line.start.y == line.end.y {
			points := line.Points()
			for j := range points {
				result[points[j]]++
			}
		}

	}

	var c int
	for _, v := range result {
		if v >= 2 {
			c++
		}
	}
	fmt.Println(c)

}

type Line struct {
	start Pos
	end   Pos
}

func (l *Line) Points() []Pos {

	var p []Pos
	xdiff := l.end.x - l.start.x
	ydiff := l.end.y - l.start.y

	xmult := 1
	if xdiff < 0 {
		xmult = -1
	}
	ymult := 1
	if ydiff < 0 {
		ymult = -1
	}

	if xdiff != 0 {
		for i := 0; i <= xmult*xdiff; i++ {
			p = append(p, Pos{x: l.start.x + xmult*i, y: l.start.y})
		}
	}

	if ydiff != 0 {
		for i := 0; i <= ymult*ydiff; i++ {
			p = append(p, Pos{x: l.start.x, y: l.start.y + ymult*i})
		}

	}

	return p
}

type Pos struct {
	x int
	y int
}

// NewPosFromString "x,y"
func NewPosFromString(s string) Pos {
	ss := strings.Split(s, ",")

	x, err := strconv.Atoi(ss[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(ss[1])
	if err != nil {
		log.Fatal(err)
	}

	return Pos{
		x: x,
		y: y,
	}

}

// NewLineFromLine "x,y -> x,y"
func NewLineFromLine(inputLine string) Line {
	f := strings.Fields(inputLine)

	return Line{
		start: NewPosFromString(f[0]),
		end:   NewPosFromString(f[2]),
	}

}
func readData() []Line {
	r := input.OpenFile("input")
	defer r.Close()
	s := bufio.NewScanner(r)
	var lines []Line
	for s.Scan() {
		l := s.Text()
		lines = append(lines, NewLineFromLine(l))
	}

	return lines
}
