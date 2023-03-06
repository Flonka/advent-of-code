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

	for i, m := range monkeys {
		fmt.Println("Monkey", i, m)
	}

	/*
		The monkeys take turns inspecting and throwing items. On a single monkey's turn, it inspects and throws all of the items it is holding one at a time and in the order listed. Monkey 0 goes first, then monkey 1, and so on until each monkey has had one turn. The process of each monkey taking a single turn is called a round.

		When a monkey throws an item to another monkey, the item goes on the end of the recipient monkey's list. A monkey that starts a round with no items could end up inspecting and throwing many items by the time its turn comes around. If a monkey is holding no items at the start of its turn, its turn ends.

	*/

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
