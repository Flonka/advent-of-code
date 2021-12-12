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
	timerMap := createTimerMap(fish)

	days := 256

	for day := 0; day < days; day++ {

		// forward timer one day
		for t := 1; t <= 9; t++ {
			timerMap[t-1] = timerMap[t]
		}

		// Spawn the new baby fish and reset timer on fish
		babies := timerMap[0]
		timerMap[9] = babies
		timerMap[7] += babies
		timerMap[0] = 0

	}

	var fishCount int
	fmt.Println(timerMap)
	for _, n := range timerMap {
		fishCount += n
	}
	fmt.Printf("Fishes after %v days: %v", days, fishCount)
}

func createTimerMap(f []Fish) [10]int {

	m := [10]int{}

	for i := range f {
		fish := &f[i]
		m[fish.timer+1]++
	}

	return m
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
