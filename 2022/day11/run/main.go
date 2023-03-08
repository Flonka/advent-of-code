package main

import (
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code/2022/day11"
	"github.com/Flonka/advent-of-code/input"
)

func main() {

	// Read monkeys
	monkeys := readData("input.txt")

	rounds := 20

	for i := 0; i < rounds; i++ {

		fmt.Println("Round", i+1)

		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			thrown := m.InspectItems()
			for _, it := range thrown {
				monkeys[it.Monkey].Items = append(monkeys[it.Monkey].Items, it.Item)
			}
			monkeys[j] = m
		}

	}
	for j := 0; j < len(monkeys); j++ {
		fmt.Println(j, monkeys[j].InspectCount)
	}
}

func readData(f string) []day11.Monkey {
	s := input.OpenFileBuffered(f)

	monkeys := make([]day11.Monkey, 0, 10)
	for s.Scan() {
		l := s.Text()

		if strings.HasPrefix(l, "Monkey") {
			monkeyLines := make([]string, 0, 5)
			for i := 0; i < 5; i++ {
				s.Scan()
				monkeyLines = append(monkeyLines, s.Text())
			}

			monkeys = append(monkeys, day11.NewMonkeyFromLines(monkeyLines))
		}
	}

	return monkeys
}
