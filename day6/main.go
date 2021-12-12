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

	fish := readFish()

	days := 256

	for day := 0; day < days; day++ {
		babyFish := 0
		for i := 0; i < len(fish); i++ {
			f := &fish[i]
			if f.PassOneDay() {
				babyFish++
			}
		}

		// spawn new fish
		for b := 0; b < babyFish; b++ {
			fish = append(fish, Fish{timer: 8})
		}
	}

	fmt.Printf("Fish after %v days %v\n", days, len(fish))
}

type Fish struct {
	timer int
}

// PassOneDay passes the fish one day returns true if it spawned a new fish
func (f *Fish) PassOneDay() bool {
	f.timer--
	if f.timer < 0 {
		f.timer = 6
		return true
	}
	return false
}

func readFish() []Fish {
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
