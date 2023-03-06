package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items        []int
	operation    func(int) int
	inspectCount int
}

/* NewMonkeyFromLines takes line input to create a new Monkey , the structure looks like this example:
Starting items: 98, 97, 98, 55, 56, 72
 Operation: new = old * 13
 Test: divisible by 11
   If true: throw to monkey 4
   If false: throw to monkey 7
*/
func NewMonkeyFromLines(lines []string) Monkey {

	// fmt.Println(lines)
	// items
	itemString := lines[0][strings.Index(lines[0], ":")+1:]
	itemStrings := strings.Split(itemString, ",")
	items := make([]int, 0, len(itemStrings))
	for _, s := range itemStrings {
		v, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		items = append(items, v)
	}
	// fmt.Println(lines[0], "->", items)

	// Operation 
	opLine := lines[1]
	fmt.Println(opLine)

	return Monkey{
		items: items,
	}
}

type test struct {
	parameter     int
	successTarget int
	failureTarget int
}
