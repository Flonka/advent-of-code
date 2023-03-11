package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Flonka/advent-of-code/2022/day11"
	"github.com/Flonka/advent-of-code/input"
)

func main() {

	// Read monkeys
	monkeys := readData("run/input.txt")

	rounds := 10000

	reducer := 1
	for _, m := range monkeys {
		reducer *= m.GetDivisor()
	}

	fmt.Println("Reducer", reducer)
	rBig := big.NewInt(int64(reducer))

	for i := 0; i < rounds; i++ {

		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]

			thrown := m.InspectItems(false)
			for _, it := range thrown {
				it.Item.Mod(&it.Item, rBig)
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
