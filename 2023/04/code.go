package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	sc := bufio.NewScanner(strings.NewReader(input))
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2f(sc)
	}
	// solve part 1 here
	return part1f(sc)
}

func part1f(sc *bufio.Scanner) int {
	lineNum := 0
	sum := 0
	for sc.Scan() {
		lineNum++
		line := sc.Text()
		winGameSplit := strings.Split(strings.Split(line, ":")[1], "|")
		winnersStr := strings.Split(strings.TrimSpace(winGameSplit[0]), " ")
		numbersStr := strings.Split(strings.TrimSpace(winGameSplit[1]), " ")

		winnersSet := make(map[string]struct{}, len(winnersStr))
		for _, winnerStr := range winnersStr {
			if winnerStr == "" {
				continue
			}
			winnersSet[winnerStr] = struct{}{}
		}

		// fmt.Println(winnersSet)

		matches := 0
		for _, numberStr := range numbersStr {
			if numberStr == "" {
				continue
			}
			if _, exists := winnersSet[numberStr]; exists {
				matches++
			}
		}

		// fmt.Println("matches", matches)

		score := 0
		for matches > 0 {
			matches--

			if score == 0 {
				score = 1
				continue
			}

			score *= 2
		}

		sum += score
		// fmt.Println(score)
	}
	return sum
}

func part2f(sc *bufio.Scanner) int {
	lineNum := 0
	cardCount := 0
	gameRepeats := make(map[int]int, 0)

	for sc.Scan() {
		lineNum++
		line := sc.Text()
		gameNameSplit := strings.Split(line, ":")
		winGameSplit := strings.Split(gameNameSplit[1], "|")
		winnersStr := strings.Split(strings.TrimSpace(winGameSplit[0]), " ")
		numbersStr := strings.Split(strings.TrimSpace(winGameSplit[1]), " ")

		gameNum := getGameNum(gameNameSplit[0])

		winnersSet := make(map[string]struct{}, len(winnersStr))
		for _, winnerStr := range winnersStr {
			if winnerStr == "" {
				continue
			}
			winnersSet[winnerStr] = struct{}{}
		}

		matches := 0
		for _, numberStr := range numbersStr {
			if numberStr == "" {
				continue
			}
			if _, exists := winnersSet[numberStr]; exists {
				matches++
			}
		}

		mult := gameRepeats[gameNum]
		cardCount += 1 + mult

		currGame := gameNum
		for matches > 0 {
			currGame++
			matches--
			if currGame > 205 {
				continue
			}

			val, exists := gameRepeats[currGame]
			if !exists {
				gameRepeats[currGame] = 1 + mult
			} else {
				gameRepeats[currGame] = val + 1 + mult
			}

		}
		fmt.Printf("%v: %v\n\n", gameNum, gameRepeats)
	}

	return cardCount
}

func getGameNum(s string) int {
	c, _ := strings.CutPrefix(s, "Card ")
	num, err := strconv.Atoi(strings.TrimSpace(c))
	if err != nil {
		log.Fatal(err)
	}

	return num
}

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int, bool) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	if l == 0 {
		return s, 0, false
	}
	return s[:l-1], s[l-1], true
}
