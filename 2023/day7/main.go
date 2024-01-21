package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type Hand struct {
	Bid   int
	Cards [5]Card
}

type Card int

const (
	T Card = iota + 10
	J
	Q
	K
	A
)

func main() {

	s := input.OpenFileBuffered("input.txt")

	hands := make([]Hand, 0, 1000)
	for s.Scan() {
		l := s.Text()
		hands = append(hands, HandFromLine(l))
	}

}

func HandFromLine(line string) Hand {

	s1 := strings.Fields(line)

	cardString := s1[0]
	bidStr := s1[1]

	bid, _ := strconv.Atoi(bidStr)
	hand := Hand{
		Bid: bid,
	}

	for i, r := range cardString {
		hand.Cards[i] = CardValueFromRune(r)
	}

	return hand
}

func CardValueFromRune(r rune) Card {
	// card value can be integer, or character
	switch r {
	case 'T':
		return T
	case 'J':
		return J
	case 'Q':
		return Q
	case 'K':
		return K
	case 'A':
		return A
	}

	n, err := strconv.Atoi(string(r))
	if err != nil {
		log.Fatal(err)
	}

	return Card(n)
}
