package main

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
)

type Bag struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id      int
	Reveals []RevealSet
}

type RevealSet struct {
	Red   int
	Green int
	Blue  int
}

func GetPowerOfSets(game *Game) int {
	r := make([]int, 0, len(game.Reveals))
	g := make([]int, 0, len(game.Reveals))
	b := make([]int, 0, len(game.Reveals))
	for _, rev := range game.Reveals {
		r = append(r, rev.Red)
		g = append(g, rev.Green)
		b = append(b, rev.Blue)
	}

	pb := Bag{
		Red:   slices.Max(r),
		Green: slices.Max(g),
		Blue:  slices.Max(b),
	}

	return pb.Red * pb.Green * pb.Blue
}

func main() {

	cli.Default()

	bag1 := Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	s := input.OpenFileBuffered("input.txt")
	part1 := 0
	part2 := 0
	for s.Scan() {

		l := s.Text()

		g := GameFromLine(l)
		if IsGamePossible(g, bag1) {
			part1 += g.Id
		}

		// Part 2
		// Find minimum of cubes needed per game
		part2 += GetPowerOfSets(&g)

	}

	slog.Info("Part 1 Result", "result", part1)
	slog.Info("Part 2 Result", "result", part2)
}

// IsGamePossible returns if the game is possible with the given bag.
func IsGamePossible(g Game, b Bag) bool {

	for _, r := range g.Reveals {
		if r.Red > b.Red || r.Green > b.Green || r.Blue > b.Blue {
			return false
		}
	}

	return true
}

func GameFromLine(line string) Game {

	s1 := strings.Split(line, ":")
	gameString := strings.Split(s1[0], " ")[1]
	gameId, err := strconv.Atoi(gameString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return Game{Id: gameId, Reveals: RevealSetsFromString(s1[1])}

}

func RevealSetsFromString(in string) []RevealSet {

	returnSets := make([]RevealSet, 0)
	sets := strings.Split(in, ";")
	for _, set := range sets {
		rSet := RevealSet{}
		set = strings.TrimSpace(set)
		for _, colString := range strings.Split(set, ",") {

			colString = strings.TrimSpace(colString)
			colS := strings.Split(colString, " ")
			count, err := strconv.Atoi(colS[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			color := colS[1]
			switch color {
			case "red":
				rSet.Red = count
			case "green":
				rSet.Green = count
			case "blue":
				rSet.Blue = count

			}
		}
		returnSets = append(returnSets, rSet)

	}

	return returnSets
}
