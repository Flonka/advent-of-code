package day2

import (
	"log"
	"strconv"
	"strings"
)

type Command struct {
	// Direction is a vector 0 = horizontal , 1 = vertical
	Direction [2]int
}

func NewCommandFromInputLine(line string) Command {

	s := strings.Split(line, " ")

	amount, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}

	dir := [2]int{0, 0}
	switch s[0] {
	case "forward":
		dir[0] = 1 * amount
	case "down":
		dir[1] = 1 * amount
	case "up":
		dir[1] = -1 * amount
	default:
		log.Fatal("not implemented direction:", s[0])
	}

	return Command{
		Direction: dir,
	}
}
