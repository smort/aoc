package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var isJWild = false

var cardValue = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type hand struct {
	Hand string
	Bid  int
}

func (h hand) Score() int {
	// sort strongest to weakest
	cards := strings.Split(h.Hand, "")
	slices.SortFunc(cards, func(a, b string) int {
		aVal := cardValue[a]
		bVal := cardValue[b]

		if bVal > aVal {
			return 1
		}

		if aVal > bVal {
			return -1
		}

		return 0
	})

	numWild := 0
	if isJWild {
		for i := len(cards) - 1; i >= 0; i-- {
			if cards[i] == "J" {
				numWild++
			} else {
				break
			}
		}
	}

	// scorekeeping
	currCount := 1
	has5 := false
	has4 := false
	hasFullHouse := false
	hasThreeofAKind := false
	numPairs := 0

	checkHand := func() {
		if currCount == 5 {
			has5 = true
		}

		if currCount == 4 {
			has4 = true
		}

		if currCount == 3 {
			hasThreeofAKind = true
		}

		if currCount == 2 {
			numPairs++
		}

		if hasThreeofAKind && numPairs > 0 {
			hasFullHouse = true
		}
	}

	prevCard := cards[0]
	for i := 1; i < len(cards); i++ {
		currCard := cards[i]
		if isJWild && currCard == "J" {
			break
		}

		if currCard != prevCard {
			checkHand()
			currCount = 1
		} else {
			currCount++
		}

		prevCard = currCard
	}

	checkHand()

	if isJWild && numWild > 0 {
		if numWild == 5 {
			has5 = true
		}

		if numWild == 4 {
			has5 = true
		}

		if numWild == 3 {
			if numPairs == 1 {
				has5 = true
			} else {
				has4 = true
			}
		}

		if numWild == 2 {
			if hasThreeofAKind {
				has5 = true
			} else if numPairs == 1 {
				has4 = true
			} else {
				hasThreeofAKind = true
			}
		}

		if numWild == 1 {
			if has4 {
				has5 = true
			} else if hasThreeofAKind {
				has4 = true
			} else if numPairs == 2 {
				hasFullHouse = true
			} else if numPairs == 1 {
				hasThreeofAKind = true
			} else {
				numPairs = 1
			}
		}
	}

	score := 1
	if has5 {
		score = 7
	} else if has4 {
		score = 6
	} else if hasFullHouse {
		score = 5
	} else if hasThreeofAKind {
		score = 4
	} else if numPairs == 2 {
		score = 3
	} else if numPairs == 1 {
		score = 2
	}

	fmt.Printf("hand: %v sorted: %v, score: %v\n", h.Hand, strings.Join(cards, ""), score)

	return score
}

func (h hand) Compare(other hand) int {
	// return -1 if this hand is weaker, 1 if it is stronger
	otherScore := other.Score()
	thisScore := h.Score()

	if otherScore > thisScore {
		return -1
	}

	if otherScore == thisScore {
		for i := 0; i < len(h.Hand); i++ {
			otherCardScore := cardValue[string(other.Hand[i])]
			thisCardScore := cardValue[string(h.Hand[i])]

			if otherCardScore == thisCardScore {
				continue
			}

			if otherCardScore > thisCardScore {
				return -1
			}

			// basically return 1
			break
		}
	}

	return 1
}

func main() {
	input := readFile("input-user.txt")
	hands := parseInput(input)

	fmt.Println("Part 1:", part1f(hands))
	fmt.Println("Part 2:", part2f(hands))
}

func part1f(hands []hand) int {
	slices.SortFunc(hands, func(a, b hand) int {
		return a.Compare(b)
	})

	score := hands[0].Bid
	for i := 1; i < len(hands); i++ {
		score += hands[i].Bid * (i + 1)
	}

	return score
}

func part2f(hands []hand) int {
	isJWild = true
	cardValue["J"] = 1

	slices.SortFunc(hands, func(a, b hand) int {
		return a.Compare(b)
	})
	fmt.Println(hands)

	score := hands[0].Bid
	for i := 1; i < len(hands); i++ {
		score += hands[i].Bid * (i + 1)
	}

	return score
}

func parseInput(input []string) []hand {
	hands := make([]hand, len(input)-1)
	for i, line := range input {
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) != 2 {
			continue
		}
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			log.Fatal(err)
		}

		hands[i] = hand{
			Hand: lineSplit[0],
			Bid:  bid,
		}
	}

	return hands
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}
