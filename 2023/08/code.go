package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const destination = "ZZZ"

func main() {
	// input := readFile("input-example.txt")
	input := readFile("input-user.txt")
	directions, layout := parseInput(input)

	fmt.Println("Part 1:", part1f(directions, layout))

	// input2 := readFile("input-example-part2.txt")
	input2 := readFile("input-user.txt")
	directions2, layout2 := parseInput(input2)
	fmt.Println("Part 2:", part2f(directions2, layout2))
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}

func parseInput(lines []string) ([]string, map[string]map[string]string) {
	directions := strings.Split(lines[0], "")

	table := make(map[string]map[string]string, len(lines)-2)

	for i := 2; i < len(lines)-1; i++ {
		currLine := lines[i]
		if currLine == "" {
			continue
		}

		equalSplit := strings.Split(currLine, " = ")
		currNode := strings.TrimSpace(equalSplit[0])
		optionsSplit := strings.Split(equalSplit[1], ",")
		left := strings.TrimSpace(strings.Replace(optionsSplit[0], "(", "", -1))
		right := strings.TrimSpace(strings.Replace(optionsSplit[1], ")", "", -1))

		table[currNode] = map[string]string{
			"L": left,
			"R": right,
		}
	}

	return directions, table
}

func part1f(directions []string, layout map[string]map[string]string) int {
	steps := 0
	i := 0
	currPos := "AAA"
	for {
		steps++
		currDirection := directions[i]
		nextNode := layout[currPos][currDirection]
		currPos = nextNode

		if nextNode == destination {
			break
		}

		i++
		if i == len(directions) {
			i = 0
		}
	}

	return steps
}

type pather struct {
	NumSteps        int
	CurrentPosition string
}

func part2f(directions []string, layout map[string]map[string]string) int {
	pathers := []*pather{}
	for k, _ := range layout {
		if strings.HasSuffix(k, "A") {
			p := pather{NumSteps: 0, CurrentPosition: k}
			pathers = append(pathers, &p)
		}
	}

	i := 0
	for {
		currDirection := directions[i]
		allDone := true
		for _, pather := range pathers {
			if strings.HasSuffix(pather.CurrentPosition, "Z") {
				continue
			}

			allDone = false
			nextPos := layout[pather.CurrentPosition][currDirection]
			pather.NumSteps = pather.NumSteps + 1
			pather.CurrentPosition = nextPos

			// if i == 2 {
			// 	fmt.Println(nextPos)
			// }

			// fmt.Printf("Step %v: %v\n", steps, nextPos)
		}

		if allDone {
			break
		}

		i++
		if i == len(directions) {
			i = 0
		}
	}

	numSteps := make([]int, 0, len(pathers))
	for _, pather := range pathers {
		numSteps = append(numSteps, pather.NumSteps)
	}
	return LCM(numSteps[0], numSteps[1:])
}

// GCD greatest common divisor via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Less Common Multiple via GCD
func LCM(first int, integers []int) int {
	result := first * integers[0] / GCD(first, integers[0])
	for i := 1; i < len(integers); i++ {
		result = LCM(result, []int{integers[i]})
	}
	return result
}
