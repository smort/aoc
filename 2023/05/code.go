package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type sourceDestMap struct {
	Source      int64
	Destination int64
	Length      int64
}

type maps []sourceDestMap

func main() {
	input := readFile("input-user.txt")
	// fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int64 {
	seeds, maps := parseInput(input)
	result := calculateScores(seeds, maps)
	return result
}

func part2(input []string) int64 {
	seeds, maps := parseInput(input)
	var minimum int64 = math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		newSeeds := make([]int64, 0)
		// fmt.Println(seeds[i], seeds[i+1])
		for j := seeds[i]; j <= (seeds[i] + seeds[i+1] - 1); j++ {
			// fmt.Println("j", j)
			newSeeds = append(newSeeds, j)
		}
		// fmt.Println("done", i)
		result := calculateScores(newSeeds, maps)
		minimum = min(minimum, result)
	}
	fmt.Println("done done")

	return minimum
}

func calculateScores(seeds []int64, data []maps) int64 {
	fmt.Println("num of seeds", len(seeds))
	var minumum int64 = math.MaxInt64
	// results := make([]int64, 0, len(seeds))
	progress := 0
	for _, seed := range seeds {
		// fmt.Println("starting seed", seed)
		val := seed
		for _, d := range data {
			var out int64 = -1
			for _, m := range d {
				low := m.Source
				high := m.Source + m.Length - 1

				if val >= low && val <= high {
					out = (val - low) + m.Destination
					break
				}
			}
			if out == -1 {
				out = val
			}
			val = out
		}
		progress++
		if progress%100_000_000 == 0 {
			fmt.Println(progress)
		}
		minumum = min(minumum, val)
	}

	return minumum
}

// first return is seeds, second is maps
func parseInput(input []string) ([]int64, []maps) {
	seedSplit := strings.Split(input[0], "seeds: ")
	seedStrs := strings.Split(strings.TrimSpace(seedSplit[1]), " ")
	seeds := make([]int64, 0, len(seedStrs))
	for _, seed := range seedStrs {
		parsed, err := strconv.ParseInt(seed, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, parsed)
	}

	m := make([]maps, 0, 7)
	var currThings maps

	i := 1
	for i < len(input) {
		line := input[i]
		if strings.TrimSpace(line) == "" || strings.Contains(line, ":") {
			if currThings != nil {
				m = append(m, currThings)
				currThings = nil
			}
			i++
			continue
		}
		if currThings == nil {
			currThings = make(maps, 0)
		}

		vals := strings.Split(line, " ")
		thing := sourceDestMap{}
		for i, val := range vals {
			valInt, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if i == 0 {
				thing.Destination = valInt
			} else if i == 1 {
				thing.Source = valInt
			} else {
				thing.Length = valInt
			}
		}
		currThings = append(currThings, thing)
		i++
	}

	return seeds, m
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}
