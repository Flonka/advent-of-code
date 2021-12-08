package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code-2021/input"
)

type BingoBoard struct {
	numbers  map[int]BingoNumberData
	rowMarks [5]int
	colMarks [5]int
}

// ProcessDrawnNumber processes the data with the drawn number
func (b *BingoBoard) ProcessDrawnNumber(n int) {
	num, hasNumber := b.numbers[n]
	if hasNumber {
		b.rowMarks[num.row]++
		b.colMarks[num.col]++
		num.called = true
		b.numbers[n] = num
	}
}

func (b *BingoBoard) HasBingo() bool {

	for i := 0; i < 5; i++ {
		if b.rowMarks[i] == 5 || b.colMarks[i] == 5 {
			return true
		}
	}

	return false
}

func (b *BingoBoard) Score() int {

	var sum int
	for n, data := range b.numbers {
		if !data.called {
			sum += n
		}
	}
	return sum
}

type BingoNumberData struct {
	row    int
	col    int
	called bool
}

func main() {

	drawNumbers, bingoBoards := ReadData()

	var bingoCount int
	for _, n := range drawNumbers {
		fmt.Println(n)
		for i := range bingoBoards {
			b := &bingoBoards[i]
			if !b.HasBingo() {
				b.ProcessDrawnNumber(n)
				if b.HasBingo() {
					fmt.Println("board#", i, b.Score()*n)
					bingoCount++
				}
			}
		}
		if bingoCount == len(bingoBoards) {
			fmt.Println("All boards bingoed")
			break
		}
	}

}

// ReadData returns draw numbers , and bingo boards from the input
func ReadData() ([]int, []BingoBoard) {

	lines := input.ReadLinesInFile("input")

	//first line is the draw numbers
	var drawNumbers []int
	for _, number := range strings.Split(lines[0], ",") {
		i, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		drawNumbers = append(drawNumbers, i)
	}

	var boards []BingoBoard
	i := 2
	for i < len(lines) {
		b := BingoBoard{}
		b.numbers = make(map[int]BingoNumberData, 25)
		for r := 0; r < 5; r++ {
			line := lines[i+r]
			numbers := strings.Fields(line)
			for c, nStr := range numbers {
				n, err := strconv.Atoi(nStr)
				if err != nil {
					fmt.Println(numbers, len(numbers))
					log.Fatal(err)
				}
				b.numbers[n] = BingoNumberData{
					row: r,
					col: c,
				}
			}
		}

		boards = append(boards, b)
		// move to next board
		i += 6
	}

	return drawNumbers, boards

}
