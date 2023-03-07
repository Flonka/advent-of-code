package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items        []int
	operation    operation
	inspectCount int
	test         test
}
type operation struct {
	fun  func(int, int) int
	oldA bool
	oldB bool
	a    int
	b    int
}

func (o *operation) Eval(old int) int {
	var a, b int
	if o.oldA {
		a = old
	} else {
		a = o.a
	}
	if o.oldB {
		b = old
	} else {
		b = o.b
	}
	return o.fun(a, b)
}

type test struct {
	parameter     int
	successTarget int
	failureTarget int
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
	operatorStrings := strings.Split(opLine[strings.Index(opLine, "= ")+2:], " ")
	fmt.Println(operatorStrings)
	operator := operatorStrings[1]
	var opFun func(int, int) int
	switch operator {
	case "*":
		opFun = multOp
	case "+":
		opFun = addOp

	default:
		fmt.Println("Unhandled operator", operator)
		os.Exit(1)
	}
	// Check if "old" for A / B
	opA, err := strconv.Atoi(operatorStrings[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	opB, err := strconv.Atoi(operatorStrings[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return Monkey{
		items: items,
		operation: operation{
			fun: opFun,
			a:   opA,
			b:   opB,
		},
	}
}

func multOp(a, b int) int {
	return a * b
}

func addOp(a, b int) int {
	return a + b
}
