package main

import (
	"bufio"
	"log"
	"regexp"
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
	games := parseGames(sc)
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2f(games)
	}

	return part1f(games)
}

type game struct {
	ID    int
	Blue  int
	Red   int
	Green int
}

func (g game) IsValid(blue, green, red int) bool {
	return blue >= g.Blue && green >= g.Green && red >= g.Red
}

func (g game) Pow() int {
	return g.Blue * g.Red * g.Green
}

var gameNum = regexp.MustCompile(`Game (\d+)`)
var blueNum = regexp.MustCompile(`(\d+) blue?`)
var greenNum = regexp.MustCompile(`(\d+) green?`)
var redNum = regexp.MustCompile(`(\d+) red?`)

func parseGames(sc *bufio.Scanner) []game {
	games := make([]game, 0)
	for sc.Scan() {
		g := game{}
		line := sc.Text()
		match := gameNum.FindStringSubmatch(line)
		i, err := strconv.Atoi(string(match[1]))
		if err != nil {
			log.Fatalf("failed getting game num. line: %s with match %v with err %v", line, match[1], err)
		}
		g.ID = i

		matches := blueNum.FindAllStringSubmatch(line, -1)
		maxBlue := -1
		for _, match := range matches {
			i, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("failed while getting blue. line: %s with match %v with err %v", line, match[1], err)
			}
			if i > maxBlue {
				maxBlue = i
			}
		}
		g.Blue = maxBlue

		greenMatches := greenNum.FindAllStringSubmatch(line, -1)
		maxGreen := -1
		for _, match := range greenMatches {
			i, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("failed while getting green. line: %s with match %v with err %v", line, match[1], err)
			}
			if i > maxGreen {
				maxGreen = i
			}
		}
		g.Green = maxGreen

		redMatches := redNum.FindAllStringSubmatch(line, -1)
		maxRed := -1
		for _, match := range redMatches {
			i, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("failed while getting red. line: %s with match %v with err %v", line, match[1], err)
			}
			if i > maxRed {
				maxRed = i
			}
		}
		g.Red = maxRed

		games = append(games, g)
	}
	return games
}

func part1f(games []game) int {
	sum := 0
	count := 0
	for _, game := range games {
		if game.IsValid(14, 13, 12) {
			count += 1
			sum += game.ID
		}
	}
	return sum
}

func part2f(games []game) int {
	sum := 0
	for _, game := range games {
		sum += game.Pow()
	}
	return sum
}
