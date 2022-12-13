package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/Flonka/advent-of-code/input"
)

type elf struct {
	Food []int
}

func (e *elf) CalorieCount() int {

	sum := 0
	for _, v := range e.Food {
		sum += v
	}

	return sum
}

func main() {

	r := input.OpenFile("input.txt")

	defer r.Close()
	s := bufio.NewScanner(r)

	elves := make([]elf, 0, 10)
	elfTmp := elf{}

	maxCal := 0
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			// new elf

			if elfTmp.CalorieCount() > maxCal {
				maxCal = elfTmp.CalorieCount()
			}

			elves = append(elves, elfTmp)
			elfTmp = elf{}

		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			elfTmp.Food = append(elfTmp.Food, i)
		}
	}

	fmt.Println("Max calories by elf:", maxCal)

	// Sort elves by calorie count
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].CalorieCount() > elves[j].CalorieCount()
	})

	var calsum int
	for i := 0; i < 3; i++ {
		calsum += elves[i].CalorieCount()
		fmt.Println(i+1, elves[i].CalorieCount())
	}

	fmt.Println("sum", calsum)

}
