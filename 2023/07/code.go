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
	Map  map[string]int
	Bid  int
}

func (h hand) IsFiveOfKind() bool {
	for _, v := range h.Map {
		return v == 5
	}

	return false
}

func (h hand) IsFourOfKind() bool {
	for _, v := range h.Map {
		if v == 4 {
			return true
		}
	}

	return false
}

func (h hand) IsFullHouse() bool {
	has3 := false
	has2 := false
	for _, v := range h.Map {
		if v == 3 {
			has3 = true
		}
		if v == 2 {
			has2 = true
		}
	}

	return has2 && has3
}

func (h hand) ThreeOfKind() bool {
	for _, v := range h.Map {
		if v == 3 {
			return true
		}
	}

	return false
}

func (h hand) TwoPair() bool {
	pairs := 0
	for _, v := range h.Map {
		if v == 2 {
			pairs++
		}
	}

	return pairs == 2
}

func (h hand) OnePair() bool {
	pairs := 0
	for _, v := range h.Map {
		if v == 2 {
			pairs++
		}
	}

	return pairs == 1
}

func (h hand) HighCard() string {
	highCard := ""

	for _, c := range h.Hand {
		val := cardValue[string(c)]
		if highCard == "" || val > cardValue[highCard] {
			highCard = string(c)
		}
	}

	return highCard
}

func (h hand) Score() int {
	if h.IsFiveOfKind() {
		return 7
	}

	if h.IsFourOfKind() {
		return 6
	}

	if h.IsFullHouse() {
		return 5
	}

	if h.ThreeOfKind() {
		return 4
	}

	if h.TwoPair() {
		return 3
	}

	if h.OnePair() {
		return 2
	}

	return 1
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
	fmt.Println(hands)
	fmt.Println("Part 1:", part1f(hands))
}

func part1f(hands []hand) int {
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

		mapped := make(map[string]int, len(lineSplit[0]))
		for _, c := range lineSplit[0] {
			s := string(c)
			if val, exists := mapped[s]; exists {
				mapped[s] = val + 1
			} else {
				mapped[s] = 1
			}
		}

		hands[i] = hand{
			Hand: lineSplit[0],
			Map:  mapped,
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
