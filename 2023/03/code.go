package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var numberRegex = regexp.MustCompile(`(\d+)`)
var symbolRegex = regexp.MustCompile(`[^0-9.]`)

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	sc := bufio.NewScanner(strings.NewReader(input))
	gb := New(sc)
	if part2 {
		return part2f(gb)
	}
	// solve part 1 here

	return part1f(gb)
}

type coordinate struct {
	X   int
	Y   int
	Val string
}

func (c *coordinate) Equals(otherC coordinate) bool {
	return c.X == otherC.X && c.Y == otherC.Y
}

type thing struct {
	Coordinates []coordinate
	Val         string
}

type gameBoard struct {
	B       [][]string
	Symbols []thing
	Numbers []thing
}

func New(sc *bufio.Scanner) gameBoard {
	gb := gameBoard{
		B:       make([][]string, 0),
		Symbols: make([]thing, 0),
		Numbers: make([]thing, 0),
	}

	lineNum := 0
	for sc.Scan() {
		line := sc.Text()

		// fill the board
		chars := make([]string, len(line))
		for i, c := range line {
			chars[i] = string(c)
		}
		gb.B = append(gb.B, chars)

		for _, match := range numberRegex.FindAllStringIndex(line, -1) {
			coords := []coordinate{}
			for i := match[0]; i < match[1]; i++ {
				coords = append(coords, coordinate{X: i, Y: lineNum})
			}
			gb.Numbers = append(gb.Numbers, thing{Coordinates: coords, Val: line[match[0]:match[1]]})
		}

		for _, match := range symbolRegex.FindAllStringIndex(line, -1) {
			coords := []coordinate{}
			for i := match[0]; i < match[1]; i++ {
				coords = append(coords, coordinate{X: i, Y: lineNum})
			}
			gb.Symbols = append(gb.Symbols, thing{Coordinates: coords, Val: line[match[0]:match[1]]})
		}

		lineNum++
	}

	return gb
}

func (gb *gameBoard) Print() {
	for _, row := range gb.B {
		fmt.Println(row)
	}

	fmt.Println("Number of numbers:", len(gb.Numbers))
	fmt.Println("Number of symbols:", len(gb.Symbols))
}

var combinations = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func (gb *gameBoard) HasThingAdjacent(x, y int, f func(string) bool) bool {
	for _, comb := range combinations {
		currX := comb[0] + x
		currY := comb[1] + y

		if currX < 0 || currY < 0 {
			continue
		}

		if currY > len(gb.B)-1 || currX > len(gb.B[y])-1 {
			continue
		}

		if f(gb.B[currY][currX]) {
			return true
		}
	}

	return false
}

func part1f(gb gameBoard) int {
	sum := 0
	for _, num := range gb.Numbers {
		coordsToCheck := make([]coordinate, 0)
		for _, coord := range num.Coordinates {
			for _, combo := range combinations {
				currX := coord.X + combo[0]
				currY := coord.Y + combo[1]
				coordsToCheck = append(coordsToCheck, coordinate{X: currX, Y: currY})
			}
		}
		hasSymbol := false
		for _, coord := range coordsToCheck {
			foundMatch := false
			for _, symb := range gb.Symbols {
				if symb.Coordinates[0].Equals(coord) {
					hasSymbol = true
					foundMatch = true
					break
				}
			}
			if foundMatch {
				break
			}
		}

		if hasSymbol {
			conv, err := strconv.Atoi(num.Val)
			if err != nil {
				log.Fatal(conv)
			}
			sum += conv
		}

	}
	return sum
}

func part2f(gb gameBoard) int {
	sum := 0
	for _, symb := range gb.Symbols {
		if symb.Val != "*" {
			continue
		}

		coord := symb.Coordinates[0]
		coordsToCheck := make([]coordinate, 0)
		for _, combo := range combinations {
			currX := coord.X + combo[0]
			currY := coord.Y + combo[1]
			coordsToCheck = append(coordsToCheck, coordinate{X: currX, Y: currY})
		}

		matchedNumbers := make([]int, 0)
		for _, num := range gb.Numbers {
			for _, numCoord := range num.Coordinates {
				didMatch := false
				for _, coordToCheck := range coordsToCheck {
					if coordToCheck.Equals(numCoord) {
						conv, err := strconv.Atoi(num.Val)
						if err != nil {
							log.Fatal(conv)
						}
						matchedNumbers = append(matchedNumbers, conv)
						didMatch = true
						break
					}
				}
				if didMatch {
					break
				}
			}
		}

		if len(matchedNumbers) == 2 {
			mult := matchedNumbers[0] * matchedNumbers[1]
			sum += mult
		}
	}

	return sum
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isSymbol(s string) bool {
	return symbolRegex.MatchString(s)
}
