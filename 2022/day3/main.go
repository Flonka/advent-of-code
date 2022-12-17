package main

import (
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input")

	for s.Scan() {
		line := s.Text()
		l := len(line)
		c1 := line[:(l / 2)]
		c2 := line[(l / 2):]

		getCommonPriority(c1, c2)
		break
	}

}

func getCommonPriority(c1 string, c2 string) int {

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	var psum int
	for _, v := range c1 {

		if strings.ContainsRune(c2, v) {

			fmt.Println(v)
			fmt.Printf("%s\n", v)
			psum += int(v)
		}

	}
	fmt.Println(psum)

	return psum
}
