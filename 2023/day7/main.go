package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

const handSize int = 5

type Hand struct {
	Bid   int
	Cards [handSize]Card
	Type  HandType
}

type HandType int

const (
	NotEvaluated HandType = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

// EvaluateHand sets the hands type from the current cards
func (h *Hand) EvaluateHand() {

	// Find relative handtype from one card comparing to other in hand
	var cardEval = [handSize]HandType{}
	highest := NotEvaluated
	for i := 0; i < handSize; i++ {
		evC := evaluateCard(h, i)
		cardEval[i] = evC
		if evC > highest {
			highest = evC
		}
	}

	// Start by setting handtype to highest found
	h.Type = highest

	// If four or five of a kind
	if highest >= FourOfAKind {
		return
	}

	if highest == ThreeOfAKind {
		// check for fullhouse
		for i := 0; i < handSize; i++ {
			if cardEval[i] == OnePair {
				h.Type = FullHouse
			}
		}
	}

	if highest == OnePair {
		// Check for two pairs

		pairCount := 0
		for i := 0; i < handSize; i++ {
			// four one pairs, means two pair
			if cardEval[i] == OnePair {
				pairCount++
			}
		}

		if pairCount == 4 {
			h.Type = TwoPair
		}
	}

}

func evaluateCard(h *Hand, cardPosition int) HandType {

	duplicates := 1
	v := h.Cards[cardPosition]
	for i := 0; i < handSize; i++ {
		if i == cardPosition {
			continue
		}

		if v == h.Cards[i] {
			duplicates++
		}

	}

	switch duplicates {
	case 1:
		return HighCard
	case 2:
		return OnePair
	case 3:
		return ThreeOfAKind
	case 4:
		return FourOfAKind
	case 5:
		return FiveOfAKind

	}
	return NotEvaluated
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
		return
	}

	// Sort all hands on the type and special rule of even , first high card.
	// Iterate and calculate the rank ( i * hand.Bid )

	fmt.Println(hands[:3])

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

	hand.EvaluateHand()

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
