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
	opA, oldA := opParse(operatorStrings[0])
	opB, oldB := opParse(operatorStrings[2])

	monkeyOp := operation{
		fun:  opFun,
		a:    opA,
		b:    opB,
		oldA: oldA,
		oldB: oldB,
	}

	// Test
	monkeyTest := test{
		parameter:     lastIntOfLine(lines[2]),
		successTarget: lastIntOfLine(lines[3]),
		failureTarget: lastIntOfLine(lines[4]),
	}

	return Monkey{
		items:     items,
		operation: monkeyOp,
		test:      monkeyTest,
	}
}

func lastIntOfLine(line string) int {
	ls := strings.Split(line, " ")
	intString := ls[len(ls)-1]
	number, err := strconv.Atoi(intString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return number

}

func opParse(op string) (int, bool) {

	if op == "old" {
		return 0, true
	} else {
		opValue, err := strconv.Atoi(op)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return opValue, false
	}

}

func multOp(a, b int) int {
	return a * b
}

func addOp(a, b int) int {
	return a + b
}
